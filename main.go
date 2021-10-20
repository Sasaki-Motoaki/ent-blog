package main

import (
	"context"
	"example/ent-blog/ent"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=db user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	usersGroup := engine.Group("/users")
	{
		usersGroup.GET("", func(c *gin.Context) {
			if users, err := client.User.Query().All(context.Background()); err != nil {
				c.Status(http.StatusInternalServerError)
			} else {
				c.IndentedJSON(http.StatusOK, gin.H{
					"users": users,
				})
			}
		})
		usersGroup.GET("/:id", func(c *gin.Context) {
			strID, ok := c.Params.Get("id")
			if !ok {
				c.Status(http.StatusBadRequest)
			}
			id, err := strconv.Atoi(strID)
			if err != nil {
				c.Status(http.StatusInternalServerError)
			}
			if user, err := client.User.Get(context.Background(), id); err != nil {
				c.Status(http.StatusNotFound)
			} else {
				c.IndentedJSON(http.StatusOK, gin.H{
					"user": user,
				})
			}
		})
		usersGroup.POST("", func(c *gin.Context) {
			if user, err := CreateUser(context.Background(), client); err != nil {
				c.Status(http.StatusInternalServerError)
			} else {
				c.IndentedJSON(http.StatusCreated, gin.H{
					"user": user,
				})
			}
		})
		usersGroup.DELETE("/:id", func(c *gin.Context) {
			strID, ok := c.Params.Get("id")
			if !ok {
				c.Status(http.StatusBadRequest)
			}
			id, err := strconv.Atoi(strID)
			if err != nil {
				c.Status(http.StatusInternalServerError)
			}
			if err := client.User.DeleteOneID(id).Exec(context.Background()); err != nil {
				c.Status(http.StatusInternalServerError)
			} else {
				c.IndentedJSON(http.StatusOK, "deleted")
			}
		})
	}
	engine.Run(":8080")
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetEmail("example@example.com").
		SetPassword("password").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}
