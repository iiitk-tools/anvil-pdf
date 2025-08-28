# 📄 Anvil-PDF

**Anvil-PDF** is a lightweight PDF toolkit written in **Go**.  
The goal is to provide a fast, developer-friendly alternative to tools like iLovePDF and Stirling-PDF.

🚀 **MVP Features (current)**

- 📷 **JPEG → PDF**: Convert multiple images into a single PDF file.  
- 📉 **PDF Compression**: Reduce the size of PDF files by downsampling/recompressing images.

---

## ⚙️ Tech Stack

- **Language**: Go (Golang)  
- **Core PDF Operations**: [pdfcpu](https://github.com/pdfcpu/pdfcpu)  
- **Hot Reload (Dev)**: [Air](https://github.com/air-verse/air)  

---

## 🛠 Development Setup

1. **Fork the repo** on GitHub.

2. Clone your fork:

   ```bash
   git clone https://github.com/<your-username>/anvil-pdf.git
   cd anvil-pdf
   ```

3. Add upstream (for syncing with original repo):

   ```bash
   git remote add upstream https://github.com/iiitk-tools/anvil-pdf.git
   ```

4. Keep your fork updated:

   ```bash
   git pull upstream main
   ```

5. Hot Reload with Air

    - We use **Air** for live reloading during development.

    #### Install Air

    ```bash
    go install github.com/air-verse/air@latest
    ```

    Ensure `$GOPATH/bin` or `$HOME/go/bin` is in your `$PATH`.

    #### Initialize Air

    ```bash
    air init
    ```

    This generates `.air.toml`.

6. Run with Hot Reload

    ```bash
    air -- --port 8081
    ```

    Whenever you modify Go files, Air rebuilds & restarts automatically.

---

## 📂 Project Structure

```
anvil-pdf/
│── cmd/
│   └── server/          # main server entrypoint
│
│── internal/
│   ├── pdf/             # PDF operations (convert, compress, etc.)
│   └── api/             # HTTP handlers
│
│── pkg/                 # reusable helpers (logger, config, etc.)
│
│── go.mod
│── go.sum
│── README.md
```

---
