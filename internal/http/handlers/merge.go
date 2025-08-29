package handlers

import (
	"io"
	"net/http"
	"os"
)

// MergeJPEGToPDFHandler handles merging uploaded JPEG images into a PDF
func MergeJPEGToPDFHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form (max 50 MB for uploads)
	err := r.ParseMultipartForm(50 << 20)
	if err != nil {
		http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve files
	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		http.Error(w, "No files uploaded", http.StatusBadRequest)
		return
	}

	// ADD BUSINESS LOGIC HERE
	// return outputPath 
	outputPath := "dummypath"

	// Serve the PDF as a download
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=merged.pdf")

	outputFile, err := os.Open(outputPath)
	if err != nil {
		http.Error(w, "Error opening PDF: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	_, err = io.Copy(w, outputFile)
	if err != nil {
		http.Error(w, "Error sending PDF: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
