// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgAuthorizedClientsApi service type
type OrgAuthorizedClientsApi datadog.Service

// DeleteOrgAuthorizedClient Delete an org authorized client.
// Disable an OAuth2 client authorization for the current organization, revoking access for all users.
func (a *OrgAuthorizedClientsApi) DeleteOrgAuthorizedClient(ctx _context.Context, orgAuthorizedClientId string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.OrgAuthorizedClientsApi.DeleteOrgAuthorizedClient")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/org_authorized_clients/{org_authorized_client_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{org_authorized_client_id}", _neturl.PathEscape(datadog.ParameterToString(orgAuthorizedClientId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "*/*"

	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return nil, err
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
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 {
			var v JSONAPIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// DeleteOrgAuthorizedClientAllUserAuthorizations Delete a user's authorizations for a client.
// Disable all authorizations for a specific user for the specified OAuth2 client in the current organization.
func (a *OrgAuthorizedClientsApi) DeleteOrgAuthorizedClientAllUserAuthorizations(ctx _context.Context, orgAuthorizedClientId string, userId string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.OrgAuthorizedClientsApi.DeleteOrgAuthorizedClientAllUserAuthorizations")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/org_authorized_clients/{org_authorized_client_id}/user/{user_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{org_authorized_client_id}", _neturl.PathEscape(datadog.ParameterToString(orgAuthorizedClientId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{user_id}", _neturl.PathEscape(datadog.ParameterToString(userId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "*/*"

	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return nil, err
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
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 {
			var v JSONAPIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// DeleteOrgAuthorizedClientUserAuthorization Delete a user authorization for a client.
// Disable a specific user authorization for the specified OAuth2 client in the current organization.
func (a *OrgAuthorizedClientsApi) DeleteOrgAuthorizedClientUserAuthorization(ctx _context.Context, orgAuthorizedClientId string, userAuthorizedClientId string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.OrgAuthorizedClientsApi.DeleteOrgAuthorizedClientUserAuthorization")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/org_authorized_clients/{org_authorized_client_id}/user_authorized_clients/{user_authorized_client_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{org_authorized_client_id}", _neturl.PathEscape(datadog.ParameterToString(orgAuthorizedClientId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{user_authorized_client_id}", _neturl.PathEscape(datadog.ParameterToString(userAuthorizedClientId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "*/*"

	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return nil, err
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
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 {
			var v JSONAPIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// GetOrgAuthorizedClientOptionalParameters holds optional parameters for GetOrgAuthorizedClient.
type GetOrgAuthorizedClientOptionalParameters struct {
	Include                                 *string
	FilterUserAuthorizedClientsDisabled     *string
	FilterUserAuthorizedClientsUserDisabled *string
}

// NewGetOrgAuthorizedClientOptionalParameters creates an empty struct for parameters.
func NewGetOrgAuthorizedClientOptionalParameters() *GetOrgAuthorizedClientOptionalParameters {
	this := GetOrgAuthorizedClientOptionalParameters{}
	return &this
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *GetOrgAuthorizedClientOptionalParameters) WithInclude(include string) *GetOrgAuthorizedClientOptionalParameters {
	r.Include = &include
	return r
}

// WithFilterUserAuthorizedClientsDisabled sets the corresponding parameter name and returns the struct.
func (r *GetOrgAuthorizedClientOptionalParameters) WithFilterUserAuthorizedClientsDisabled(filterUserAuthorizedClientsDisabled string) *GetOrgAuthorizedClientOptionalParameters {
	r.FilterUserAuthorizedClientsDisabled = &filterUserAuthorizedClientsDisabled
	return r
}

// WithFilterUserAuthorizedClientsUserDisabled sets the corresponding parameter name and returns the struct.
func (r *GetOrgAuthorizedClientOptionalParameters) WithFilterUserAuthorizedClientsUserDisabled(filterUserAuthorizedClientsUserDisabled string) *GetOrgAuthorizedClientOptionalParameters {
	r.FilterUserAuthorizedClientsUserDisabled = &filterUserAuthorizedClientsUserDisabled
	return r
}

// GetOrgAuthorizedClient Get an org authorized client.
// Get a single OAuth2 client authorized for the current organization.
func (a *OrgAuthorizedClientsApi) GetOrgAuthorizedClient(ctx _context.Context, orgAuthorizedClientId string, o ...GetOrgAuthorizedClientOptionalParameters) (OrgAuthorizedClientResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue OrgAuthorizedClientResponse
		optionalParams      GetOrgAuthorizedClientOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type GetOrgAuthorizedClientOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.OrgAuthorizedClientsApi.GetOrgAuthorizedClient")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/org_authorized_clients/{org_authorized_client_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{org_authorized_client_id}", _neturl.PathEscape(datadog.ParameterToString(orgAuthorizedClientId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
	if optionalParams.FilterUserAuthorizedClientsDisabled != nil {
		localVarQueryParams.Add("filter[user_authorized_clients][disabled]", datadog.ParameterToString(*optionalParams.FilterUserAuthorizedClientsDisabled, ""))
	}
	if optionalParams.FilterUserAuthorizedClientsUserDisabled != nil {
		localVarQueryParams.Add("filter[user_authorized_clients][user][disabled]", datadog.ParameterToString(*optionalParams.FilterUserAuthorizedClientsUserDisabled, ""))
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
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 {
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

// ListOrgAuthorizedClientUserAuthorizationsOptionalParameters holds optional parameters for ListOrgAuthorizedClientUserAuthorizations.
type ListOrgAuthorizedClientUserAuthorizationsOptionalParameters struct {
	PageSize           *int64
	PageNumber         *int64
	Sort               *OrgAuthorizedClientUserAuthorizationsSort
	FilterDisabled     *string
	FilterUserName     *string
	FilterUserEmail    *string
	FilterUserDisabled *string
}

// NewListOrgAuthorizedClientUserAuthorizationsOptionalParameters creates an empty struct for parameters.
func NewListOrgAuthorizedClientUserAuthorizationsOptionalParameters() *ListOrgAuthorizedClientUserAuthorizationsOptionalParameters {
	this := ListOrgAuthorizedClientUserAuthorizationsOptionalParameters{}
	return &this
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListOrgAuthorizedClientUserAuthorizationsOptionalParameters) WithPageSize(pageSize int64) *ListOrgAuthorizedClientUserAuthorizationsOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListOrgAuthorizedClientUserAuthorizationsOptionalParameters) WithPageNumber(pageNumber int64) *ListOrgAuthorizedClientUserAuthorizationsOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *ListOrgAuthorizedClientUserAuthorizationsOptionalParameters) WithSort(sort OrgAuthorizedClientUserAuthorizationsSort) *ListOrgAuthorizedClientUserAuthorizationsOptionalParameters {
	r.Sort = &sort
	return r
}

// WithFilterDisabled sets the corresponding parameter name and returns the struct.
func (r *ListOrgAuthorizedClientUserAuthorizationsOptionalParameters) WithFilterDisabled(filterDisabled string) *ListOrgAuthorizedClientUserAuthorizationsOptionalParameters {
	r.FilterDisabled = &filterDisabled
	return r
}

// WithFilterUserName sets the corresponding parameter name and returns the struct.
func (r *ListOrgAuthorizedClientUserAuthorizationsOptionalParameters) WithFilterUserName(filterUserName string) *ListOrgAuthorizedClientUserAuthorizationsOptionalParameters {
	r.FilterUserName = &filterUserName
	return r
}

// WithFilterUserEmail sets the corresponding parameter name and returns the struct.
func (r *ListOrgAuthorizedClientUserAuthorizationsOptionalParameters) WithFilterUserEmail(filterUserEmail string) *ListOrgAuthorizedClientUserAuthorizationsOptionalParameters {
	r.FilterUserEmail = &filterUserEmail
	return r
}

// WithFilterUserDisabled sets the corresponding parameter name and returns the struct.
func (r *ListOrgAuthorizedClientUserAuthorizationsOptionalParameters) WithFilterUserDisabled(filterUserDisabled string) *ListOrgAuthorizedClientUserAuthorizationsOptionalParameters {
	r.FilterUserDisabled = &filterUserDisabled
	return r
}

// ListOrgAuthorizedClientUserAuthorizations List user authorizations for a client.
// Get a list of user authorizations for the specified OAuth2 client in the current organization.
func (a *OrgAuthorizedClientsApi) ListOrgAuthorizedClientUserAuthorizations(ctx _context.Context, orgAuthorizedClientId string, o ...ListOrgAuthorizedClientUserAuthorizationsOptionalParameters) (UserAuthorizedClientsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue UserAuthorizedClientsResponse
		optionalParams      ListOrgAuthorizedClientUserAuthorizationsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListOrgAuthorizedClientUserAuthorizationsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.OrgAuthorizedClientsApi.ListOrgAuthorizedClientUserAuthorizations")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/org_authorized_clients/{org_authorized_client_id}/user_authorized_clients"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{org_authorized_client_id}", _neturl.PathEscape(datadog.ParameterToString(orgAuthorizedClientId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("page[size]", datadog.ParameterToString(*optionalParams.PageSize, ""))
	}
	if optionalParams.PageNumber != nil {
		localVarQueryParams.Add("page[number]", datadog.ParameterToString(*optionalParams.PageNumber, ""))
	}
	if optionalParams.Sort != nil {
		localVarQueryParams.Add("sort", datadog.ParameterToString(*optionalParams.Sort, ""))
	}
	if optionalParams.FilterDisabled != nil {
		localVarQueryParams.Add("filter[disabled]", datadog.ParameterToString(*optionalParams.FilterDisabled, ""))
	}
	if optionalParams.FilterUserName != nil {
		localVarQueryParams.Add("filter[user][name]", datadog.ParameterToString(*optionalParams.FilterUserName, ""))
	}
	if optionalParams.FilterUserEmail != nil {
		localVarQueryParams.Add("filter[user][email]", datadog.ParameterToString(*optionalParams.FilterUserEmail, ""))
	}
	if optionalParams.FilterUserDisabled != nil {
		localVarQueryParams.Add("filter[user][disabled]", datadog.ParameterToString(*optionalParams.FilterUserDisabled, ""))
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
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 {
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

// ListOrgAuthorizedClientUserAuthorizationsWithPagination provides a paginated version of ListOrgAuthorizedClientUserAuthorizations returning a channel with all items.
func (a *OrgAuthorizedClientsApi) ListOrgAuthorizedClientUserAuthorizationsWithPagination(ctx _context.Context, orgAuthorizedClientId string, o ...ListOrgAuthorizedClientUserAuthorizationsOptionalParameters) (<-chan datadog.PaginationResult[UserAuthorizedClientData], func()) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int64(10)
	if len(o) == 0 {
		o = append(o, ListOrgAuthorizedClientUserAuthorizationsOptionalParameters{})
	}
	if o[0].PageSize != nil {
		pageSize_ = *o[0].PageSize
	}
	o[0].PageSize = &pageSize_
	page_ := int64(0)
	o[0].PageNumber = &page_

	items := make(chan datadog.PaginationResult[UserAuthorizedClientData], pageSize_)
	go func() {
		for {
			resp, _, err := a.ListOrgAuthorizedClientUserAuthorizations(ctx, orgAuthorizedClientId, o...)
			if err != nil {
				var returnItem UserAuthorizedClientData
				items <- datadog.PaginationResult[UserAuthorizedClientData]{Item: returnItem, Error: err}
				break
			}
			respData, ok := resp.GetDataOk()
			if !ok {
				break
			}
			results := *respData

			for _, item := range results {
				select {
				case items <- datadog.PaginationResult[UserAuthorizedClientData]{Item: item, Error: nil}:
				case <-ctx.Done():
					close(items)
					return
				}
			}
			if len(results) < int(pageSize_) {
				break
			}
			pageOffset_ := *o[0].PageNumber + 1
			o[0].PageNumber = &pageOffset_
		}
		close(items)
	}()
	return items, cancel
}

// ListOrgAuthorizedClientsOptionalParameters holds optional parameters for ListOrgAuthorizedClients.
type ListOrgAuthorizedClientsOptionalParameters struct {
	PageSize               *int64
	PageNumber             *int64
	Sort                   *string
	Filter                 *string
	FilterOauth2ClientName *string
	FilterDisabled         *string
	Include                *string
}

// NewListOrgAuthorizedClientsOptionalParameters creates an empty struct for parameters.
func NewListOrgAuthorizedClientsOptionalParameters() *ListOrgAuthorizedClientsOptionalParameters {
	this := ListOrgAuthorizedClientsOptionalParameters{}
	return &this
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListOrgAuthorizedClientsOptionalParameters) WithPageSize(pageSize int64) *ListOrgAuthorizedClientsOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListOrgAuthorizedClientsOptionalParameters) WithPageNumber(pageNumber int64) *ListOrgAuthorizedClientsOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *ListOrgAuthorizedClientsOptionalParameters) WithSort(sort string) *ListOrgAuthorizedClientsOptionalParameters {
	r.Sort = &sort
	return r
}

// WithFilter sets the corresponding parameter name and returns the struct.
func (r *ListOrgAuthorizedClientsOptionalParameters) WithFilter(filter string) *ListOrgAuthorizedClientsOptionalParameters {
	r.Filter = &filter
	return r
}

// WithFilterOauth2ClientName sets the corresponding parameter name and returns the struct.
func (r *ListOrgAuthorizedClientsOptionalParameters) WithFilterOauth2ClientName(filterOauth2ClientName string) *ListOrgAuthorizedClientsOptionalParameters {
	r.FilterOauth2ClientName = &filterOauth2ClientName
	return r
}

// WithFilterDisabled sets the corresponding parameter name and returns the struct.
func (r *ListOrgAuthorizedClientsOptionalParameters) WithFilterDisabled(filterDisabled string) *ListOrgAuthorizedClientsOptionalParameters {
	r.FilterDisabled = &filterDisabled
	return r
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *ListOrgAuthorizedClientsOptionalParameters) WithInclude(include string) *ListOrgAuthorizedClientsOptionalParameters {
	r.Include = &include
	return r
}

// ListOrgAuthorizedClients List org authorized clients.
// Get a list of all OAuth2 clients authorized for the current organization.
func (a *OrgAuthorizedClientsApi) ListOrgAuthorizedClients(ctx _context.Context, o ...ListOrgAuthorizedClientsOptionalParameters) (OrgAuthorizedClientsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue OrgAuthorizedClientsResponse
		optionalParams      ListOrgAuthorizedClientsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListOrgAuthorizedClientsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.OrgAuthorizedClientsApi.ListOrgAuthorizedClients")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/org_authorized_clients"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("page[size]", datadog.ParameterToString(*optionalParams.PageSize, ""))
	}
	if optionalParams.PageNumber != nil {
		localVarQueryParams.Add("page[number]", datadog.ParameterToString(*optionalParams.PageNumber, ""))
	}
	if optionalParams.Sort != nil {
		localVarQueryParams.Add("sort", datadog.ParameterToString(*optionalParams.Sort, ""))
	}
	if optionalParams.Filter != nil {
		localVarQueryParams.Add("filter", datadog.ParameterToString(*optionalParams.Filter, ""))
	}
	if optionalParams.FilterOauth2ClientName != nil {
		localVarQueryParams.Add("filter[oauth2_client][name]", datadog.ParameterToString(*optionalParams.FilterOauth2ClientName, ""))
	}
	if optionalParams.FilterDisabled != nil {
		localVarQueryParams.Add("filter[disabled]", datadog.ParameterToString(*optionalParams.FilterDisabled, ""))
	}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
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
		if localVarHTTPResponse.StatusCode == 403 {
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

// ListOrgAuthorizedClientsWithPagination provides a paginated version of ListOrgAuthorizedClients returning a channel with all items.
func (a *OrgAuthorizedClientsApi) ListOrgAuthorizedClientsWithPagination(ctx _context.Context, o ...ListOrgAuthorizedClientsOptionalParameters) (<-chan datadog.PaginationResult[OrgAuthorizedClientData], func()) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int64(10)
	if len(o) == 0 {
		o = append(o, ListOrgAuthorizedClientsOptionalParameters{})
	}
	if o[0].PageSize != nil {
		pageSize_ = *o[0].PageSize
	}
	o[0].PageSize = &pageSize_
	page_ := int64(0)
	o[0].PageNumber = &page_

	items := make(chan datadog.PaginationResult[OrgAuthorizedClientData], pageSize_)
	go func() {
		for {
			resp, _, err := a.ListOrgAuthorizedClients(ctx, o...)
			if err != nil {
				var returnItem OrgAuthorizedClientData
				items <- datadog.PaginationResult[OrgAuthorizedClientData]{Item: returnItem, Error: err}
				break
			}
			respData, ok := resp.GetDataOk()
			if !ok {
				break
			}
			results := *respData

			for _, item := range results {
				select {
				case items <- datadog.PaginationResult[OrgAuthorizedClientData]{Item: item, Error: nil}:
				case <-ctx.Done():
					close(items)
					return
				}
			}
			if len(results) < int(pageSize_) {
				break
			}
			pageOffset_ := *o[0].PageNumber + 1
			o[0].PageNumber = &pageOffset_
		}
		close(items)
	}()
	return items, cancel
}

// UpdateOrgAuthorizedClient Update an org authorized client.
// Enable or disable an OAuth2 client authorization for the current organization.
func (a *OrgAuthorizedClientsApi) UpdateOrgAuthorizedClient(ctx _context.Context, orgAuthorizedClientId string, body OrgAuthorizedClientUpdateRequest) (OrgAuthorizedClientResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue OrgAuthorizedClientResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.OrgAuthorizedClientsApi.UpdateOrgAuthorizedClient")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/org_authorized_clients/{org_authorized_client_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{org_authorized_client_id}", _neturl.PathEscape(datadog.ParameterToString(orgAuthorizedClientId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = &body
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 422 {
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

// NewOrgAuthorizedClientsApi Returns NewOrgAuthorizedClientsApi.
func NewOrgAuthorizedClientsApi(client *datadog.APIClient) *OrgAuthorizedClientsApi {
	return &OrgAuthorizedClientsApi{
		Client: client,
	}
}
