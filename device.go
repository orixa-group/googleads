package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/common"
	"github.com/shenzhencenter/google-ads-pb/enums"
)

type DeviceType func(device *common.DeviceInfo)

func DeviceMobile(device *common.DeviceInfo) {
	device.Type = enums.DeviceEnum_MOBILE
}

func DeviceTablet(device *common.DeviceInfo) {
	device.Type = enums.DeviceEnum_TABLET
}

func DeviceDesktop(device *common.DeviceInfo) {
	device.Type = enums.DeviceEnum_DESKTOP
}

func DeviceConnectedTV(device *common.DeviceInfo) {
	device.Type = enums.DeviceEnum_CONNECTED_TV
}
