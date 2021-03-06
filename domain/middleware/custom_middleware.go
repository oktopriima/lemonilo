/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/10/2020, 20:51
 * Copyright (c) 2020
 */

package middleware

import (
	"context"
	"errors"
	"fmt"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/lemonilo/domain/config"
	"github.com/oktopriima/lemonilo/domain/response"
	"net/http"
	"strings"
)

var jwtMiddleware *jwtmiddleware.JWTMiddleware
var signingKey []byte
var myRole map[string][]string

func NewMiddlewareConfig(cfg config.Config) (error, interface{}) {
	var err error
	app := strings.Split(cfg.GetString("permission.app"), ",")
	signature := cfg.GetString("app.signature")

	role := make(map[string][]string)
	role["app"] = app

	InitRole(role)
	InitJWTMiddlewareCustom([]byte(signature), jwt.SigningMethodHS512)

	return err, role
}

func InitRole(roles map[string][]string) {
	myRole = roles
}

func MyAuth(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := checkJWTToken(ctx.Request); err != nil {
			abortMission(ctx, err)
			return
		}

		for _, role := range roles {
			if err := checkRole(ctx.Request, role); err != nil {
				abortMission(ctx, err)
				return
			}
		}
	}
}

func InitJWTMiddlewareCustom(secret []byte, signingMethod jwt.SigningMethod) {
	signingKey = secret
	jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return signingKey, nil
		},
		SigningMethod: signingMethod,
	})
}

func ExtractToken(r *http.Request, key string) (interface{}, error) {
	tokenStr, err := jwtMiddleware.Options.Extractor(r)
	if err != nil {
		return "", err
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims[key], nil
	} else {
		return "", nil
	}
}

func abortMission(ctx *gin.Context, err error) {
	resp := new(response.Error)
	resp.ErrorCode = http.StatusUnauthorized
	resp.Message = err.Error()
	resp.Status = http.StatusUnauthorized

	ctx.AbortWithStatusJSON(http.StatusUnauthorized, &resp)
	return
}

func checkJWTToken(r *http.Request) error {
	if !jwtMiddleware.Options.EnableAuthOnOptions {
		if r.Method == "OPTIONS" {
			return nil
		}
	}

	token, err := jwtMiddleware.Options.Extractor(r)
	if err != nil {
		eExtractor := errors.New("400")
		return eExtractor
	}

	if token == "" {
		if jwtMiddleware.Options.CredentialsOptional {
			return nil
		}
		eReqiredToken := errors.New("required authorization token not found")
		return eReqiredToken
	}

	parsedToken, err := jwt.Parse(token, jwtMiddleware.Options.ValidationKeyGetter)
	if err != nil {
		ePassingToken := errors.New("Error parsing token: " + err.Error())
		return ePassingToken
	}

	if jwtMiddleware.Options.SigningMethod != nil && jwtMiddleware.Options.SigningMethod.Alg() != parsedToken.Header["alg"] {
		errorMsg := fmt.Sprintf("Expected %s signing method but token specified %s",
			jwtMiddleware.Options.SigningMethod.Alg(),
			parsedToken.Header["alg"])
		eTokenSpecified := errors.New(errorMsg)
		return eTokenSpecified
	}

	if !parsedToken.Valid {
		eInvalidToken := errors.New("token invalid")
		return eInvalidToken
	}

	newRequest := r.WithContext(context.WithValue(r.Context(), jwtMiddleware.Options.UserProperty, parsedToken))
	*r = *newRequest
	return nil
}

func checkRole(r *http.Request, roles string) (err error) {
	tokenRole, err := ExtractToken(r, "role")
	if err != nil || tokenRole == nil {
		err = errors.New("you don't have permission to access this route")
		return err
	}

	if roles == "*" {
		return nil
	}

	for k, r := range myRole {
		if k == roles {
			for _, c := range r {
				if c == tokenRole {
					return nil
				}
			}
			break
		}
	}

	err = errors.New("access denied")
	return err
}

func GetAuthenticatedUser(r *http.Request) (int64, error) {
	userID, err := ExtractToken(r, "user_id")
	if err != nil {
		return 0, err
	}
	return int64(userID.(float64)), nil
}
