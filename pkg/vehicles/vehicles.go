// package vehicles contains vehicle types.
package vehicles

import (
	"github.com/komish/vendor-demo-mod-3/pkg/colors"
)

// vehicle represents a vehicle of some sort.
type vehicle struct {
	Type   string
	Wheels uint8
	Seats  uint8
	Color  colors.Color
}

// GetSedan returns a sedan vehicle that seats 4 in the color provided.
func GetSedan(color colors.Color) vehicle {
	return vehicle{
		Type:   "sedan",
		Wheels: 4,
		Seats:  4,
		Color:  colors.Color,
	}
}

// GetMotorcycle returns a motorcycle vehicle that seats 1 in the color provided.
func GetMotorcycle(color colors.Color) vehicle {
	return vehicle{
		Type:   "motorcycle",
		Wheels: 2,
		Seats:  1,
		Color:  colors.Color,
	}
}

// GetCoupe returns a coupe vehicle that seats 2 in the color provided.
func GetCoupe(color colors.Color) vehicle {
	return vehicle{
		Type:   "coupe",
		Wheels: 4,
		Seats:  2,
		Color:  colors.Color,
	}
}
