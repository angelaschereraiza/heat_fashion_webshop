package aliexpress

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetAliExpressProductDetails get product details from AliExpress with the aliExpressID
func GetAliExpressProductDetails(aliExpressID int, apiKey string) (*models.ProductDetails, error) {
	url := fmt.Sprintf("%s?product_id=%d", "https://api.example.com/aliexpress/product/details", aliExpressID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(body))
	}

	var productDetails models.ProductDetails
	if err := json.Unmarshal(body, &productDetails); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return &productDetails, nil
}
