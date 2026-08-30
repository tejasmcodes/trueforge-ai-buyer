package mcpserver

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type SearchProductsInput struct {
	Query    string  `json:"query" jsonschema:"the product the user is looking for"`
	MaxPrice float64 `json:"max_price,omitempty" jsonschema:"maximum acceptable price"`
	Category string  `json:"category,omitempty" jsonschema:"product category"`
}

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}

type SearchProductsOutput struct {
	Products []Product `json:"products"`
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

	return server
}
