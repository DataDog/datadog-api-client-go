// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WebIntegrationAccountsResponse Response containing a list of web integration accounts.
type WebIntegrationAccountsResponse struct {
	//
	Data []WebIntegrationAccountResponseData `json:"data,omitempty"`
	// The name of the integration.
	IntegrationName *string `json:"integration_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWebIntegrationAccountsResponse instantiates a new WebIntegrationAccountsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWebIntegrationAccountsResponse() *WebIntegrationAccountsResponse {
	this := WebIntegrationAccountsResponse{}
	return &this
}

// NewWebIntegrationAccountsResponseWithDefaults instantiates a new WebIntegrationAccountsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWebIntegrationAccountsResponseWithDefaults() *WebIntegrationAccountsResponse {
	this := WebIntegrationAccountsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *WebIntegrationAccountsResponse) GetData() []WebIntegrationAccountResponseData {
	if o == nil || o.Data == nil {
		var ret []WebIntegrationAccountResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebIntegrationAccountsResponse) GetDataOk() (*[]WebIntegrationAccountResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WebIntegrationAccountsResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []WebIntegrationAccountResponseData and assigns it to the Data field.
func (o *WebIntegrationAccountsResponse) SetData(v []WebIntegrationAccountResponseData) {
	o.Data = v
}

// GetIntegrationName returns the IntegrationName field value if set, zero value otherwise.
func (o *WebIntegrationAccountsResponse) GetIntegrationName() string {
	if o == nil || o.IntegrationName == nil {
		var ret string
		return ret
	}
	return *o.IntegrationName
}

// GetIntegrationNameOk returns a tuple with the IntegrationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebIntegrationAccountsResponse) GetIntegrationNameOk() (*string, bool) {
	if o == nil || o.IntegrationName == nil {
		return nil, false
	}
	return o.IntegrationName, true
}

// HasIntegrationName returns a boolean if a field has been set.
func (o *WebIntegrationAccountsResponse) HasIntegrationName() bool {
	return o != nil && o.IntegrationName != nil
}

// SetIntegrationName gets a reference to the given string and assigns it to the IntegrationName field.
func (o *WebIntegrationAccountsResponse) SetIntegrationName(v string) {
	o.IntegrationName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WebIntegrationAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.IntegrationName != nil {
		toSerialize["integration_name"] = o.IntegrationName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WebIntegrationAccountsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data            []WebIntegrationAccountResponseData `json:"data,omitempty"`
		IntegrationName *string                             `json:"integration_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "integration_name"})
	} else {
		return err
	}
	o.Data = all.Data
	o.IntegrationName = all.IntegrationName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
