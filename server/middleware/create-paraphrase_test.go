package middleware

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"server/context"
	"server/core/models"
	"strings"
	"testing"
	"time"
)

// Mock Paraphrase Repository
type MockParaphraseRepository struct {
	paraphrases map[int64]models.Paraphrase
}

func (mpr MockParaphraseRepository) GetById(id int64) (models.Paraphrase, error) {
	paraphrase, ok := mpr.paraphrases[id]
	if !ok {
		return paraphrase, errors.New("Paraphrase not found")
	}
	return paraphrase, nil
}

func (mpr MockParaphraseRepository) Create(paraphrase models.Paraphrase) (models.Paraphrase, error) {
	paraphrase.Id = int64(len(mpr.paraphrases) + 1) // autoincrement ID
	mpr.paraphrases[paraphrase.Id] = paraphrase
	return paraphrase, nil
}

func (mpr MockParaphraseRepository) Update(paraphrase models.Paraphrase) (models.Paraphrase, error) {
	if _, ok := mpr.paraphrases[paraphrase.Id]; !ok {
		return models.Paraphrase{}, errors.New(fmt.Sprintf("Cannot Update: Paraphrase with id %d does not exist.", paraphrase.Id))
	}
	mpr.paraphrases[paraphrase.Id] = paraphrase
	return paraphrase, nil
}

func (mpr MockParaphraseRepository) AddRating(paraphraseId int64, ratingId int64) (models.Paraphrase, error) {
	paraphrase, ok := mpr.paraphrases[paraphraseId]
	if !ok {
		return models.Paraphrase{}, errors.New(fmt.Sprintf("Cannot add rating, Parraphrase with id %d does not exist", paraphraseId))
	}
	paraphrase.RatingId = sql.NullInt64{Int64: ratingId, Valid: true}
	mpr.paraphrases[paraphraseId] = paraphrase
	return paraphrase, nil
}

// Mock Storage
type MockStorageDoc struct {
	key string
	url string
	doc string
}

type MockStorageClient struct {
	docs map[string]MockStorageDoc
}

func (msc MockStorageClient) Save(key string, doc string) (string, error) {
	url := fmt.Sprintf("https://s3.aws.com/bucket/%s", key)
	msc.docs[url] = MockStorageDoc{
		key: key,
		url: url,
		doc: doc,
	}
	return url, nil
}

// Mock Ai Model
type MockAiModel struct {
	StartTime time.Time
	EndTime   time.Time
}

func (mam MockAiModel) RequestParaphrase(originalText string) (models.ParaphraseResponse, error) {
	return models.ParaphraseResponse{
		StartTime:  mam.StartTime,
		EndTime:    mam.EndTime,
		Paraphrase: originalText[1:], // just for now to test
	}, nil
}

func createMockedAppContext(withStorage bool) context.AppContext {
	mockUserRepository := MockUserRepository{
		user: models.User{Id: 5, SessionToken: "token_5"},
	}
	mockParaphraseRepository := MockParaphraseRepository{
		paraphrases: make(map[int64]models.Paraphrase),
	}
	mockStorageClient := MockStorageClient{
		docs: make(map[string]MockStorageDoc),
	}
	mockAiModel := MockAiModel{
		StartTime: time.Now(),
		EndTime:   time.Now().Add(100),
	}

	mockContext := context.AppContext{
		UserRepository:        mockUserRepository,
		ParaphraseRespository: mockParaphraseRepository,
		WithStorage:           withStorage,
		Storage:               mockStorageClient,
		ParaphraseApi:         mockAiModel,
	}

	return mockContext
}

func TestCreateParaphrasePassingWithStorage(t *testing.T) {
	mockContext := createMockedAppContext(true)

	bodyText := `{
		"session_token": "token_5",
		"original_text": "What is love, baby don't hurt me no more"
	}`

	req, err := http.NewRequest("POST", "/api/paraphrase/create", strings.NewReader(bodyText))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	CreateParaphrase(mockContext).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code OK, received %v", status)
	}

	cpRes := createParaphraseRes{}
	json.NewDecoder(rr.Body).Decode(&cpRes)

	if cpRes.Result != "hat is love, baby don't hurt me no more" {
		t.Errorf("Expected paraphrase, received %s", cpRes.Result)
	}
}

func TestCreateParaphrasePassingWithoutStorage(t *testing.T) {
	mockContext := createMockedAppContext(false)

	bodyText := `{
		"session_token": "token_5",
		"original_text": "I've Got a feeling, that its going to be a good time."
	}`

	req, err := http.NewRequest("POST", "/api/paraphrase/create", strings.NewReader(bodyText))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	CreateParaphrase(mockContext).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code OK, received %v", status)
	}

	cpRes := createParaphraseRes{}
	json.NewDecoder(rr.Body).Decode(&cpRes)

	if cpRes.Result != "'ve Got a feeling, that its going to be a good time." {
		t.Errorf("Expected paraphrase, received %s", cpRes.Result)
	}
}

func TestCreateParaphraseInvalidToken(t *testing.T) {
	mockContext := createMockedAppContext(true)

	bodyText := `{
		"session_token": "invalid_token",
		"originalText": "What is love, baby don't hurt me no more"
	}`

	req, err := http.NewRequest("POST", "/api/paraphrase/create", strings.NewReader(bodyText))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	CreateParaphrase(mockContext).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Expected status code Unauthorized, received %v", status)
	}

	if body := rr.Body.String(); body != "" {
		t.Errorf("Expected body to be empty, received: %s", body)
	}

}

func TestCreateParaphraseNoOriginalText(t *testing.T) {
	mockContext := createMockedAppContext(true)

	bodyText := `{
		"session_token": "token_5"
	}`

	req, err := http.NewRequest("POST", "/api/paraphrase/create", strings.NewReader(bodyText))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	CreateParaphrase(mockContext).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status Bad Request, received: %v", status)
	}

	if body := rr.Body.String(); body != "" {
		t.Errorf("Expected body to be empty, received: %s", body)
	}
}

func TestCreateParaphraseNonJsonContent(t *testing.T) {
	mockContext := createMockedAppContext(true)

	bodyText := "Create a paraphrase for me please"

	req, err := http.NewRequest("POST", "/api/paraphrase/create", strings.NewReader(bodyText))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	CreateParaphrase(mockContext).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status Bad Request, received: %v", status)
	}

	if body := rr.Body.String(); body != "" {
		t.Errorf("Expected body to be empty, received: %s", body)
	}
}
