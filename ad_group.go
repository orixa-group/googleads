package googleads

import (
	"context"
	"fmt"
	"strconv"

	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

type AdGroup struct {
	*resources.AdGroup
	Campaign *Campaign
	Criteria AdGroupCriteria
	Assets   AdGroupAssets
	Ads      AdGroupAds
}

func NewAdGroup() *AdGroup {
	return &AdGroup{
		AdGroup:  &resources.AdGroup{},
		Campaign: NewCampaign(),
		Criteria: NewAdGroupCriteria(),
		Assets:   NewAdGroupAssets(),
		Ads:      NewAdGroupAds(),
	}
}

func (ag AdGroup) GetId() string {
	return strconv.Itoa(int(ag.AdGroup.GetId()))
}

func (ag *AdGroup) SetId(id string) {
	i, _ := strconv.ParseInt(id, 10, 64)
	ag.AdGroup.Id = Int64(i)
}

func (ag AdGroup) GetName() string {
	return ag.AdGroup.GetName()
}

func (ag *AdGroup) SetName(name string) {
	ag.AdGroup.Name = String(name)
}

func (ag AdGroup) GetEnabled() bool {
	return ag.AdGroup.GetStatus() == enums.AdGroupStatusEnum_ENABLED
}

func (ag *AdGroup) SetEnabled(enabled bool) {
	if enabled {
		ag.AdGroup.Status = enums.AdGroupStatusEnum_ENABLED
	} else {
		ag.AdGroup.Status = enums.AdGroupStatusEnum_PAUSED
	}
}

func (ag AdGroup) GetCpcBidCents() int {
	return int(ag.AdGroup.GetCpcBidMicros() / 10_000)
}

func (ag *AdGroup) SetCpcBidCents(cents int) {
	ag.AdGroup.CpcBidMicros = Int64(int64(cents * 10_000))
}

func (ag AdGroup) IsOptimizedTargetingEnabled() bool {
	return ag.AdGroup.GetOptimizedTargetingEnabled()
}

func (ag *AdGroup) SetOptimizedTargetingEnabled(enabled bool) {
	ag.AdGroup.OptimizedTargetingEnabled = enabled
}

func (ag AdGroup) GetTrackingUrl() string {
	return ag.AdGroup.GetTrackingUrlTemplate()
}

func (ag *AdGroup) SetTrackingUrl(url string) {
	ag.AdGroup.TrackingUrlTemplate = String(url)
}

func (ag AdGroup) GetFinalUrlSuffix() string {
	return ag.AdGroup.GetFinalUrlSuffix()
}

func (ag *AdGroup) SetFinalUrlSuffix(suffix string) {
	ag.AdGroup.FinalUrlSuffix = String(suffix)
}

func (ag AdGroup) GetType() enums.AdGroupTypeEnum_AdGroupType {
	return ag.AdGroup.GetType()
}

func (ag AdGroup) GetAdRotationMode() enums.AdGroupAdRotationModeEnum_AdGroupAdRotationMode {
	return ag.AdGroup.GetAdRotationMode()
}

type AdGroupType func(ag *resources.AdGroup)

func (ag *AdGroup) SetType(fn AdGroupType) {
	fn(ag.AdGroup)
}

type AdGroupAdRotationMode func(ag *resources.AdGroup)

func (ag *AdGroup) SetAdRotationMode(fn AdGroupAdRotationMode) {
	fn(ag.AdGroup)
}

func AdGroupTypeSearchStandard(ag *resources.AdGroup) {
	ag.Type = enums.AdGroupTypeEnum_SEARCH_STANDARD
}

func AdGroupTypeSearchDynamic(ag *resources.AdGroup) {
	ag.Type = enums.AdGroupTypeEnum_SEARCH_DYNAMIC_ADS
}

func AdGroupTypeDisplayStandard(ag *resources.AdGroup) {
	ag.Type = enums.AdGroupTypeEnum_DISPLAY_STANDARD
}

func AdGroupRotationOptimize(ag *resources.AdGroup) {
	ag.AdRotationMode = enums.AdGroupAdRotationModeEnum_OPTIMIZE
}

func AdGroupRotationRotateForever(ag *resources.AdGroup) {
	ag.AdRotationMode = enums.AdGroupAdRotationModeEnum_ROTATE_FOREVER
}

func (ag *AdGroup) createOperation(tempId tempIdGenerator) *services.MutateOperation {
	ag.ResourceName = fmt.Sprintf("customers/%s/adGroups/%s", ag.Campaign.Customer.GetId(), tempId())
	ag.AdGroup.Campaign = String(ag.Campaign.GetResourceName())

	return &services.MutateOperation{
		Operation: &services.MutateOperation_AdGroupOperation{
			AdGroupOperation: &services.AdGroupOperation{
				Operation: &services.AdGroupOperation_Create{
					Create: ag.AdGroup,
				},
			},
		},
	}
}

func (ag *AdGroup) Create(ctx context.Context) error {
	tempId := newTempIdGenerator()

	ops := make([]*services.MutateOperation, 0)

	ops = append(ops, ag.createOperation(tempId))
	ops = append(ops, ag.Criteria.createOperations(ag)...)
	ops = append(ops, ag.Assets.createOperations(ag, tempId)...)
	ops = append(ops, ag.Ads.createOperations(ag)...)

	_, err := services.NewGoogleAdsServiceClient(instance.conn).Mutate(ctx, &services.MutateGoogleAdsRequest{
		CustomerId:       ag.Campaign.Customer.GetId(),
		MutateOperations: ops,
	})
	return err
}

type AdGroups []*AdGroup

func NewAdGroups() AdGroups {
	return make(AdGroups, 0)
}
