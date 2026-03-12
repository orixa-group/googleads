package googleads

import (
	"fmt"
	"strings"

	"github.com/shenzhencenter/google-ads-pb/errors"
	"google.golang.org/grpc/status"
)

type AdsServiceError struct {
	Message   string
	FieldPath []string
}

func (ase AdsServiceError) Error() string {
	return fmt.Sprintf("error: %s\npath: %s\n", ase.Message, strings.Join(ase.FieldPath, " -> "))
}

type AdsServiceErrors []*AdsServiceError

func (ases AdsServiceErrors) Error() string {
	return strings.Join(Map(ases, func(ase *AdsServiceError) string {
		return ase.Error()
	}), "\n")
}

func DecodeAdsServiceError(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return err
	}

	var ases AdsServiceErrors
	for _, detail := range st.Details() {
		switch t := detail.(type) {
		case *errors.GoogleAdsFailure:
			for _, adsErr := range t.GetErrors() {
				ases = append(ases, &AdsServiceError{
					Message: adsErr.GetMessage(),
					FieldPath: Map(adsErr.GetLocation().GetFieldPathElements(), func(field *errors.ErrorLocation_FieldPathElement) string {
						return field.GetFieldName()
					}),
				})
			}
		}
	}

	return ases
}
