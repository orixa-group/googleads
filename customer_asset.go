package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/common"
	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

type CustomerAsset struct {
	*resources.CustomerAsset
	Asset *Asset
}

func NewCustomerAsset() *CustomerAsset {
	return &CustomerAsset{
		CustomerAsset: &resources.CustomerAsset{},
		Asset:         NewAsset(),
	}
}

func (c *CustomerAsset) createOperations(customer *Customer, tempId tempIdGenerator) []*services.MutateOperation {
	aop := c.Asset.createOperation(customer, tempId)

	c.CustomerAsset.Asset = c.Asset.GetResourceName()

	return []*services.MutateOperation{
		aop,
		{
			Operation: &services.MutateOperation_CustomerAssetOperation{
				CustomerAssetOperation: &services.CustomerAssetOperation{
					Operation: &services.CustomerAssetOperation_Create{
						Create: c.CustomerAsset,
					},
				},
			},
		},
	}
}

type CustomerAssets []*CustomerAsset

func NewCustomerAssets() CustomerAssets {
	return make(CustomerAssets, 0)
}

func (c *CustomerAssets) Add(asset *CustomerAsset) {
	*c = append(*c, &CustomerAsset{&resources.CustomerAsset{
		FieldType: asset.GetFieldType(),
	}, &Asset{&resources.Asset{
		AssetData: asset.Asset.GetAssetData(),
		FinalUrls: asset.Asset.GetFinalUrls(),
	}}})
}

func (c *CustomerAssets) AddSitelink(text, description1, description2 string, finalUrls ...string) {
	c.Add(&CustomerAsset{
		CustomerAsset: &resources.CustomerAsset{
			FieldType: enums.AssetFieldTypeEnum_SITELINK,
		},
		Asset: &Asset{
			Asset: &resources.Asset{
				FinalUrls: finalUrls,
				AssetData: &resources.Asset_SitelinkAsset{
					SitelinkAsset: &common.SitelinkAsset{
						LinkText:     text,
						Description1: description1,
						Description2: description2,
					},
				},
			},
		},
	})
}

func (c *CustomerAssets) AddCallout(text string) {
	c.Add(&CustomerAsset{
		CustomerAsset: &resources.CustomerAsset{
			FieldType: enums.AssetFieldTypeEnum_CALLOUT,
		},
		Asset: &Asset{
			Asset: &resources.Asset{
				AssetData: &resources.Asset_CalloutAsset{
					CalloutAsset: &common.CalloutAsset{
						CalloutText: text,
					},
				},
			},
		},
	})
}

func (c *CustomerAssets) AddCall(countryCode, phoneNumber string) {
	c.Add(&CustomerAsset{
		CustomerAsset: &resources.CustomerAsset{
			FieldType: enums.AssetFieldTypeEnum_CALL,
		},
		Asset: &Asset{
			Asset: &resources.Asset{
				AssetData: &resources.Asset_CallAsset{
					CallAsset: &common.CallAsset{
						CountryCode: countryCode,
						PhoneNumber: phoneNumber,
					},
				},
			},
		},
	})
}

func (c *CustomerAssets) AddStructuredSnippet(header string, values ...string) {
	c.Add(&CustomerAsset{
		CustomerAsset: &resources.CustomerAsset{
			FieldType: enums.AssetFieldTypeEnum_STRUCTURED_SNIPPET,
		},
		Asset: &Asset{
			Asset: &resources.Asset{
				AssetData: &resources.Asset_StructuredSnippetAsset{
					StructuredSnippetAsset: &common.StructuredSnippetAsset{
						Header: header,
						Values: values,
					},
				},
			},
		},
	})
}

func (c CustomerAssets) createOperations(customer *Customer, tempId tempIdGenerator) []*services.MutateOperation {
	return Flatten(Map(c, func(item *CustomerAsset) []*services.MutateOperation {
		return item.createOperations(customer, tempId)
	}))
}
