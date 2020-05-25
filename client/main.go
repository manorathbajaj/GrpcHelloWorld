package main

import (
	"fmt"
	"helloWorld/proto"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())

	if err != nil {
		fmt.Printf("failed to connect to localhost")
		panic(err)
	}

	client := proto.NewCRUDSreviceClient(conn)

	g := gin.Default()
	// Bad way of doing it but demo meh.. Should be a post with info in body
	g.GET("/CREATE/:id/:name/:age/:address/:salary", func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter id"})
			return
		}

		age, err := strconv.ParseUint(ctx.Param("age"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter age"})
			return
		}

		salary, err := strconv.ParseUint(ctx.Param("salary"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter salary"})
			return
		}

		name := ctx.Param("name")
		address := ctx.Param("address")

		req := &proto.Create{Id: int64(id), Name: string(name), Age: int64(age), Salary: int64(salary), Address: string(address)}

		if response, err := client.CreateEmp(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Done),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	// Actual get method here
	g.GET("/RET/:id", func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter id"})
			return
		}
		req := &proto.Retrieve{Id: int64(id)}

		if response, err := client.RetrieveEmp(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"id":      fmt.Sprint(response.GetId()),
				"name":    fmt.Sprint(response.GetName()),
				"age":     fmt.Sprint(response.GetAge()),
				"address": fmt.Sprint(response.GetAddress()),
				"salary":  fmt.Sprint(response.GetSalary()),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	// Should be a put but ¯\_(ツ)_/¯ demo
	g.GET("/UPDATE/:id/:name/:age/:address/:salary", func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter id"})
			return
		}

		age, err := strconv.ParseUint(ctx.Param("age"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter age"})
			return
		}

		salary, err := strconv.ParseUint(ctx.Param("salary"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter salary"})
			return
		}

		name := ctx.Param("name")
		address := ctx.Param("address")

		req := &proto.Create{Id: int64(id), Name: string(name), Age: int64(age), Salary: int64(salary), Address: string(address)}

		if response, err := client.UpdateEmp(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Done),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	// ¯\_(ツ)_/¯ demo ¯\_(ツ)_/¯ DELETE

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
