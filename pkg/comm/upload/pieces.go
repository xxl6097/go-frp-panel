package upload

import (
	"fmt"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-service/pkg/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type pieces struct {
	outDir string
}

func NewPieces(outDir string) *pieces {
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outDir, 0755); err != nil {
			errs := fmt.Sprintf("Failed to create upload directory: %v", err)
			fmt.Println("Failed to create upload directory", errs)
		}
	}
	return &pieces{outDir: outDir}
}
func (this *pieces) Upload(w http.ResponseWriter, r *http.Request) {
	this.UploadHandler(w, r)

}
func (this *pieces) UploadHandler(w http.ResponseWriter, r *http.Request) (error, string) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return fmt.Errorf("invalid request method: %v", http.StatusMethodNotAllowed), ""
	}

	fileName := r.Header.Get("X-File-Name")
	if fileName == "" {
		http.Error(w, "Missing file name", http.StatusBadRequest)
		return fmt.Errorf("missing file name %v", http.StatusBadRequest), ""
	}

	chunkIndexStr := r.Header.Get("X-Chunk-Index")
	if chunkIndexStr == "" {
		http.Error(w, "Missing chunk index", http.StatusBadRequest)
		return fmt.Errorf("missing chunk index %v", http.StatusBadRequest), ""
	}

	chunkIndex, err := strconv.Atoi(chunkIndexStr)
	if err != nil {
		http.Error(w, "Invalid chunk index", http.StatusBadRequest)
		return fmt.Errorf("invalid chunk index %v", http.StatusBadRequest), ""
	}

	totalChunksStr := r.Header.Get("X-Total-Chunks")
	if totalChunksStr == "" {
		http.Error(w, "Missing total chunks", http.StatusBadRequest)
		return fmt.Errorf("missing total chunks %v", http.StatusBadRequest), ""
	}

	totalChunks, err := strconv.Atoi(totalChunksStr)
	if err != nil {
		http.Error(w, "Invalid total chunks", http.StatusBadRequest)
		return fmt.Errorf("invalid total chunks %v", http.StatusBadRequest), ""
	}

	hash := r.Header.Get("X-File-Hash")
	if hash == "" {
		http.Error(w, "File hash missed", http.StatusBadRequest)
		return fmt.Errorf("missing file hash %v", http.StatusBadRequest), ""
	}

	appDir := glog.AppHome("chunk")
	chunkDir := filepath.Join(appDir, hash, "chunk")
	if _, err := os.Stat(chunkDir); os.IsNotExist(err) {
		if err := os.MkdirAll(chunkDir, 0755); err != nil {
			errs := fmt.Sprintf("Failed to create upload directory: %v", err)
			http.Error(w, errs, http.StatusBadRequest)
			return fmt.Errorf("failed to create upload directory: %v", errs), ""
		}
	}

	chunkPath := filepath.Join(chunkDir, fmt.Sprintf("chunk_%d", chunkIndex))
	file, err := os.OpenFile(chunkPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to open chunk file: %v", err), http.StatusInternalServerError)
		return fmt.Errorf("failed to open chunk file: %v", err), ""
	}
	defer file.Close()

	_, err = io.Copy(file, r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to write chunk: %v", err), http.StatusInternalServerError)
		defer utils.DeleteAllDirector(appDir)
		return fmt.Errorf("failed to write chunk: %v", err), ""
	}

	if chunkIndex == totalChunks-1 {
		//outDir := filepath.Join(chunkDir, "../")
		//if this.OutDir != "" {
		//	outDir = this.OutDir
		//}
		defer utils.DeleteAllDirector(appDir)
		// 合并所有块
		if err := this.mergeChunks(this.outDir, chunkDir, fileName, totalChunks); err != nil {
			http.Error(w, fmt.Sprintf("Failed to merge chunks: %v", err), http.StatusInternalServerError)
			return fmt.Errorf("failed to merge chunks: %v", err), ""
		}
		return nil, filepath.Join(this.outDir, fileName)
	}

	w.WriteHeader(http.StatusOK)
	return fmt.Errorf("chunk index is out of range: %d", chunkIndex), ""
}

func (this *pieces) mergeChunks(outDir, chunkDir, fileName string, totalChunks int) error {
	outputPath := filepath.Join(outDir, fileName)
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	for i := 0; i < totalChunks; i++ {
		chunkPath := filepath.Join(chunkDir, fmt.Sprintf("chunk_%d", i))
		chunkFile, err := os.Open(chunkPath)
		if err != nil {
			fmt.Printf("Failed to open chunk file: %v\n", err)
			return err
		}

		_, err = io.Copy(outputFile, chunkFile)
		chunkFile.Close()
		if err != nil {
			fmt.Println("Failed to copy chunk:", err)
			return err
		}
	}

	return nil
}
