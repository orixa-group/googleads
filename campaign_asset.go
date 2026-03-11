package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/common"
	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

type CampaignAsset struct {
	*resources.CampaignAsset
	Asset *Asset
}

func (c *CampaignAsset) createOperations(customer *Customer, campaign *Campaign, tempId tempIdGenerator) []*services.MutateOperation {
	aop := c.Asset.createOperation(customer, tempId)

	c.CampaignAsset.Campaign = String(campaign.GetResourceName())
	c.CampaignAsset.Asset = String(c.Asset.GetResourceName())

	return []*services.MutateOperation{
		aop,
		{
			Operation: &services.MutateOperation_CampaignAssetOperation{
				CampaignAssetOperation: &services.CampaignAssetOperation{
					Operation: &services.CampaignAssetOperation_Create{
						Create: c.CampaignAsset,
					},
				},
			},
		},
	}
}

type CampaignAssets []*CampaignAsset

func (c *CampaignAssets) Add(asset *CampaignAsset, options ...AssetOption) {
	a := &resources.Asset{
		AssetData: asset.Asset.GetAssetData(),
		FinalUrls: asset.Asset.GetFinalUrls(),
		Name:      String(asset.Asset.GetName()),
	}
	for _, opt := range options {
		opt(a)
	}

	*c = append(*c, &CampaignAsset{&resources.CampaignAsset{
		FieldType: asset.GetFieldType(),
	}, &Asset{a}})
}

func (c *CampaignAssets) AddSitelink(text, description1, description2 string, finalUrls ...string) {
	c.Add(&CampaignAsset{
		CampaignAsset: &resources.CampaignAsset{
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

func (c *CampaignAssets) AddCallout(text string) {
	c.Add(&CampaignAsset{
		CampaignAsset: &resources.CampaignAsset{
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

func (c *CampaignAssets) AddCall(countryCode, phoneNumber string) {
	c.Add(&CampaignAsset{
		CampaignAsset: &resources.CampaignAsset{
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

func (c *CampaignAssets) AddImage(source AssetImageSource) error {
	ia := &common.ImageAsset{}
	if err := source(ia); err != nil {
		return err
	}

	c.Add(&CampaignAsset{
		CampaignAsset: &resources.CampaignAsset{
			FieldType: enums.AssetFieldTypeEnum_AD_IMAGE,
		},
		Asset: &Asset{
			Asset: &resources.Asset{
				AssetData: &resources.Asset_ImageAsset{
					ImageAsset: ia,
				},
			},
		},
	})

	return nil
}

func (c CampaignAssets) createOperations(customer *Customer, campaign *Campaign, tempId tempIdGenerator) []*services.MutateOperation {
	return Flatten(Map(c, func(item *CampaignAsset) []*services.MutateOperation {
		return item.createOperations(customer, campaign, tempId)
	}))
}
