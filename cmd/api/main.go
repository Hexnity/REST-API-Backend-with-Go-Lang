package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"hexnity/internal/delivery/http"
	"hexnity/internal/repository/postgres"
	"hexnity/internal/usecase"
)

func main() {
	// 1. Database Connection
	// Replace with your actual database credentials
	connStr := "postgresql://postgres:password@localhost:5432/hexnity_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close() // Best practice: close DB when main exits

	// 2. Wiring the Layers (Dependency Injection)
	userRepo := postgres.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := &http.UserHandler{UseCase: userUseCase}

	// 3. Setup Gin
	r := gin.Default()

	// 4. Routes
	r.POST("/register", userHandler.Register)

	// 5. Start Server
	log.Println("Hexnity Platform API running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
