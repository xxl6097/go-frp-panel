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

// 存储文件的目录
const uploadDir = "./uploads"

func main() {
	// 创建上传目录
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.MkdirAll(uploadDir, 0755)
		if err != nil {
			log.Fatalf("Failed to create upload directory: %v", err)
		}
	}

	// 处理分片上传
	http.HandleFunc("/upload", handleChunkUpload)
	// 处理合并分片
	http.HandleFunc("/merge", handleMergeChunks)
	// 提供前端页面
	http.HandleFunc("/", serveIndexPage)

	log.Println("Server started on http://127.0.0.1:8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

// 处理分片上传
func handleChunkUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// 获取文件信息
	file, handler, err := r.FormFile("chunk")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting file: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded file: %+v\n", handler.Filename)

	// 获取文件唯一标识和分片序号
	fileId := r.FormValue("fileId")
	chunkIndexStr := r.FormValue("chunkIndex")
	chunkIndex, err := strconv.Atoi(chunkIndexStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid chunk index: %v", err), http.StatusBadRequest)
		return
	}

	// 创建分片文件
	chunkFileName := fmt.Sprintf("%s_%d", fileId, chunkIndex)
	chunkFilePath := filepath.Join(uploadDir, chunkFileName)
	chunkFile, err := os.Create(chunkFilePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating chunk file: %v", err), http.StatusInternalServerError)
		return
	}
	defer chunkFile.Close()

	// 将分片内容写入文件
	_, err = io.Copy(chunkFile, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error writing chunk file: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Chunk uploaded successfully"})
}

// 处理合并分片
func handleMergeChunks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// 获取文件唯一标识和分片总数
	fileId := r.FormValue("fileId")
	totalChunksStr := r.FormValue("totalChunks")
	totalChunks, err := strconv.Atoi(totalChunksStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid total chunks: %v", err), http.StatusBadRequest)
		return
	}

	// 创建合并后的文件
	originalFileName := r.FormValue("originalFileName")
	mergedFilePath := filepath.Join(uploadDir, originalFileName)
	mergedFile, err := os.Create(mergedFilePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating merged file: %v", err), http.StatusInternalServerError)
		return
	}
	defer mergedFile.Close()

	// 按顺序合并分片
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
		// 删除已合并的分片
		os.Remove(chunkFilePath)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "File merged successfully"})
}

// 提供前端页面
func serveIndexPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
