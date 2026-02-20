package googleads

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

var ErrResourceNotFound = errors.New("resource not found")

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

func Create[
	R any,
	Req interface {
		GetCustomerId() string
	},
	Res any,
	Filter ~func() string,
	Builder interface {
		Where(...Filter) Builder
		Build() string
	},
](ctx context.Context,
	createResource func(context.Context, Req, ...grpc.CallOption) (Res, error),
	req Req,
	getResourceName func(string, Res) string,
	service services.GoogleAdsServiceClient,
	builder Builder,
	byResourceName func(string) Filter,
	createInstance func(*services.GoogleAdsRow) *R,
) (*R, error) {
	res, err := createResource(ctx, req)
	if err != nil {
		grpcErr, ok := status.FromError(err)
		if !ok {
			return nil, fmt.Errorf("failed to create %T: %v", new(R), err)
		}

		return nil, fmt.Errorf("failed to create %T: %s", new(R), strings.Join(Map(grpcErr.Proto().GetDetails(), func(detail *anypb.Any) string {
			return string(detail.GetValue())
		}), ", "))
	}

	return Fetch(ctx, service, req.GetCustomerId(), builder.Where(byResourceName(getResourceName(req.GetCustomerId(), res))).Build(), createInstance)
}
