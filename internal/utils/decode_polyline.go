package utils

import (
	"errors"
	"explorer/internal/domain/models"
	"log"

	"github.com/twpayne/go-polyline"
)

func DecodeShapes(shapeData []models.Shape) ([][][]float64, error) {
	if len(shapeData) == 0 {
		return nil, errors.New("no shape data available")
	}

	// Decode the polylines
	var decodedShapes [][][]float64
	for _, shape := range shapeData {
		polylineData := shape.Attributes.PolyLine
		if polylineData != "" {
			decoded, _, err := polyline.DecodeCoords([]byte(polylineData))
			if err != nil {
				log.Println("Error decoding polyline:", err)
				continue
			}
			decodedShapes = append(decodedShapes, decoded)
		}
	}

	return decodedShapes, nil
}
