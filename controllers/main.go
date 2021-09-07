package controllers

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func Status(ctx *fiber.Ctx) error {

	return ctx.JSON(fiber.Map{"time": time.Now().Format("2006-02-01 15:04:05 MST"), "status": "ok"})
}
