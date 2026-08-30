package mcpserver

import "testing"

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
				MaxPrice: 5000,
			},
			wantCount: 2,
		},
		{
			name: "category filters products",
			input: SearchProductsInput{
				Category: "running",
			},
			wantCount: 3,
		},
		{
			name: "combined filters",
			input: SearchProductsInput{
				Query:    "swift",
				MaxPrice: 5000,
				Category: "running",
			},
			wantCount: 1,
		},
		{
			name: "no matches",
			input: SearchProductsInput{
				Query: "laptop",
			},
			wantCount: 0,
		},
		{
			name: "empty input returns all products",
			input: SearchProductsInput{},
			wantCount: 3,
		},
		{
			name: "query is case insensitive",
			input: SearchProductsInput{
				Query: "SWIFTRUN",
			},
			wantCount: 1,
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
			wantCount: 3,
		},
		{
			name: "price equal to max price is included",
			input: SearchProductsInput{
				MaxPrice: 3499,
			},
			wantCount: 1,
		},
		{
			name: "price above max price is excluded",
			input: SearchProductsInput{
				MaxPrice: 3498,
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