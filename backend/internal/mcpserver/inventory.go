package mcpserver

var inventory = []Inventory{
	{
		ProductID: "shoe-001",
		Quantity: 5,
	},
	{
		ProductID: "shoe-002",
		Quantity: 3,
	},
	{
		ProductID: "shoe-003",
		Quantity: 1,
	},
	{
		ProductID: "shoe-004",
		Quantity: 0,
	},
}



func getInventory(productID string)(Inventory, bool){
	for _, product := range inventory {
		if product.ProductID == productID{
			return product, true
		}
	}

	return Inventory{}, false

}