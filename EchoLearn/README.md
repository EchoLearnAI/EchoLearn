# EchoLearn API

A Go backend service for an English learning app that uses multiple-choice questions for grammar practice.

## Features

- User Management (simple version, no auth for now)
- Question Management with grammar rules and categories
- Game Sessions with three modes:
  - Survival Mode: Stop after 3 mistakes
  - Topic Mode: 50 questions, 10 per topic
  - Infinite Mode: Stops when user taps "Finish"

## Prerequisites

- Go 1.21 or higher
- SQLite (default) or PostgreSQL

## Project Structure

```
EchoLearn/
├── cmd/
│   └── server/           # Main application entry point
├── internal/
│   ├── models/           # Data models
│   ├── controllers/      # Request handlers
│   ├── routes/           # API routes
│   └── db/               # Database connection
├── scripts/              # Utility scripts
├── data/                 # SQLite database storage
├── config/               # Configuration files
└── docs/                 # Documentation
```

## Getting Started

1. Clone the repository
2. Copy `.env.example` to `.env` and adjust the settings if needed
3. Run the server:

```bash
go run cmd/server/main.go
```

4. To seed the database with sample data:

```bash
go run scripts/seed.go
```

## API Endpoints

### User Management

- `POST /users` - Create a new user
- `GET /users` - Get all users
- `GET /users/:id` - Get user by ID
- `PUT /users/:id` - Update user
- `DELETE /users/:id` - Delete user

### Question Management

- `GET /questions/random` - Get a random question
- `GET /questions/category/:name` - Get 10 questions by category
- `GET /questions/:id` - Get question by ID
- `POST /questions` - Create a new question
- `PUT /questions/:id` - Update question
- `DELETE /questions/:id` - Delete question

### Game Sessions

- `POST /sessions/start` - Start a new session
- `POST /sessions/submit` - Submit an answer
- `POST /sessions/finish` - Finish a session
- `GET /sessions/:id/summary` - Get session summary

## Database Configuration

The application supports both SQLite and PostgreSQL databases. By default, it uses SQLite.

To switch to PostgreSQL, update the `.env` file:

```
DB_TYPE=postgres
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=echolearn
DB_SSLMODE=disable
```

## License

MIT 