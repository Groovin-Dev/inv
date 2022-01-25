package state

import (
	"groovin.dev/inv-api/pkg/models"
)

// Get all build orders
func GetBuildOrders(s *StateManager) []models.BuildOrder {
	// Return all build orders
	return s.GetBuildOrders()
}

// Get build order by ID
func GetBuildOrder(s *StateManager, id int) models.BuildOrder {
	// Get the build orders from the state manager
	buildOrders := s.GetBuildOrders()

	// Loop through build orders
	for _, buildOrder := range buildOrders {
		// If build order ID matches
		if buildOrder.ID == id {
			// Return build order
			return buildOrder
		}
	}

	// Return empty build order
	return models.BuildOrder{}
}

// Create build order
func CreateBuildOrder(s *StateManager, buildOrder models.BuildOrder) models.BuildOrder {
	// Get the build orders from the state manager
	buildOrders := s.GetBuildOrders()

	// Add build order to array
	buildOrders = append(buildOrders, buildOrder)

	// Return build order
	return buildOrder
}

// Update build order
func UpdateBuildOrder(s *StateManager, id int, buildOrder models.BuildOrder) models.BuildOrder {
	// Get the build orders from the state manager
	buildOrders := s.GetBuildOrders()

	// Loop through build orders
	for i, bo := range buildOrders {
		// If build order ID matches
		if bo.ID == id {
			// Update build order
			buildOrders[i] = buildOrder

			// Return build order
			return buildOrder
		}
	}

	// Return empty build order
	return models.BuildOrder{}
}

// Delete build order
func DeleteBuildOrder(s *StateManager, id int) models.BuildOrder {
	// Get the build orders from the state manager
	buildOrders := s.GetBuildOrders()

	// Loop through build orders
	for i, bo := range buildOrders {
		// If build order ID matches
		if bo.ID == id {
			// Delete build order
			buildOrders = append(buildOrders[:i], buildOrders[i+1:]...)

			// Return build order
			return bo
		}
	}

	// Return empty build order
	return models.BuildOrder{}
}
