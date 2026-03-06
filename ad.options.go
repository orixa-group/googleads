package googleads

import (
	"strings"

	"github.com/shenzhencenter/google-ads-pb/resources"
)

type AdOption func(*resources.Ad)

func ReplaceWordInAd(old, new string) AdOption {
	return func(a *resources.Ad) {
		if data, ok := a.GetAdData().(*resources.Ad_ResponsiveSearchAd); ok {
			for _, h := range data.ResponsiveSearchAd.GetHeadlines() {
				h.Text = String(strings.ReplaceAll(h.GetText(), old, new))
			}
			for _, d := range data.ResponsiveSearchAd.GetDescriptions() {
				d.Text = String(strings.ReplaceAll(d.GetText(), old, new))
			}
		}
	}
}

func ReplaceURLInAd(url string) AdOption {
	return func(a *resources.Ad) {
		a.FinalUrls = []string{url}
	}
}
