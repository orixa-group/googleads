package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/resources"
)

type NetworkSetting func(settings *resources.Campaign_NetworkSettings)

func NetworkTargetGoogleSearch(target bool) NetworkSetting {
	return func(settings *resources.Campaign_NetworkSettings) {
		settings.TargetGoogleSearch = Bool(target)
	}
}

func NetworkTargetSearchNetwork(target bool) NetworkSetting {
	return func(settings *resources.Campaign_NetworkSettings) {
		settings.TargetSearchNetwork = Bool(target)
	}
}

func NetworkTargetContentNetwork(target bool) NetworkSetting {
	return func(settings *resources.Campaign_NetworkSettings) {
		settings.TargetContentNetwork = Bool(target)
	}
}

func NetworkTargetPartnerSearchNetwork(target bool) NetworkSetting {
	return func(settings *resources.Campaign_NetworkSettings) {
		settings.TargetPartnerSearchNetwork = Bool(target)
	}
}

func NetworkDefaultSearch(settings *resources.Campaign_NetworkSettings) {
	settings.TargetGoogleSearch = Bool(true)
	settings.TargetSearchNetwork = Bool(true)
	settings.TargetContentNetwork = Bool(false)
	settings.TargetPartnerSearchNetwork = Bool(false)
}

func NetworkDefaultDisplay(settings *resources.Campaign_NetworkSettings) {
	settings.TargetGoogleSearch = Bool(false)
	settings.TargetSearchNetwork = Bool(false)
	settings.TargetContentNetwork = Bool(true)
	settings.TargetPartnerSearchNetwork = Bool(false)
}
