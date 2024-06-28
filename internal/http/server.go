package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"github.com/misterpuffin/go-rest-api-boilerplate/internal/db/sqlc"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/users"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/util"
)

func RunServer(config util.Config) {

	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", config.Postgres.Username, config.Postgres.Password, config.Postgres.Host, config.Postgres.Port, config.Postgres.DBName)
	fmt.Printf("Connecting to %s", connectionString)

	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	q := db.New(conn)

	userService := users.NewService(config, q)
	fmt.Println(userService)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
