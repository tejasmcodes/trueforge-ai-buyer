package mcpserver

import "testing"

func TestGetInventory(t *testing.T) {
	tests := []struct {
		name              string
		input             string
		expectedFound     bool
		expectedInventory Inventory
	}{
		{
			name:          "product exists with stock",
			input:         "shoe-001",
			expectedFound: true,
			expectedInventory: Inventory{
				ProductID: "shoe-001",
				Quantity:  5,
			},
		},
		{
			name:          "product exists with zero stock",
			input:         "shoe-004",
			expectedFound: true,
			expectedInventory: Inventory{
				ProductID: "shoe-004",
				Quantity:  0,
			},
		},
		{
			name:              "product doesn't exist",
			input:             "shoe-111",
			expectedFound:     false,
			expectedInventory: Inventory{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productID := tt.input
			inventory, found := getInventory(productID)
			if inventory != tt.expectedInventory {
				t.Fatalf("expected %+v, but got %+v",
					tt.expectedInventory,
					inventory,
				)
			}
			if found != tt.expectedFound {
				t.Fatalf("expected %v, but got %v",
					tt.expectedFound,
					found,
				)
			}
		})
	}
}
