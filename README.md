# TrackIT - AI-Powered Task Management System

## 📋 Overview

TrackIT is a cutting-edge task management solution infused with AI capabilities, designed to optimize workflows through smart organization and seamless collaboration. Leveraging advanced technology, it enables both individuals and teams to manage tasks efficiently with AI-driven insights, real-time updates, and intelligent task breakdowns.

## ✨ Key Features

### Core Functionalities
- 🔐 Secure JWT-based authentication
- ✅ Effortless task creation and tracking
- 🤖 AI-powered task breakdown and recommendations
- 👥 Task assignment and collaboration tools

## 🛠️ Tech Stack

### Backend
- **Language:** Go (Golang)
- **Framework:** Gin
- **Database:** MongoDB
- **Authentication:** JWT
- **AI Integration:** OpenRouter Deepseek-r1

### Frontend
- **Framework:** Next.js 14
- **Language:** TypeScript
- **Styling:** Tailwind CSS
- **Component Library:** Shadcn UI

### DevOps
- **Containerization:** Docker
- **Orchestration:** Kubernetes

## 📦 Installation Guide

### Prerequisites
Ensure you have the following installed before proceeding:
- Go 1.21+
- Node.js 18+
- MongoDB (or Atlas)
- Docker
- OpenRouter API Key

## 🚀 Local Development Setup

### Clone the Repository

```bash
git clone https://github.com/SoNiC-HeRE/TrackIt.git
cd TrackIt
```

---

## 🔧 Backend Setup

1. Navigate to the backend directory:

```bash
cd backend-trackit
```

2. Create a new environment configuration file:

```bash
touch .env
```

3. Populate `.env` with necessary details:
   - **MongoDB Connection String**
   - **JWT Secret Key**
   - **OpenRouter API Key**

4. Install required Go dependencies:

```bash
go mod tidy
```

5. Start MongoDB (or use MongoDB Atlas):

```bash
docker run -d --name mongodb -p 27017:27017 mongo
```

6. Run the backend server:

```bash
go run main.go
```

---

## 🌐 Frontend Setup

1. Navigate to the frontend directory:

```bash
cd frontend-trackit
```

2. Install project dependencies:

```bash
npm install
```

3. Start the development server:

```bash
npm run dev
```

---

## 🐳 Docker Deployment

### Running Services with Docker Compose

1. Build and start all containers:

```bash
docker-compose up --build
```

2. Stop and remove all services:

```bash
docker-compose down
```

---

## 🤝 Contribution Guidelines

We welcome contributions! Follow these steps to contribute:

1. Fork the repository.
2. Create a new feature branch:
   ```bash
   git checkout -b feature/your-feature
   ```
3. Commit your changes:
   ```bash
   git commit -m 'Added new feature'
   ```
4. Push the branch to your fork:
   ```bash
   git push origin feature/your-feature
   ```
5. Submit a Pull Request for review.

---

Thank you for supporting and contributing to TrackIT! 🚀