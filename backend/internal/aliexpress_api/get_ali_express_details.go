package aliexpress_api

import (
	"backend/internal/config"
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetAliExpressDetails fügt einem Produkt Details von der AliExpress API hinzu
func GetAliExpressDetails(cfg *config.Config, product *models.Product) error {
	if cfg.AliExpressAPI.AccessToken == nil {
		err := GetAliExpressToken(cfg)
		if err != nil {
			return fmt.Errorf("failed to refresh token: %v", err)
		}
	}

	url := fmt.Sprintf("https://api.aliexpress.com/v1/item/get?product_id=%d", product.AliExpressID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cfg.AliExpressAPI.AccessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %v", err)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		err := GetAliExpressToken(cfg)
		if err != nil {
			return fmt.Errorf("failed to refresh token: %v", err)
		}
		return GetAliExpressDetails(cfg, product)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(body))
	}

	var details models.AliExpressDetails
	if err := json.Unmarshal(body, &details); err != nil {
		return fmt.Errorf("failed to parse response: %v", err)
	}

	product.AliExpressDetails = details
	return nil
}
