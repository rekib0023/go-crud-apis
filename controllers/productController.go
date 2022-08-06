package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/rekib0023/go-crud-apis/database"
	"github.com/rekib0023/go-crud-apis/models"
	"github.com/rekib0023/go-crud-apis/serializers"
)

func ProductResponse(product models.Product) serializers.ProductSerializer {
	return serializers.ProductSerializer{ID: product.ID, Name: product.Name, SerialNumber: product.SerialNumber, CreatedAt: product.CreatedAt, UpdatedAt: product.UpdatedAt}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(err.Error())
	}

	database.Database.Db.Create(&product)

	responseProduct := ProductResponse(product)

	return c.Status(fiber.StatusCreated).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	database.Database.Db.Find(&products)
	responseProducts := []serializers.ProductSerializer{}

	for _, product := range products {
		responseProduct := ProductResponse(product)
		responseProducts = append(responseProducts, responseProduct)
	}

	return c.Status(fiber.StatusOK).JSON(responseProducts)
}

func findProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("product does not exist")
	}
	return nil
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON("Please ensure that :id is an integer")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(err.Error())
	}

	responseProduct := ProductResponse(product)

	return c.Status(fiber.StatusOK).JSON(responseProduct)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON("Please ensure that :id is an integer")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(err.Error())
	}

	type UpdateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	var updateProduct UpdateProduct

	if err := c.BodyParser(&updateProduct); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(err.Error())
	}

	product.Name = updateProduct.Name
	product.SerialNumber = updateProduct.SerialNumber

	database.Database.Db.Save(&product)
	responseProduct := ProductResponse(product)

	return c.Status(fiber.StatusOK).JSON(responseProduct)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON("Please ensure that :id is an integer")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(err.Error())
	}

	return c.Status(fiber.StatusNoContent).JSON("Successfully Deleted Product")
}
