package main

import (
	"bytes"
	_ "embed"
	"time"

	i18n "voucher/lang"

	"github.com/go-pdf/fpdf"
)

// NOTE: Capitalized function name to export it from the package

/**
* GeneratePDFVoucher generates a PDF voucher and returns it as a byte slice.
 */
func GeneratePDFVoucher(vr VoucherRequest, lang string) ([]byte, error) {
	p := fpdf.New("P", "mm", "A4", "")
	p.SetMargins(15, 15, 15)
	p.AddPage()

	tr := p.UnicodeTranslatorFromDescriptor("") 			// maps UTF-8 -> CP1252

	dict := i18n.Translate(lang)
	dateFmt := "02 Jan 2006"
	if lang == "de" { dateFmt = "02.01.2006" }

	// Header
	p.SetFont("Arial", "B", 20)
	p.Cell(0, 12, tr(dict["voucher"] + " - " + vr.Activity.Name))
	p.Ln(14)

	// Body
	p.SetFont("Arial", "", 12)

	startStr := vr.Activity.StartDate
    endStr := vr.Activity.EndDate

    if t, err := time.Parse("2006-01-02", vr.Activity.StartDate); err == nil {
        startStr = t.Format(dateFmt) // e.g. "05 Feb 2025"
    }
    if t, err := time.Parse("2006-01-02", vr.Activity.EndDate); err == nil {
        endStr = t.Format(dateFmt)
    }

	row := func(label, value string) {
		p.CellFormat(40, 8, label+":", "", 0, "", false, 0, "")
		p.MultiCell(0, 8, value, "", "L", false)
	}

	row(tr(dict["booking_id"]),			vr.BookingID)
	row(tr(dict["client_name"]),    	vr.Client.Salutation + " " + vr.Client.FirstName + " " + vr.Client.LastName)
	row(tr(dict["activity"]),      		vr.Activity.Name)
    row(tr(dict["start_date"]),			startStr)
    row(tr(dict["end_date"]), 			endStr)
	row(tr(dict["agency_name"]),       	vr.Agency.Name)
	row(tr(dict["agency_contact"]),    	vr.Agency.Email + " | " + vr.Agency.Phone)
	row(tr(dict["agency_addr"]),    	vr.Agency.StreetAndNumber + ", " + vr.Agency.Postcode + " " + vr.Agency.City + ", " + vr.Agency.Country)

	var buf bytes.Buffer

	if err := p.Output(&buf); err != nil { 
		return nil, err
	}

	return buf.Bytes(), nil
}