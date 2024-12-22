package aliexpress_api

import (
	"backend/internal/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAliExpressToken(cfg *config.Config) error {
	url := "https://oauth.aliexpress.com/token"
	payload := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     cfg.AliExpressAPI.ClientID,
		"client_secret": cfg.AliExpressAPI.ClientSecret,
	}
	jsonData, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get token, status code: %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	*cfg.AliExpressAPI.AccessToken = result["access_token"].(string)
	return nil
}
