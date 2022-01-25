package state

import (
	"groovin.dev/inv-api/pkg/models"
)

// State manager class
type StateManager struct {
	Employees   []models.Employee
	Assets      []models.Asset
	BuildOrders []models.BuildOrder
}

// Get all employees
func (s StateManager) GetEmployees() []models.Employee {
	return s.Employees
}

// Set employees
func (s *StateManager) SetEmployees(employees []models.Employee) {
	s.Employees = employees
}

// Get all assets
func (s StateManager) GetAssets() []models.Asset {
	return s.Assets
}

// Set assets
func (s *StateManager) SetAssets(assets []models.Asset) {
	s.Assets = assets
}

// Get all build orders
func (s StateManager) GetBuildOrders() []models.BuildOrder {
	return s.BuildOrders
}

// Set build orders
func (s *StateManager) SetBuildOrders(buildOrders []models.BuildOrder) {
	s.BuildOrders = buildOrders
}

// Init the state manager
func (s *StateManager) Init() {
	// Employee array
	// 10 employees in different departments
	var employees = []models.Employee{
		{ID: 1, Name: "John", Dept: "IT"},
		{ID: 2, Name: "Jane", Dept: "IT"},
		{ID: 3, Name: "Joe", Dept: "Sales"},
		{ID: 4, Name: "Jack", Dept: "Sales"},
		{ID: 5, Name: "Jill", Dept: "Sales"},
		{ID: 6, Name: "Jenny", Dept: "Management"},
		{ID: 7, Name: "Juan", Dept: "Management"},
		{ID: 8, Name: "Jenny", Dept: "Management"},
		{ID: 9, Name: "Juan", Dept: "Management"},
		{ID: 10, Name: "Jenny", Dept: "CEO"},
	}

	// Asset array
	// 10 assets in different types
	var assets = []models.Asset{
		{ID: 1, Name: "Laptop", Type: "Computer", AssignedTo: 1},
		{ID: 2, Name: "Monitor", Type: "Monitor", AssignedTo: 1},
		{ID: 3, Name: "Printer", Type: "Printer", AssignedTo: 1},
		{ID: 4, Name: "Scanner", Type: "Printer", AssignedTo: 1},
		{ID: 5, Name: "Phone", Type: "Phone", AssignedTo: 1},
		{ID: 6, Name: "Tablet", Type: "Phone", AssignedTo: 1},
		{ID: 7, Name: "Laptop", Type: "Computer", AssignedTo: 2},
		{ID: 8, Name: "Monitor", Type: "Monitor", AssignedTo: 2},
		{ID: 9, Name: "Printer", Type: "Printer", AssignedTo: 2},
		{ID: 10, Name: "Scanner", Type: "Printer", AssignedTo: 2},
	}

	// BuildOrder array
	// 10 build orders
	var buildOrders = []models.BuildOrder{
		// Build order for John
		// 1 Laptop, 1 Monitor, 1 Printer, 1 Scanner
		{ID: 1, Name: "Build Laptop", Description: "Build a laptop", CreatedFor: 1, Assets: []models.Asset{
			{ID: 1, Name: "Laptop", Type: "Computer", AssignedTo: 1},
			{ID: 2, Name: "Monitor", Type: "Monitor", AssignedTo: 1},
			{ID: 3, Name: "Printer", Type: "Printer", AssignedTo: 1},
			{ID: 4, Name: "Scanner", Type: "Printer", AssignedTo: 1},
		}},
	}

	// Set employees
	s.Employees = employees

	// Set assets
	s.Assets = assets

	// Set build orders
	s.BuildOrders = buildOrders
}
