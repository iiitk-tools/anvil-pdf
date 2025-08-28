package pdffunc_test

import (
	"io"
	"os"
	"testing"

	pdffunc "github.com/iiitk-tools/anvil-pdf/internal/pdf_func"
)

func TestConvertImagesToPDF(t *testing.T) {
	// Open two sample image files (ensure both exist in the tmp directory)
	imgFile1, err := os.Open("/Users/vijayvenkatj/Projects/anvil/anvil-pdf/tmp/sample.png")
	if err != nil {
		t.Fatalf("failed to open first test image: %v", err)
	}
	defer imgFile1.Close()

	imgFile2, err := os.Open("/Users/vijayvenkatj/Projects/anvil/anvil-pdf/tmp/sample2.png")
	if err != nil {
		t.Fatalf("failed to open second test image: %v", err)
	}
	defer imgFile2.Close()

	imgs := []io.Reader{imgFile1, imgFile2}

	pdfBuff, err := pdffunc.ConvertImagesToPDF(imgs, "portrait")
	if err != nil {
		t.Fatalf("ConvertImagesToPDF failed: %v", err)
	}
	if pdfBuff == nil || pdfBuff.Len() == 0 {
		t.Fatal("expected non-empty PDF buffer")
	}

	// Optionally, write the PDF to disk for manual inspection
	os.WriteFile("/Users/vijayvenkatj/Projects/anvil/anvil-pdf/tmp/out.pdf", pdfBuff.Bytes(), 0644)
}
