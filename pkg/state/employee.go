package state

import (
	"groovin.dev/inv-api/pkg/models"
)

// Get all employees
func GetEmployees(s *StateManager) []models.Employee {
	// Return all employees
	return s.GetEmployees()
}

// Get employee by ID
func GetEmployee(s *StateManager, id int) models.Employee {
	// Get the employees from the state manager
	employees := s.GetEmployees()

	// Loop through employees
	for _, employee := range employees {
		// If employee ID matches
		if employee.ID == id {
			// Return employee
			return employee
		}
	}

	// Return empty employee
	return models.Employee{}
}

// Create employee
func CreateEmployee(s *StateManager, employee models.Employee) models.Employee {
	// Get the employees from the state manager
	employees := s.GetEmployees()

	// Add employee to array
	employees = append(employees, employee)

	// Return employee
	return employee
}

// Update employee
func UpdateEmployee(s *StateManager, id int, employee models.Employee) models.Employee {
	// Get the employees from the state manager
	employees := s.GetEmployees()

	// Loop through employees
	for i, emp := range employees {
		// If employee ID matches
		if emp.ID == id {
			// Update employee
			employees[i] = employee

			// Return employee
			return employee
		}
	}

	// Return empty employee
	return models.Employee{}
}

// Delete employee
func DeleteEmployee(s *StateManager, id int) models.Employee {
	// Get the employees from the state manager
	employees := s.GetEmployees()

	// Loop through employees
	for i, emp := range employees {
		// If employee ID matches
		if emp.ID == id {
			// Delete employee
			employees = append(employees[:i], employees[i+1:]...)

			// Return employee
			return emp
		}
	}

	// Return empty employee
	return models.Employee{}
}
