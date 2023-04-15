package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Anv3sh/go-fiber-postgres/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct{
	Author string `json:"author"`
	Title string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct{
	DB *gorm.DB
}

func (r *Repository) CreateBook(context *fiber.Ctx) error{
	book := Book{}
	err := context.BodyParser(&book)
	 if err != nil{
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message":"request failed"})
			return err
	 }

	 err = r.DB.Create(&book).Error
	 if err!=nil{
		  context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"mesage":"could not create book"}
		  )
		  return err
	 }
	 context.Status(http.StatusOK).JSON(
		&fiber.Map{"message":"book has been created"}
	)
	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error{
		bookModels := &[]models.Book{}
		err := r.DB.Find(bookModels).Error
		if err != nil{
			context.Satus(http.StatusBadRequest).JSON(
				&fiber.Map{"message":"could not get books"}
			)
			return err
		}
		context.Status(http.StatusOK).JSON(
			&fiber.Map{"message":"books fetched successfully",
						"data":bookModels,
		})
		return nil
}

func (r *Repository) SetupRoutes(app *fiber.App){
	api := app.Group("/api")
	api.Post("/create_books",r.CreateBook)
	api.Delete("delete_book/:id",e,DeleteBook)
	api.Get("/get_books/:id",r.GetBooks)
	api.Get("/books",r.GetBooks) 
}

func main(){
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal(err)
	}

	config = &storage.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS")
		User: os.Getenv("DB_USER")
		DBName: os.Getenv("DB_NAME")
		SSLMode: os.Getenv("DB_SSLMODE")
	}

	db,err := storage.NewConnection(config.storage)
	if err != nil{
		log.Fatal("could not load the database")
	}
	r := Repository(
		DB : db,
	)
	app := Fiber.new()
	r.SetupRoutes(app)

	app.listen(":8080")
}