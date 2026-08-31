package mcpserver

import (
	"context"
	"reflect"
	"testing"
)

func TestSearchProducts(t *testing.T) {
	tests := []struct {
		name             string
		input            SearchProductsInput
		expectedProducts []Product
	}{
		{
			name: "search returns matching products",
			input: SearchProductsInput{
				Query: "SwiftRun",
			},
			expectedProducts: []Product{
				{
					ID:       "shoe-001",
					Name:     "SwiftRun Pro",
					Category: "running",
					Price:    3499,
					Currency: "INR",
				},
				{
					ID:       "shoe-004",
					Name:     "SwiftRun ProMax",
					Category: "running",
					Price:    4499,
					Currency: "INR",
				},
			},
		},
		{
			name: "search returns no matches",
			input: SearchProductsInput{
				Query: "lAmBoRgHiNi",
			},
			expectedProducts: []Product{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, productsOutput, err := searchProducts(context.Background(), nil, tt.input)

			if err != nil {
				t.Fatalf("expected no error, but got: %v", err)
			}

			if result != nil {
				t.Fatalf("expected nil as the result, but got: %v", result)
			}

			if !reflect.DeepEqual(productsOutput.Products, tt.expectedProducts) {
				t.Fatalf("expected products %+v, but got %+v",
					tt.expectedProducts,
					productsOutput.Products,
				)
			}
		})
	}
}

func TestGetProductByID(t *testing.T) {
	tests := []struct {
		name            string
		input           GetProductInput
		expectedProduct Product
		expectedError   string
	}{
		{
			name: "product exists",
			input: GetProductInput{
				ProductID: "shoe-001",
			},
			expectedProduct: Product{
				ID:       "shoe-001",
				Name:     "SwiftRun Pro",
				Category: "running",
				Price:    3499,
				Currency: "INR",
			},
			expectedError: "",
		},
		{
			name: "product doesn't exist",
			input: GetProductInput{
				ProductID: "shoe-111",
			},
			expectedProduct: Product{},
			expectedError:   `product "shoe-111" not found`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, productOutput, err := getProductByID(context.Background(), nil, tt.input)
			if result != nil {
				t.Fatalf("expected nil as the result, but got: %v", result)
			}

			if tt.expectedError == "" {
				if err != nil {
					t.Fatalf("expected no error, but got: %v", err)
				}
			} else {
				if err == nil {
					t.Fatalf("expected error %q, but got nil", tt.expectedError)
				}

				if err.Error() != tt.expectedError {
					t.Fatalf(
						"expected error %q, but got %q",
						tt.expectedError,
						err.Error(),
					)
				}
			}
			if productOutput.Product != tt.expectedProduct {
				t.Fatalf("expected product: %+v, but got: %+v",
					tt.expectedProduct,
					productOutput.Product,
				)
			}
		})
	}
}

func TestCheckInventory(t *testing.T) {
	tests := []struct {
		name             string
		input            CheckInventoryInput
		expectedQuantity int
		expectedError    string
	}{
		{
			name: "product is in stock",
			input: CheckInventoryInput{
				ProductID: "shoe-001",
			},
			expectedQuantity: 5,
			expectedError:    "",
		},
		{
			name: "product is out of stock",
			input: CheckInventoryInput{
				ProductID: "shoe-004",
			},
			expectedQuantity: 0,
			expectedError:    "",
		},
		{
			name: "product doesn't exist",
			input: CheckInventoryInput{
				ProductID: "shoe-111",
			},
			expectedQuantity: 0,
			expectedError:    `product "shoe-111" not found`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, inventoryOutput, err := checkInventory(context.Background(), nil, tt.input)

			if result != nil {
				t.Fatalf("expected nil as the result, but got %v", result)
			}

			if tt.expectedError == "" {
				if err != nil {
					t.Fatalf("expected no error, but got: %v", err)
				}
			} else {
				if err == nil {
					t.Fatalf("expected error %q, but got nil", tt.expectedError)
				}

				if err.Error() != tt.expectedError {
					t.Fatalf(
						"expected error %q, but got %q",
						tt.expectedError,
						err.Error(),
					)
				}
			}

			if inventoryOutput.Quantity != tt.expectedQuantity {
				t.Fatalf("expected %d, but got %d",
					tt.expectedQuantity,
					inventoryOutput.Quantity,
				)
			}
		})
	}
}
