package core

import (
	"golang-crud-rest-api/entities"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
)

// Reference https://github.com/alexsergivan/blog-examples/tree/master/authentication
const (
	accessTokenCookieName  = "access-token"
	refreshTokenCookieName = "refresh-token"
	tokenTimeExp           = 1 * time.Millisecond
	refreshTokenTimeExp    = 24 * time.Hour
)

func GetJWTSecret() string {
	return AppConfig.Secret
}

func GetRefreshJWTSecret() string {
	return AppConfig.Secret
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	ID       int16  `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func JwtMiddleware() {
	// jwtMiddleware := middleware.
}

func GenerateAccessToken(user *entities.User) (string, time.Time, error) {
	// Declare the expiration time of the token
	expirationTime := time.Now().Add(tokenTimeExp)

	return generateToken(user, expirationTime, []byte(GetJWTSecret()))
}

func GenerateRefreshToken(user *entities.User) (string, time.Time, error) {
	// Declare the expiration time of the token
	expirationTime := time.Now().Add(refreshTokenTimeExp)

	return generateToken(user, expirationTime, []byte(GetRefreshJWTSecret()))
}

func generateToken(user *entities.User, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

// TokenRefresherMiddleware middleware, which refreshes JWT tokens if the access token is about to expire.
func TokenRefresherMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// If the user is not authenticated (no user token data in the context), don't do anything.
		if c.Get("user") == nil {
			return next(c)
		}
		// Gets user token from the context.
		u := c.Get("user").(*jwt.Token)

		claims := u.Claims.(*Claims)

		// We ensure that a new token is not issued until enough time has elapsed
		// In this case, a new token will only be issued if the old token is within
		// 15 mins of expiry.
		if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 15*time.Minute {
			// Gets the refresh token from the cookie.
			rc, err := c.Cookie(refreshTokenCookieName)
			if err == nil && rc != nil {
				// Parses token and checks if it valid.
				tkn, err := jwt.ParseWithClaims(rc.Value, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(GetRefreshJWTSecret()), nil
				})
				if err != nil {
					if err == jwt.ErrSignatureInvalid {
						c.Response().Writer.WriteHeader(http.StatusUnauthorized)
					}
				}

				if tkn != nil && tkn.Valid {
					// If everything is good, update tokens.
					// _ = GenerateTokensAndSetCookies(&entities.User{
					// 	FullName: claims.Name,
					// }, c)
				}
			}
		}

		return next(c)
	}
}
