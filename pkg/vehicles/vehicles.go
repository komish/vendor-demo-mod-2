// Package vehicles contains vehicle types.
package vehicles

import (
	"github.com/komish/vendor-demo-mod-3/pkg/colors"
)

// Vehicle represents a vehicle of some sort.
type Vehicle struct {
	Type   string
	Wheels uint8
	Seats  uint8
	Color  colors.Color
}

// GetSedan returns a sedan vehicle that seats 4 in the color provided.
func GetSedan(color colors.Color) Vehicle {
	return Vehicle{
		Type:   "sedan",
		Wheels: 4,
		Seats:  4,
		Color:  color,
	}
}

// GetMotorcycle returns a motorcycle vehicle that seats 1 in the color provided.
func GetMotorcycle(color colors.Color) Vehicle {
	return Vehicle{
		Type:   "motorcycle",
		Wheels: 2,
		Seats:  1,
		Color:  color,
	}
}

// GetCoupe returns a coupe vehicle that seats 2 in the color provided.
func GetCoupe(color colors.Color) Vehicle {
	return Vehicle{
		Type:   "coupe",
		Wheels: 4,
		Seats:  2,
		Color:  color,
	}
}
