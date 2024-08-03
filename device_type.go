package udetect

// DeviceType object declaration
type DeviceType int

// RTB 5.17 Device Type
const (
	DeviceTypeUnknown   DeviceType = 0
	DeviceTypeMobile    DeviceType = 1
	DeviceTypePC        DeviceType = 2
	DeviceTypeTV        DeviceType = 3
	DeviceTypePhone     DeviceType = 4
	DeviceTypeTablet    DeviceType = 5
	DeviceTypeConnected DeviceType = 6
	DeviceTypeSetTopBox DeviceType = 7
	DeviceTypeWatch     DeviceType = 8
	DeviceTypeGlasses   DeviceType = 9
	DeviceTypeOOH       DeviceType = 10
)
