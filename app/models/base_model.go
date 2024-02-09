package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BaseModel struct {
	ID         uint `json:"id,omitempty" gorm:"primarykey"`
	CreateTime int64  `json:"create_time"`
	ModifyTime int64  `json:"modify_time"`
}

// check json .. json, required key, number key
func CheckPayload(c *fiber.Ctx, reqKeys *[]string, keysNumber *[]string) (string, error) {
	var p map[string]interface{}
	if err := json.Unmarshal(c.Body(), &p); err != nil {
		return "unmarshal body failed", errors.New("invalid")
	}

	// check string
	buf := bytes.NewBufferString("")
	for _, k := range *reqKeys {
		value := p[k]
		// log.Println(">>>>>>>>> K", k, p[k], reflect.TypeOf(p[k]))
		if reflect.TypeOf(value) == nil || value == "" {
			buf.WriteString("[" + k + " is required]")
		}
	}

	// check number
	for _, k := range *keysNumber {
		value := p[k]
		if reflect.TypeOf(value) == nil || value == "" {
			buf.WriteString("[" + k + " is required]")
		} else {
			_, err := strconv.Atoi(fmt.Sprint(value))
			if err != nil {
				buf.WriteString("[" + k + " is not number]")
				buf.WriteString(", ")
			}
		}
	}

	if buf.String() != "" {
		return buf.String(), errors.New("invalid")
	}
	return "", nil
}
