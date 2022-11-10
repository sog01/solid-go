package model

type Employee struct {
	Id      int     `json:"id"`
	Name    string  `json:"name,omitempty"`
	Address string  `json:"address,omitempty"`
	Salary  float64 `json:"salary,omitempty"`
}
