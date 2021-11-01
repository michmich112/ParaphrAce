package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"server/context"
	"server/core/models"
	"testing"
)

type MockUserRepository struct {
	user models.User
}

func (mur MockUserRepository) Create(user models.User) (models.User, error) {
	mur.user = user
	mur.user.Id = 10 //act as the auto increment of the id
	return mur.user, nil
}

func (mur MockUserRepository) GetBySessionToken(token string) (models.User, error) {
	if mur.user.SessionToken == token {
		return mur.user, nil
	}
	return models.User{}, errors.New("User not found")
}

func mockAppContext() context.AppContext {
	return context.AppContext{
		UserRepository: MockUserRepository{},
	}
}

func TestCreateUserPassing(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/user/create", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	CreateUser(mockAppContext()).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Expected status code Created, received %v", status)
	}

	user := models.User{}
	err = json.NewDecoder(rr.Body).Decode(&user)

	if err != nil {
		t.Fatal(err)
	}

	if user.Id == 0 || user.SessionToken == "" {
		t.Errorf("Invalid User Data. Received user.Id: %d, user.SessionToken: %s", user.Id, user.SessionToken)
	}

}
