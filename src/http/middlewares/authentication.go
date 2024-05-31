package middlewares

import (
	"belimang/src/entities"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func RequireAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, entities.ErrorResponse{
					Status:  false,
					Message: "Unauthorized",
				})
			}

			// parts := strings.SplitN(authHeader, " ", 2)
			// if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			// 	return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization format")
			// }
			// if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			// 	c.JSON(http.StatusUnauthorized, echo.Map{"error": "missing or malformed Authorization header"})
			// 	return nil
			// }

			// Extract the token from the header
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			// Parse the JWT token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Provide the key for verifying the token's signature
				// (replace with your actual key)
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil || !token.Valid {
				c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
				return nil
			}
			// If the token is valid, proceed with the next middleware/handler

			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				structClaims, _ := MapClaimsToCustomClaims(claims)
				c.Set("jwtClaims", structClaims)
			}
			return next(c)
		}
	}
	// Get the Authorization header value

	// Check if the header is empty or doesn't start with "Bearer "
}

func AuthWithRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := c.Get("jwtClaims").(*entities.CustomClaims)
			// structClaims, _ := MapClaimsToCustomClaims(claims)
			if claims.Role != role {
				return c.JSON(http.StatusUnauthorized, entities.ErrorResponse{
					Status:  false,
					Message: "Unauthorized",
				})
			}
			return next(c)
		}
		// parts := strings.SplitN(authHeader, " ", 2)
		// if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		// 	return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization format")
		// }
		// if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		// 	c.JSON(http.StatusUnauthorized, echo.Map{"error": "missing or malformed Authorization header"})
		// 	return nil
		// }

	}
}

func MapClaimsToCustomClaims(mapClaims jwt.MapClaims) (*entities.CustomClaims, error) {
	userId, ok := mapClaims["userId"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid userId")
	}

	username, ok := mapClaims["username"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid username")
	}

	role, ok := mapClaims["role"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid role")
	}

	expiresAt, ok := mapClaims["exp"].(float64) // JWT timestamps are usually float64
	if !ok {
		return nil, fmt.Errorf("invalid exp")
	}

	return &entities.CustomClaims{
		UserId:   userId,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(expiresAt),
		},
	}, nil
}
