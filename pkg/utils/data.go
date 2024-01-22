package utils

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// get begin and end from json payload
func GetBeginAndEndTime(c *fiber.Ctx) (time.Time, time.Time, error) {
	payload := struct {
		BeginAt time.Time `json:"startDate"`
		EndAt   time.Time `json:"endDate"`
	}{}
	err := c.BodyParser(&payload)

	return payload.BeginAt, payload.EndAt, err
}

// get begin and end from json payload
func GetBeginAndEndDateString(c *fiber.Ctx) (string, string, error) {
	payload := struct {
		BeginAt string `json:"startDate"`
		EndAt   string `json:"endDate"`
	}{}
	if err := c.BodyParser(&payload); err != nil {
		return "", "", err
	}

	if payload.BeginAt == "" || payload.EndAt == "" {
		return "", "", &fiber.Error{Message: "invalid date"}
	}

	return payload.BeginAt, payload.EndAt, nil
}

// uint to string
func UintToStr(value uint) string {
	return strconv.FormatInt(int64(value), 10)
}

// uint to string
func IntToStr(value int) string {
	return strconv.FormatInt(int64(value), 10)
}

// uint to string
func StrToInt(value string) int {
	vInt, _ := strconv.Atoi(value)
	return vInt
}
