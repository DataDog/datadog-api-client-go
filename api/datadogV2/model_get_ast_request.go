// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetAstRequest The request payload for parsing source code into an abstract syntax tree.
type GetAstRequest struct {
	// CSRF token for security, sent by browser-based clients. Ignored by the API when absent.
	AuthenticationToken *string `json:"_authentication_token,omitempty"`
	// The primary data object in the get-AST request.
	Data GetAstRequestData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetAstRequest instantiates a new GetAstRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetAstRequest(data GetAstRequestData) *GetAstRequest {
	this := GetAstRequest{}
	this.Data = data
	return &this
}

// NewGetAstRequestWithDefaults instantiates a new GetAstRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetAstRequestWithDefaults() *GetAstRequest {
	this := GetAstRequest{}
	return &this
}

// GetAuthenticationToken returns the AuthenticationToken field value if set, zero value otherwise.
func (o *GetAstRequest) GetAuthenticationToken() string {
	if o == nil || o.AuthenticationToken == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationToken
}

// GetAuthenticationTokenOk returns a tuple with the AuthenticationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAstRequest) GetAuthenticationTokenOk() (*string, bool) {
	if o == nil || o.AuthenticationToken == nil {
		return nil, false
	}
	return o.AuthenticationToken, true
}

// HasAuthenticationToken returns a boolean if a field has been set.
func (o *GetAstRequest) HasAuthenticationToken() bool {
	return o != nil && o.AuthenticationToken != nil
}

// SetAuthenticationToken gets a reference to the given string and assigns it to the AuthenticationToken field.
func (o *GetAstRequest) SetAuthenticationToken(v string) {
	o.AuthenticationToken = &v
}

// GetData returns the Data field value.
func (o *GetAstRequest) GetData() GetAstRequestData {
	if o == nil {
		var ret GetAstRequestData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetAstRequest) GetDataOk() (*GetAstRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *GetAstRequest) SetData(v GetAstRequestData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetAstRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AuthenticationToken != nil {
		toSerialize["_authentication_token"] = o.AuthenticationToken
	}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetAstRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthenticationToken *string            `json:"_authentication_token,omitempty"`
		Data                *GetAstRequestData `json:"data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"_authentication_token", "data"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AuthenticationToken = all.AuthenticationToken
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
