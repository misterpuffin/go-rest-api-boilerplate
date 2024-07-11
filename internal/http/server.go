package http

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"github.com/misterpuffin/go-rest-api-boilerplate/internal/config"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/db/sqlc"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/http/controllers"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/patterns"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/users"
)

func RunServer(config config.Config) {

	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", config.Postgres.Username, config.Postgres.Password, config.Postgres.Host, config.Postgres.Port, config.Postgres.DBName)
	fmt.Printf("Connecting to %s", connectionString)

	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	q := db.New(conn)

	r := gin.Default()

	r.Use(ErrorHandler)

	// Initialize services
	userService := users.NewService(config, q)
	patternService := patterns.NewService(config, q)

	// Initialize controllers
	authController := controllers.NewAuthController(userService)
	patternController := controllers.NewPatternController(patternService)

	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	pattern := r.Group("/pattern")
	pattern.Use(AuthHandler(config))
	{
		pattern.POST("/", patternController.Post)
		pattern.GET("/", patternController.Get)
	}

	r.Run()
}
