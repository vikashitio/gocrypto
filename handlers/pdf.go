package handlers

import (
	"bytes"
	"strconv"
	"template/database"
	"template/models"

	"github.com/gofiber/fiber/v2"
	"github.com/jung-kurt/gofpdf"
)

// User struct represents a user record in the database
type User struct {
	ID    int
	Name  string
	Email string
}

func TransactionsPDF(c *fiber.Ctx) error {

	s, _ := store.Get(c)

	// Get value
	LoginMerchantID := s.Get("LoginMerchantID")
	if LoginMerchantID == nil {
		return c.Redirect("/login")
	}

	transactionList := []models.TransactionList{}
	database.DB.Db.Table("transactions").Order("transactionid desc").Where("client_id = ?", LoginMerchantID).Find(&transactionList)

	//fmt.Println("DATA ARE =>", transactionList)

	// // Generate PDF with dynamic data
	pdf, err := generatePDF(transactionList)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate PDF")
	}
	//mid := strconv.FormatUint(uint64(LoginMerchantID), 10)
	// // Set the response headers to serve the PDF file
	c.Response().Header.Set("Content-Type", "application/pdf")
	c.Response().Header.Set("Content-Disposition", "attachment; filename=Transaction.pdf")
	c.Response().SetBody(pdf)

	return nil
}

// generatePDF generates a PDF with the given user data
func generatePDF(transactionList []models.TransactionList) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)

	// Title
	pdf.Cell(40, 10, "Transactions List")

	pdf.Ln(12) // New line

	// Column headers
	pdf.SetFont("Arial", "B", 8)
	pdf.Cell(20, 10, "TransID")
	pdf.Cell(20, 10, "WalletID")
	pdf.Cell(30, 10, "Amount")
	pdf.Cell(40, 10, "Type")
	pdf.Cell(50, 10, "Status")
	pdf.Cell(60, 10, "Timestamp")
	pdf.Ln(10)

	// Table content
	pdf.SetFont("Arial", "", 8)

	for _, row := range transactionList {
		mid := strconv.FormatUint(uint64(row.Transactionid), 10)
		pdf.Cell(20, 10, mid)
		pdf.Cell(20, 10, row.Walletid)
		pdf.Cell(30, 10, row.Amount)
		pdf.Cell(40, 10, row.Transactiontype)
		pdf.Cell(50, 10, row.Status)
		pdf.Cell(60, 10, row.Timestamp)
		pdf.Ln(10)
	}

	// Create a buffer to write the PDF content
	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
