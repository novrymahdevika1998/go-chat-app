package main

import (
	"chat-app/routes"
	"database/sql"
	"fmt"
	_ "github.com/gin-contrib/cors"
	_ "github.com/glebarez/go-sqlite"
	"os"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Channel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	ID        int    `json:"id"`
	ChannelID int    `json:"channel_id"`
	UserID    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	Text      string `json:"text"`
}

func main() {
	wd, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Working directory:", wd)

	db, err := sql.Open("sqlite", wd+"/database.db")

	defer func(db *sql.DB) {
		err := db.Close()

		if err != nil {
			fmt.Println(err)
		}
	}(db)

	r := routes.SetupRouter(db)

	// Run the server
	err = r.Run(":8080")

	if err != nil {
		fmt.Println(err)
	}
}
