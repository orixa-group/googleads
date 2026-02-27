package googleads

import (
	"fmt"
	"io"
	"net/http"

	"github.com/shenzhencenter/google-ads-pb/common"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

type Asset struct {
	*resources.Asset
}

func NewAsset() *Asset {
	return &Asset{&resources.Asset{}}
}

func (a *Asset) createOperation(customer *Customer, tempId tempIdGenerator) *services.MutateOperation {
	a.ResourceName = fmt.Sprintf("customers/%s/assets/%s", customer.GetId(), tempId())

	return &services.MutateOperation{
		Operation: &services.MutateOperation_AssetOperation{
			AssetOperation: &services.AssetOperation{
				Operation: &services.AssetOperation_Create{
					Create: a.Asset,
				},
			},
		},
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
