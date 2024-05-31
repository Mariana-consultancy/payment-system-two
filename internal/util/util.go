package util

import (
	//"encoding/binary"
	//"crypto/rand"
	"math/rand"
	"net/http"
	"net/mail"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Response is customized to help return all responses need
func Response(c *gin.Context, message string, status int, data interface{}, errs []string) {
	responsedata := gin.H{
		"message":   message,
		"data":      data,
		"errors":    errs,
		"status":    http.StatusText(status),
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
	}

	c.IndentedJSON(status, responsedata)
}

// Hash password changes the password to bytes, it sorts of encrypts the password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Function for
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

const (
	min = 11111111
	max = 99999999
)

// function to Generate unique Account numbers
func GenerateAccountNo() (int, error) {

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min, nil

}
