// ```go
package mcpserver

import "testing"

func float64Ptr(value float64) *float64 {
	return &value
}

func TestFilterProducts(t *testing.T) {
	tests := []struct {
		name      string
		input     SearchProductsInput
		wantCount int
	}{
		{
			name: "query filters products",
			input: SearchProductsInput{
				Query: "trail",
			},
			wantCount: 1,
		},
		{
			name: "max price filters products",
			input: SearchProductsInput{
				MaxPrice: float64Ptr(5000),
			},
			wantCount: 3,
		},
		{
			name: "category filters products",
			input: SearchProductsInput{
				Category: "running",
			},
			wantCount: 4,
		},
		{
			name: "combined filters",
			input: SearchProductsInput{
				Query:    "swift",
				MaxPrice: float64Ptr(5000),
				Category: "running",
			},
			wantCount: 2,
		},
		{
			name: "no matches",
			input: SearchProductsInput{
				Query: "laptop",
			},
			wantCount: 0,
		},
		{
			name:      "empty input returns all products",
			input:     SearchProductsInput{},
			wantCount: 4,
		},
		{
			name: "query is case insensitive",
			input: SearchProductsInput{
				Query: "SWIFTRUN",
			},
			wantCount: 2,
		},
		{
			name: "query ignores surrounding whitespace",
			input: SearchProductsInput{
				Query: "  trail  ",
			},
			wantCount: 1,
		},
		{
			name: "category is case insensitive",
			input: SearchProductsInput{
				Category: "RUNNING",
			},
			wantCount: 4,
		},
		{
			name: "price equal to max price is included",
			input: SearchProductsInput{
				MaxPrice: float64Ptr(3499),
			},
			wantCount: 1,
		},
		{
			name: "price above max price is excluded",
			input: SearchProductsInput{
				MaxPrice: float64Ptr(3498),
			},
			wantCount: 0,
		},
		{
			name: "zero max price returns no products",
			input: SearchProductsInput{
				MaxPrice: float64Ptr(0),
			},
			wantCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := filterProducts(tt.input)

			if len(got) != tt.wantCount {
				t.Fatalf(
					"expected %d products, got %d",
					tt.wantCount,
					len(got),
				)
			}
		})
	}
}

func TestGetProduct(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		expectedFound   bool
		expectedProduct Product
	}{
		{
			name:          "product exists",
			input:         "shoe-001",
			expectedFound: true,
			expectedProduct: Product{
				ID:       "shoe-001",
				Name:     "SwiftRun Pro",
				Category: "running",
				Price:    3499,
				Currency: "INR",
			},
		},
		{
			name:            "product doesn't exist",
			input:           "shoe-111",
			expectedFound:   false,
			expectedProduct: Product{},
		},
		{
			name:            "Empty input",
			input:           "",
			expectedFound:   false,
			expectedProduct: Product{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, found := getProduct(tt.input)

			if tt.expectedFound != found {
				t.Fatalf(
					"expected %v, but got %v",
					tt.expectedFound,
					found,
				)
			}

			if found && product != tt.expectedProduct {
				t.Fatalf(
					"expected product %+v, but got %+v",
					tt.expectedProduct,
					product,
				)
			}
		})
	}
}

