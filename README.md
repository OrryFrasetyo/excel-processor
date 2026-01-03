# 🚀 High Performance Excel/CSV Processor (Go Backend)

![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)
![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED?style=flat&logo=docker)
![Architecture](https://img.shields.io/badge/Architecture-Clean%20%2F%20Hexagonal-orange)

A high-performance backend service designed for **Bulk Import** data processing using Golang. Capable of processing **10,000+ rows** asynchronously without blocking the main thread, leveraging **Worker Pool Pattern** and **Goroutines**.

## 🌟 Key Features (Why this project matters?)

- **⚡ Concurrency & Parallelism:** Uses **Worker Pool Pattern** to keep CPU/RAM usage low even under high workloads.
- **🛡️ Thread-Safe Operations:** Implements **Mutex** to prevent _Race Conditions_ when tracking status (Success/Fail counters).
- **🏗️ Simplified Clean Architecture:** Modular code structure (Handler -> Worker -> Repository) that simplifies testing and maintenance.
- **🐳 Dockerized Environment:** Ready-to-use PostgreSQL database and application setup with a single command.
- **🧪 Unit Tested:** Worker logic is fully tested using `Testify` and `Mocking` to simulate success and _partial failure_ scenarios.

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Framework:** Gin Gonic (HTTP Web Framework)
- **Database:** PostgreSQL 15
- **ORM:** GORM (with optimized `pgx` driver)
- **Concurrency:** Native Goroutines, Channels, Sync.WaitGroup, Sync.Mutex
- **Config:** Godotenv (.env)
- **Testing:** Testify (Assert & Mock)

## 📂 Project Structure

```text
├── cmd
│   └── api/          # Entry point (Wiring Dependency Injection)
├── internal
│   ├── entity/       # Database Models
│   ├── handler/      # HTTP Handlers (Gin)
│   ├── repository/   # Database Access Layer
│   └── worker/       # CORE LOGIC (Concurrency & Worker Pool)
├── docker-compose.yml
└── ...
```

## 🚀 How to Run

### Prerequisite

- Docker & Docker Compose

### Quick Start (Docker)

1. Clone this repository.
2. Create .env file (copy from .env.example if available).
3. Jalankan perintah:
   
   ```bash
   docker-compose up -d --build
   ```
4. The server will run at `http://localhost:8080`

### Manual Run (Local)

If you want to run without Docker (ensure your local PostgreSQL is running):

```bash
go run cmd/api/main.go
```

## 📡 API Endpoints

### 1. Upload & Process Data

Uploads a CSV file for background processing.

- URL: `/process-data`
- Method: `POST`
- Response:

```json
{
  "message": "Permintaan diterima! Data sedang diproses di background.",
  "file": "large_students.csv",
  "status": "Check terminal logs for progress"
}
```

### 2. Health Check

- URL: `/ping`
- Method: `GET`

## 🧪 Running Tests
This project uses testify for Unit Testing on the Worker layer.

```bash
go test ./internal/worker -v
```

#### Author: Orry Frasetyo - Software Engineer
