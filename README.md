

---

## **Step-by-Step: Setup `wkhtmltopdf` with Golang on Windows**

### **1. Install `wkhtmltopdf` on Windows**

1. Go to the [official wkhtmltopdf download page](https://wkhtmltopdf.org/downloads.html).
2. Download the **Windows 64-bit (MSVC)** version (recommended).
3. Run the installer and follow the setup instructions.
4. Make sure to **check the box** to add `wkhtmltopdf` to your system's **PATH** during installation.

   * If you forgot: manually add the install path to your system PATH, usually:

     ```
     C:\Program Files\wkhtmltopdf\bin
     ```

### **2. Verify Installation**

Open Command Prompt and run:

```bash
wkhtmltopdf --version
```

You should see something like:

```
wkhtmltopdf 0.12.6 (with patched qt)
```

---

### **3. Set Up Your Go Project**

Create a new folder and initialize your Go module:

```bash
mkdir mypdfgen
cd mypdfgen
go mod init mypdfgen
```

### **4. Install a Go Wrapper for wkhtmltopdf**

Install a wrapper like [`github.com/SebastiaanKlippert/go-wkhtmltopdf`](https://github.com/SebastiaanKlippert/go-wkhtmltopdf):

```bash
go get github.com/SebastiaanKlippert/go-wkhtmltopdf
```

---

### **5. Sample Go Code to Generate PDF**

Here’s a simple example:

```go
package main

import (
    "log"
    "os"

    pdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
    // Create new PDF generator
    pdfg, err := pdf.NewPDFGenerator()
    if err != nil {
        log.Fatal(err)
    }

    // Set global options
    pdfg.Dpi.Set(300)
    pdfg.Orientation.Set(pdf.OrientationPortrait)
    pdfg.PageSize.Set(pdf.PageSizeA4)

    // Create a new input page from HTML string
    page := pdf.NewPageReader(os.Stdin)
    page.Input = "<h1>Hello from Golang</h1>"
    pdfg.AddPage(page)

    // Create the PDF
    err = pdfg.Create()
    if err != nil {
        log.Fatal(err)
    }

    // Write PDF to file
    err = pdfg.WriteFile("./output.pdf")
    if err != nil {
        log.Fatal(err)
    }

    log.Println("PDF created successfully.")
}
```

You can replace `page.Input` with `pdf.NewPage("https://example.com")` if you want to generate a PDF from a URL.

---

### **6. Run the Program**

Run it from terminal:

```bash
go run main.go
```

Make sure `wkhtmltopdf.exe` is accessible via PATH — or set `pdfg.BinPath` manually to the executable.

---


