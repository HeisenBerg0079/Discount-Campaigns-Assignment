package handler

import (
	"github.com/Earth/discount-campaigns/models"
	"github.com/Earth/discount-campaigns/services"
	"github.com/gofiber/fiber/v2"
)

func CalculateDiscountedPrice(c *fiber.Ctx) error {
	var req models.Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	total := services.CalculateTotal(req.Items)

	total = services.ApplyCoupon(total, req.Discount.Coupon.Campaigns, req.Discount.Coupon.Parameters)
	total = services.ApplyOnTop(req.Items, total, req.Discount.OnTop.Campaigns, req.Discount.OnTop.Parameters)
	total = services.ApplySeasonal(total, req.Discount.Seasonal.Parameters)

	return c.JSON(fiber.Map{
		"final_price": total,
	})
}