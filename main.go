package main

import (
	"fmt"
	"finalproject/database"
	"finalproject/router"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	database.StartDB()
	r := router.StartApp()
	r.Run(fmt.Sprintf(":%s", port))
}
