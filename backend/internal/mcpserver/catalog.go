package mcpserver

import (
	"strings"
)

var products = []Product{
	{
		ID:       "shoe-001",
		Name:     "SwiftRun Pro",
		Category: "running",
		Price:    3499,
		Currency: "INR",
	},
	{
		ID:       "shoe-002",
		Name:     "RoadRunner Lite",
		Category: "running",
		Price:    4299,
		Currency: "INR",
	},
	{
		ID:       "shoe-003",
		Name:     "TrailMax 2",
		Category: "running",
		Price:    5799,
		Currency: "INR",
	},
	{
		ID:       "shoe-004",
		Name:     "SwiftRun ProMax",
		Category: "running",
		Price:    4499,
		Currency: "INR",
	},
}

func filterProducts(input SearchProductsInput) []Product {
	query := strings.ToLower(strings.TrimSpace(input.Query))
	category := strings.ToLower(strings.TrimSpace(input.Category))

	result := make([]Product, 0)

	for _, product := range products {
		if query != "" {
			nameMatch := strings.Contains(strings.ToLower(product.Name), query)
			if !nameMatch {
				continue
			}
		}

		if category != "" &&
			!strings.EqualFold(product.Category, category) {
			continue
		}

		if input.MaxPrice != nil &&
			product.Price > *input.MaxPrice {
			continue
		}

		result = append(result, product)

	}

	return result
}

func getProduct(productID string) (Product, bool) {
	for _, product := range products {
		if product.ID == productID {
			return product, true
		}
	}

	return Product{}, false
}
