package main

import (
	"fmt"
	"template/database"
	"template/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

var store = session.New()

func init() {

	// Init sessions store
	//handlers.InitSessionsStore()
}

func main() {
	database.ConnectDb()
	//store := session.New()

	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("ENV not Found")
		return
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	// Default middleware config will allow all origins
	app.Use(cors.New())

	// To specify allowed origins:
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
		MaxAge:           86400,
	}))

	setUpRoutes(app)

	//app.Static("/", "./public")
	app.Static("/views", "./views")
	//app.Static("/static", "./static")

	app.Listen(":3000")
}

func setUpRoutes(app *fiber.App) {

	//With Session Opening
	app.Get("/", handlers.IndexView)
	app.Get("/profile", handlers.ProfileView)
	app.Post("/profilePost", handlers.ProfilePost)
	app.Get("/login-history", handlers.Loginhistory)
	app.Get("/vault", handlers.VoltView)
	app.Get("/wallet-list", handlers.WalletListView)
	app.Get("/update-wallet-balance/:VID/:WID/:AID", handlers.UpdateWalletBalance)
	app.Get("/wallet/:VID/:WID", handlers.WalletView)
	app.Get("/generate-new-wallet-address/:VID/:WID", handlers.CreateVaultWalletAddress)
	app.Get("/generate-new-wallet-address/:VID/", handlers.CreateVaultWalletView)
	app.Post("/generate-new-wallet-address", handlers.CreateVaultWallet)
	app.Get("/generate-new-vault", handlers.CreateNewVault)
	app.Get("/fireblocks-users", handlers.UsersView)
	app.Get("/vault-accounts", handlers.VaultAccountsView)
	app.Get("/qrcode", handlers.QrcodeView)
	//Without Session Open
	app.Get("/login", handlers.LoginView)
	app.Post("/loginPost", handlers.LoginPost)
	app.Get("/registration", handlers.RegistrationView)
	app.Post("/registrationPost", handlers.RegistrationPost)
	app.Get("/logout", handlers.LogOut)
	//app.Get("/items", handlers.GetPaginatedItems)
	app.Post("/exchanGeneratePost", handlers.PostExchangeRate)
	app.Get("/exchange-rate", handlers.ExchangeRateView)
	//app.Get("/exchangerate/:asset_id_base", handlers.GetExchangeList)
	app.Get("/exchangerate", handlers.GetExchangeList)
	app.Get("/coin-list", handlers.GetCoinList)
	app.Get("/coin-list/edit/:TID", handlers.EditCoin)
	app.Get("/coin-list/delete/:TID", handlers.DeleteCoin)
	app.Get("/add-coin", handlers.AddCoinView)
	app.Post("/coinPost", handlers.CoinPost)
	app.Get("/transactions", handlers.TransactionsView)
	app.Get("/pdf-transactions", handlers.TransactionsPDF)
	app.Get("/zoksh", handlers.ZokshView)
}
