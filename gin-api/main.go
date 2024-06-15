package main

import (
	// "net/http"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
func main() {
	r := gin.Default()

	err := godotenv.Load(fmt.Sprintf("./env/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r.GET("/",func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("my-health", func(c *gin.Context) {
		// FitbitAPIへのリクエストをする
		c.Request.URL.Path = "https://api.fitbit.com/1/user/-/profile.json"
		c.Request.Header.Add("Authorization", "Bearer " + os.Getenv("FITBIT_AUTH"))
		// FitbitAPIからのレスポンスを受け取る
		res, err := http.DefaultClient.Do(c.Request)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(200,gin.H {
			"result": string(body),
		})
	})

	r.Run("localhost:3300") 
}