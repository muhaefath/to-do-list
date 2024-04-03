package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhaefath/to-do-list/internals/models"
	"github.com/muhaefath/to-do-list/internals/repository"
	"github.com/muhaefath/to-do-list/internals/repository/product"
	"github.com/muhaefath/to-do-list/internals/request"
)

type MarketplaceHandler struct {
	productRepo product.ProductRepository
	orderRepo   repository.OrderRepository
}

func NewMarketplaceHandler(productRepo product.ProductRepository, orderRepo repository.OrderRepository) *MarketplaceHandler {
	return &MarketplaceHandler{productRepo: productRepo, orderRepo: orderRepo}
}

func (h *MarketplaceHandler) GetAllProduct(c *fiber.Ctx) error {
	products, err := h.productRepo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve todos: " + err.Error())
	}

	return c.JSON(products)
}

func (h *MarketplaceHandler) GetAllOrder(c *fiber.Ctx) error {
	orders, err := h.orderRepo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve todos: " + err.Error())
	}

	return c.JSON(orders)
}

func (h *MarketplaceHandler) GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")

	todo, err := h.productRepo.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve todos: " + err.Error())
	}

	return c.JSON(todo)
}

func (h *MarketplaceHandler) CreateOrder(c *fiber.Ctx) error {
	var createOrderBody request.CreateOrder
	if err := c.BodyParser(&createOrderBody); err != nil {
		return err
	}

	product, err := h.productRepo.GetByID(createOrderBody.ProductID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve todos: " + err.Error())
	}

	productQty := product.Quantity - createOrderBody.Quantity
	if productQty < 0 {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to create order because qty less than product qty")
	}

	order := models.Order{
		ProductID: createOrderBody.ProductID,
		Quantity:  createOrderBody.Quantity,
	}
	if err := h.orderRepo.CreateTodo(&order); err != nil {
		return err
	}

	product.Quantity = productQty
	if err := h.productRepo.UpdateQuantityByID(createOrderBody.ProductID, product); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}

func (h *MarketplaceHandler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return err
	}

	if err := h.productRepo.UpdateQuantityByID(id, &product); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *MarketplaceHandler) CancelOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.orderRepo.DeleteTodoByID(id); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
