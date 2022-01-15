package routes

import (
	"errors"

	"github.com/Chaksack/centri-backend-web/database"
	"github.com/Chaksack/centri-backend-web/models"
	"github.com/gofiber/fiber/v2"
)

type Category struct {
	//this is not the model Product, see this as the serializer
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func CreateResponseCategory(categoryModel models.Category) Category {
	return Category{ID: categoryModel.ID, Name: categoryModel.Name}
}

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&category)
	responseCategory := CreateResponseCategory(category)

	return c.Status(200).JSON(responseCategory)
}

func GetCategorys(c *fiber.Ctx) error {
	categorys := []models.Category{}

	database.Database.Db.Find(&categorys)
	responseCategorys := []Category{}
	for _, category := range categorys {
		responseCategory := CreateResponseCategory(category)
		responseCategorys = append(responseCategorys, responseCategory)
	}

	return c.Status(200).JSON(responseCategorys)
}

func findCategory(id int, category *models.Category) error {
	database.Database.Db.Find(&category, "id = ?", id)
	if category.ID == 0 {
		return errors.New("category does not exist")
	}
	return nil
}

func GetCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var category models.Category

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findCategory(id, &category); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseCategory := CreateResponseCategory(category)

	return c.Status(200).JSON(responseCategory)
}

func UpdateCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var category models.Category

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findCategory(id, &category); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateCategory struct {
		Name string `json:"name"`
	}
	var updateData UpdateCategory

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	category.Name = updateData.Name

	database.Database.Db.Save(&category)

	responseCategory := CreateResponseCategory(category)
	return c.Status(200).JSON(responseCategory)
}

func DeleteCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var category models.Category

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findCategory(id, &category); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&category).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Deleted Category")
}
