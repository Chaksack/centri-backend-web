package routes

import (
	"errors"
	"time"

	"github.com/Chaksack/centri-backend-web/database"
	"github.com/Chaksack/centri-backend-web/models"
	"github.com/gofiber/fiber/v2"
)

type Invoice struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Category  Category  `json:"category"`
	Product   Product   `json:"product"`
	CreatedAt time.Time `json:"invoice_date"`
}

func CreateResponseInvoice(invoice models.Invoice, user User, category Category, product Product) Invoice {
	return Invoice{ID: invoice.ID, User: user, Category: category, Product: product, CreatedAt: invoice.CreatedAt}
}

func CreateInvoice(c *fiber.Ctx) error {
	var invoice models.Invoice

	if err := c.BodyParser(&invoice); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := findUser(invoice.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var category models.Category
	if err := findCategory(invoice.CategoryRefer, &category); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := findProduct(invoice.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&invoice)

	responseUser := CreateResponseUser(user)
	responseCategory := CreateResponseCategory(category)
	responseProduct := CreateResponseProduct(product)
	respoonseInvoice := CreateResponseInvoice(invoice, responseUser, responseCategory, responseProduct)

	return c.Status(200).JSON(respoonseInvoice)

}

func GetInvoices(c *fiber.Ctx) error {
	invoices := []models.Invoice{}
	database.Database.Db.Find(&invoices)
	responseInvoices := []Invoice{}

	for _, invoice := range invoices {
		var user models.User
		var category models.Category
		var product models.Product
		database.Database.Db.Find(&user, "id = ?", invoice.UserRefer)
		database.Database.Db.Find(&category, "id = ?", invoice.CategoryRefer)
		database.Database.Db.Find(&product, "id = ?", invoice.ProductRefer)
		responseInvoice := CreateResponseInvoice(invoice, CreateResponseUser(user), CreateResponseCategory(category), CreateResponseProduct(product))
		responseInvoices = append(responseInvoices, responseInvoice)
	}

	return c.Status(200).JSON(responseInvoices)
}

func FindInvoice(id int, invoice *models.Invoice) error {
	database.Database.Db.Find(&invoice, "id = ?", id)
	if invoice.ID == 0 {
		return errors.New("invoice does not exist")
	}
	return nil
}

func GetInvoice(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var invoice models.Invoice

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindInvoice(id, &invoice); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	var category models.Category
	var product models.Product

	database.Database.Db.First(&user, invoice.UserRefer)
	database.Database.Db.First(&category, invoice.UserRefer)
	database.Database.Db.First(&product, invoice.UserRefer)
	responseUser := CreateResponseUser(user)
	responseCategory := CreateResponseCategory(category)
	responseProduct := CreateResponseProduct(product)

	responseInvoice := CreateResponseInvoice(invoice, responseUser, responseCategory, responseProduct)

	return c.Status(200).JSON(responseInvoice)
}
