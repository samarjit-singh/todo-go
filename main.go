package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	Completed bool `json:"completed"` // by default it will be false
	Body string `json:"body"`
}

var collection *mongo.Collection

func main () {
	app := fiber.New()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file:",err)
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client,err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())
	
	err = client.Ping(context.Background(),nil)
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MONGODB ATLAS")


	PORT := os.Getenv("PORT")

	collection = client.Database("golang_db").Collection(("todos"))

	// todos := []Todo{} // array of Todo structure

	// folloing : is the pointer and address concept
	// var x int = 5;
	// var p *int = &x;

	// fmt.Println("prints the address", p)
	// fmt.Println("prints the actual value", *p)
	
	app.Get("/api/todos", getTodos)
	app.Post("/api/todos", createTodos)
	app.Patch("/api/todos/:id", updateTodos)

	// Create todo
	// app.Post("/api/todos", func(c *fiber.Ctx) error {
	// 	todo := &Todo{}

	// 	if err := c.BodyParser(todo); err != nil {
	// 		return err
	// 	}

		// if todo.Body == "" {
		// 	return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		// }

	// 	todo.ID = len(todos) + 1
	// 	todos = append(todos, *todo)

	// 	return c.Status(201).JSON(todo)
	// })

	// Update a todo
	// app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
	// 	id := c.Params("id")

	// 	for i, todo := range todos {
	// 		if fmt.Sprint(todo.ID) == id {
	// 			todos[i].Completed = true
	// 			return c.Status(200).JSON(todos[i])
	// 		}
	// 	}

	// 	return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	// })

	// Delete a Todo
	// app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
	// 	id := c.Params("id")

	// 	for i,todo := range todos {
	// 		if fmt.Sprint(todo.ID) == id {
	// 			todos = append(todos[:1],todos[i+1:]...)
	// 			return c.Status(200).JSON(fiber.Map{"success": "Todo deleted"})
	// 		}
	// 	}

	// 	return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	// })

	log.Fatal(app.Listen(":"+PORT))
}

func getTodos(c *fiber.Ctx) error {
	var todos []Todo

	cursor, err := collection.Find(context.Background(),bson.M{})

	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo

		if err := cursor.Decode(&todo); err!=nil {
			return nil
		}

		todos = append(todos, todo)
	}

	return c.JSON(todos)
}

func createTodos(c *fiber.Ctx) error {
	todo := new(Todo)

	if err := c.BodyParser(todo); err != nil {
		return err
	}	

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
	}

	insertResult, err := collection.InsertOne(context.Background(), todo)

	if err != nil {
		return err
	}

	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(todo)
}

func updateTodos(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}

	_, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"success": true})
}


// func deleteTodos(c *fiber.Ctx) error {}