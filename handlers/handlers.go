package handlers

import (
	"fmt"
	"strconv"
	"template/database"
	"template/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store = session.New()

func IndexView(c *fiber.Ctx) error {

	s, _ := store.Get(c)

	// For check session
	//keys := s.Keys()
	//fmt.Println("Keys = > ", keys)

	// Get value
	LoginMerchantID := s.Get("LoginMerchantID")
	if LoginMerchantID == nil {
		return c.Redirect("/login")
	}

	userProfileData, _ := GetUserSessionData(c)

	return c.Render("index", fiber.Map{
		"Title":          "Dashboard",
		"Subtitle":       "Home",
		"CurrentSession": userProfileData.LoginMerchantName,
		"Sessions":       userProfileData.Sessions,
	})
}

func GetUserSessionData(c *fiber.Ctx) (*models.UserSession, error) {
	// Get current session

	s, err := store.Get(c)

	if err != nil {
		//U := &models.UserSession{Session: "Error"}
		//return U, nil
		//function.CheckSession()

		fmt.Print("Session Store")
	}

	// Get value
	LoginMerchantName := s.Get("LoginMerchantName").(string)
	LoginMerchantID := s.Get("LoginMerchantID").(uint)
	LoginMerchantEmail := s.Get("LoginMerchantEmail").(string)
	LoginMerchantStatus := s.Get("LoginMerchantStatus").(int)
	LoginIP := s.Get("LoginIP").(string)
	LoginTime := s.Get("LoginTime").(string)
	LoginAgent := s.Get("LoginAgent").(string)
	// If there is a valid session
	if len(s.Keys()) > 0 {

		// Get profile info
		U := &models.UserSession{
			LoginMerchantName:   LoginMerchantName,
			LoginMerchantID:     LoginMerchantID,
			LoginMerchantEmail:  LoginMerchantEmail,
			LoginMerchantStatus: LoginMerchantStatus,
			Session:             "Test",
		}

		// Append session
		U.Sessions = append(
			U.Sessions,
			models.UserSessionOther{
				LoginIP:    LoginIP,
				LoginTime:  LoginTime,
				LoginAgent: LoginAgent,
			},
		)
		//fmt.Println(U)
		return U, nil
	}

	return nil, nil
}

func TransactionsView(c *fiber.Ctx) error {

	// check session
	s, _ := store.Get(c)
	// Get value
	LoginMerchantID := s.Get("LoginMerchantID")
	if LoginMerchantID == nil {
		return c.Redirect("/login")
	}

	// Get query parameters for page and limit
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid page number")
	}
	pageLimit := "10"
	limit, err := strconv.Atoi(c.Query("limit", pageLimit))
	if err != nil || limit < 1 {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid limit number")
	}

	// Calculate offset
	offset := (page - 1) * limit

	transactionList := []models.TransactionList{}
	database.DB.Db.Table("transactions").Order("transactionid desc").Limit(limit).Offset(offset).Where("client_id = ?", LoginMerchantID).Find(&transactionList)

	var total int64
	database.DB.Db.Table("transactions").Count(&total)

	fmt.Println(total)

	// Prepare pagination data
	totalPage := total / 10
	fmt.Println(totalPage)
	nextPage := page + 1
	prevPage := page - 1
	if page == 1 {
		prevPage = 0
	}

	if page >= int(totalPage+1) {
		nextPage = 0
	}
	//.Select("login_time")
	userProfileData, err := GetUserSessionData(c)
	if err != nil {
		panic(err)
	}

	//fmt.Println(loginHistory)
	return c.Render("transactions", fiber.Map{
		"Title":           "Transactions",
		"Subtitle":        "Transactions",
		"TransactionList": transactionList,
		"CurrentSession":  userProfileData.LoginMerchantName,
		"Sessions":        userProfileData.Sessions,
		"NextPage":        nextPage,
		"PrevPage":        prevPage,
		"Limit":           limit,
		"Count":           total,
	})
}
