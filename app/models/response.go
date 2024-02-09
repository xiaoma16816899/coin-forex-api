package models

import (
	"github.com/gofiber/fiber/v2"
	"server.com/pkg/utils"
)

type ResponseData struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type ResponseMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// Define Code status
const (
	SUCCESS               = 0
	INVALID_JSON          = -1
	ERROR                 = -2
	INVALID_ROLE          = -3
	NOT_FOUND             = -4
	CONFLICT_STATUS       = -5
	ERROR_INPUT           = -6
	BAD_REQUEST           = -5000
	OWNER_NOT_FOUND       = -10000
	INVALID_TOKEN         = -20000
	INVALID_IP_ADDRESS    = -30000
	INTERNAL_SERVER_ERROR = -50000
)

func Result(code int, data interface{}, msg string, c *fiber.Ctx) error {
	// 开始时间
	// c.JSON(http.StatusOK,Response{ code, data, msg })
	var response = ResponseData{
		Code: code,
		Data: data,
		Msg:  msg,
	}
	c.Request().Header.Add("Backed-Log", "yes")
	return c.JSON(response)
}

func ResultWithMsg(code int, msg string, c *fiber.Ctx) error {
	// 开始时间
	// c.JSON(http.StatusOK,Response{ code, data, msg })
	var response = ResponseMsg{
		Code: code,
		Msg:  msg,
	}
	c.Request().Header.Add("Backed-Log", "yes")
	c.Request().Header.Add("Api-Code", utils.IntToStr(code))
	c.Request().Header.Add("Api-Msg", msg)
	return c.JSON(response)
}

func ErrorResult(code int, msg string, c *fiber.Ctx) error {
	// 开始时间
	var response = ResponseData{
		Code: code,
		Msg:  msg,
	}
	c.Request().Header.Add("Backed-Log", "yes")
	c.Request().Header.Add("Api-Code", utils.IntToStr(code))
	c.Request().Header.Add("Api-Msg", msg)
	return c.JSON(response)
}

func Ok(c *fiber.Ctx) {
	c.Request().Header.Add("Backed-Log", "yes")
	c.Request().Header.Add("Api-Code", utils.IntToStr(SUCCESS))
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *fiber.Ctx) error {
	return ResultWithMsg(SUCCESS, message, c)
}

func OkWithData(data interface{}, c *fiber.Ctx) error {
	return Result(SUCCESS, data, "查询成功", c)
}

func OkWithDetailed(data interface{}, message string, c *fiber.Ctx) error {
	return Result(SUCCESS, data, message, c)
}

func Fail(c *fiber.Ctx) error {
	return Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(code int, message string, c *fiber.Ctx) error {
	return ErrorResult(code, message, c)
}

// forboded
func Forbidden(code int, message string, c *fiber.Ctx) error {
	c.Status(416)
	c.Request().Header.Add("Backed-Log", "yes")
	return ErrorResult(code, message, c)
}

func FailWithDetailed(data interface{}, message string, c *fiber.Ctx) error {
	return Result(ERROR, data, message, c)
}

// set failed response status code
func SetFailedStatus(c *fiber.Ctx) {
	c.Status(400)
}

// set success response status code
func SetSuccessStatus(c *fiber.Ctx) {
	c.Status(200)
}
