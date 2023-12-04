package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var u CreateUserReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := h.Service.CreateUser(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) Login(c *gin.Context) {
	var u LoginUserReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	r, err := h.Service.Login(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("jwt", r.accessToken, 3600, "/", "localhost", false, true)
	res := &LoginUserRes{
		// accessToken: r.accessToken,
		Username: r.Username,
		ID: r.ID,
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message":"logout successful!"})
}


func validateToken(tokenString string) (*JWTClaims, error) {
	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
			}
			return []byte(secretKey), nil
	})

	if err != nil {
			return nil, err
	}

	if !token.Valid {
			return nil, errors.New("invalid token")
	}

	return claims, nil
}

func (h *Handler) Profile(c *gin.Context) {
	// Retrieve the JWT token from the cookie
	token, err := c.Cookie("jwt")
	if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "you must be logged in to get your profile"})
			return
	}

	claims, err := validateToken(token)
	if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
	}

	// Assuming you want to use the ID from the token claims for fetching the profile
	res, err := h.Service.Profile(c.Request.Context(), claims.ID)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) ChangeProfile(c *gin.Context) {
	token, err := c.Cookie("jwt")
	if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "you must be logged in to change your profile"})
			return
	}

	claim, err := validateToken(token)
	if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
	}

	// Parse the request data
	var req CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	// Update the profile using the service layer
	err = h.Service.ChangeProfile(c.Request.Context(), claim.ID, &req)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}

	c.JSON(http.StatusOK, gin.H{"message": "profile updated successfully"})
}