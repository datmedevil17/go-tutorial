# Simple JS Server

A basic Express.js server with CORS support and simple API endpoints.

## Getting Started

### Prerequisites
- Node.js (v14 or higher)
- npm

### Installation

1. Navigate to the project directory:
```bash
cd js-server
```

2. Install dependencies:
```bash
npm install
```

### Running the Server

#### Development mode (with auto-restart):
```bash
npm run dev
```

#### Production mode:
```bash
npm start
```

The server will start on `http://localhost:3000`

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Welcome message with timestamp |
| GET | `/api/posts` | Get all posts |
| GET | `/api/posts/:id` | Get a specific post by ID |
| POST | `/api/posts` | Create a new post |

### Example Requests

#### Get all posts:
```bash
curl http://localhost:3000/api/posts
```

#### Get specific post:
```bash
curl http://localhost:3000/api/posts/1
```

#### Create new post:
```bash
curl -X POST http://localhost:3000/api/posts \
  -H "Content-Type: application/json" \
  -d '{"title": "My Post", "body": "Post content", "author": "Me"}'
```

## Project Structure

```
js-server/
├── package.json    # Dependencies and scripts
├── server.js       # Main server file
└── README.md       # This file
```
