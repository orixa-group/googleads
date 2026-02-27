package googleads

import (
	"strconv"

	"github.com/shenzhencenter/google-ads-pb/resources"
)

type BillingSetup struct {
	*resources.BillingSetup
}

func NewBillingSetup() *BillingSetup {
	return &BillingSetup{&resources.BillingSetup{}}
}

func (bs *BillingSetup) GetId() string {
	return strconv.Itoa(int(bs.BillingSetup.GetId()))
}

func (bs *BillingSetup) SetId(id string) {
	i, _ := strconv.ParseInt(id, 10, 64)
	bs.BillingSetup.Id = Int64(i)
}
