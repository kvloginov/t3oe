package base

type VolumeObject struct {
	Positional
	PivotRelativeX float64 // from 0 to 1
	PivotRelativeY float64 // from 0 to 1
	Width          float64 // in units
	Height         float64 // in units
}
