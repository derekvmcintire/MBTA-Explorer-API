package pkg

import (
	"errors" // Import the errors package to create and handle errors
	"explorer/internal/core/domain/models"
	"log" // Import the log package for logging information

	"github.com/twpayne/go-polyline" // Import the go-polyline package to decode polyline strings
)

// DecodeShapes takes a slice of Shape data and decodes the encoded polylines into latitude and longitude coordinates
// The function returns a 3D slice of float64 values representing decoded shapes and an error if decoding fails
func DecodeShapes(shapeData []models.Shape) ([][][]float64, error) {
	// Check if the input shapeData is empty
	if len(shapeData) == 0 {
		return nil, errors.New("no shape data available") // Return an error indicating no shape data is provided
	}

	// Initialize a 3D slice to hold the decoded shapes
	var decodedShapes [][][]float64

	// Iterate over the shape data to decode each polyline
	for _, shape := range shapeData {
		polylineData := shape.Attributes.PolyLine // Get the encoded polyline data from shape attributes

		// Check if the polyline data is non-empty
		if polylineData != "" {
			// Decode the polyline data into a slice of latitude/longitude pairs
			decoded, _, err := polyline.DecodeCoords([]byte(polylineData))
			if err != nil {
				// Log the decoding error and skip the current polyline
				log.Println("Error decoding polyline:", err)
				continue
			}
			// Append the successfully decoded shape to the result
			decodedShapes = append(decodedShapes, decoded)
		}
	}

	// Return the decoded shapes and no error
	return decodedShapes, nil
}
