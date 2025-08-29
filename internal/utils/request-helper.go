package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)



func GetReadersFromUpload(r *http.Request) (string,[]io.Reader, error) {
    // Parse multipart form (50 MB max)
    if err := r.ParseMultipartForm(50 << 20); err != nil {
        return "",nil, err
    }

	orientation := r.FormValue("orientation")
	if orientation == "" {
		orientation = "portrait"
	}

    files := r.MultipartForm.File["files"]
    if len(files) == 0 {
        return "",nil, fmt.Errorf("no files uploaded")
    }

    readers := make([]io.Reader, 0, len(files))

    for _, fh := range files {
        f, err := fh.Open()
        if err != nil {
            return "",nil, err
        }

        // We need to read once, so slurp it into memory
        data, err := io.ReadAll(f)
        f.Close() // safe to close immediately since we copied it
        if err != nil {
            return "",nil, err
        }

        readers = append(readers, bytes.NewReader(data))
    }

    return orientation, readers, nil
}
