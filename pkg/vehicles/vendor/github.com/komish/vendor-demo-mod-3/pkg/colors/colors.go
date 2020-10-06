// Package colors provides access to the Color type
package colors

// Color is the type for a color which contains english and spanish names for that color.
type Color struct {
	NameEnglish string
	NameSpanish string
}

// GetColor is an accessor method for the Color struct
func GetColor(en, es string) Color {
	return Color{
		NameEnglish: en,
		NameSpanish: es,
	}
}
