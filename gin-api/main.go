package main

import (
	// "net/http"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"encoding/json"

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
		err := godotenv.Load(fmt.Sprintf("./env/%s.env", os.Getenv("GO_ENV")))
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		fitbit_auth := os.Getenv("FITBIT_AUTH")
		fitbit_api_url_core := "https://api.fitbit.com/1/"

		client := &http.Client {}

		req, err := http.NewRequest("GET", fitbit_api_url_core + "user/-/profile.json", nil)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Authorization", "Bearer " + fitbit_auth)

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(body))
		
		// 次の課題
		// https://maku77.github.io/p/dsbs9p5/#google_vignette
		
		health := []byte{}
		
		health,err = json.MarshalIndent(string(body), "", "    ")
		if err != nil {
			fmt.Println(err)
			return
		}

		c.JSON(200,gin.H {
			"result": health, // string(body),
		})
	})

	r.Run("localhost:3300") 
}