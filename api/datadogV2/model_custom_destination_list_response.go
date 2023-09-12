// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationListResponse Response containing information about multiple custom destinations.
type CustomDestinationListResponse struct {
	// Array of returned custom destinations.
	Data []CustomDestinationWithId `json:"data"`
	// The metadata relevant to your configuration or recent request.
	Meta *CustomDestinationMetadata `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewCustomDestinationListResponse instantiates a new CustomDestinationListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomDestinationListResponse(data []CustomDestinationWithId) *CustomDestinationListResponse {
	this := CustomDestinationListResponse{}
	this.Data = data
	return &this
}

// NewCustomDestinationListResponseWithDefaults instantiates a new CustomDestinationListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomDestinationListResponseWithDefaults() *CustomDestinationListResponse {
	this := CustomDestinationListResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *CustomDestinationListResponse) GetData() []CustomDestinationWithId {
	if o == nil {
		var ret []CustomDestinationWithId
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationListResponse) GetDataOk() (*[]CustomDestinationWithId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *CustomDestinationListResponse) SetData(v []CustomDestinationWithId) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CustomDestinationListResponse) GetMeta() CustomDestinationMetadata {
	if o == nil || o.Meta == nil {
		var ret CustomDestinationMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDestinationListResponse) GetMetaOk() (*CustomDestinationMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CustomDestinationListResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given CustomDestinationMetadata and assigns it to the Meta field.
func (o *CustomDestinationListResponse) SetMeta(v CustomDestinationMetadata) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomDestinationListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomDestinationListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]CustomDestinationWithId `json:"data"`
		Meta *CustomDestinationMetadata `json:"meta,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
