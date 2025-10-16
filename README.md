
# 🥡 BiteRush

**BiteRush** is a backend service written in **Go** using **Go-Kit** for service architecture, **Gorilla Mux** for routing, and **Jet ORM** for type-safe SQL.  
It simulates a basic food-delivery system managing **users**, **riders**, **restaurants**, **menus**, **orders**, and **ratings**.

---

## 🚀 Features
- RESTful HTTP API built with **Go-Kit** + **Gorilla Mux**
- **MySQL** database with indexed, relational schema
- **Jet ORM** for generating and executing SQL queries safely
- Modular structure (transport / endpoint / database layers)
- Configurable via `config.yaml`
- Dockerized setup for MySQL
- Graceful shutdown handling via OS signals

---

## 🧱 Architecture Overview

| Layer | Package | Responsibility |
|-------|----------|----------------|
| **cmd/server** | `cmd/server/main.go` | Application entry point, config load, server start |
| **endpoint** | `internal/transport/endpoint` | Business endpoints wrapping Bite manager |
| **service (manager)** | `bite/` | Business logic implementation |
| **database** | `internal/db` | Database operations using Jet |
| **http transport** | `internal/transport/httptransport` | HTTP routing + decode/encode functions |
| **config** | `cmd/config.yaml` | Application configuration |
| **schema.sql** | root | Database schema definition |

---

## ⚙️ Tech Stack
- **Language:** Go 1.21+
- **Frameworks/Libraries:**
  - `go-kit/kit` – service & transport architecture
  - `gorilla/mux` – HTTP routing
  - `go-jet/jet` – SQL builder & ORM
  - `k8s.io/klog/v2` – structured logging
  - `spf13/viper` – configuration management
- **Database:** MySQL 8.0
- **Containerization:** Docker + Docker Compose

---

## 🧩 Directory Structure
```
biterush/
├── cmd/
│   └── server/
│       └── main.go
├── bite/
│   └── manager.go
├── internal/
│   ├── db/
│   ├── transport/
│   │   ├── endpoint/
│   │   └── httptransport/
│   └── internal.go
├── config.yaml
├── schema.sql
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
```

---

## 🛠️ Setup & Installation

### 1️⃣ Prerequisites
- Go 1.25.3 or newer  
- Docker and Docker Compose  
- Make (optional but recommended)

### 2️⃣ Start MySQL via Docker
```bash
make up-db
```
This will:
- Launch MySQL 8.0 container
- Apply the schema defined in `schema.sql`

### 3️⃣ Generate Jet ORM code
```bash
make gen-jet
```

### 4️⃣ Build the application
```bash
make build
```

### 5️⃣ Run the service locally
```bash
./.builds/biterush
```
or via Docker:
```bash
make dockerbuild
docker run -p 8080:8080 biterush:latest
```

### 6️⃣ Access API
Once running, visit:
```
http://localhost:8080/biterush
```

---

## 🌐 Example Routes

| Method | Endpoint | Description |
|--------|-----------|-------------|
| `POST` | `/biterush/users` | Create a new user |
| `POST` | `/biterush/riders` | Create a new rider |
| `PATCH` | `/biterush/riders/{id}/location` | Update rider’s location |
| `POST` | `/biterush/restaurants` | Create a new restaurant |
| `GET` | `/biterush/restaurants/{id}/menu` | Get restaurant menu |
| `POST` | `/biterush/orders` | Create new order |
| `GET` | `/biterush/orders` | List orders (filtered by user/rider/restaurant) |
| `POST` | `/biterush/orders/{id}/accept` | Accept an order |

---

## 🧰 Makefile Commands

| Command | Description |
|----------|-------------|
| `make fmt` | Format Go code |
| `make clean` | Clean build and generated files |
| `make build` | Compile app binary |
| `make dockerbuild` | Build Docker image |
| `make dockerpush` | Push Docker image |
| `make up-db` | Start MySQL via docker-compose |
| `make down-db` | Stop database |
| `make gen-jet` | Generate Jet ORM models |

---

## 📦 Configuration (`config.yaml`)
Example:
```yaml
server:
  port: 8080

database:
  user: root
  password: root
  host: 127.0.0.1
  port: 3306
  name: biterush
```

---

## 🧪 Example Request
```bash
curl -X POST http://localhost:8080/biterush/users   -H "Content-Type: application/json"   -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "9999999999",
    "address": "Delhi"
  }'
```

Response:
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com"
}
```
