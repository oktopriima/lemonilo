/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/10/2020, 20:52
 * Copyright (c) 2020
 */

package helper

import (
	"errors"
	"github.com/oktopriima/lemonilo/domain/middleware"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	HourDays = 24
	WeekDays = 7
	AllValue = "*"
	StrWeeks = "weeks"
	StrDays  = "days"
)

func ResponseToByte(response *http.Response) ([]byte, error) {
	var bodyBytes []byte
	bodyBytes, err := ioutil.ReadAll(response.Body)
	return bodyBytes, err
}

func RandString(n int) string {
	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

func GetAuthenticatedUser(r *http.Request) (int64, error) {
	userID, err := middleware.ExtractToken(r, "user_id")
	if err != nil {
		return 0, err
	}
	return int64(userID.(float64)), nil
}

func GetLimitOffset(page, size int) (limit int, offset int) {
	if page == 0 || size == 0 {
		// using -1 to disable gorm size and offset in case page and size not set
		size = -1
		offset = -1
		return size, offset
	}
	offset = (page - 1) * size
	return size, offset
}

func GetTransferCode() float64 {
	rand.Seed(time.Now().UnixNano())
	val := rand.Intn(400-1) + 1
	return float64(val)
}

type UnitTime struct {
	Value string `json:"value"`
}

func GetUnitTime(interval int, unit string) (*UnitTime, error) {

	res := new(UnitTime)
	if unit == "" {
		return nil, errors.New("missing time unit " + unit)
	}

	switch StrDays {
	case "days":
		res.Value = strconv.Itoa(interval*HourDays) + "h"
		return res, nil
	case StrWeeks:
		week := interval * WeekDays
		res.Value = strconv.Itoa(week*HourDays) + "h"
	default:
		return nil, errors.New("missing time unit " + unit)
	}

	return nil, errors.New("missing time unit " + unit)
}
