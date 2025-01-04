package mw

import (
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

const RoleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		req := ctx.Request().Header.Get("User-Role")
		if strings.EqualFold(req, RoleAdmin) {
			log.Println("admin-user detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
