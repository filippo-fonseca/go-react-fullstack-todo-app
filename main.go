package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID int `json:"id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

func main() {
	fmt.Println("Hello, sssssWorsld!")
	app := fiber.New()

	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	//create a todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{} // { id: 0, completed: false, body: "" }
		
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{ "msg": "Todo body is required" })
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)

	})

	//update a todo
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{ "msg": "Todo not found" })

	})

	//delete a todo
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{ "success": true })
			}
		}

		return c.Status(404).JSON(fiber.Map{ "msg": "Todo not found" })
	})



	log.Fatal(app.Listen(":4000"))
}