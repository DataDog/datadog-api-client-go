// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayProviderResponse Response containing one AI Gateway provider.
type AiGatewayProviderResponse struct {
	// AI Gateway provider resource.
	Data AiGatewayProviderResource `json:"data"`
	// Model catalog entries created with the provider.
	Included []AiGatewayModelCatalogEntryResource `json:"included,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayProviderResponse instantiates a new AiGatewayProviderResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayProviderResponse(data AiGatewayProviderResource) *AiGatewayProviderResponse {
	this := AiGatewayProviderResponse{}
	this.Data = data
	return &this
}

// NewAiGatewayProviderResponseWithDefaults instantiates a new AiGatewayProviderResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayProviderResponseWithDefaults() *AiGatewayProviderResponse {
	this := AiGatewayProviderResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *AiGatewayProviderResponse) GetData() AiGatewayProviderResource {
	if o == nil {
		var ret AiGatewayProviderResource
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderResponse) GetDataOk() (*AiGatewayProviderResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *AiGatewayProviderResponse) SetData(v AiGatewayProviderResource) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AiGatewayProviderResponse) GetIncluded() []AiGatewayModelCatalogEntryResource {
	if o == nil || o.Included == nil {
		var ret []AiGatewayModelCatalogEntryResource
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderResponse) GetIncludedOk() (*[]AiGatewayModelCatalogEntryResource, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AiGatewayProviderResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []AiGatewayModelCatalogEntryResource and assigns it to the Included field.
func (o *AiGatewayProviderResponse) SetIncluded(v []AiGatewayModelCatalogEntryResource) {
	o.Included = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayProviderResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *AiGatewayProviderResource           `json:"data"`
		Included []AiGatewayModelCatalogEntryResource `json:"included,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data
	o.Included = all.Included

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
