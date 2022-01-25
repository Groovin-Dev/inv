package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"groovin.dev/inv-api/pkg/models"
	"groovin.dev/inv-api/pkg/state"
)

// Define the home route
func Home(w http.ResponseWriter, r *http.Request) {
	// Print the home page
	fmt.Fprintf(w, "Welcome to the home page!")

	// Print to the server that we are done
	fmt.Println("Home page served")
}

// RegisterRouter is a function that registers all the routes
func RegisterRouter(s *state.StateManager) *mux.Router {

	// Create a new mux router
	router := mux.NewRouter()

	// Handle GET /
	router.HandleFunc("/", Home)

	// Handle employee routes
	GetEmployees(s, router)
	GetEmployee(s, router)
	CreateEmployee(s, router)
	UpdateEmployee(s, router)
	DeleteEmployee(s, router)

	// Handle asset routes
	GetAssets(s, router)
	GetAsset(s, router)
	CreateAsset(s, router)
	UpdateAsset(s, router)
	DeleteAsset(s, router)

	// Handle build order routes
	GetBuildOrders(s, router)
	GetBuildOrder(s, router)
	CreateBuildOrder(s, router)
	UpdateBuildOrder(s, router)
	DeleteBuildOrder(s, router)

	// Return the router
	return router
}

// Define all the routes

// Employee routes
// Get all employees
func GetEmployees(s *state.StateManager, rut *mux.Router) {
	rut.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		// Get all employees
		employees := s.GetEmployees()

		// Write the OK header
		w.WriteHeader(http.StatusOK)

		// Return a json object of all employees
		json.NewEncoder(w).Encode(employees)

		// Print to the server that we are done
		fmt.Println("Employees served")

	}).Methods("GET")
}

// Get a single employee
func GetEmployee(s *state.StateManager, rut *mux.Router) {
	rut.HandleFunc("/employees/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Get the employee id
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		// Get the employee
		employee := state.GetEmployee(s, id)

		// Write the OK header
		w.WriteHeader(http.StatusOK)

		// Return a json object of the employee
		json.NewEncoder(w).Encode(employee)

		// Print to the server that we are done
		fmt.Println("Employee served")

	}).Methods("GET")
}

// Create a new employee
func CreateEmployee(s *state.StateManager, rut *mux.Router) {
	rut.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		// Create an empty employee
		employee := models.Employee{}

		// Decode the request body
		json.NewDecoder(r.Body).Decode(&employee)

		// Create the employee
		employee = state.CreateEmployee(s, employee)

		// Write the OK header
		w.WriteHeader(http.StatusOK)

		// Return a json object of the employee
		json.NewEncoder(w).Encode(employee)

		// Print to the server that we are done
		fmt.Println("Employee served")

	}).Methods("POST")
}

// Update an employee
func UpdateEmployee(s *state.StateManager, rut *mux.Router) {
	rut.HandleFunc("/employees/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Get the employee id
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		// Create an empty employee
		employee := models.Employee{}

		// Decode the request body
		json.NewDecoder(r.Body).Decode(&employee)

		// Update the employee
		employee = state.UpdateEmployee(s, id, employee)

		// Write the OK header
		w.WriteHeader(http.StatusOK)

		// Return a json object of the employee
		json.NewEncoder(w).Encode(employee)

		// Print to the server that we are done
		fmt.Println("Employee served")

	}).Methods("PUT")
}

// Delete an employee
func DeleteEmployee(s *state.StateManager, rut *mux.Router) {
	rut.HandleFunc("/employees/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Get the employee id
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		// Delete the employee
		employee := state.DeleteEmployee(s, id)

		// Write the OK header
		w.WriteHeader(http.StatusOK)

		// Return a json object of the employee
		json.NewEncoder(w).Encode(employee)

		// Print to the server that we are done
		fmt.Println("Employee served")

	}).Methods("DELETE")
}

// Asset Routes
// Get all assets
func GetAssets(s *state.StateManager, rut *mux.Router) {
	rut.HandleFunc("/assets", func(w http.ResponseWriter, r *http.Request) {
		// Get all assets
		assets := s.GetAssets()

		// Write the OK header
		w.WriteHeader(http.StatusOK)

		// Return a json object of all assets
		json.NewEncoder(w).Encode(assets)

		// Print to the server that we are done
		fmt.Println("Assets served")

	}).Methods("GET")
}

// Get a single asset
func GetAsset(s *state.StateManager, rut *mux.Router) {
	rut.HandleFunc("/assets/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Get the asset id
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		// Get the asset
		asset := state.GetAsset(s, id)

		// Write the OK header
		w.WriteHeader(http.StatusOK)

		// Return a json object of the asset
		json.NewEncoder(w).Encode(asset)

		// Print to the server that we are done
		fmt.Println("Asset served")

	}).Methods("GET")
}

