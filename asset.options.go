package googleads

import (
	"io"
	"net/http"

	"github.com/shenzhencenter/google-ads-pb/common"
	"github.com/shenzhencenter/google-ads-pb/resources"
)

type AssetOption func(*resources.Asset)

func ChangeAssetPhoneNumber(number string) AssetOption {
	return func(a *resources.Asset) {
		if data, ok := a.GetAssetData().(*resources.Asset_CallAsset); ok {
			data.CallAsset.PhoneNumber = number
		}
	}
}

type AssetImageSource func(asset *common.ImageAsset) error

func AssetImageFromBytes(data []byte) AssetImageSource {
	return func(asset *common.ImageAsset) error {
		asset.Data = data
		return nil
	}
}

func AssetImageFromUrl(url string) AssetImageSource {
	return func(asset *common.ImageAsset) error {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		asset.Data = data
		return nil
	}
}
