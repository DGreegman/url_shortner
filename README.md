# URL Shortener API (Go + Fiber + PostgreSQL)

A lightweight URL shortener backend built with Go, Fiber, and PostgreSQL.

This project provides:
- Short URL creation endpoint
- Redirect endpoint for short codes
- Optional expiration support per URL
- Basic click counting on redirects
- Auto-migration for the `urls` table on startup

## Tech Stack

- Go (module-based project)
- Fiber v2 (web framework)
- PostgreSQL
- pgx v5 (PostgreSQL driver)
- godotenv (environment variable loading)

## Project Structure

```text
cmd/
  main.go                  # App bootstrap and server startup
internal/
  database/
    database.go            # DB connection setup
    migrate.go             # Schema migration
    query.go               # DB query helpers
  handlers/
    shorten.go             # POST /api/v1/shorten
    redirect.go            # GET /:code
  models/
    requests.go            # Request DTOs
    url.go                 # URL model
  routes/
    routes.go              # Route registration
  utils/
    code.go                # Short code generator
```

## Features

- `POST /api/v1/shorten`
  - Accepts a target URL and optional expiration time in seconds
  - Generates a random 6-character short code
  - Stores the record in PostgreSQL
  - Returns a full short URL (currently `http://localhost:3000/{code}`)

- `GET /:code`
  - Looks up short code
  - Rejects missing/expired links
  - Increments click count
  - Redirects to original URL with HTTP `301 Moved Permanently`

## Prerequisites

- Go (latest stable recommended)
- PostgreSQL (local or remote)
- Git (for version control and GitHub push)

## Environment Variables

The app loads environment variables from `.env` in the project root.

Create a `.env` file with:

```env
DB_USER=postgres
DB_PASSWORD=your_password
DB_HOST=localhost
DB_PORT=5432
DB_NAME=url_shortner
DB_SSLMODE=disable
```

Notes:
- `DB_SSLMODE=disable` is typical for local development.
- For managed databases, use the provider-recommended SSL mode.

## Database Setup

1. Create the database:

```sql
CREATE DATABASE url_shortner;
```

2. Start the app. On startup, migration runs automatically and creates table `urls` if it does not exist.

Current schema created by migration:

```sql
CREATE TABLE IF NOT EXISTS urls(
  id SERIAL PRIMARY KEY,
  code TEXT NOT NULL UNIQUE,
  target_url TEXT NOT NULL,
  clicks INT DEFAULT 0,
  created_at TIMESTAMP DEFAULT NOW(),
  expire_at TIMESTAMP NOT NULL
);
```

## Installation & Run

1. Clone repository and enter it:

```bash
git clone <your-repo-url>
cd url_shortner
```

2. Install dependencies:

```bash
go mod tidy
```

3. Configure environment variables in `.env`.

4. Run the API:

```bash
go run cmd/main.go
```

Server starts on:
- `http://localhost:3000`

## API Reference

### Health Check

- **GET** `/health`

Response:

```json
{
  "status": "ok",
  "message": "API is healthy"
}
```

---

### Create Short URL

- **POST** `/api/v1/shorten`
- **Content-Type:** `application/json`

Request body:

```json
{
  "url": "https://example.com/some/long/path",
  "expire_in": 3600,
  "redirect_type" : "301"
}
```

Fields:
- `url` (string, required): original URL
- `expire_in` (int, optional): expiration in seconds from now

Success response example:

```json
{
  "short_url": "http://localhost:3000/Ab3XyZ"
}
```

Possible errors:
- `400 Bad Request` for invalid JSON body
- `400 Bad Request` when `url` is empty
- `500 Internal Server Error` for DB insert failure

---

### Redirect

- **GET** `/:code`

Behavior:
- Finds URL by code
- Rejects if not found (`404`)
- Rejects if expired (`410`)
- Increments `clicks`
- Redirects to target URL (`301`)

Possible errors:
- `400 Bad Request` for missing code
- `404 Not Found` if code does not exist
- `410 Gone` if URL is expired

## Quick cURL Examples

Create short URL:

```bash
curl -X POST http://localhost:3000/api/v1/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://golang.org","expire_in":3600}'
```

Open short URL in browser:

```bash
http://localhost:3000/<code>
```

Check health:

```bash
curl http://localhost:3000/health
```

## Notes and Limitations

- Generated short code length is fixed to 6 characters.
- Current code generation uses pseudo-random `math/rand`.
- No URL ownership/authentication yet.
- No custom alias support yet.
- Base URL is hardcoded (`http://localhost:3000`) in response.

## Troubleshooting

- `Error loading .env file`
  - Ensure `.env` exists in project root.

- `Unable to connect to Database`
  - Verify DB credentials/host/port/db name in `.env`.
  - Confirm PostgreSQL is running.

- `Migration Failed`
  - Confirm DB user has permission to create tables.

- API returns `410 URL Expired` unexpectedly
  - Check the `expire_in` value used when creating the link.

## Roadmap Ideas

- Add request validation middleware (URL format, expiration bounds)
- Add collision handling/retry for short code uniqueness
- Make base URL configurable via environment variable
- Add unit/integration tests
- Add Docker and docker-compose setup
- Add rate limiting and analytics endpoints

## License

Choose a license before public release (for example: MIT).
