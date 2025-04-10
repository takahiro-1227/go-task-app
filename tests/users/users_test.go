package users

import (
	"encoding/json"
	"go-task-app/internal/config"
	usersTypes "go-task-app/internal/users/types"
	"go-task-app/tests/helpers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func requestSignUp(signUpInput *usersTypes.SignUpInput) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodPost, "/sign-up", helpers.CreateReaderFromStruct(signUpInput))
	return helpers.Request(req, nil)
}
func TestSignUp(t *testing.T) {
	helpers.InitIntegrationTest()

	httpRecorder := requestSignUp(&usersTypes.SignUpInput{
		Name:     "test1",
		Password: "test1234---2A",
	})
	assert.Equal(t, http.StatusCreated, httpRecorder.Code)

	var users []usersTypes.User
	config.DB.Find(&users)
	assert.Equal(t, users[0].Name, "test1")
}

func TestSignIn(t *testing.T) {
	helpers.InitIntegrationTest()
	requestSignUp(&usersTypes.SignUpInput{
		Name:     "test1",
		Password: "test1234---2A",
	})

	signInInput := &usersTypes.SignInInput{
		Name:     "test1",
		Password: "test1234---2A",
	}
	req, _ := http.NewRequest(http.MethodPost, "/sign-in", helpers.CreateReaderFromStruct(signInInput))
	httpRecorder := helpers.Request(req, nil)
	assert.Equal(t, http.StatusOK, httpRecorder.Code)

	var res usersTypes.SignInResponse
	err := json.NewDecoder(httpRecorder.Body).Decode(&res)
	if err != nil {
		panic(err)
	}
	assert.NotNil(t, res.AccessToken)

	signInInput = &usersTypes.SignInInput{
		Name:     "test1",
		Password: "test12345---2A",
	}
	req, _ = http.NewRequest(http.MethodPost, "/sign-in", helpers.CreateReaderFromStruct(signInInput))
	httpRecorder = helpers.Request(req, nil)
	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
}
