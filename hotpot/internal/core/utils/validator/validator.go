// Package validator provides utilities for validating data transfer objects (DTOs)
// and mapping generic responses to specific structures.
package validator

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

// ValidateDTO validates a given data transfer object (DTO) using the go-playground/validator library.
//
// Arguments:
//
//	dto - The DTO to validate (must be a struct).
//
// Returns:
//
//	An error if validation fails, or nil if the DTO is valid.
func ValidateDTO(dto any) error {
	v := validator.New() // Create a new validator instance.
	return v.Struct(dto) // Validate the struct and return any errors.
}

// MapResponse maps a generic response to a target DTO by serializing and deserializing the data.
//
// Arguments:
//
//	response - The generic response to map (e.g., map[string]any, struct).
//	target - A pointer to the target DTO where the response should be mapped.
//
// Behavior:
//   - The function marshals the response into JSON format.
//   - Then, it unmarshals the JSON into the target structure.
//
// Returns:
//
//	An error if marshaling or unmarshaling fails, or nil if the operation is successful.
func MapResponse(response any, target any) error {
	// Convert the response to JSON.
	data, err := json.Marshal(response)
	if err != nil {
		return err
	}

	// Map the JSON data to the target DTO.
	return json.Unmarshal(data, target)
}
