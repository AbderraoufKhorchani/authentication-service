package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary Register a new user and return a JWT token
// @Description Register a new user with the provided credentials and return a JWT token upon successful registering.
// @Accept json
// @Produce json
// @Tags Registering
// @Param body body SignupRequest true "User registration details in JSON format"
// @Success 200 {string} string "User registered successfully"
// @Failure 400 {string} string "Bad Request - Invalid request payload"
// @Failure 500 {string} string "Internal Server Error"
// @Router /signup [post]
func Register(c *gin.Context) {

	userJSON := SignupRequest{}

	if err := c.ShouldBindJSON(&userJSON); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Request - Invalid request payload")
		return
	}

	user := User{

		UserName:  userJSON.UserName,
		FirstName: userJSON.FirstName,
		LastName:  userJSON.LastName,
		Password:  userJSON.Password,
	}

	token, err := GenerateToken(user.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = Insert(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Respond with the generated token
	c.JSON(http.StatusOK, token)
}

// Login godoc
// @Summary Authenticates a user and return a JWT token
// @Description Authenticates a user based on the provided credentials (username and password) and return a JWT token upon successful authentication.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body LoginRequest true "Login request payload"
// @Success 200 {string} string "User successfully authenticated"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 401 {string} string "Invalid credentials"
// @Failure 500 {string} string "Internal Server Error"
// @Router /login [post]
func Login(c *gin.Context) {

	loginRequest := LoginRequest{}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Retrieve the user by username
	user, err := GetByUserName(loginRequest.UserName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Compare the provided password with the stored hashed password
	err = ComparePassword(user.Password, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Generate a JWT token
	token, err := GenerateToken(user.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to generate token")
		return
	}

	user.Password = ""

	// Respond with the generated token
	c.JSON(http.StatusOK, token)

}

// @Summary Reset user's password
// @Description Reset the password for a user.
// @Tags Reset password
// @Accept json
// @Produce json
// @Param requestPayload body ResetPasswordRequest true "Request payload for password reset"
// @Success 200 {string} string "Password reset successful"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 401 {string} string "Invalid credentials"
// @Failure 500 {string} string "Internal Server Error"
// @Router /rest_password [post]
func ResetPassword(c *gin.Context) {

	requestPayload := ResetPasswordRequest{}

	if err := c.ShouldBindJSON(&requestPayload); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := GetByUserName(requestPayload.UserName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Invalid credentials")
		return
	}

	err = ComparePassword(user.Password, requestPayload.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Invalid credentials")
		return
	}

	err = ResetPasswordBD(requestPayload.NewPassword, user.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Invalid credentials")
		return
	}

	resonse := user.UserName + "'s Password resetted"

	c.JSON(http.StatusOK, resonse)

}
