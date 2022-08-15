package models

type Order struct {
	ID int `json:"id" gorm:"primary_key:auto_increment"`
	// Product_Order []ProductOrderResponse `json:"product_order"`
	// Title  string `json:"title" form:"title" gorm:"type: varchar(255)"`
	// Price  int    `json:"price" form:"price" gorm:"type: int"`
	// Image  string `json:"image" form:"image" gorm:"type: varchar(255)"`
	// Qty    int    `json:"-" form:"qty"`
	UserID int `json:"user_id" form:"user_id"`
	// User       UsersProfileResponse `json:"user"`
	// Category   []Category           `json:"category" gorm:"many2many:product_categories"`
	// CategoryID []int                `json:"-" form:"category_id" gorm:"-"`
	// CreatedAt time.Time `json:"-"`
	// UpdatedAt time.Time `json:"-"`
}

type OrderTransactionResponse struct {
	ID int `json:"id"`
	// Product_Order []ProductOrderResponse `json:"product_order"`
	// Category   []Category           `json:"category" gorm:"many2many:product_categories"`
	// CategoryID []int                `json:"category_id" form:"category_id" gorm:"-"`
}

func (OrderTransactionResponse) TableName() string {
	return "products"
}
