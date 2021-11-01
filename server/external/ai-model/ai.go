package aimodel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/core/infrastructure"
	"server/core/models"
	"time"
)

type PegasusParaphraseApi struct {
	url string
}

type pegasusNewParaphraseApiResponse struct {
	StartTime string   `json:"start_time"`
	EndTime   string   `json:"end_time"`
	Result    []string `json:"result"`
}

const UtcTimestampLayout = "2006-01-02T15:04:05.000Z"

func New() infrastructure.ParaphrasingApi {
	url := os.Getenv("PEGASUS_API_URL")
	if url == "" {
		log.Fatalln("[PegasusParaphraseApi][Error] - PEGASUS_API_URL environment variable not found.")
	}
	return PegasusParaphraseApi{
		url: url,
	}
}

func (ppa PegasusParaphraseApi) RequestParaphrase(originalText string) (models.ParaphraseResponse, error) {
	body := struct {
		Original string `json:"original"`
	}{
		Original: originalText,
	}

	postBody, _ := json.Marshal(body)

	resp, err := http.Post(fmt.Sprintf("%s/paraphrase", ppa.url), "application/json", bytes.NewBuffer(postBody))

	if err != nil {
		log.Println("[PegasusParaphraseApi][Error] - Error Getting paraphrase")
		return models.ParaphraseResponse{}, err
	}

	defer resp.Body.Close()
	var apiResp pegasusNewParaphraseApiResponse

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		log.Printf("[PegasusParaphraseApi][Error] - Error parsing json %v", err)
		return models.ParaphraseResponse{}, err
	}

	start, errS := time.Parse(UtcTimestampLayout, apiResp.StartTime)
	end, errE := time.Parse(UtcTimestampLayout, apiResp.EndTime)

	if errS != nil && errE != nil {
		log.Printf("[PegasusParaphraseApi][Error] - Error parsing start & end times")
		return models.ParaphraseResponse{}, err
	}

	return models.ParaphraseResponse{
		StartTime:  start,
		EndTime:    end,
		Paraphrase: apiResp.Result[0], // get the first value of the result
	}, nil
}
