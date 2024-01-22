package googleChat

import (
	"errors"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/nleeper/goment"
	logger "github.com/xiaoma/trading/pkg/utils"
)

type Payload struct {
	Title   string
	Message string
}

type GoogleChat struct {
	Url string
}

func GetGoogleChat() *GoogleChat {
	url := os.Getenv("GOOGLE_CHAT_WEBHOOK")
	return &GoogleChat{
		Url: url,
	}
}

func (gc GoogleChat) SendRaw(text string) error {
	if gc.Url == "" {
		return errors.New("empty google chat webhook url")
	}

	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod(fiber.MethodPost)
	req.SetRequestURI(gc.Url)
	a.JSON(fiber.Map{"text": text})

	if err := a.Parse(); err != nil {
		logger.Info(err)
	}

	if code, bytes, _ := a.Bytes(); code != 200 {
		logger.Error("SendRaw GoogleChat:", gc.Url, " ", string(bytes))
	}

	return nil
}

func (gc GoogleChat) Send(p Payload) error {
	appMode := "dev"
	if mode := os.Getenv("APP_ENV"); mode != "" {
		appMode = mode
	}
	now, _ := goment.New()
	text := "==================================\n"
	text += p.Title + "\n_" + now.Format("YYYY-MM-DD hh:mm:ss z") + "_\n"
	text += `[` + appMode + `]` + "\n"
	text += "```\n"
	text += p.Message + "\n"
	text += "```\n"
	return gc.SendRaw(text)
}
