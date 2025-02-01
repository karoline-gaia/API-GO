package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Raimundo", "teste@teste.com", "abcsffgsh")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Raimundo", user.Name)
	assert.Equal(t, "teste@teste.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Raimundo", "teste@teste.com", "abcsffgsh")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("abcsffgsh"))
	assert.False(t, user.ValidatePassword("abcsffgsht"))
	assert.NotEqual(t, "abcsffgsh", user.Password)
}
