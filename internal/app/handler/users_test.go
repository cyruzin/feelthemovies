package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/google/uuid"
)

func TestGetUsersSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.HandleFunc("/v1/users", h.handler.GetUsers)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
func TestGetUserSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.HandleFunc("/v1/user/{id}", h.handler.GetUser)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestCreateUserSuccess(t *testing.T) {
	var newUser = []byte(`{
		"name":"Travis Fox", 
		"email":"travis_fox@outlook.com",
		"password": "qw12erty"
		}`)

	token, err := h.handler.GenerateToken(info)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/v1/user", bytes.NewBuffer(newUser))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	router.With(h.handler.AuthMiddleware).HandleFunc("/v1/user", h.handler.CreateUser)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusCreated, status)
	}
}

func TestUpdateUserSuccess(t *testing.T) {
	var upUser = []byte(`{
		"name":"Travis Fox Jr.", 
		"email":"travis_fox_jr@outlook.com",
		"password": "qw12erty"
		}`)

	token, err := h.handler.GenerateToken(info)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/v1/user/1", bytes.NewBuffer(upUser))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	router.With(h.handler.AuthMiddleware).HandleFunc("/v1/user/{id}", h.handler.UpdateUser)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestDeleteUserSuccess(t *testing.T) {
	token, err := h.handler.GenerateToken(info)
	if err != nil {
		t.Fatal(err)
	}

	newUser := model.User{
		Name:      "John Silver",
		Email:     "johnsilver@test.org",
		Password:  "qw12erty",
		APIToken:  uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	queryUserInsert := `
		INSERT INTO users (
		name, 
		email, 
		password,
		api_token, 
		created_at, 
		updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := h.database.Exec(
		queryUserInsert,
		newUser.Name,
		newUser.Email,
		newUser.Password,
		newUser.APIToken,
		newUser.CreatedAt,
		newUser.UpdatedAt,
	)
	if err != nil {
		t.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("DELETE", "/v1/user/"+strconv.FormatInt(id, 10), nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	router.With(h.handler.AuthMiddleware).HandleFunc("/v1/user/{id}", h.handler.DeleteUser)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
