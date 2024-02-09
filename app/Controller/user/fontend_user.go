package user

import (
	"github.com/gofiber/fiber/v2"
	"server.com/app/models"
	"server.com/app/service"
	"server.com/pkg"
	"server.com/pkg/pagination"
	"server.com/pkg/utils"
)

// geoIPRequest := new(models.GeoIPRequest)
// geoIPRequest.PrivateKey = os.Getenv("PRIVATE_KEY")
// geoIPRequest.IP = clientIP
// res, ok := openAPI.RouteGEO_IP_API(os.Getenv("GEO_IP_END_POINT"), *geoIPRequest)
//
//	if ok != nil {
//		return models.FailWithMessage(models.ERROR, ok.Error(), c)
//	}
type userParams struct {
	UserID int `json:"userId"`
}

func GetUsers(c *fiber.Ctx) error {
	page, limit := pagination.GetPaginationParams(c)
	users, _ := service.GetUserPaginated(page, limit)
	return models.OkWithDetailed(users, "操作成功", c)
}

func GetUser(c *fiber.Ctx) error {
	json := new(userParams)
	if err := c.BodyParser(json); err != nil {
		return models.FailWithMessage(-400, "Invalid Json", c)
	}

	user, err := service.GetUser(json.UserID)

	if err != nil {
		return models.FailWithMessage(-100000, err.Error(), c)
	}

	return models.OkWithDetailed(user, "Get successful", c)
}

func CreateUser(c *fiber.Ctx) error {

	json := new(models.User)
	if ok := c.BodyParser(json); ok != nil {
		return models.FailWithMessage(-400, "Invalid Json", c)
	}

	hash, _ := service.HashPassword(json.Password)
	json.Password = hash

	ok := service.CreateUser(*json)
	if ok != nil {
		return models.FailWithMessage(-50000, "internal server error", c)
	}

	return models.OkWithMessage("create new user is successfully", c)
}

func UpdateUser(c *fiber.Ctx) error {
	json := new(models.User)
	if err := c.BodyParser(json); err != nil {
		return models.FailWithMessage(-400, "Invalid Json", c)
	}

	ok := service.UpdateUser(*json)
	if ok != nil {
		return models.FailWithMessage(-50000, "Internal Server Error", c)
	}

	return models.OkWithMessage("update user successfully", c)
}

func DeleteUser(c *fiber.Ctx) error {
	json := new(userParams)
	if ok := c.BodyParser(json); ok != nil {
		return models.FailWithMessage(-400, "Invalid Json", c)
	}

	user, ok := service.GetUser(json.UserID)

	if ok != nil {
		return models.FailWithMessage(-10000, ok.Error(), c)
	}

	ok = service.DeleteUser(int64(user.ID))
	if ok != nil {
		return models.FailWithMessage(-50000, ok.Error(), c)
	}
	return models.FailWithMessage(0, "Delete successfully", c)
}

func SendCodeMobileVerify(c *fiber.Ctx) error {
	type bodySendCode struct {
		AreaCode string `json:"area_code"`
		Mobile   string `json:"mobile"`
	}
	body := new(bodySendCode)

	if ok := c.BodyParser(body); ok != nil {
		return models.FailWithMessage(models.INVALID_JSON, ok.Error(), c)
	}

	clientIP := utils.GetHeaderPeekByKey(c, utils.X_Real_IP)

	telecom, ok := utils.GetTelecomCompany(clientIP)
	if ok != nil {
		return models.FailWithMessage(models.ERROR, "invalid ip info", c)
	}

	sendCodeObj := new(models.VerifyCode)
	sendCodeObj.IP = clientIP
	sendCodeObj.IPAddress = telecom.Country + " " + telecom.City
	sendCodeObj.Code = pkg.GenerateOTP(6)
	sendCodeObj.Content = "Your verify code is " + sendCodeObj.Code
	sendCodeObj.TelecomName = telecom.Hostname
	sendCodeObj.CreateTime = utils.GetCurrentTimeStamp()
	sendCodeObj.ModifyTime = utils.GetCurrentTimeStamp()

	ok = service.SubmitVerifyCode(*sendCodeObj)

	if ok != nil {
		return models.FailWithMessage(models.INTERNAL_SERVER_ERROR, ok.Error(), c)
	}

	return models.OkWithMessage("Successful", c)
}

func UserRegister(c *fiber.Ctx) error {
	body := new(models.User)
	switch body.Type {
	case 1:
		if body.Mobile == "" {
			return models.FailWithMessage(models.ERROR_INPUT, "Please input the mobile", c)
		}
	case 2:
		if body.Email == "" {
			return models.FailWithMessage(models.ERROR_INPUT, "Please enter the email", c)
		}
	case 3:

	}
	return models.OkWithMessage("Register is successful", c)
}
