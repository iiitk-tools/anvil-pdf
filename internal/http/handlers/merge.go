package handlers

import (
	"bytes"
	"net/http"

	pdffunc "github.com/iiitk-tools/anvil-pdf/internal/pdf_func"
	"github.com/iiitk-tools/anvil-pdf/internal/utils"
)

// MergeJPEGToPDFHandler handles merging uploaded JPEG images into a PDF
func MergeJPEGToPDFHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	orientation, fileReader, err := utils.GetReadersFromUpload(r)
	if err != nil {
		http.Error(w, "Error reading Images : "+err.Error(), http.StatusInternalServerError)
		return
	}

	var pdfBuf *bytes.Buffer

	pdfBuf, err = pdffunc.ConvertImagesToPDF(fileReader,orientation)
	if err != nil {
		http.Error(w, "Error converting images to PDF: "+err.Error(), http.StatusInternalServerError)
		return
	}


	if pdfBuf == nil {
		http.Error(w, "Error generating PDF", http.StatusInternalServerError)
		return
	}

	// Serve the PDF as a download
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=merged.pdf")
	w.Header().Set("Content-Length", string(len(pdfBuf.Bytes())))

	// Write directly from buffer to response
	_, err = w.Write(pdfBuf.Bytes())
	if err != nil {
		http.Error(w, "Error sending PDF: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
