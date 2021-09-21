package model

import ()
//Dashboard representation
type Dashboard struct {
	Sales          Module         `json:"sales,omitempty"`
	Purchases      Module         `json:"purchases,omitempty"`
	Payments       Module         `json:"payments,omitempty"`
	Receipts       Module         `json:"receipts,omitempty"`
	Expences       Module         `json:"expences,omitempty"`
	Customers      Module         `json:"customers,omitempty"`
	Suppliers      Module         `json:"suppliers,omitempty"`
	Cost           Module        `json:"cost,omitempty"`
	Profit         Module        `json:"profit,omitempty"`
	Expence        float64        `json:"expence,omitempty"`
	Dashlinechart  Dashlinechart  `json:"dashlinechart,omitempty"`
	Dashbarchart   Dashbarchart   `json:"dashbarchart,omitempty"`
	Movingproducts []Movingproducts `json:"movingproducts,omitempty"`
}
//Module structure of dashboard items
type Module struct{
	Name string
	Total float64
	Description string
	Icon string
}
//Email structure
type Email struct{
	Email string
	To string
	Subject string
	Message string
	Customers []Customer
}
type Dashlinechart struct {
	One   Linechart `json:"one,omitempty"`
	Two   Linechart `json:"two,omitempty"`
	Three Linechart `json:"three,omitempty"`
	Four  Linechart `json:"four,omitempty"`
	Five  Linechart `json:"five,omitempty"`
	Six   Linechart `json:"six,omitempty"`
	Seven Linechart `json:"seven,omitempty"`
}
type Dashbarchart struct {
	One   Linechart `json:"one,omitempty"`
	Two   Linechart `json:"two,omitempty"`
	Three Linechart `json:"three,omitempty"`
}
type Linechart struct {
	Sales     float64 `json:"sales,omitempty"`
	Purchases float64 `json:"purchases,omitempty"`
	Expences  float64 `json:"expences,omitempty"`
	Profit    float64 `json:"profit,omitempty"`
	Cost      float64 `json:"cost,omitempty"`
	Listo     string  `json:"listo,omitempty"`
}
type Movingproducts struct {
	Productcode string  `json:"productcode,omitempty"`
	Productname string  `json:"productname,omitempty"`
	Quantity    int64   `json:"quantity,omitempty"`
	TotalSales  float64 `json:"total_sales,omitempty"`
}
