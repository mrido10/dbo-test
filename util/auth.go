package util

import (
	"database/sql"
	"dbo-test/config"
	"dbo-test/model/dto"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
)

type Auth struct {
	dto.DataIN
}

func Authorization(c *gin.Context) (jwt.MapClaims, error) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("SIGNING METHOD INVALID")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("SIGNING METHOD INVALID")
		}

		return []byte(config.Configure.Jwt.Key), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func (e Auth) JWTValidations(c *gin.Context, serve func(*gin.Context, dto.DataIN) (interface{}, error)) {
	claims, err := Authorization(c)
	if err != nil {
		Response(c, 401, "Unauthorized", nil)
		return
	}

	e.UserID = strconv.FormatFloat(claims["id"].(float64), 'f', 0, 64)
	e.UserName = claims["name"].(string)
	e.setResponse(c, serve)
}

func (e Auth) IgnoreValidateJWT(c *gin.Context, serve func(*gin.Context, dto.DataIN) (interface{}, error)) {
	e.setResponse(c, serve)
}

func (e Auth) setResponse(c *gin.Context, serve func(*gin.Context, dto.DataIN) (interface{}, error)) {
	respCode := 200
	respMsg := "success"
	var respData interface{}
	var tx *sql.Tx
	db, err := ConnectPostgres()
	defer func() {
		if errs := recover(); errs != nil {
			respCode = 500
			err = errors.New(fmt.Sprintf("%s", errs))
		}
		if err != nil {
			_ = tx.Rollback()
			respMsg = err.Error()
			if strings.Contains(err.Error(), config.Configure.Source.Name) {
				respCode = 400
			} else {
				respCode = 500
				respMsg = "internal server error"
			}
			log.Println(err)
		} else {
			_ = tx.Commit()
		}
		Response(c, respCode, respMsg, respData)
	}()
	if err != nil {
		return
	}
	tx, err = db.Begin()
	if err != nil {
		return
	}
	e.DataIN.Tx = tx
	respData, err = serve(c, e.DataIN)
}
