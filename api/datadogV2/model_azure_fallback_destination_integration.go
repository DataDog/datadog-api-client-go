// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureFallbackDestinationIntegration The Azure archive's integration destination.
type AzureFallbackDestinationIntegration struct {
	// A client ID.
	ClientId string `json:"clientId"`
	// A tenant ID.
	TenantId string `json:"tenantId"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewAzureFallbackDestinationIntegration instantiates a new AzureFallbackDestinationIntegration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAzureFallbackDestinationIntegration(clientId string, tenantId string) *AzureFallbackDestinationIntegration {
	this := AzureFallbackDestinationIntegration{}
	this.ClientId = clientId
	this.TenantId = tenantId
	return &this
}

// NewAzureFallbackDestinationIntegrationWithDefaults instantiates a new AzureFallbackDestinationIntegration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAzureFallbackDestinationIntegrationWithDefaults() *AzureFallbackDestinationIntegration {
	this := AzureFallbackDestinationIntegration{}
	return &this
}

// GetClientId returns the ClientId field value.
func (o *AzureFallbackDestinationIntegration) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *AzureFallbackDestinationIntegration) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value.
func (o *AzureFallbackDestinationIntegration) SetClientId(v string) {
	o.ClientId = v
}

// GetTenantId returns the TenantId field value.
func (o *AzureFallbackDestinationIntegration) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *AzureFallbackDestinationIntegration) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value.
func (o *AzureFallbackDestinationIntegration) SetTenantId(v string) {
	o.TenantId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AzureFallbackDestinationIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["clientId"] = o.ClientId
	toSerialize["tenantId"] = o.TenantId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AzureFallbackDestinationIntegration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClientId *string `json:"clientId"`
		TenantId *string `json:"tenantId"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ClientId == nil {
		return fmt.Errorf("required field clientId missing")
	}
	if all.TenantId == nil {
		return fmt.Errorf("required field tenantId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"clientId", "tenantId"})
	} else {
		return err
	}
	o.ClientId = *all.ClientId
	o.TenantId = *all.TenantId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
