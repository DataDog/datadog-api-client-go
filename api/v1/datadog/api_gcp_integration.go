/*
 * (C) Datadog, Inc. 2019
 * All rights reserved
 * Licensed under a 3-clause BSD style license (see LICENSE)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// GCPIntegrationApiService GCPIntegrationApi service
type GCPIntegrationApiService service

/*
CreateGCPIntegration Add a GCP integration to your Datadog account.
### Overview Create a Datadog-GCP integration. ### Arguments * **&#x60;type&#x60;** [*required*, *default* &#x3D; **None**]: The value for service_account found in your JSON service account key. * **&#x60;project_id&#x60;** [*required*, *default* &#x3D; **None**]: Your Google Cloud project ID found in your JSON service account key. * **&#x60;private_key_id&#x60;** [*required*, *default* &#x3D; **None**]: Your private key ID found in your JSON service account key. * **&#x60;private_key&#x60;** [*required*, *default* &#x3D; **None**]: Your private key name found in your JSON service account key. * **&#x60;client_email&#x60;** [*required*, *default* &#x3D; **None**]: Your email found in your JSON service account key. * **&#x60;client_id&#x60;** [*required*, *default* &#x3D; **None**]: Your ID found in your JSON service account key. * **&#x60;auth_uri&#x60;** [*required*, *default* &#x3D; **None**]: Should be https://accounts.google.com/o/oauth2/auth. * **&#x60;token_uri&#x60;** [*required*, *default* &#x3D; **None**]: Should be https://accounts.google.com/o/oauth2/token. * **&#x60;auth_provider_x509_cert_url&#x60;** [*required*, *default* &#x3D; **None**]: Should be https://www.googleapis.com/oauth2/v1/certs. * **&#x60;client_x509_cert_url&#x60;** [*required*, *default* &#x3D; **None**]: Should be https://www.googleapis.com/robot/v1/metadata/x509/&lt;CLIENT_EMAIL&gt; where &lt;CLIENT_EMAIL&gt; is the email found in your JSON service account key. * **&#x60;host_filters&#x60;** [*optional*, *default* &#x3D; **None**]: Limit the GCE instances that are pulled into Datadog by using tags. Only hosts that match one of the defined tags are imported into Datadog. * **&#x60;automute&#x60;** [*optional*, *default* &#x3D; false]: Silence monitors for expected GCE instance shutdowns.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param gcpAccount Create a Datadog-Azure integration.
@return interface{}
*/
func (a *GCPIntegrationApiService) CreateGCPIntegration(ctx _context.Context, gcpAccount GcpAccount) (interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(ctx, "GCPIntegrationApiService.CreateGCPIntegration")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/integration/gcp"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &gcpAccount
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["api_key"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarQueryParams.Add("api_key", key)
			}
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["application_key"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarQueryParams.Add("application_key", key)
			}
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
DeleteGCPIntegration Delete a GCP Integration from your Datadog account.
### Overview Delete a given Datadog-GCP integration. ### Arguments * **&#x60;project_id&#x60;** [*required*, *default* &#x3D; **None**]: Your Google Cloud project ID to be deleted. * **&#x60;client_email&#x60;** [*required*, *default* &#x3D; **None**]: Your client email associated with the integration to be deleted.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param gcpAccount Delete a given Datadog-GCP integration.
@return interface{}
*/
func (a *GCPIntegrationApiService) DeleteGCPIntegration(ctx _context.Context, gcpAccount GcpAccount) (interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(ctx, "GCPIntegrationApiService.DeleteGCPIntegration")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/integration/gcp"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &gcpAccount
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["api_key"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarQueryParams.Add("api_key", key)
			}
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["application_key"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarQueryParams.Add("application_key", key)
			}
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
ListGCPIntegration List configured GCP integrations.
### Overview List all Datadog-GCP integrations configured in your Datadog account.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return []GcpAccount
*/
func (a *GCPIntegrationApiService) ListGCPIntegration(ctx _context.Context) ([]GcpAccount, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []GcpAccount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(ctx, "GCPIntegrationApiService.ListGCPIntegration")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/integration/gcp"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["api_key"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarQueryParams.Add("api_key", key)
			}
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["application_key"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarQueryParams.Add("application_key", key)
			}
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v []GcpAccount
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
UpdateGCPIntegration Update a GCP integration in your Datadog account.
### Overview Update a Datadog-GCP integrations host_filters and/or automute. Requires a project_id and client_email, however these fields cannot be updated. If you need to update these fields please delete and use the create (POST) endpoint. The unspecified fields will keep their original values. ### Arguments * **&#x60;project_id&#x60;** [*required*, *default* &#x3D; **None**]: Your Google Cloud project ID associated with the integration to be updated. * **&#x60;client_email&#x60;** [*required*, *default* &#x3D; **None**]: Your client email associated with the integration to be updated. * **&#x60;host_filters&#x60;** [*optional*, *default* &#x3D; **None**]: Limit the GCE instances that are pulled into Datadog by using tags. Only hosts that match one of the defined tags are imported into Datadog. * **&#x60;automute&#x60;** [*optional*, *default* &#x3D; false]: Silence monitors for expected GCE instance shutdowns.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param gcpAccount Update a Datadog-GCP integration.
@return interface{}
*/
func (a *GCPIntegrationApiService) UpdateGCPIntegration(ctx _context.Context, gcpAccount GcpAccount) (interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(ctx, "GCPIntegrationApiService.UpdateGCPIntegration")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/integration/gcp"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &gcpAccount
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["api_key"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarQueryParams.Add("api_key", key)
			}
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["application_key"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarQueryParams.Add("application_key", key)
			}
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
