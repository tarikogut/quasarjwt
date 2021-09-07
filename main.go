package main

import (
	"estpguiapi/config"
	"estpguiapi/database"
	"estpguiapi/router"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	database.ConnectDB()
	database.InitTables()
	router.SetupRoutes(app)

	applistenIP := fmt.Sprintf(":%v", config.Config("SERVERPORT"))
	log.Fatal(app.Listen(applistenIP))

}
//func requestEstp() {
//	a := fiber.AcquireAgent()
//	a.UserAgent("ESTPGUI-API v1.0")
//	req := a.Request()
//	req.Header.SetMethod(fiber.MethodPost)
//	req.SetRequestURI("http://estp-dev.fink-telecom.com:8082/api/mtp3-list")
//
//	args := fasthttp.AcquireArgs()
//	args.Set("username", "tarik")
//	args.Set("password", "tdqNrEuB")
//	args.Set("details", "1")
//	a.Form(args)
//	a.Parse()
//	code, body, errs := a.Bytes()
//	fmt.Println(code, string(body), errs)
//}
