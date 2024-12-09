package services

import (
	"github.com/Earth/discount-campaigns/models"
	"math"
)

func CalculateTotal(items []models.Item) float64 {
	total := 0.0
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}

func ApplyCoupon(total float64, couponType string, parameter int) float64 {
	if couponType == "Fixed" {
		total = total - float64(parameter)
	} else if couponType == "Percentage" {
		total = total - (float64(parameter) * total / 100)
	}
	return total
}

func ApplyOnTop(items []models.Item, total float64, onTopType string, parameter models.OnTopParameters) float64 {
	if onTopType == "PercenByCate" {
		discount := 0.0
		for _, item := range items {
			if parameter.Category == "clothing" {
				if item.Category == "clothing" {
					discount += item.Price * float64(item.Quantity) * float64(parameter.Discount)/100
				}
			} else if parameter.Category == "accessories" {
				if item.Category == "accessories" {
					discount += item.Price * float64(item.Quantity) * float64(parameter.Discount)/100
				}
			} else if parameter.Category == "electronics" {
				if item.Category == "electronics" {
					discount += item.Price * float64(item.Quantity) * float64(parameter.Discount)/100
				}
			}
		}
		return total - discount
	} else if onTopType == "ByPoints" {
		points := float64(parameter.Discount)
		maxDiscount := total * float64(parameter.Max) / 100
		return total - math.Min(points, maxDiscount)
	}
	return total
}

func ApplySeasonal(total float64, seasonal models.SeasonalParameters) float64 {
	count := int(total) / seasonal.Every
	totalDiscount := float64(count) * float64(seasonal.Discount)
	total = total - totalDiscount
	return total
}
