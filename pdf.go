package main

import (
	"bytes"
	"time"

	"github.com/go-pdf/fpdf"
)

// NOTE: Capitalized function name to export it from the package

/**
* GeneratePDFVoucher generates a PDF voucher and returns it as a byte slice.
 */
func GeneratePDFVoucher(vr VoucherRequest) ([]byte, error) {
	p := fpdf.New("P", "mm", "A4", "")
	p.SetMargins(15, 15, 15)
	p.AddPage()

	// Header
	p.SetFont("Arial", "B", 20)
	p.Cell(0, 12, "Voucher - " + vr.Activity.Name)
	p.Ln(14)

	// Body
	p.SetFont("Arial", "", 12)

	startStr := vr.Activity.StartDate
    endStr := vr.Activity.EndDate

    if t, err := time.Parse("2006-01-02", vr.Activity.StartDate); err == nil {
        startStr = t.Format("02 Jan 2006") // e.g. "05 Feb 2025"
    }
    if t, err := time.Parse("2006-01-02", vr.Activity.EndDate); err == nil {
        endStr = t.Format("02 Jan 2006")
    }

	row := func(label, value string) {
		p.CellFormat(40, 8, label+":", "", 0, "", false, 0, "")
		p.MultiCell(0, 8, value, "", "L", false)
	}

	row("Booking ID", 		  	vr.BookingID)
	row("Client Name",       	vr.Client.Salutation + " " + vr.Client.FirstName + " " + vr.Client.LastName)
	row("Activity Name",      	vr.Activity.Name)
    row("Start Date", 			startStr)
    row("End Date", 			endStr)
	row("Agency Name",       	vr.Agency.Name)
	row("Agency Contact",    	vr.Agency.Email + " | " + vr.Agency.Phone)
	row("Agency Address",    	vr.Agency.StreetAndNumber + ", " + vr.Agency.Postcode + " " + vr.Agency.City + ", " + vr.Agency.Country)

	var buf bytes.Buffer

	if err := p.Output(&buf); err != nil { 
		return nil, err
	}

	return buf.Bytes(), nil
}