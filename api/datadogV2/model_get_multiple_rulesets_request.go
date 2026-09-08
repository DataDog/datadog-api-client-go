// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetMultipleRulesetsRequest The request payload for retrieving rules for multiple rulesets in a single batch call.
type GetMultipleRulesetsRequest struct {
	// CSRF token for security, sent by browser-based clients. Ignored by the API when absent.
	AuthenticationToken *string `json:"_authentication_token,omitempty"`
	// The primary data object in the get-multiple-rulesets request, containing request attributes and resource type.
	Data *GetMultipleRulesetsRequestData `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetMultipleRulesetsRequest instantiates a new GetMultipleRulesetsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetMultipleRulesetsRequest() *GetMultipleRulesetsRequest {
	this := GetMultipleRulesetsRequest{}
	return &this
}

// NewGetMultipleRulesetsRequestWithDefaults instantiates a new GetMultipleRulesetsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetMultipleRulesetsRequestWithDefaults() *GetMultipleRulesetsRequest {
	this := GetMultipleRulesetsRequest{}
	return &this
}

// GetAuthenticationToken returns the AuthenticationToken field value if set, zero value otherwise.
func (o *GetMultipleRulesetsRequest) GetAuthenticationToken() string {
	if o == nil || o.AuthenticationToken == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationToken
}

// GetAuthenticationTokenOk returns a tuple with the AuthenticationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsRequest) GetAuthenticationTokenOk() (*string, bool) {
	if o == nil || o.AuthenticationToken == nil {
		return nil, false
	}
	return o.AuthenticationToken, true
}

// HasAuthenticationToken returns a boolean if a field has been set.
func (o *GetMultipleRulesetsRequest) HasAuthenticationToken() bool {
	return o != nil && o.AuthenticationToken != nil
}

// SetAuthenticationToken gets a reference to the given string and assigns it to the AuthenticationToken field.
func (o *GetMultipleRulesetsRequest) SetAuthenticationToken(v string) {
	o.AuthenticationToken = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetMultipleRulesetsRequest) GetData() GetMultipleRulesetsRequestData {
	if o == nil || o.Data == nil {
		var ret GetMultipleRulesetsRequestData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsRequest) GetDataOk() (*GetMultipleRulesetsRequestData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetMultipleRulesetsRequest) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given GetMultipleRulesetsRequestData and assigns it to the Data field.
func (o *GetMultipleRulesetsRequest) SetData(v GetMultipleRulesetsRequestData) {
	o.Data = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetMultipleRulesetsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AuthenticationToken != nil {
		toSerialize["_authentication_token"] = o.AuthenticationToken
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetMultipleRulesetsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthenticationToken *string                         `json:"_authentication_token,omitempty"`
		Data                *GetMultipleRulesetsRequestData `json:"data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"_authentication_token", "data"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AuthenticationToken = all.AuthenticationToken
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
