package tasks

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go-task-app/internal/config"
	"go-task-app/internal/routes"
	usersTypes "go-task-app/internal/users/types"
	"go-task-app/tests/helpers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSignUp(t *testing.T) {
	config.LoadEnv("../../.env")
	helpers.InitDB()
	config.ConnectDB()

	router := routes.SetUpRouter()
	httpRecorder := httptest.NewRecorder()

	signUpInput := &usersTypes.SignUpInput{
		Name:     "test1",
		Password: "test1234---2A",
	}

	signUpInputJson, _ := json.Marshal(signUpInput)
	req, _ := http.NewRequest("POST", "/sign-up", strings.NewReader(string(signUpInputJson)))

	router.ServeHTTP(httpRecorder, req)

	assert.Equal(t, http.StatusCreated, httpRecorder.Code)

	var users []usersTypes.User

	config.DB.Find(&users)

	assert.Equal(t, users[0].Name, "test1")
}
