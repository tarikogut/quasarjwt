package controllers

import (
	"estpguiapi/database"
	"estpguiapi/models"
	"github.com/gofiber/fiber/v2"
)

//func EstpControllers(c *fiber.Ctx) error {
//	id := c.Params("id")
//	service := c.Params("service")
//	action := c.Params("action")
//	var estp_servers = models.EstpInstances{}
//	database.DB.Model(models.EstpInstances{}).Where("id", id).First(&estp_servers)
//	if estp_servers.Hostname == "" {
//		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No estp servers found with ID", "data": nil})
//	}
//
//	url := ""
//	if estp_servers.Https {
//		url = fmt.Sprintf("https://%v:%v/api/%v?%v", estp_servers.Hostname, estp_servers.Port, service, action)
//	} else {
//		url = fmt.Sprintf("http://%v:%v/api/%v?%v", estp_servers.Hostname, estp_servers.Port, service, action)
//	}
//	fmt.Println(url)
//	a := fiber.AcquireAgent()
//	a.UserAgent("ESTPGUI-API v1.0")
//	req := a.Request()
//	req.Header.SetMethod(fiber.MethodPost)
//	req.SetRequestURI(url)
//
//	args := fasthttp.AcquireArgs()
//	args.Set("username", "tarik")
//	args.Set("password", "tdqNrEuB")
//
//	a.Form(args)
//	a.Parse()
//	_, body, errs := a.Bytes()
//	if errs!=nil{
//		return c.Status(403).JSON(fiber.Map{"status": "error", "message": "Connection Error", "data": nil})
//	}
//	return c.SendString(string(body))
//}
func EstpAdd(c *fiber.Ctx) error {
	estp_servers := new(models.EstpInstances)

	if err := c.BodyParser(estp_servers); err != nil {
		return err
	}
	database.DB.Model(estp_servers).Create(&estp_servers)
	return c.JSON(fiber.Map{"status": "success", "data": estp_servers, "message": "success"})

}
func EstpUpdate(c *fiber.Ctx) error {
	return nil
}
func EstpDelete(c *fiber.Ctx) error {
	return nil
}
func EstpList(c *fiber.Ctx) error {
	var estp_servers = []models.EstpInstances{}
	database.DB.Model(&models.EstpInstances{}).Find(&estp_servers)
	return c.JSON(fiber.Map{"status": "success", "data": estp_servers, "message": "success"})
}
func EstpDetails(c *fiber.Ctx) error {
	return nil
}
