package main

import (
	"biterush/cmd"
	"biterush/internal"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-jet/jet/v2/generator/metadata"
	"github.com/go-jet/jet/v2/generator/mysql"
	"github.com/go-jet/jet/v2/generator/template"
	mysql2 "github.com/go-jet/jet/v2/mysql"
)

func main() {

	cfg := createDatabaseConfig()
	databaseDSN := cfg.CreateDSN()

	internal.MigrateSQLs(databaseDSN)

	err := mysql.GenerateDSN(
		databaseDSN,
		"./generated",
		template.Default(mysql2.Dialect).
			UseSchema(func(schemaMetaData metadata.Schema) template.Schema {
				return template.DefaultSchema(schemaMetaData).
					UsePath("database").
					UseModel(template.DefaultModel().UseTable(func(table metadata.Table) template.TableModel {
						builder := template.DefaultTableModel(table)
						builder.Skip = table.Name == "schema_migrations"

						if !builder.Skip {
							log.Printf("\t > Generating code for model %q\n", builder.TypeName)
						}

						return builder
					})).
					UseSQLBuilder(
						template.DefaultSQLBuilder().UseTable(func(table metadata.Table) template.TableSQLBuilder {
							builder := template.DefaultTableSQLBuilder(table)
							builder.Skip = table.Name == "schema_migrations"

							if !builder.Skip {
								fmt.Printf("\t > Generating code for table %q\n", table.Name)
							}

							return builder
						}),
					)
			}),
	)
	if err != nil {
		log.Fatalf("cannot generate data sources: %v", err)
	}
}

func createDatabaseConfig() cmd.DatabaseConfiguration {

	cfg := cmd.ResolveConfig()

	port := 3306
	portStr := defaultEnvVar("DB_PORT", "3306")
	if p, err := strconv.Atoi(portStr); err == nil && p > 0 {
		port = p
	}

	return cmd.DatabaseConfiguration{
		Name:     defaultEnvVar("DB_NAME", cfg.Database.Name),
		User:     defaultEnvVar("DB_USER", cfg.Database.User),
		Password: defaultEnvVar("DB_PASS", cfg.Database.Password),
		Host:     defaultEnvVar("DB_HOST", cfg.Database.Host),
		Port:     port,
	}
}

func defaultEnvVar(key, defaultVal string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}

	return defaultVal
}
