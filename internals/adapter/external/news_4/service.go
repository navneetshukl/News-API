package news_4

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func NewsSvc() (*NewsAPIResponse, error) {

	currentDate := time.Now()
	yesterday := currentDate.AddDate(0, 0, -1)
	date := yesterday.Format("2006-01-02")

	apiKey := os.Getenv("NEWS_API_KEY")
	BaseURL := "https://newsapi.org/v2/everything?q=apple"
	URL := fmt.Sprintf("%s&from=%s&to=%s&sortBy=popularity&apiKey=%s", BaseURL, date, date, apiKey)

	req, err := http.NewRequest("GET", URL, nil) // Changed to GET as the NYT API usually uses GET for this endpoint
	if err != nil {
		log.Println("error in creating the request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Request failed with status: %s\n", resp.Status)
		return nil, errors.New("request failed with status: " + resp.Status)
	}

	var response NewsAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Println("Error decoding response:", err)
		return nil, err
	}

	return &response, nil
}
