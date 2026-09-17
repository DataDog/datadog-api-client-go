// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioIntegrationAccountsResponse Response payload for a list of Twilio integration accounts.
type TwilioIntegrationAccountsResponse struct {
	// List of Twilio integration accounts.
	Data []TwilioIntegrationAccountResponseData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTwilioIntegrationAccountsResponse instantiates a new TwilioIntegrationAccountsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTwilioIntegrationAccountsResponse(data []TwilioIntegrationAccountResponseData) *TwilioIntegrationAccountsResponse {
	this := TwilioIntegrationAccountsResponse{}
	this.Data = data
	return &this
}

// NewTwilioIntegrationAccountsResponseWithDefaults instantiates a new TwilioIntegrationAccountsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTwilioIntegrationAccountsResponseWithDefaults() *TwilioIntegrationAccountsResponse {
	this := TwilioIntegrationAccountsResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *TwilioIntegrationAccountsResponse) GetData() []TwilioIntegrationAccountResponseData {
	if o == nil {
		var ret []TwilioIntegrationAccountResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationAccountsResponse) GetDataOk() (*[]TwilioIntegrationAccountResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *TwilioIntegrationAccountsResponse) SetData(v []TwilioIntegrationAccountResponseData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TwilioIntegrationAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TwilioIntegrationAccountsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]TwilioIntegrationAccountResponseData `json:"data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
	o.Data = *all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
