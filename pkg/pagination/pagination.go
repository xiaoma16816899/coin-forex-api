package pagination

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"server.com/pkg/types"
)

func GetPaginateMetaFromParamsWithoutDefault(c *fiber.Ctx) types.ResultOrError[types.PaginationMeta] {
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	meta := types.PaginationMeta{}
	meta.Page = 1
	meta.Limit = 10

	page, err := strconv.Atoi(pageStr)
	if err == nil {
		meta.Page = uint64(page)
	} else {
		return types.ResultOrError[types.PaginationMeta]{Data: nil, Error: errors.New("no page param")}
	}

	limit, err := strconv.Atoi(limitStr)
	if err == nil {
		meta.Limit = uint64(limit)
	}

	return types.ResultOrError[types.PaginationMeta]{Data: &meta, Error: nil}
}

func GetPaginateMetaFromParams(c *fiber.Ctx) types.Result[types.PaginationMeta] {
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	meta := types.PaginationMeta{}
	meta.Page = 1
	meta.Limit = 10

	page, err := strconv.Atoi(pageStr)
	if err == nil {
		meta.Page = uint64(page)
	}

	limit, err := strconv.Atoi(limitStr)
	if err == nil {
		meta.Limit = uint64(limit)
	}

	return types.Result[types.PaginationMeta]{Data: meta, Error: nil}
}

// get pagination params
func GetPaginationParams(c *fiber.Ctx) (int, int) {
	page, limit := 0, 10
	payload := struct {
		Page int `json:"page"`
	}{}
	if err := c.BodyParser(&payload); err != nil {
		page = 0
	} else {
		page = payload.Page
	}
	payload1 := struct {
		Limit int `json:"limit"`
	}{}

	if err := c.BodyParser(&payload1); err != nil {
		limit = 10
	} else {
		if payload1.Limit > 0 {
			limit = payload1.Limit
		}
	}
	return page, limit
}
