# Job Scraper

<div align="center">

# **Job Scraper**

### Go-based Job Scraping API — Aggregate job listings from Ashby, Lever, Amazon, and Atlassian.

[![Go](https://img.shields.io/badge/Go_1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Gin](https://img.shields.io/badge/Gin-009688?style=for-the-badge&logo=gin&logoColor=white)](https://gin-gonic.com/)
[![SQLite](https://img.shields.io/badge/SQLite-003B57?style=for-the-badge&logo=sqlite&logoColor=white)](https://www.sqlite.org/)

</div>

---

## ⚡ Overview

Job Scraper is a Go-based web scraping service that aggregates job listings from multiple ATS (Applicant Tracking System) platforms. It scrapes jobs from Ashby, Lever, Amazon, and Atlassian, normalizes the data, and stores it in SQLite for easy querying via a REST API.

---

## ✨ Features

| Feature | Description |
|---------|-------------|
| 🌐 **Multi-Platform Scraping** | Aggregate from Ashby, Lever, Amazon, and Atlassian |
| 📦 **SQLite Storage** | Persistent job storage with full-text search support |
| 🔄 **Parallel Sync** | Concurrent scraping with rate limiting per platform |
| 🔍 **Filtered Queries** | Search by title, company, location with pagination |
| 🔐 **Protected Sync** | Bearer token authentication for sync endpoints |
| 🔔 **Change Detection** | Track new, updated, and removed jobs |

---

## 📊 Supported Platforms

| Platform | Scraper Type | Companies |
|----------|--------------|-----------|
| **Ashby** | API-based | 150+ companies |
| **Lever** | API-based | CRED, ShieldAI |
| **Amazon** | Custom | Amazon |
| **Atlassian** | Custom | Atlassian |

---

## 🤖 API Endpoints

### 📋 Query Jobs

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/` | Health check |
| `GET` | `/getallJobsFromSQL` | Get paginated jobs with filters |
| `GET` | `/companies` | Get all companies with active jobs |
| `GET` | `/locations` | Get all unique locations |
| `GET` | `/job/:id` | Get job by ID |

### 🔄 Sync Jobs

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/syncall` | Trigger full sync (requires auth) |
| `POST` | `/sync` | Trigger sync with password in body |

---

## 💻 API Usage

### Get All Jobs

```bash
GET /getallJobsFromSQL?search=engineer&company=vercel&location=Remote&sort=newest&limit=20&offset=0
```

**Response:**

```json
{
  "jobs": [
    {
      "id": 1,
      "jobName": "Senior Software Engineer",
      "companyName": "Vercel",
      "location": "Remote",
      "description": "...",
      "applyLink": "https://vercel.com/careers/...",
      "meta": {
        "department": "Engineering",
        "team": "Platform",
        "employmentType": "Full-time",
        "remote": true,
        "source": "ashby"
      }
    }
  ],
  "offset": 0,
  "limit": 20,
  "total": 150
}
```

### Get Companies

```bash
GET /companies
```

**Response:**

```json
{
  "companies": ["1Password", "Abridge", "Airtable", "Alan", ...]
}
```

### Get Locations

```bash
GET /locations
```

**Response:**

```json
{
  "locations": ["Remote", "San Francisco", "New York", "London", ...]
}
```

### Sync All Jobs

```bash
# Using Authorization header
curl -X GET https://your-api.com/syncall \
  -H "Authorization: Bearer your_password"

# Or using JSON body
curl -X POST https://your-api.com/sync \
  -H "Content-Type: application/json" \
  -d '{"password": "your_password"}'
```

**Response:**

```json
{
  "message": "synced successfully",
  "count": 1250,
  "results": [
    {"company": "Amazon", "status": "success", "count": 45},
    {"company": "Atlassian", "status": "success", "count": 32},
    {"company": "1Password", "status": "success", "count": 12},
    ...
  ]
}
```

---

## ⚙️ Configuration

Config file: `.env` (copy from `.env.example`)

```bash
# Required
SYNC_PASSWORD=your_secure_password

# CORS (defaults to http://localhost:3000)
CORS_ALLOWED_ORIGIN=https://your-frontend.com

# Database (defaults to ./jobs.db)
DB_PATH=./jobs.db

# Company Configuration (optional - defaults to companies.json)
ASHBY_COMPANIES='[{"Company":"Vercel","AshbySlug":"vercel","Enabled":true}]'
ASHBY_COMPANIES_COMMA="Vercel:vercel,Linear:linear"
```

---

## 🚀 Deployment

### Docker

```bash
# Build
docker build -t jobscraper .

# Run
docker run -p 8080:8080 \
  -e SYNC_PASSWORD=your_password \
  -e CORS_ALLOWED_ORIGIN=https://your-frontend.com \
  -v $(pwd)/data:/data \
  jobscraper
```

### Railway

1. Connect your GitHub repository
2. Set environment variables (`SYNC_PASSWORD`, `CORS_ALLOWED_ORIGIN`)
3. Deploy

---

## 📂 Project Structure

```
jobscraper/
├── main.go                  # Entry point, Gin router setup
├── common/
│   └── payload.go          # JobPayload, JobMeta types
├── db/
│   └── sqlite.go          # SQLite operations
├── internal/
│   ├── handler/
│   │   ├── jobs.go        # GET /getallJobsFromSQL, /companies, /locations
│   │   └── sync.go        # POST /sync, GET /syncall
│   └── scraper/
│       ├── scraper.go     # Pool runner with concurrency control
│       └── adapters.go    # Platform-specific scraper adapters
├── scrapers/
│   ├── ashby/
│   │   ├── fetch/fetch.go # Ashby API client
│   │   ├── normalize/normalize.go
│   │   └── ...
│   ├── lever/
│   ├── amazon/
│   └── atlassian/
└── target/
    └── target.go          # Company configuration management
```

---

## 👨‍💻 Tech Stack

| Tech | Use Case |
|------|----------|
| **Go 1.25** | Core backend |
| **Gin** | HTTP framework |
| **SQLite** | Persistent storage |
| **go-sqlite3** | SQLite driver |
| **godotenv** | Environment variables |

---

## 🔧 Development

```bash
# Install dependencies
go mod download

# Run locally
go run main.go

# Run with docker
docker build -f Dockerfile -t jobscraper .
docker run -p 8080:8080 jobscraper
```

---

## 📝 Notes

- Ashby scraper uses a semaphore limit of 4 concurrent requests
- Global scraper limit is 15 concurrent requests
- Jobs are deduplicated by `job_id` (content hash)
- Inactive jobs are marked as `removed` instead of deleted