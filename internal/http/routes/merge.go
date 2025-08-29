package routes

import (
	"net/http"

	h "github.com/iiitk-tools/anvil-pdf/internal/http/handlers"
)

func NewPDF() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /merge-images", h.MergeJPEGToPDFHandler)

	return mux
}