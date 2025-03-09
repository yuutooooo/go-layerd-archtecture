package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"layerd-archtecture/infrastructure/db"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	println("DB接続を試みています")
	_, err := db.InitDB()
	if err != nil {
		println(err.Error())
		println("DB接続に失敗しました")
		return
	}
	println("DB接続に成功しました")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("こんにちは、Goの世界へようこそ!"))
	})

	log.Println("サーバーを起動します。http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
