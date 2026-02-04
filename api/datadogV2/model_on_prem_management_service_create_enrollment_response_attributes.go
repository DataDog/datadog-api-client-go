// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnPremManagementServiceCreateEnrollmentResponseAttributes Attributes for the created enrollment.
type OnPremManagementServiceCreateEnrollmentResponseAttributes struct {
	// The enrollment token.
	Token string `json:"token"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnPremManagementServiceCreateEnrollmentResponseAttributes instantiates a new OnPremManagementServiceCreateEnrollmentResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnPremManagementServiceCreateEnrollmentResponseAttributes(token string) *OnPremManagementServiceCreateEnrollmentResponseAttributes {
	this := OnPremManagementServiceCreateEnrollmentResponseAttributes{}
	this.Token = token
	return &this
}

// NewOnPremManagementServiceCreateEnrollmentResponseAttributesWithDefaults instantiates a new OnPremManagementServiceCreateEnrollmentResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnPremManagementServiceCreateEnrollmentResponseAttributesWithDefaults() *OnPremManagementServiceCreateEnrollmentResponseAttributes {
	this := OnPremManagementServiceCreateEnrollmentResponseAttributes{}
	return &this
}

// GetToken returns the Token field value.
func (o *OnPremManagementServiceCreateEnrollmentResponseAttributes) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *OnPremManagementServiceCreateEnrollmentResponseAttributes) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value.
func (o *OnPremManagementServiceCreateEnrollmentResponseAttributes) SetToken(v string) {
	o.Token = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnPremManagementServiceCreateEnrollmentResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnPremManagementServiceCreateEnrollmentResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Token *string `json:"token"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Token == nil {
		return fmt.Errorf("required field token missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"token"})
	} else {
		return err
	}
	o.Token = *all.Token

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
