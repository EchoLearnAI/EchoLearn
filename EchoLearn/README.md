# EchoLearn Backend

This is the backend service for the EchoLearn English learning mobile app.

## Prerequisites

- Go (version 1.21 or higher recommended)

## Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/echolearn.git
   cd echolearn
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Create .env file:**
   Copy `.env.example` to `.env` and customize if needed (e.g., `SERVER_PORT`).
   ```bash
   cp .env.example .env
   ```

4. **Run the server:**
   ```bash
   go run main.go
   ```
   The server will start on `http://localhost:8080` by default.

## Project Structure

```
EchoLearn/
├── api/
│  ├── handlers/
│  │  ├── question_handler.go
│  │  ├── session_handler.go
│  │  └── user_handler.go
│  └── routes.go
├── config/
│  └── config.go
├── data/
│  └── seed.go
│  └── questions.json // Example seed data
├── db/
│  └── database.go
├── models/
│  ├── question.go
│  ├── session.go
│  └── user.go
├── utils/
│  └── errors.go 
├── main.go
├── go.mod
├── go.sum
└── README.md
```

## API Endpoints

### User Management

- `POST /users` - Create a new user
  - Body: `{"name": "John Doe", "email": "john.doe@example.com"}`
- `GET /users/:id` - Get user by ID

### Question Management

- `GET /questions/random` - Get a random question
- `GET /questions/category/:name` - Get 10 questions by category
- `POST /questions` - Add a new question (admin use)
- `GET /questions/:id` - Get question with detailed feedback

### Game Modes

- `POST /session/start` - Initialize a session
  - Body: `{"user_id": "uuid", "mode": "survival"}`
- `POST /session/submit` - Submit an answer
  - Body: `{"session_id": "uuid", "question_id": "uuid", "selected_option_id": "uuid"}`
- `GET /session/:id/summary` - Get session results

## Database

- Uses SQLite for local development (`echolearn.db`).
- The database is automatically created and seeded with sample questions if it doesn't exist.

## Future Enhancements

- Authentication (JWT)
- More complex user progress tracking
- Leaderboards
- More game modes
- Swagger documentation
- Unit and integration tests 