# Task Manager

A full-stack task management application with Go backend and React frontend.

## Architecture

### System Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                        Client Layer                          │
│  ┌──────────────────────────────────────────────────────┐  │
│  │         React Frontend (Vite + React 18)              │  │
│  │  - Task List UI                                        │  │
│  │  - Task Creation/Editing                              │  │
│  │  - Real-time Updates                                   │  │
│  └──────────────────────────────────────────────────────┘  │
└───────────────────────┬─────────────────────────────────────┘
                        │ HTTP/REST API
                        │
┌───────────────────────▼─────────────────────────────────────┐
│                     Application Layer                        │
│  ┌──────────────────────────────────────────────────────┐  │
│  │         Go Backend (HTTP Server)                       │  │
│  │  - RESTful API Endpoints                              │  │
│  │  - Business Logic                                     │  │
│  │  - Request Validation                                 │  │
│  └──────────────────────────────────────────────────────┘  │
└───────────────────────┬─────────────────────────────────────┘
                        │
                        │ (Future: Database Layer)
                        │
┌───────────────────────▼─────────────────────────────────────┐
│                    Data Layer (Future)                         │
│  ┌──────────────────────────────────────────────────────┐  │
│  │         Database (PostgreSQL/MongoDB)                  │  │
│  │  - Task Storage                                       │  │
│  │  - User Management                                    │  │
│  └──────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
```

### Component Architecture

**Frontend Components:**
- `App.jsx` - Main application component
- `TaskList` - Displays list of tasks
- `TaskForm` - Create/edit tasks
- `TaskItem` - Individual task component

**Backend Components:**
- `main.go` - HTTP server and routing
- `handlers/` - Request handlers
- `models/` - Data models
- `middleware/` - Authentication, logging, CORS

## Design Decisions

### Frontend Design
- **Framework**: React 18 with Vite for fast development and optimized builds
- **State Management**: React hooks (useState, useContext) for local state
- **Styling**: CSS Modules or Tailwind CSS (to be added)
- **HTTP Client**: Fetch API or Axios for API communication

### Backend Design
- **Language**: Go for high performance and concurrency
- **HTTP Framework**: Standard `net/http` package (lightweight)
- **API Style**: RESTful API with JSON responses
- **Error Handling**: Structured error responses with HTTP status codes

### Architecture Patterns
- **Layered Architecture**: Separation of concerns (UI, Business Logic, Data)
- **RESTful Design**: Stateless API with standard HTTP methods
- **Microservices Ready**: Backend can be containerized and scaled independently

## End-to-End Flow

### Flow 1: Create a New Task

```
1. User Interaction
   └─> User fills task form in React frontend
       └─> Form validation (client-side)
           └─> Submit button clicked

2. Frontend Request
   └─> React component calls API service
       └─> HTTP POST request to /api/tasks
           └─> Request body: { title, description, priority, dueDate }

3. Network Layer
   └─> Request travels over HTTP/HTTPS
       └─> Reaches Go backend server (port 8080)

4. Backend Processing
   └─> Go HTTP server receives request
       └─> Router matches POST /api/tasks endpoint
           └─> Middleware chain:
               ├─> CORS middleware (adds headers)
               ├─> Logging middleware (logs request)
               └─> Authentication middleware (validates token)
           └─> Handler function:
               ├─> Parse JSON request body
               ├─> Validate task data
               ├─> Generate unique task ID
               ├─> Create task object
               └─> Store in database (future) or memory

5. Response Generation
   └─> Handler creates response:
       ├─> HTTP Status: 201 Created
       ├─> Response body: { id, title, description, status: "created" }
       └─> JSON encoding

6. Network Response
   └─> Response sent back to frontend
       └─> HTTP response received by React app

7. Frontend Update
   └─> React component receives response
       └─> Update local state with new task
           └─> Re-render UI to show new task in list
               └─> Show success notification to user
```

### Flow 2: View All Tasks

```
1. Component Mount
   └─> React TaskList component mounts
       └─> useEffect hook triggers

2. API Request
   └─> HTTP GET request to /api/tasks
       └─> No request body needed

3. Backend Processing
   └─> Go server receives GET request
       └─> Handler function:
           ├─> Query database for all tasks
           ├─> Filter by user (future: authentication)
           ├─> Sort by creation date
           └─> Format as JSON array

4. Response
   └─> HTTP 200 OK with task array
       └─> [{ id, title, status, ... }, ...]

5. Frontend Update
   └─> React receives task array
       └─> Map tasks to TaskItem components
           └─> Render list in UI
```

### Flow 3: Update Task Status

```
1. User Action
   └─> User clicks "Complete" button on a task

2. Frontend Request
   └─> HTTP PATCH /api/tasks/{id}
       └─> Request body: { status: "completed" }

3. Backend Processing
   └─> Handler:
       ├─> Validate task ID exists
       ├─> Update task status in database
       ├─> Update timestamp
       └─> Return updated task

4. Response & UI Update
   └─> Frontend receives updated task
       └─> Update task in local state
           └─> Re-render with new status (visual indicator)
```

## Data Flow

```
┌──────────┐     JSON Request      ┌──────────┐
│  React   │ ───────────────────> │   Go     │
│ Frontend │                       │ Backend  │
└──────────┘                       └──────────┘
      │                                   │
      │     JSON Response                │
      │ <─────────────────────────────────│
      │                                   │
      │                                   ▼
      │                            ┌──────────┐
      │                            │ Database │
      │                            │ (Future) │
      └──────────────────────────>└──────────┘
```

## API Endpoints

### Task Management
- `GET /api/tasks` - List all tasks
- `GET /api/tasks/:id` - Get task by ID
- `POST /api/tasks` - Create new task
- `PATCH /api/tasks/:id` - Update task
- `DELETE /api/tasks/:id` - Delete task

### Health & Status
- `GET /health` - Health check endpoint

## Deployment Architecture

### Development
```
Frontend (Vite Dev Server) :3000
    │
    └─> Backend (Go Server) :8080
```

### Production (Docker)
```
┌─────────────────────────────────────┐
│     Docker Compose Network          │
│                                     │
│  ┌──────────────┐  ┌─────────────┐ │
│  │   Frontend   │  │   Backend    │ │
│  │   (Nginx)    │  │   (Go App)   │ │
│  │   Port 3000  │  │   Port 8080 │ │
│  └──────────────┘  └─────────────┘ │
│                                     │
└─────────────────────────────────────┘
```

## Build & Run

### Prerequisites
- Go 1.21+
- Node.js 18+
- Docker (optional)

### Backend
```bash
cd backend
go mod download
go run main.go
# Server runs on :8080
```

### Frontend
```bash
cd frontend
npm install
npm run dev
# App runs on :3000
```

### Docker
```bash
docker-compose up --build
```

## Future Enhancements

- [ ] Database integration (PostgreSQL)
- [ ] User authentication (JWT)
- [ ] Real-time updates (WebSockets)
- [ ] Task categories and tags
- [ ] File attachments
- [ ] Search and filtering
- [ ] Task reminders and notifications
