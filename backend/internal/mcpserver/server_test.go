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
