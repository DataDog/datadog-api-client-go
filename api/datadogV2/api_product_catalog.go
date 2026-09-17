// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	_context "context"
	_fmt "fmt"
	_log "log"
	_nethttp "net/http"
	_neturl "net/url"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductCatalogApi service type
type ProductCatalogApi datadog.Service

// ListProductCatalogSKUsOptionalParameters holds optional parameters for ListProductCatalogSKUs.
type ListProductCatalogSKUsOptionalParameters struct {
	AsOfDate *time.Time
}

// NewListProductCatalogSKUsOptionalParameters creates an empty struct for parameters.
func NewListProductCatalogSKUsOptionalParameters() *ListProductCatalogSKUsOptionalParameters {
	this := ListProductCatalogSKUsOptionalParameters{}
	return &this
}

// WithAsOfDate sets the corresponding parameter name and returns the struct.
func (r *ListProductCatalogSKUsOptionalParameters) WithAsOfDate(asOfDate time.Time) *ListProductCatalogSKUsOptionalParameters {
	r.AsOfDate = &asOfDate
	return r
}

// ListProductCatalogSKUs List SKUs.
// Get every generally available Datadog SKU, with the pricing and allotment metadata that
// applies to it, for the Datadog site serving the request. A SKU is generally available
// when it is billed through a metered commitment or through automatic billing; SKUs in any
// other phase are not returned.
//
// Prices, allotments, and pricing tiers are returned as they were in effect on
// `as_of_date`, which defaults to the date of the request. Prices are public list prices:
// they do not reflect discounts, commitments, or negotiated rates on an account.
//
// Each SKU is a separate resource in `data`, identified by its SKU code, and sorted by
// that code in ascending order. The whole catalog is returned in a single response, so
// this endpoint is not paginated.
func (a *ProductCatalogApi) ListProductCatalogSKUs(ctx _context.Context, version ProductCatalogSKUsAPIVersion, o ...ListProductCatalogSKUsOptionalParameters) (ProductCatalogSKUsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ProductCatalogSKUsResponse
		optionalParams      ListProductCatalogSKUsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListProductCatalogSKUsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.ListProductCatalogSKUs"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.ProductCatalogApi.ListProductCatalogSKUs")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/product-catalog/skus"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("version", datadog.ParameterToString(version, ""))
	if optionalParams.AsOfDate != nil {
		localVarQueryParams.Add("as_of_date", datadog.ParameterToString(*optionalParams.AsOfDate, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return localVarReturnValue, nil, err
		}
	} else {
		datadog.SetAuthKeys(
			ctx,
			&localVarHeaderParams,
			[2]string{"apiKeyAuth", "DD-API-KEY"},
			[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
		)
	}
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 {
			var v JSONAPIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// NewProductCatalogApi Returns NewProductCatalogApi.
func NewProductCatalogApi(client *datadog.APIClient) *ProductCatalogApi {
	return &ProductCatalogApi{
		Client: client,
	}
}
