// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetAgentsV2ResponseMeta Metadata for the v2 list of agents, including pagination information.
type FleetAgentsV2ResponseMeta struct {
	// Pagination details for the v2 list of agents.
	Page *FleetAgentsV2Page `json:"page,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetAgentsV2ResponseMeta instantiates a new FleetAgentsV2ResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetAgentsV2ResponseMeta() *FleetAgentsV2ResponseMeta {
	this := FleetAgentsV2ResponseMeta{}
	return &this
}

// NewFleetAgentsV2ResponseMetaWithDefaults instantiates a new FleetAgentsV2ResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetAgentsV2ResponseMetaWithDefaults() *FleetAgentsV2ResponseMeta {
	this := FleetAgentsV2ResponseMeta{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *FleetAgentsV2ResponseMeta) GetPage() FleetAgentsV2Page {
	if o == nil || o.Page == nil {
		var ret FleetAgentsV2Page
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentsV2ResponseMeta) GetPageOk() (*FleetAgentsV2Page, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *FleetAgentsV2ResponseMeta) HasPage() bool {
	return o != nil && o.Page != nil
}

// SetPage gets a reference to the given FleetAgentsV2Page and assigns it to the Page field.
func (o *FleetAgentsV2ResponseMeta) SetPage(v FleetAgentsV2Page) {
	o.Page = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetAgentsV2ResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetAgentsV2ResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Page *FleetAgentsV2Page `json:"page,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"page"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Page != nil && all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Page = all.Page

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
