# TrackIT - AI-Driven Task Management System

## 📋 Overview

TrackIT is an innovative, AI-enhanced task management platform designed to streamline workflows with intelligent organization and real-time collaboration. Utilizing modern technology, it empowers teams and individuals to efficiently manage tasks with AI-based suggestions, real-time synchronization, and smart task structuring.

## ✨ Key Features

### Core Functionalities
- 🔐 Secure authentication using JWT
- ✅ Seamless task creation and management
- 🤖 AI-assisted task recommendations and breakdowns
- 👥 Task delegation and collaboration

## 🛠️ Technology Stack

### Backend
- **Programming Language:** Go (Golang)
- **Framework:** Gin
- **Database:** MongoDB
- **Authentication:** JWT
- **AI Services:** OpenRouter Deepseek-r1

### Frontend
- **Framework:** Next.js 14
- **Language:** TypeScript
- **Styling:** Tailwind CSS
- **UI Library:** Shadcn UI

### DevOps
- **Containerization:** Docker
- **Orchestration:** Kubernetes

## 📦 Installation Guide

### Prerequisites
Before proceeding, ensure you have the following installed:
- Go 1.21+
- Node.js 18+
- MongoDB
- Docker
- OpenRouter API Key

## Setting Up for Local Development

Follow the instructions below to configure and run TrackIt on your local machine.

### Cloning the Repository

```bash
git clone 
cd task-management-system
```

---

## Backend Configuration

1. Navigate to the backend directory:

```bash
cd backend
```

2. Duplicate the sample environment file:

```bash
create .env
```

3. Update `.env` with your configuration details:
   - **MongoDB Connection URI**
   - **JWT Secret Key**
   - **OpenRouter API Key**

4. Install Go dependencies:

```bash
go mod tidy
```

5. Start MongoDB using Docker:

```bash
docker run -d --name mongodb -p 27017:27017 mongo or can skip if using atlas
```

6. Launch the backend server:

```bash
go run main.go
```

---

## Frontend Configuration

1. Move to the frontend directory:

```bash
cd frontend
```

2. Install necessary dependencies:

```bash
npm i
```

3. Start the frontend development server:

```bash
npm run dev
```

---

## Docker Deployment

### Running with Docker Compose

1. Build and launch all services:

```bash
docker-compose up --build
```

2. Shut down all services:

```bash
docker-compose down
```

---

## 🤝 Contributing

We welcome contributions! Follow these steps to get started:

1. Fork the repository.
2. Create a new feature branch:  
   ```bash
   git checkout -b feature/your-feature
   ```
3. Commit your changes:  
   ```bash
   git commit -m 'Implemented new feature'
   ```
4. Push your changes:  
   ```bash
   git push origin feature/your-feature
   ```
5. Open a Pull Request.

---

Thank you for contributing to TrackIt! 🚀