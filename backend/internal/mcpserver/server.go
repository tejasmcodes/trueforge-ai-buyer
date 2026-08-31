package mcpserver

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type SearchProductsInput struct {
	Query    string   `json:"query" jsonschema:"the product the user is looking for"`
	MaxPrice *float64 `json:"max_price,omitempty" jsonschema:"maximum acceptable price"`
	Category string   `json:"category,omitempty" jsonschema:"product category"`
}

type SearchProductsOutput struct {
	Products []Product `json:"products"`
}

type GetProductInput struct {
	ProductID string `json:"product_id" jsonschema:"the ID of the product to retrieve"`
}

type GetProductOutput struct {
	Product Product `json:"product"`
}

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}

type Inventory struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type CheckInventoryInput struct {
	ProductID string `json:"product_id" jsonschema:"availability of the product in inventory"`
}

type CheckInventoryOutput struct {
	Quantity int `json:"quantity"`
}

func searchProducts(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input SearchProductsInput,
) (*mcp.CallToolResult, SearchProductsOutput, error) {

	return nil, SearchProductsOutput{
		Products: filterProducts(input),
	}, nil
}

func getProductByID(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input GetProductInput,
) (*mcp.CallToolResult, GetProductOutput, error) {
	product, found := getProduct(input.ProductID)
	if !found {
		return nil, GetProductOutput{}, fmt.Errorf("product %q not found", input.ProductID)
	}
	return nil, GetProductOutput{
		Product: product,
	}, nil
}

func checkInventory(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input CheckInventoryInput,
) (*mcp.CallToolResult, CheckInventoryOutput, error) {
	inventory, found := getInventory(input.ProductID)

	if !found {
		return nil, CheckInventoryOutput{}, fmt.Errorf("product %q not found", input.ProductID)
	}
	return nil, CheckInventoryOutput{
		Quantity: inventory.Quantity,
	}, nil

}

func NewServer() *mcp.Server {
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    "trueforge-buyer",
			Version: "0.1.0",
		},
		nil,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "search_products",
			Description: "Search the buyer's product catalog for products matching the user's requirements.",
		},
		searchProducts,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_product",
			Description: "Retrieve a product from the buyer's catalog by its product ID.",
		},
		getProductByID,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_inventory",
			Description: "Retrieve a product quantity from the inventory by its product ID.",
		},
		checkInventory,
	)

	return server
}
