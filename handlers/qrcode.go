package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	qrcode "github.com/skip2/go-qrcode"
)

func QrcodeView(c *fiber.Ctx) error {
	// Generate QR code with content "Hello, World!"
	qr, err := qrcode.Encode("Hello, World!", qrcode.Medium, 256)
	if err != nil {
		log.Println("Error generating QR code:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error generating QR code")
	}

	// Set response headers
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Send the QR code image as response
	return c.Send(qr)

	// return c.Render("login", fiber.Map{
	// 	"Title": "Login Form",
	// 	"Alert": "",
	// })

}
