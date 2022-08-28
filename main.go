package main

import (
	"bytes"
	"fmt"
	initialenviroment "go-gin-web/initialEnviroment"
	"go-gin-web/models"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func init() {
	initialenviroment.EnvInit()
	initialenviroment.DbInit()
}

type Diary struct {
	Date    string `json:"date"`
	Content string `json:"content"`
}

func main() {
	fmt.Println("hello")

	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Title": "My first day ",
		})
	})

	r.POST("/post", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Result": "Got it",
		})
	})

	r.POST("/create", func(c *gin.Context) {
		var body struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		}

		if err := c.Bind(&body); err != nil {
			log.Fatal(err)
		}

		add := models.Blog{Title: body.Title, Content: body.Content}

		log.Printf("%#v", add)
		result := initialenviroment.DB.Create(&add)
		if result.Error != nil {
			log.Print(result.Error.Error())
			c.JSON(http.StatusInternalServerError, "failed")
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "create ok"})
	})

	r.GET("/list", func(c *gin.Context) {
		blog := []models.Blog{}
		initialenviroment.DB.Find(&blog)
		c.JSON(http.StatusOK, blog)
	})

	r.POST("/exec", func(c *gin.Context) {
		command := c.Query("command")
		log.Println(command)
		cmd := exec.Command(command)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, out.String())

	})

	r.Run()
}
