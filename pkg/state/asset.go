package state

import (
	"groovin.dev/inv-api/pkg/models"
)

// Get all assets
func GetAssets(s *StateManager) []models.Asset {
	// Return all assets
	return s.GetAssets()
}

// Get asset by ID
func GetAsset(s *StateManager, id int) models.Asset {
	// Get the assets from the state manager
	assets := s.GetAssets()

	// Loop through assets
	for _, asset := range assets {
		// If asset ID matches
		if asset.ID == id {
			// Return asset
			return asset
		}
	}

	// Return empty asset
	return models.Asset{}
}

// Create asset
func CreateAsset(s *StateManager, asset models.Asset) models.Asset {
	// Get the assets from the state manager
	assets := s.GetAssets()

	// Add asset to array
	assets = append(assets, asset)

	// Return asset
	return asset
}

// Update asset
func UpdateAsset(s *StateManager, id int, asset models.Asset) models.Asset {
	// Get the assets from the state manager
	assets := s.GetAssets()

	// Loop through assets
	for i, a := range assets {
		// If asset ID matches
		if a.ID == id {
			// Update asset
			assets[i] = asset

			// Return asset
			return asset
		}
	}

	// Return empty asset
	return models.Asset{}
}

// Delete asset
func DeleteAsset(s *StateManager, id int) models.Asset {
	// Get the assets from the state manager
	assets := s.GetAssets()

	// Loop through assets
	for i, a := range assets {
		// If asset ID matches
		if a.ID == id {
			// Delete asset
			assets = append(assets[:i], assets[i+1:]...)

			// Return asset
			return a
		}
	}

	// Return empty asset
	return models.Asset{}
}
