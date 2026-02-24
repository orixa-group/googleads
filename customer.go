package googleads

import (
	"strconv"

	"github.com/shenzhencenter/google-ads-pb/resources"
)

type Customer struct {
	*resources.Customer
}

func (c Customer) GetId() string {
	return strconv.Itoa(int(c.Customer.GetId()))
}
