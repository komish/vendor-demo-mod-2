package vehicles

import (
	"github.com/komish/vendor-demo-mod-3/pkg/colors"
)

// GetColor is an accessor method for the Color struct
func GetColor(en, es string) colors.Color {
	return colors.GetColor(en, es)
}
