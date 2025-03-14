package entity

type (
	//Product
	Product struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		CategoryID   string `json:"category_id"`
		ShortInfo    string `json:"short_info"`
		Description  string `json:"description"`
		Cost         int    `json:"cost"`
		Count        int    `json:"count"`
		DiscountCost int    `json:"discount_cost"`
		Discount     int    `json:"discount"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
	}
)

type (
	//Category
	Category struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
)

type (
	// Attribute

	Attribute struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		CategoryID string `json:"category_id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
)

type (
	//Order
	Order struct {
		ID                string `json:"id"`
		UserID            string `json:"user_id"`
		IntegrationID     string `json:"integration_id"`
		Status            string `json:"status"`
		StatusChangedTime string `json:"status_chaged_time"`
		TotalCost         int    `json:"total_cost"`
		CreatedAt         string `json:"created_at"`
		UpdatedAt         string `json:"updated_at"`
	}
)

type (
	// OrderProducts
	OrderProducts struct {
		ID        string `json:"id"`
		OrderID   string `json:"order_id"`
		ProductID string `json:"product_id"`
		Count     int    `json:"count"`
		Cost      int    `json:"cost"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
)
