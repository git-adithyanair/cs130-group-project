package util

const (
	NUMERICAL = "numerical"
	OZ        = "oz"
	LBS       = "lbs"
	FL_OZ     = "fl_oz"
	GAL       = "gal"
	LITRES    = "litres"
)

func IsValidItemQuantityType(itemQuantityType string) bool {
	switch itemQuantityType {
	case NUMERICAL, OZ, LBS, FL_OZ, GAL, LITRES:
		return true
	default:
		return false
	}
}