// Create a new asset
func CreateAsset(s *state.StateManager, rut *mux.Router) {
	rut.HandleFunc("/assets", func(w http.ResponseWriter, r *http.Request) {
		// Create an empty asset
		asset := models.Asset{}

		// Decode the request body
		json.NewDecoder(r.Body).Decode(&asset)

		// Create the asset
		asset = state.CreateAsset(s, asset)

		// Write the OK header
		w.WriteHeader(http.StatusOK)

		// Return a json object of the asset
		json.NewEncoder(w).Encode(asset)

		// Print to the server that we are done
		fmt.Println("Asset served")

	}).Methods("POST")
}

// Update an asset
func UpdateAsset(s *state.StateManager, rut *mux.Router) {
	rut.HandleFunc("/assets/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Get the asset id
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		// Create an empty asset
		asset := models.Asset{}

		// Decode the request body
		json.NewDecoder(r.Body).Decode(&asset)

		// Update the asset
		asset = state.UpdateAsset(s, id, asset)

		// Write the OK header
		w.WriteHeader(http.StatusOK)

		// Return a json object of the asset
		json.NewEncoder(w).Encode(asset)

		// Print to the server that we are done
		fmt.Println("Asset served")

	}).Methods("PUT")
}

// Delete an asset
func DeleteAsset(s *state.StateManager, rut *mux.Router) {
	rut.HandleFunc("/assets/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Get the asset id
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		// Delete the asset
		asset := state.DeleteAsset(s, id)

		// Write the OK header
		w.WriteHeader(http.StatusOK)

		// Return a json object of the asset
		json.NewEncoder(w).Encode(asset)

		// Print to the server that we are done
		fmt.Println("Asset served")

	}).Methods("DELETE")
}

// Build Orders
// Get all build orders
func GetBuildOrders(s *state.StateManager, rut *mux.Router) {
	rut.HandleFunc("/buildorders", func(w http.ResponseWriter, r *http.Request) {
		// Get all build orders
		buildorders := s.GetBuildOrders()

		// Write the OK header
		w.WriteHeader(http.StatusOK)

		// Return a json object of all build orders
		json.NewEncoder(w).Encode(buildorders)

		// Print to the server that we are done
		fmt.Println("Build orders served")

	}).Methods("GET")
}

// Get a single build order
func GetBuildOrder(s *state.StateManager, rut *mux.Router) {
	rut.HandleFunc("/buildorders/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Get the build order id
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		// Get the build order
		buildorder := state.GetBuildOrder(s, id)

		// Write the OK header
		w.WriteHeader(http.StatusOK)

		// Return a json object of the build order
		json.NewEncoder(w).Encode(buildorder)

		// Print to the server that we are done
		fmt.Println("Build order served")

	}).Methods("GET")
}

// Create a new build order
func CreateBuildOrder(s *state.StateManager, rut *mux.Router) {
	rut.HandleFunc("/buildorders", func(w http.ResponseWriter, r *http.Request) {
		// Create an empty build order
		buildorder := models.BuildOrder{}

		// Decode the request body
		json.NewDecoder(r.Body).Decode(&buildorder)

		// Create the build order
		buildorder = state.CreateBuildOrder(s, buildorder)

		// Write the OK header
		w.WriteHeader(http.StatusOK)

		// Return a json object of the build order
		json.NewEncoder(w).Encode(buildorder)

		// Print to the server that we are done
		fmt.Println("Build order served")

	}).Methods("POST")
}

// Update a build order
func UpdateBuildOrder(s *state.StateManager, rut *mux.Router) {
	rut.HandleFunc("/buildorders/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Get the build order id
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		// Create an empty build order
		buildorder := models.BuildOrder{}

		// Decode the request body
		json.NewDecoder(r.Body).Decode(&buildorder)

		// Update the build order
		buildorder = state.UpdateBuildOrder(s, id, buildorder)

		// Write the OK header
		w.WriteHeader(http.StatusOK)

		// Return a json object of the build order
		json.NewEncoder(w).Encode(buildorder)

		// Print to the server that we are done
		fmt.Println("Build order served")

	}).Methods("PUT")
}

// Delete a build order
func DeleteBuildOrder(s *state.StateManager, rut *mux.Router) {
	rut.HandleFunc("/buildorders/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Get the build order id
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		// Delete the build order
		buildorder := state.DeleteBuildOrder(s, id)

		// Write the OK header
		w.WriteHeader(http.StatusOK)

		// Return a json object of the build order
		json.NewEncoder(w).Encode(buildorder)

		// Print to the server that we are done
		fmt.Println("Build order served")

	}).Methods("DELETE")
}
