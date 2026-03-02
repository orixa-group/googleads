package googleads

import (
	"context"
	"errors"
	"fmt"

	"github.com/shenzhencenter/google-ads-pb/services"
)

var ErrResourceNotFound = errors.New("resource not found")

type resource struct {
	updatedFields map[string]struct{}
}

func (r *resource) addUpdatedField(field string) {
	if r.updatedFields == nil {
		r.updatedFields = make(map[string]struct{})
	}
	r.updatedFields[field] = struct{}{}
}

func (r *resource) isNew(id string) bool {
	return id == "0" || id == ""
}

func (r *resource) getUpdatedFields() []string {
	paths := make([]string, 0, len(r.updatedFields))
	for path := range r.updatedFields {
		paths = append(paths, path)
	}
	return paths
}

func Fetch[R any](ctx context.Context, service services.GoogleAdsServiceClient, customerId, query string, createInstance func(*services.GoogleAdsRow) *R) (*R, error) {
	resp, err := service.Search(ctx, &services.SearchGoogleAdsRequest{
		CustomerId: customerId,
		Query:      query,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to search %T: %v", new(R), err)
	}

	if len(resp.Results) <= 0 {
		return nil, fmt.Errorf("failed to search %T: %v", new(R), ErrResourceNotFound)
	}

	return createInstance(resp.Results[0]), nil
}

func List[R any](ctx context.Context, service services.GoogleAdsServiceClient, customerId, query string, createInstance func(*services.GoogleAdsRow) *R) ([]*R, error) {
	resp, err := service.Search(ctx, &services.SearchGoogleAdsRequest{
		CustomerId: customerId,
		Query:      query,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to search %T: %v", new(R), err)
	}

	return Map(resp.Results, func(result *services.GoogleAdsRow) *R {
		return createInstance(result)
	}), nil
}
