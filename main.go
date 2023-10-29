package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Division string `json:"division"`
	Salary   int64  `json:"salary"`
}

var obj []Employee

func getEmployee(c *fiber.Ctx) error {
	if len(obj) < 1 {
		obj = append(obj, Employee{ID: 1, Name: "Muhammad", Division: "CAKE", Salary: 12345})
		obj = append(obj, Employee{ID: 2, Name: "Jeff", Division: "CAKE", Salary: 15161})
		obj = append(obj, Employee{ID: 3, Name: "Doug", Division: "Data Zoo", Salary: 15498})
		obj = append(obj, Employee{ID: 4, Name: "Rofek", Division: "CAKE", Salary: 495595})
		obj = append(obj, Employee{ID: 5, Name: "Melia", Division: "Data Zoo", Salary: 8461})
	}

	return c.Status(fiber.StatusOK).JSON(obj)
}

func getEmployeeById(c *fiber.Ctx) error {
	paramsId := c.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		log.Fatal(err)
	}

	var res Employee
	for _, x := range obj {
		if x.ID == id {
			res = x
			break
		}
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func createEmployee(c *fiber.Ctx) error {
	body := new(Employee)
	err := c.BodyParser(body)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}

	res := Employee{
		ID:       len(obj) + 1,
		Name:     body.Name,
		Division: body.Division,
		Salary:   body.Salary,
	}
	obj = append(obj, res)
	return c.Status(fiber.StatusOK).JSON(obj)
}

func updateEmployee(c *fiber.Ctx) error {
	body := new(Employee)
	err := c.BodyParser(body)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}

	paramsId := c.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}

	for i, x := range obj {
		if x.ID == id {
			x.Name = body.Name
			x.Division = body.Division
			x.Salary = body.Salary
			obj[i] = x
		}
	}
	return c.Status(fiber.StatusOK).JSON(obj)
}

func deleteEmployee(c *fiber.Ctx) error {
	paramsId := c.Params("id")
	id, _ := strconv.Atoi(paramsId)

	var res []Employee
	for _, x := range obj {
		if x.ID != id {
			res = append(res, x)
		}
	}
	obj = res
	return c.Status(fiber.StatusOK).JSON(res)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to fiber")
	})

	// some necessary imports
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/emp", getEmployee)
	app.Get("/emp/:id", getEmployeeById)
	app.Post("/emp", createEmployee)
	app.Put("/emp/:id", updateEmployee)
	app.Delete("/emp/:id", deleteEmployee)

	log.Fatal(app.Listen(":8081"))
}
