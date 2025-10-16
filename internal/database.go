package internal

import (
	"biterush/cmd"
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"log"

	"github.com/go-jet/jet/v2/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"k8s.io/klog/v2"
)

//go:embed migrations/*.sql
var sqlFiles embed.FS

const SQLLogLevel = 6

const TransactionContextKey = ":biterush:transaction"

func LogStatement(ctx context.Context, statement mysql.Statement) {

	logger := klog.FromContext(ctx)
	if !logger.V(SQLLogLevel).Enabled() {
		return
	}

	stmt, p := statement.Sql()
	logger.Info("executing statement", "statement", stmt, "params", p)
}

func GetTransaction(ctx context.Context) (*sql.Tx, error) {
	tx, ok := ctx.Value(TransactionContextKey).(*sql.Tx)
	if ok {
		return tx, nil
	}

	return nil, errors.New("no transaction found")
}

func NewTransactionContext(ctx context.Context, db *sql.DB, opts *sql.TxOptions) (context.Context, *sql.Tx, error) {
	tx, ok := ctx.Value(TransactionContextKey).(*sql.Tx)
	if ok {
		return ctx, tx, nil
	}

	tx, err := db.BeginTx(ctx, opts)
	if err != nil {
		return nil, nil, err
	}

	return context.WithValue(ctx, TransactionContextKey, tx), tx, nil
}

func CreateMysqlConnection(cfg cmd.Config) (*sql.DB, error) {
	dsn := cfg.Database.CreateDSN()
	MigrateSQLs(dsn)

	return sql.Open("mysql", dsn)
}

func MigrateSQLs(dsn string) {

	dsn = fmt.Sprintf("%s&multiStatements=true", dsn)

	source, err := iofs.New(sqlFiles, "migrations")
	if err != nil {
		log.Fatalf("cannot prepare migration source: %v", err)
	}

	klog.V(3).Infof("opening mysql connection for migration...")
	mig, err := migrate.NewWithSourceInstance("iofs", source, "mysql://"+dsn)
	if err != nil {
		log.Fatalf("cannot prepare migration instance: %v", err)
	}

	logMigrationVersions("[before migration]", mig)

	if err := mig.Up(); err != nil && err != migrate.ErrNoChange {
		klog.Errorf("up migration failed: %v", err)

		if err := mig.Down(); err != nil {
			klog.Errorf("migration rollback failed: %v", err)
		} else {
			klog.Info("rolled back to last version")
		}

		panic(err)
	}

	logMigrationVersions("[after migration] ", mig)
}

func logMigrationVersions(logHead string, mig *migrate.Migrate) {
	if v, d, err := mig.Version(); err != nil && err != migrate.ErrNilVersion {
		klog.Fatalf("cannot load migration version: %v", err)
	} else {
		klog.Infof("%s: current database version is %d (dirty=%v)", logHead, v, d)
	}
}
