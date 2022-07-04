// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadog

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"os"
)

// OrganizationsApiService OrganizationsApi service.
type OrganizationsApiService service

type apiUploadIdPMetadataRequest struct {
	ctx        _context.Context
	ApiService *OrganizationsApiService
	idpFile    **os.File
}

// UploadIdPMetadataOptionalParameters holds optional parameters for UploadIdPMetadata.
type UploadIdPMetadataOptionalParameters struct {
	IdpFile **os.File
}

// NewUploadIdPMetadataOptionalParameters creates an empty struct for parameters.
func NewUploadIdPMetadataOptionalParameters() *UploadIdPMetadataOptionalParameters {
	this := UploadIdPMetadataOptionalParameters{}
	return &this
}

// WithIdpFile sets the corresponding parameter name and returns the struct.
func (r *UploadIdPMetadataOptionalParameters) WithIdpFile(idpFile *os.File) *UploadIdPMetadataOptionalParameters {
	r.IdpFile = &idpFile
	return r
}

func (a *OrganizationsApiService) buildUploadIdPMetadataRequest(ctx _context.Context, o ...UploadIdPMetadataOptionalParameters) (apiUploadIdPMetadataRequest, error) {
	req := apiUploadIdPMetadataRequest{
		ApiService: a,
		ctx:        ctx,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type UploadIdPMetadataOptionalParameters is allowed")
	}

	if o != nil {
		req.idpFile = o[0].IdpFile
	}
	return req, nil
}

// UploadIdPMetadata Upload IdP metadata.
// Endpoint for uploading IdP metadata for SAML setup.
//
// Use this endpoint to upload or replace IdP metadata for SAML login configuration.
func (a *OrganizationsApiService) UploadIdPMetadata(ctx _context.Context, o ...UploadIdPMetadataOptionalParameters) (*_nethttp.Response, error) {
	req, err := a.buildUploadIdPMetadataRequest(ctx, o...)
	if err != nil {
		return nil, err
	}

	return req.ApiService.uploadIdPMetadataExecute(req)
}

// uploadIdPMetadataExecute executes the request.
func (a *OrganizationsApiService) uploadIdPMetadataExecute(r apiUploadIdPMetadataRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPost
		localVarPostBody   interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.UploadIdPMetadata")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/saml_configurations/idp_metadata"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	formFile := FormFile{}
	formFile.formFileName = "idp_file"
	var localVarFile *os.File
	if r.idpFile != nil {
		localVarFile = *r.idpFile
	}
	if localVarFile != nil {
		fbs, _ := _ioutil.ReadAll(localVarFile)
		formFile.fileBytes = fbs
		formFile.fileName = localVarFile.Name()
		localVarFile.Close()
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-API-KEY"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["appKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-APPLICATION-KEY"] = key
			}
		}
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, &formFile)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
