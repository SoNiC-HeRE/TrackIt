package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "task-management/internal/handlers"
    "task-management/internal/middleware"
    "task-management/internal/database"
)
func main() {
    fmt.Println("Hello, Go!")
}
