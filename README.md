# StackOverflow Clone Backend (Golang + PostgreSQL + Gemini LLM)

This is a backend system for a StackOverflow-style Q\&A platform. It is built using **Golang**, **PostgreSQL** on aws ec2, and integrates **Gemini** for summarization.


---

## Repository Structure

```
.
├── cmd/api/             # Entry point for the server (main.go)
├── config/              # Config files for the project
├── controllers/         # Request handlers
├── database/            # DB connection setup
├── models/              # GORM models for all tables
├── routes/              # Router setup
├── utils/               # JWT, Gemini and other utility functions
├── Dockerfile
├── docker-compose.yml
├── .env
└── README.md
```

---
## Database Schema and Indexes
```
users - id, username, email, password, created_at, updated_at  
questions - id, user_id, title, body, created_at, updated_at  
answers - id, question_id, user_id, body, created_at, updated_at  
votes - id, user_id, item_id, item_type, vote_type, created_at, updated_at  

```

Indexes:
- questions(user_id) – fetch questions by user
- answers(question_id) – fetch answers for a question
- answers(user_id) – (optional) fetch answers by user
- votes(item_id, item_type) – get votes for a question/answer
- votes(user_id, item_id, item_type) – check if user already voted (enforces uniqueness, speeds up vote updates)
---

## Setup & Run

### Prerequisites

* Go >= 1.24
* Docker + Docker Compose

### Clone and Configure

```bash
git clone https://github.com/ashblend17/stackoverflow-sample.git
cd stackoverflow-sample

# Create .env file and set your variables
cp .env.example .env
```

### Run with Docker Compose

```bash
docker-compose up --build
```

### Server will be live at:

```
http://localhost:8080
```

---

## API Endpoints & Usage Examples

> All authenticated routes require an `Authorization: Bearer <token>` header. Only /register and /login are unprotected routes

### Register

```
POST /api/register
{
  "username": "sefghuio",
  "email": "you@example.com",
  "password": "verystrong"
}
```
```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"username":"sefghuio","email":"you@example.com","password":"verystrong"}'

```
Response:
```
{ "message": "User registered successfully" }
```


### Login

```
POST /api/login
{
  "email": "you@example.com",
  "password": "verystrong"
}
```
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"you@example.com","password":"verystrong"}'

```
Response:
```
{
  "token": "jwt-token-here"
}
```

### Auth test (with token)
```
GET /api/test
```
```bash
curl -X GET http://localhost:8080/api/test \
  -H "Authorization: Bearer <your-token>"
```
Response:
```
{
    "status": "Auth passed"
}
```

### Post Question

```
POST /api/createQuestion
{
  "title": "How is this even a question?",
  "body": "This is beyond me"
}
```
```bash
curl -X POST http://localhost:8080/api/createQuestion \
  -H "Authorization: Bearer <your-token>" \
  -H "Content-Type: application/json" \
  -d '{"title":"How is this even a question?","body":"This is beyond me..."}'

```
Response:
```
{
  "id": 1,
  "message": "Question created"
}
```

### Post Answer

```
POST /api/question/:id/createAnswer
{
  "body": "This cannot be an answer."
}
```
```bash
curl -X POST http://localhost:8080/api/question/1/createAnswer \
  -H "Authorization: Bearer <your-token>" \
  -H "Content-Type: application/json" \
  -d '{"body":"This cannot be an answer."}'

```


### Fetch question and answers
```bash
curl -X GET http://localhost:8080/api/getQnA/1 \
-H "Authorization: Bearer <token-here>";

```

### Vote on Question or Answer

```
POST /api/question/:id/vote
{
  "vote": "upvote" | "downvote" | "remove"
}

POST /api/answer/:id/vote
{
  "vote": "upvote" | "downvote" | "remove"
}
```
```bash
curl -X POST http://localhost:8080/api/question/1/vote \
  -H "Authorization: Bearer <your-token>" \
  -H "Content-Type: application/json" \
  -d '{"vote":"upvote"}' # or "downvote" or "remove"


curl -X POST http://localhost:8080/api/answer/4/vote \
  -H "Authorization: Bearer <your-token>" \
  -H "Content-Type: application/json" \
  -d '{"vote":"downvote"}'

```

Response:
```
{
    "message": "Vote updated"
}
```

### Get Summary

```
GET /api/question/:id/summary
```
```bash
curl -X GET http://localhost:8080/api/question/1/summary \
  -H "Authorization: Bearer <your-token>"

```

Response:
```
{
  "question_id": <id>, 
  "summary": "summary-here"
}
```

---

## Assumptions Made

* Only credentials type auth supported.
* Gemini is the default LLM for summarization.
* Each user can only vote once per question/answer.
* User passwords are stored as bcrypt hashes.
* Summaries are generated on-the-fly (no caching).
* Schema strictly matches the provided SQL structure.
