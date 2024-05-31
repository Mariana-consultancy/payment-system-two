package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"payment-system-one/internal/middleware"
	"payment-system-one/internal/models"
	"payment-system-one/internal/util"
	"golang.org/x/crypto/bcrypt"
)

//Create admin function for handler

func (u *HTTPHandler) CreateAdmin(c *gin.Context) {
	var admin *models.Admin
	if err := c.ShouldBind(&admin); err != nil {
		util.Response(c, "invalid request", 400, "bad request body", nil)
		return
	}
		// check if admin email is valid
	if !util.IsValidEmail(admin.Email) {
		util.Response(c, "invalid email", 400, "Bad request", nil)
		return
	}

	//validate admin email
	_, err := u.Repository.FindAdminByEmail(admin.Email)
	if err == nil {
		util.Response(c, "admin does exist", 404, "admin already exists", nil)
		return
	}
	
		// hashPassword
	hashPass, err := util.HashPassword(admin.Password)
	if err != nil {
		util.Response(c, "Password not hashed", 404, "not hashed", nil)
		return
		}
		admin.Password = hashPass
	
	err = u.Repository.CreateAdmin(admin)
	if err != nil {
		util.Response(c, "admin not created", 400, err.Error(), nil)
		return
	}
	util.Response(c, "admin created", 200, "success", nil)
}

// admin login
func (u *HTTPHandler) LoginAdmin(c *gin.Context) {
	var adminLoginRequest *models.AdminRequest
	if err := c.ShouldBind(&adminLoginRequest); err != nil {
		util.Response(c, "invalid request", 400, "bad request body", nil)
		return
	}
	if adminLoginRequest.Email == "" || adminLoginRequest.Password == "" {
		util.Response(c, "Please enter your email or password", 400, "bad request body", nil)
		return
	}
	admin, err := u.Repository.FindAdminByEmail(adminLoginRequest.Email)
	if err != nil {
		util.Response(c, "user does not exist", 404, "user not found", nil)
		return
	}

	if admin.Password != adminLoginRequest.Password {
		if err != nil {
			util.Response(c, "internal server error", 500, "user not found", nil)
			return
		}
		util.Response(c, "password mismatch", 404, "user not found", nil)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(adminLoginRequest.Password)); err != nil {
		util.Response(c, "invalid password", 400, "invalid request", nil)
	}

	if admin.Password != adminLoginRequest.Password {

		util.Response(c, "password mismatch", 404, "user not found", nil)
		return
	}

	// Generate token for access and refresh
	accessClaims, refreshClaims := middleware.GenerateClaims(admin.Email)

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
		"admin":         admin,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil)
}
