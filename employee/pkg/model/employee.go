package model

type Employee struct {
	Name   string  `json:"name"`
	Job    string  `json:"job"`
	Salary float64 `json:"salary"`
	Age    uint    `json:"age"`
}
