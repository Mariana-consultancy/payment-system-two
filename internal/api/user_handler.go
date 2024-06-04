package api

import (
	"net/http"
	"os"
	"payment-system-one/internal/middleware"
	"payment-system-one/internal/models"
	"payment-system-one/internal/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Create a user
func (u *HTTPHandler) CreateUser(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBind(&user); err != nil {
		util.Response(c, "invalid request", 400, "bad request body", nil)
		return
	}

	//validate user email
	// check if user email is valid
	if !util.IsValidEmail(user.Email) {
		util.Response(c, "invalid email", 400, "Bad request", nil)
		return
	}

	//validate user email
	_, err := u.Repository.FindUserByEmail(user.Email)
	if err == nil {
		util.Response(c, "user does exist", 404, "user already exists", nil)
		return
	}
	// hashing passwords to encrypt
	hashPass, err := util.HashPassword((user.Password))
	if err != nil {
		util.Response(c, "Password not hashed", 404, "not hashed", nil)
		return
	}
	user.Password = hashPass
	// Creating Account number
	accNo, err := util.GenerateAccountNo()
	if err != nil {
		util.Response(c, "couldnt generate acc", 500, "server issue", nil)
		return
	}
	user.AccountNo = accNo

	//persist information in the data base
	err = u.Repository.CreateUser(user)
	if err != nil {
		util.Response(c, "user not created", 400, err.Error(), nil)
		return
	}
	util.Response(c, "user created", 200, "success", nil)
}

// login user
func (u *HTTPHandler) LoginUser(c *gin.Context) {
	var loginRequest *models.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		util.Response(c, "invalid request", 400, "bad request body", nil)
		return
	}

	if loginRequest.Email == "" || loginRequest.Password == "" {
		util.Response(c, "Please enter your email or password", 400, "bad request body", nil)
		return
	}

	// check if user already exists
	user, err := u.Repository.FindUserByEmail(loginRequest.Email)
	if err != nil {
		util.Response(c, "user does not exist", 404, "user not found", nil)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		util.Response(c, "invalid password", 400, "invalid request", nil)
	}

	/*if user.Password != loginRequest.Password {

		util.Response(c, "password mismatch", 404, "user not found", nil)
		return
	}
	*/
	//Generate token
	accessClaims, refreshClaims := middleware.GenerateClaims(user.Email)

	secret := os.Getenv("JWT_SECRET")

	accessToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)
	if err != nil {
		util.Response(c, "error generating access token", 500, "error generating access token", nil)
		return
	}
	refreshToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, refreshClaims, &secret)
	if err != nil {
		util.Response(c, "error generating refresh token", 500, "error generating refresh token", nil)
		return
	}
	c.Header("access_token", *accessToken)
	c.Header("refresh_token", *refreshToken)

	util.Response(c, "login successful", http.StatusOK, gin.H{
		"user":          user,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil)
}

// call a protected route
func (u *HTTPHandler) GetUserByEmail(c *gin.Context) {
	_, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "User not logged in", 500, "user not found", nil)
		return
	}

	email := c.Query("email")

	if email == "" {
		util.Response(c, "email is required", 400, "email is required", nil)
		return
	}

	user, err := u.Repository.FindUserByEmail(email)
	if err != nil {
		util.Response(c, "user not fount", 500, "user not found", nil)
		return
	}

	util.Response(c, "user found", 200, user, nil)
}

// transfer Funds

func (u *HTTPHandler) TransferFunds(c *gin.Context) {
	// declare request
	var transferMoney *models.TransferMoney

	// bind data to struct json data to the struct
	if err := c.ShouldBind(&transferMoney); err != nil {
		util.Response(c, "Invalid Transfer", 400, "bad Request", nil)
		return
	}

	// get user from context(make sure the person logged is the owner of the ACC)
	user, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "user not fount", 500, "user not found", nil)
		return
	}

	// validate the amount
	if transferMoney.Amount <= 0 {
		util.Response(c, "invalid amount, must be more than 0 ", 400, "Bad Request", nil)
		return
	}

	// check if the account number exists
	/// db method of finding user by acount, and actual user that exist
	// find user by account method in the repository user.go
	recipient, err := u.Repository.FindUserByAccNo(transferMoney.RecipiencACC)
	if err != nil {
		util.Response(c, "user not fount", 404, "Recipient account not found", nil)
		return
	}

	// make sure the amount they are transfering is less than the curent blance
	if transferMoney.Amount >= user.AccountBalance {
		util.Response(c, "insufficient funds", 400, "bad Request", nil)
		return
	}

	// persist the data into the DB, take the money away from payer and give it to the recipient
	err = u.Repository.Transferfunds(user, recipient, transferMoney.Amount)
	if err != nil {
		util.Response(c, "transfer failed", 500, "transer failed", nil)
		return
	}
	util.Response(c, "transfer successful", 200, "transfer successful", nil)

}

// ADD mmoney

func (u *HTTPHandler) AddFunds(c *gin.Context) {

	// declare request
	var addFunds models.AddMoney

	// bind data to struct data struct
	if err := c.ShouldBind(&addFunds); err != nil {
		util.Response(c, "Invalid Transer", 400, "bad Request", nil)
	}

	// get user from context(make sure the person logged is the owner of the ACC)
	user, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "user not fount", 500, "user not found", nil)
		return
	}

	// validate amount make sure its more than zero
	if addFunds.Amount <= 0 {
		util.Response(c, "insufficient funds", 404, "bad Request", nil)
		return
	}

	// add the money to the account and update the acc
	// repository code and persist the data into the database

	err = u.Repository.ADDfunds(user, addFunds.Amount)
	if err != nil {
		util.Response(c, "transfer failed", 500, "transer failed", nil)
		return
	}
	util.Response(c, "Adding funds successful", 200, "successful", nil)

}

// Current Balance
func (u *HTTPHandler) BalanceCheck(c *gin.Context) {

	// get user from context
	user, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "user not fount", 500, "user not found", nil)
		return
	}
	// checking balance
	util.Response(c, "Balance retrieved successfully", 200, "sucess", nil)
	c.IndentedJSON(200, gin.H{"balance": user.AccountBalance})

}

// Transaction history

func (u *HTTPHandler) TransactionHistory(c *gin.Context) {

	var transactionHistory models.Transaction
	// get user from context
	user, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "user not fount", 500, "user not found", nil)
		return
	}
	// ADd/ create a  function to the repository  to retrive the transacton history

	//

}
