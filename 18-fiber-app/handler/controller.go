package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jabhishek87/fiberapp/model"
)

type handler struct {
	DB map[string]model.Course
}

func RegisterRoutes(app *fiber.App, db map[string]model.Course) {
	globalData := &handler{
		DB: db,
	}
	routes := app.Group("/v1")

	routes.Get("/records", globalData.GetRecords)
	routes.Get("/record/:id", globalData.GetRecord)
	routes.Post("/record", globalData.AddRecord)
	routes.Put("/record/:id", globalData.UpdateRecord)
	routes.Delete("/record/:id", globalData.DeleteRecord)
}

func (globalData handler) GetRecords(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&globalData.DB)
}

func (globalData handler) GetRecord(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	rec, exist := globalData.DB[id]
	if !exist {
		return ctx.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	}
	return ctx.Status(fiber.StatusOK).JSON(&rec)
}

func (globalData handler) AddRecord(ctx *fiber.Ctx) error {
	body := model.Course{}

	// parse body, attach to AddBookRequestBody struct
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	body.CourseId = model.GetNewID()
	globalData.DB[body.CourseId] = body
	// fmt.Printf("%+v", globalData.DB)
	return ctx.Status(fiber.StatusCreated).JSON(&body)
}

func (globalData handler) UpdateRecord(ctx *fiber.Ctx) error {
	body := model.Course{}
	id := ctx.Params("id")

	_, exist := globalData.DB[id]
	if !exist {
		return ctx.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	}

	// parse body, attach to AddBookRequestBody struct
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	body.CourseId = id
	//globalData.DB[body.CourseId] = body
	// fmt.Printf("%+v", globalData.DB)
	// fmt.Printf("%+v", body)
	return ctx.Status(fiber.StatusAccepted).JSON(&body)
}

func (globalData handler) DeleteRecord(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	rec, exist := globalData.DB[id]
	if !exist {
		return ctx.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	}
	delete(globalData.DB, id)
	return ctx.Status(fiber.StatusAccepted).JSON(&rec)
}
