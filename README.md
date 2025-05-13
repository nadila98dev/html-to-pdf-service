

# PDF Generation in Go using wkhtmltopdf

This project uses `wkhtmltopdf` to generate PDFs from HTML content in a Go application. `wkhtmltopdf` is a command-line tool that renders HTML into PDF using WebKit.

## Prerequisites

### 1. Install wkhtmltopdf

#### macOS (Homebrew)
```bash
brew install wkhtmltopdf
Windows
Download the installer from the official site:
https://wkhtmltopdf.org/downloads.html

Choose the appropriate version for your system (usually the 64-bit one with the installer), and run the installer.

Note: On Windows, install it to a path without spaces (e.g., C:\wkhtmltopdf) to avoid command line issues.

Setup Instructions
Go Dependencies
Install the Go wrapper for wkhtmltopdf:

bash
Copy code
go get github.com/SebastiaanKlippert/go-wkhtmltopdf
Import it in your Go file:

go
Copy code
import "github.com/SebastiaanKlippert/go-wkhtmltopdf"
Environment Variable Configuration
Your Go application needs to know where to find the wkhtmltopdf binary.

macOS/Linux
If installed via Homebrew, it's usually in /usr/local/bin. You can confirm by running:

bash
Copy code
which wkhtmltopdf
Set the path in your environment:

bash
Copy code
export WKHTMLTOPDF_PATH=/usr/local/bin/wkhtmltopdf
You can add this line to your .bashrc, .zshrc, or .profile to make it permanent.

Windows
Find the path where wkhtmltopdf.exe is installed (e.g., C:\wkhtmltopdf\bin\wkhtmltopdf.exe).

Set the environment variable in PowerShell or Command Prompt:

powershell
Copy code
$env:WKHTMLTOPDF_PATH="C:\wkhtmltopdf\bin\wkhtmltopdf.exe"
To make it permanent:

Open System Properties > Environment Variables.

Add a new user variable WKHTMLTOPDF_PATH with the full path to wkhtmltopdf.exe.

Sample Usage
go
Copy code
package main

import (
    "log"
    "os"

    wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
    pdfg, err := wkhtml.NewPDFGenerator()
    if err != nil {
        log.Fatal(err)
    }

    page := wkhtml.NewPageReader(strings.NewReader("<html><body><h1>Hello PDF</h1></body></html>"))
    pdfg.AddPage(page)

    err = pdfg.Create()
    if err != nil {
        log.Fatal(err)
    }

    err = pdfg.WriteFile("output.pdf")
    if err != nil {
        log.Fatal(err)
    }

    log.Println("PDF created successfully!")
}
Notes
Make sure the binary is accessible and executable.

If you deploy to other environments (e.g., Docker), install wkhtmltopdf in the image and set the correct path.

yaml
Copy code

---



