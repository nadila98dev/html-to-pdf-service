package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

// Helper function to handle errors
func handleError(w http.ResponseWriter, logMessage string, clientMessage string, statusCode int, err error) {
    log.Println(logMessage, err)
    http.Error(w, clientMessage, statusCode)
}

func convertHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
        return
    }

    os.Setenv("WKHTMLTOPDF_PATH", `C:\Program Files\wkhtmltopdf\bin\wkhtmltopdf.exe`)
    log.Println("WKHTMLTOPDF_PATH set to:", os.Getenv("WKHTMLTOPDF_PATH"))

    htmlData, err := io.ReadAll(r.Body)
    if err != nil {
        handleError(w, "Error reading HTML:", "Error reading HTML", http.StatusBadRequest, err)
        return
    }

    pdfg, err := wkhtmltopdf.NewPDFGenerator()
    if err != nil {
        handleError(w, "PDFGenerator init error:", "Failed to create PDF generator", http.StatusInternalServerError, err)
        return
    }

    page := wkhtmltopdf.NewPageReader(bytes.NewReader(htmlData))
    page.EnableLocalFileAccess.Set(true) // Allow access to local files
    pdfg.AddPage(page)

    // Set the page size to A4
    pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

    if err := pdfg.Create(); err != nil {
        handleError(w, "PDF generation error:", "PDF generation failed", http.StatusInternalServerError, err)
        return
    }

    w.Header().Set("Content-Type", "application/pdf")
    w.Header().Set("Content-Disposition", "attachment; filename=output.pdf")
    w.Write(pdfg.Bytes())
}

func main() {
    http.HandleFunc("/convert", convertHandler)
    log.Println("Server is running on http://localhost:6060")
    log.Fatal(http.ListenAndServe(":6060", nil))
}