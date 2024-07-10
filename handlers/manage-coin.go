package handlers

import (
	"fmt"
	"os"
	"strconv"
	"template/database"
	"template/models"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// For Manage Coin

func GetCoinList(c *fiber.Ctx) error {

	// check session
	s, _ := store.Get(c)
	// Get value
	LoginMerchantID := s.Get("LoginMerchantID")
	if LoginMerchantID == nil {
		return c.Redirect("/login")
	}

	Alerts := s.Get("Alerts")
	s.Delete("Alerts")
	if err := s.Save(); err != nil {
		panic(err)
	}

	// Get query parameters for page and limit
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid page number")
	}
	limit, err := strconv.Atoi(c.Query("limit", "4"))
	if err != nil || limit < 1 {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid limit number")
	}

	// Calculate offset
	offset := (page - 1) * limit

	coinList := []models.CoinList{}
	database.DB.Db.Table("coin_list").Order("coin ASC").Limit(limit).Offset(offset).Find(&coinList)
	//.Where("client_id = ?", LoginMerchantID)

	var total int64
	database.DB.Db.Table("coin_list").Find(&coinList).Count(&total)

	fmt.Println(total)
	//fmt.Println(loginHistory)
	// Prepare pagination data
	nextPage := page + 1
	prevPage := page - 1
	if page == 1 {
		prevPage = 0
	}
	//.Select("login_time")
	userProfileData, err := GetUserSessionData(c)
	if err != nil {
		panic(err)
	}

	//fmt.Println(coinList)
	return c.Render("coin-list", fiber.Map{
		"Title":          "Coin List",
		"Subtitle":       "Coin List",
		"Action":         "List",
		"Alert":          Alerts,
		"CoinList":       coinList,
		"CurrentSession": userProfileData.LoginMerchantName,
		"Sessions":       userProfileData.Sessions,
		"NextPage":       nextPage,
		"PrevPage":       prevPage,
		"Limit":          limit,
		"Count":          total,
	})
}

func AddCoinView(c *fiber.Ctx) error {

	// check session
	s, _ := store.Get(c)
	// Get value
	LoginMerchantID := s.Get("LoginMerchantID")
	if LoginMerchantID == nil {
		return c.Redirect("/login")
	}

	//.Select("login_time")
	userProfileData, err := GetUserSessionData(c)
	if err != nil {
		panic(err)
	}

	return c.Render("coin-list", fiber.Map{
		"Title":          "Coin List",
		"Subtitle":       "Coin List",
		"Action":         "Add",
		"CurrentSession": userProfileData.LoginMerchantName,
		"Sessions":       userProfileData.Sessions,
	})
}

func EditCoin(c *fiber.Ctx) error {

	TID := c.Params("TID")

	// check session
	s, _ := store.Get(c)
	// Get value
	LoginMerchantID := s.Get("LoginMerchantID")
	if LoginMerchantID == nil {
		return c.Redirect("/login")
	}

	coinData := models.CoinList{}
	database.DB.Db.Table("coin_list").Where("coin_id = ?", TID).Find(&coinData)
	//.Select("login_time")
	fmt.Println(coinData)

	//.Select("login_time")
	userProfileData, err := GetUserSessionData(c)
	if err != nil {
		panic(err)
	}

	return c.Render("coin-list", fiber.Map{
		"Title":          "Coin List",
		"Subtitle":       "Coin List",
		"Action":         "Edit",
		"CoinID":         coinData.Coin_id,
		"CoinTitle":      coinData.Coin,
		"CoinIcon":       coinData.Icon,
		"CoinStatus":     coinData.Status,
		"CurrentSession": userProfileData.LoginMerchantName,
		"Sessions":       userProfileData.Sessions,
	})
}

func CoinPost(c *fiber.Ctx) error {
	// Parses the request body
	//coin_id := c.FormValue("coin_id")
	//mode := c.FormValue("mode")
	coin := c.FormValue("coin")
	//icon := c.FormValue("icon")

	status1, err := strconv.ParseInt(c.FormValue("status"), 10, 32)
	if err != nil {
		fmt.Println("Error 104")
		//return c.Status(fiber.StatusBadRequest).SendString("Invalid number format 11")
	}
	status := int(status1)
	vvv := c.FormValue("coinId")
	cid, err := strconv.ParseUint(vvv, 10, 32)
	if err != nil {
		fmt.Println("Error 105")
		//return c.Status(fiber.StatusBadRequest).SendString("Invalid number format 22")
	}
	coin_id := uint(cid)
	s, _ := store.Get(c)

	///////////
	file, err := c.FormFile("icon")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to read file")
	}

	uploadDir := "./views/images"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	filePath := fmt.Sprintf("%s/%s", uploadDir, file.Filename)
	err = c.SaveFile(file, filePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save file")
	}

	//////////

	// if GET ID than work update else insert
	// for Full path use- filePath & only file name use file.Filename
	result := database.DB.Db.Table("coin_list").Save(&models.CoinList{Coin_id: coin_id, Coin: coin, Icon: file.Filename, Status: status})

	//fmt.Println(loginList.Status)
	Alerts := "Coin Updated successfully"
	if result.Error != nil {
		//fmt.Println("ERROR in QUERY")
		Alerts = "Coin Not Updated"
	}

	// check session

	s.Set("Alerts", Alerts)
	if err := s.Save(); err != nil {
		//panic(err)
		fmt.Println("Session Save Issue")
	}

	return c.Redirect("/coin-list")

}

func DeleteCoin(c *fiber.Ctx) error {

	s, _ := store.Get(c)

	id := c.Params("TID")
	var item models.CoinList
	database.DB.Db.Table("coin_list").First(&item, id)
	database.DB.Db.Table("coin_list").Delete(&item)

	Alerts := "Coin Deleted successfully"

	// check session

	s.Set("Alerts", Alerts)
	if err := s.Save(); err != nil {
		//panic(err)
		fmt.Println("Session Save Issue")
	}

	return c.Redirect("/coin-list")

}
