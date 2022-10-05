package main

import (
	"awesomeProject2/internal/db"
	"awesomeProject2/routes"
	_ "database/sql"
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
)

func main() {
	// app logging to file
	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    2, // megabytes
		MaxBackups: 30,
		MaxAge:     40,   // days
		Compress:   true, // disabled by default
	})
	log.Println("---------------Start logging---------------")

	db, err := db.InitDB()
	if err != nil {
		log.Println("Connection err", err)
		return
	}
	defer db.Close()

	fmt.Println("Server listening")
	http.ListenAndServe(":8080", routes.Routers())

}
