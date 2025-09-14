package controllers

import (
	"net/http"

	"os"
	"reflect"
	"strings"
	"time"

	"github.com/binoymanoj/jwt-auth-go/initializers"
	"github.com/binoymanoj/jwt-auth-go/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type SignUpRequest struct {
	FirstName       string `json:"first_name" binding:"required,min=3"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// helper function for Validation
func ValidationErrorMessages(err error, obj interface{}) map[string]string {
	ve, ok := err.(validator.ValidationErrors)
	if !ok {
		return map[string]string{"error": err.Error()}
	}

	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	out := make(map[string]string)
	for _, fe := range ve {
		field, _ := t.FieldByName(fe.Field())
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = strings.ToLower(fe.Field()) // fallback
		} else {
			jsonTag = strings.Split(jsonTag, ",")[0] // remove options like `omitempty`
		}

		tag := fe.Tag()
		param := fe.Param()
		var msg string

		switch tag {
		case "required":
			msg = jsonTag + " is required"
		case "email":
			msg = jsonTag + " must be a valid email address"
		case "min":
			msg = jsonTag + " must be at least " + param + " characters long"
		case "eqfield":
			msg = jsonTag + " must be equal to " + param
		default:
			msg = jsonTag + " is invalid"
		}

		out[jsonTag] = msg
	}
	return out
}

func SignUp(c *gin.Context) {
	var req SignUpRequest

	// Validation
	if err := c.ShouldBindJSON(&req); err != nil {
		errs := ValidationErrorMessages(err, req)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to create user",
			"errors":  errs,
		})
		return
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to hash password",
		})
		return
	}

	// Create User
	user := models.User{Email: req.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Failed to create user",
			// "message": result.Error,  // for logging
		})
		return
	}

	// Response
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Signup Successful",
	})
}

func Login(c *gin.Context) {
	var req LoginRequest

	// Validation
	if err := c.ShouldBindJSON(&req); err != nil {
		errs := ValidationErrorMessages(err, req)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to login",
			"errors":  errs,
		})
		return
	}

	if c.Bind(&req) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Failed to read body",
		})

		return
	}

	// Lookup in DB
	var user models.User
	initializers.DB.First(&user, "email = ?", req.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid Email or Password",
		})

		return
	}

	// Compare passwords
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid Email or Password",
		})

		return
	}

	// Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get encoded token
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Failed to create token",
		})

		return
	}

	// Send as reponse
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Cookie Passed",
		"token":   tokenString,
	})
}

func Validate(c *gin.Context) {
	userInterface, _ := c.Get("user")

	user := userInterface.(models.User)

	// response to pass in API
	response := gin.H{
		"id":        user.ID,
		"name":      user.Name,
		"email":     user.Email,
		"createdAt": user.CreatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Logged In",
		"user":    response, // passing only the filtered data
		// "user":    userInterface, 		// passing all data from user model (this includes hashed password field)
	})
}
