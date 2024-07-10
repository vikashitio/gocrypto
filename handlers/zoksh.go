package handlers

import (
	"fmt"
	zkosh "template/zoksh"

	"github.com/gofiber/fiber/v2"
)

var SecretKey = "sk_test_7NtaDatZ/bCX5+4xJQWXMQ=="

func ZokshView(c *fiber.Ctx) error {

	tokenProvider, err := zkosh.SignatureZoksh(SecretKey)
	fmt.Println(tokenProvider)
	fmt.Println("Errors : -> ", err)
	return c.Render("fireblocks-users", fiber.Map{
		"Title":    "Fire Blocked User List",
		"Subtitle": "User List",
	})
}
