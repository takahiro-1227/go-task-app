package users

import (
	"encoding/json"
	"go-task-app/internal/config"
	usersTypes "go-task-app/internal/users/types"
	"go-task-app/tests/globals"
	"go-task-app/tests/helpers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	helpers.InitTest()
	httpRecorder := httptest.NewRecorder()

	signUpInput := &usersTypes.SignUpInput{
		Name:     "test1",
		Password: "test1234---2A",
	}

	signUpInputJson, _ := json.Marshal(signUpInput)
	req, _ := http.NewRequest(http.MethodPost, "/sign-up", strings.NewReader(string(signUpInputJson)))

	globals.Router.ServeHTTP(httpRecorder, req)

	assert.Equal(t, http.StatusCreated, httpRecorder.Code)

	var users []usersTypes.User

	config.DB.Find(&users)

	assert.Equal(t, users[0].Name, "test1")
}

func TestSignIn(t *testing.T) {
	httpRecorder := httptest.NewRecorder()
	signInInput := &usersTypes.SignInInput{
		Name:     "test1",
		Password: "test1234---2A",
	}
	signInInputJson, _ := json.Marshal(signInInput)
	req, _ := http.NewRequest(http.MethodPost, "/sign-in", strings.NewReader(string(signInInputJson)))

	globals.Router.ServeHTTP(httpRecorder, req)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
}
