package main

import (
	"bytes"

	"github.com/go-pdf/fpdf"
)

// NOTE: Capitalized function name to export it from the package

/**
* GeneratePDFVoucher generates a PDF voucher and returns it as a byte slice.
 */
// TODO: add parameters for dynamic content "data VoucherData"
func GeneratePDFVoucher() ([]byte, error) {
	p := fpdf.New("P", "mm", "A4", "")
	p.SetMargins(15, 15, 15)
	p.AddPage()
	p.SetFont("Arial", "B", 20)
	p.Cell(0, 12, "Voucher Content Here")			// Placeholder content

	var buf bytes.Buffer

	if err := p.Output(&buf); err != nil { 
		return nil, err
	}

	return buf.Bytes(), nil
}