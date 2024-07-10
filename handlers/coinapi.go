package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ExchangeRateResponse struct {
	AssetIDBase  string  `json:"asset_id_base"`
	AssetIDQuote string  `json:"asset_id_quote"`
	Rate         float64 `json:"rate"`
}

type ExchangeRate struct {
	Time         string  `json:"time"`
	AssetIDBase  string  `json:"asset_id_base"`
	AssetIDQuote string  `json:"asset_id_quote"`
	Rate         float64 `json:"rate"`
}

type CoinPayCurrentRate struct {
	AssetIDBase string        `json:"asset_id_base"`
	Rates       []CoinPayRate `json:"rates"`
}

type CoinPayRate struct {
	Time         string  `json:"time"`
	AssetIDQuote string  `json:"asset_id_quote"`
	Rate         float64 `json:"rate"`
}

// Sample route to fetch exchange rates
func PostExchangeRate(c *fiber.Ctx) error {
	base := c.FormValue("fromcurrency")
	quote := c.FormValue("tocurrency")

	s, _ := store.Get(c)

	// Get value
	LoginMerchantID := s.Get("LoginMerchantID")

	if LoginMerchantID == nil {
		return c.Redirect("/login")
	}

	//fmt.Print("==>", base, quote)
	rate, err := getExchangeRate(base, quote)
	//fmt.Println(rate)
	//fmt.Println(rate.AssetIDBase)
	//fmt.Println(rate.AssetIDQuote)
	//fmt.Println(rate.Rate)
	if err != nil {
		fmt.Println("Data Not Found")

	}
	///////////////////////////////////
	userProfileData, err := GetUserSessionData(c)
	if err != nil {
		fmt.Println(err)
	}

	return c.Render("exchange-rate", fiber.Map{
		"Title":          "Exchange Rate",
		"Subtitle":       "Exchange Rate",
		"AssetIDBase":    rate.AssetIDBase,
		"AssetIDQuote":   rate.AssetIDQuote,
		"Rates":          rate.Rate,
		"CurrentSession": userProfileData.LoginMerchantName,
		"Sessions":       userProfileData.Sessions,
	})
}

func ExchangeRateView(c *fiber.Ctx) error {

	// check session
	s, _ := store.Get(c)
	// Get value
	LoginMerchantID := s.Get("LoginMerchantID")

	if LoginMerchantID == nil {
		return c.Redirect("/login")
	}

	userProfileData, err := GetUserSessionData(c)
	if err != nil {
		panic(err)
	}

	return c.Render("exchange-rate", fiber.Map{
		"Title":          "Exchange Rate",
		"Subtitle":       "Exchange Rate",
		"CurrentSession": userProfileData.LoginMerchantName,
		"Sessions":       userProfileData.Sessions,
	})
}

// Sample route to fetch exchange rates
func GetExchangeList(c *fiber.Ctx) error {
	s, _ := store.Get(c)

	// Get value
	LoginMerchantID := s.Get("LoginMerchantID")

	if LoginMerchantID == nil {
		return c.Redirect("/login")
	}
	url := "https://rest.coinapi.io/v1/exchangerate/USD"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		//return nil
	}
	req.Header.Add("Accept", "text/plain")
	req.Header.Add("X-CoinAPI-Key", "A871B600-DBB3-4E5E-8FAB-A7D106310C4C")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		//return nil
	}

	// Parse the JSON data into the struct
	var fireblocksData CoinPayCurrentRate
	if err := json.Unmarshal(body, &fireblocksData); err != nil {
		fmt.Println(err)
	}

	///////////////////////////////////
	userProfileData, err := GetUserSessionData(c)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(fireblocksData)
	return c.Render("exchangerate", fiber.Map{
		"Title":          "Fire Blocked User List",
		"Subtitle":       "User List",
		"Coin":           fireblocksData.AssetIDBase,
		"Rates":          fireblocksData.Rates,
		"CurrentSession": userProfileData.LoginMerchantName,
		"Sessions":       userProfileData.Sessions,
	})

}

func getExchangeRate(base string, quote string) (*ExchangeRateResponse, error) {
	//apiKey := os.Getenv("A871B600-DBB3-4E5E-8FAB-A7D106310C4C")
	apiKey := "A871B600-DBB3-4E5E-8FAB-A7D106310C4C"
	url := fmt.Sprintf("https://rest.coinapi.io/v1/exchangerate/%s/%s", base, quote)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	//req.Header.Set("X-CoinAPI-Key", apiKey)
	req.Header.Add("Accept", "text/plain")
	req.Header.Add("X-CoinAPI-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	//fmt.Println("xxxxxxxxxxxxxxx")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error from CoinAPI: %s", string(body))
	}

	var exchangeRateResponse ExchangeRateResponse
	if err := json.Unmarshal(body, &exchangeRateResponse); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return &exchangeRateResponse, nil
}
