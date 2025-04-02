package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const uploadDir = "./uploads"

func main() {
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.MkdirAll(uploadDir, 0755)
		if err != nil {
			log.Fatalf("Failed to create upload directory: %v", err)
		}
	}

	http.HandleFunc("/upload", handleChunkUpload)
	http.HandleFunc("/merge", handleMergeChunks)
	http.HandleFunc("/", serveIndexPage)

	log.Println("Server started on http://127.0.0.1:8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func handleChunkUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	file, handler, err := r.FormFile("chunk")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting file: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Printf("%v\n", handler.Filename)

	fileId := r.FormValue("fileId")
	chunkIndexStr := r.FormValue("chunkIndex")
	chunkIndex, err := strconv.Atoi(chunkIndexStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid chunk index: %v", err), http.StatusBadRequest)
		return
	}

	chunkFileName := fmt.Sprintf("%s_%d", fileId, chunkIndex)
	chunkFilePath := filepath.Join(uploadDir, chunkFileName)
	chunkFile, err := os.Create(chunkFilePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating chunk file: %v", err), http.StatusInternalServerError)
		return
	}
	defer chunkFile.Close()

	_, err = io.Copy(chunkFile, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error writing chunk file: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Chunk uploaded successfully"})
}

func handleMergeChunks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fileId := r.FormValue("fileId")
	totalChunksStr := r.FormValue("totalChunks")
	totalChunks, err := strconv.Atoi(totalChunksStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid total chunks: %v", err), http.StatusBadRequest)
		return
	}

	originalFileName := r.FormValue("originalFileName")
	mergedFilePath := filepath.Join(uploadDir, originalFileName)
	mergedFile, err := os.Create(mergedFilePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating merged file: %v", err), http.StatusInternalServerError)
		return
	}
	defer mergedFile.Close()

	for i := 0; i < totalChunks; i++ {
		chunkFileName := fmt.Sprintf("%s_%d", fileId, i)
		chunkFilePath := filepath.Join(uploadDir, chunkFileName)
		chunkFile, err := os.Open(chunkFilePath)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error opening chunk file: %v", err), http.StatusInternalServerError)
			return
		}
		_, err = io.Copy(mergedFile, chunkFile)
		chunkFile.Close()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error merging chunk file: %v", err), http.StatusInternalServerError)
			return
		}
		os.Remove(chunkFilePath)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "File merged successfully"})
}

func serveIndexPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
