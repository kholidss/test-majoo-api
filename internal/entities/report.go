package entities

import "time"

type Transaction struct {
	Id        int       `json:"id"`
	Merchant  Merchant  `json:"merchant" gorm:"embedded"`
	Outlet    Outlet    `json:"outlet" gorm:"embedded"`
	BillTotal int       `json:"bill_total"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

type Transactions []Transaction

type Merchant struct {
	Id           int       `json:"id"`
	UserId       int       `json:"user_id"`
	MerchantName string    `json:"merchant_name"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    int       `json:"created_by"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    int       `json:"updated_by"`
}

type Outlet struct {
	Id         int       `json:"id"`
	MerchantId int       `json:"user_id"`
	OutletName string    `json:"outlet_name"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  int       `json:"created_by"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  int       `json:"updated_by"`
}
