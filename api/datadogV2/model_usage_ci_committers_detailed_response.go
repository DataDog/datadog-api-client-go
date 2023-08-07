// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageCICommittersDetailedResponse Response containing the user email for each CI Committed for each hour for a given organization.
type UsageCICommittersDetailedResponse struct {
	// An array of objects regarding hourly usage of CI Committers Detailed response.
	Data []UsageCICommittersDetailedHour `json:"data,omitempty"`
	// The object containing document metadata.
	Meta *HourlyUsageMetadata `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUsageCICommittersDetailedResponse instantiates a new UsageCICommittersDetailedResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageCICommittersDetailedResponse() *UsageCICommittersDetailedResponse {
	this := UsageCICommittersDetailedResponse{}
	return &this
}

// NewUsageCICommittersDetailedResponseWithDefaults instantiates a new UsageCICommittersDetailedResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageCICommittersDetailedResponseWithDefaults() *UsageCICommittersDetailedResponse {
	this := UsageCICommittersDetailedResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UsageCICommittersDetailedResponse) GetData() []UsageCICommittersDetailedHour {
	if o == nil || o.Data == nil {
		var ret []UsageCICommittersDetailedHour
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCICommittersDetailedResponse) GetDataOk() (*[]UsageCICommittersDetailedHour, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UsageCICommittersDetailedResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []UsageCICommittersDetailedHour and assigns it to the Data field.
func (o *UsageCICommittersDetailedResponse) SetData(v []UsageCICommittersDetailedHour) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *UsageCICommittersDetailedResponse) GetMeta() HourlyUsageMetadata {
	if o == nil || o.Meta == nil {
		var ret HourlyUsageMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCICommittersDetailedResponse) GetMetaOk() (*HourlyUsageMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *UsageCICommittersDetailedResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given HourlyUsageMetadata and assigns it to the Meta field.
func (o *UsageCICommittersDetailedResponse) SetMeta(v HourlyUsageMetadata) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageCICommittersDetailedResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageCICommittersDetailedResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data []UsageCICommittersDetailedHour `json:"data,omitempty"`
		Meta *HourlyUsageMetadata            `json:"meta,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = all.Data
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
