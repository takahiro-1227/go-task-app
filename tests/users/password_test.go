package users

import (
	"go-task-app/internal/users/constants"
	"go-task-app/internal/users/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword (t *testing.T) {
	err := services.ValidatePassword("testt1---_2")
	assert.Equal(t, err, nil)

	err = services.ValidatePassword("test1---_")
	assert.Equal(t, err, constants.ErrLessPasswordLength)

	err = services.ValidatePassword("testT1-1testT2-2testT3-3testT4-4")
	assert.Equal(t, err, nil)

	err = services.ValidatePassword("testT1-1testT2-2testT3-3testT4-4T")
	assert.Equal(t, err, constants.ErrOverPasswordLength)

	err = services.ValidatePassword("testT1testT")
	assert.Equal(t, err, constants.ErrPasswordCharacterCategory)

	err = services.ValidatePassword("123654789078_")
	assert.Equal(t, err, constants.ErrPasswordCharacterCategory)

	err = services.ValidatePassword("testttesttt_")
	assert.Equal(t, err, constants.ErrPasswordCharacterCategory)
}