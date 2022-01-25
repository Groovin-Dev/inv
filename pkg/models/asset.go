package models

// Asset struct
type Asset struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	AssignedTo int    `json:"assignedTo"`
}
