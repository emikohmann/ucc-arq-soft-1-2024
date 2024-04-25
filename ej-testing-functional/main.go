package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	itemID := "MLA904895322"

	item, err := GetItem(itemID)
	if err != nil {
		panic(fmt.Errorf("error getting item %s: %w", itemID, err))
	}

	fmt.Printf("successfully obtained item %s: %s\n", itemID, item["title"])
}

func GetItem(itemID string) (map[string]interface{}, error) {
	response, err := http.Get(fmt.Sprintf("https://api.mercadolibre.com/items/%s", itemID))
	if err != nil {
		return nil, fmt.Errorf("error executing http.Get: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected %d in http.Get", response.StatusCode)
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error executing io.ReadAll: %w", err)
	}
	output := make(map[string]interface{})
	if err := json.Unmarshal(bytes, &output); err != nil {
		return nil, fmt.Errorf("error executing json.Unmarshal: %w", err)
	}
	return output, nil
}
