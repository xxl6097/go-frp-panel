package main

import (
	"github.com/xxl6097/go-frp-panel/internal/comm/upload"
	"log"
	"net/http"
)

func serveIndexPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
func main() {
	http.HandleFunc("/upload", upload.NewPieces("./uploads").Upload)
	http.HandleFunc("/", serveIndexPage)
	log.Println("Server started on http://127.0.0.1:8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
