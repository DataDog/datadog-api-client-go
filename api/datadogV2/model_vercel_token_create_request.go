// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// VercelTokenCreateRequest Request to exchange a Vercel marketplace authorization code for a Datadog-managed access token.
type VercelTokenCreateRequest struct {
	// OAuth authorization code received from the Vercel marketplace flow.
	AuthGrantCode string `json:"authGrantCode"`
	// Vercel configuration identifier returned by the marketplace flow.
	VercelConfigurationId string `json:"vercelConfigurationId"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVercelTokenCreateRequest instantiates a new VercelTokenCreateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVercelTokenCreateRequest(authGrantCode string, vercelConfigurationId string) *VercelTokenCreateRequest {
	this := VercelTokenCreateRequest{}
	this.AuthGrantCode = authGrantCode
	this.VercelConfigurationId = vercelConfigurationId
	return &this
}

// NewVercelTokenCreateRequestWithDefaults instantiates a new VercelTokenCreateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVercelTokenCreateRequestWithDefaults() *VercelTokenCreateRequest {
	this := VercelTokenCreateRequest{}
	return &this
}

// GetAuthGrantCode returns the AuthGrantCode field value.
func (o *VercelTokenCreateRequest) GetAuthGrantCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AuthGrantCode
}

// GetAuthGrantCodeOk returns a tuple with the AuthGrantCode field value
// and a boolean to check if the value has been set.
func (o *VercelTokenCreateRequest) GetAuthGrantCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthGrantCode, true
}

// SetAuthGrantCode sets field value.
func (o *VercelTokenCreateRequest) SetAuthGrantCode(v string) {
	o.AuthGrantCode = v
}

// GetVercelConfigurationId returns the VercelConfigurationId field value.
func (o *VercelTokenCreateRequest) GetVercelConfigurationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.VercelConfigurationId
}

// GetVercelConfigurationIdOk returns a tuple with the VercelConfigurationId field value
// and a boolean to check if the value has been set.
func (o *VercelTokenCreateRequest) GetVercelConfigurationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VercelConfigurationId, true
}

// SetVercelConfigurationId sets field value.
func (o *VercelTokenCreateRequest) SetVercelConfigurationId(v string) {
	o.VercelConfigurationId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o VercelTokenCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["authGrantCode"] = o.AuthGrantCode
	toSerialize["vercelConfigurationId"] = o.VercelConfigurationId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *VercelTokenCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthGrantCode         *string `json:"authGrantCode"`
		VercelConfigurationId *string `json:"vercelConfigurationId"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthGrantCode == nil {
		return fmt.Errorf("required field authGrantCode missing")
	}
	if all.VercelConfigurationId == nil {
		return fmt.Errorf("required field vercelConfigurationId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"authGrantCode", "vercelConfigurationId"})
	} else {
		return err
	}
	o.AuthGrantCode = *all.AuthGrantCode
	o.VercelConfigurationId = *all.VercelConfigurationId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
