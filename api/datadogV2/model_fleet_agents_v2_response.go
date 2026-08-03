// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetAgentsV2Response Response containing a paginated list of Datadog Agents.
type FleetAgentsV2Response struct {
	// Array of agents matching the query criteria.
	Data []FleetAgentV2 `json:"data"`
	// Metadata for the v2 list of agents, including pagination information.
	Meta *FleetAgentsV2ResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetAgentsV2Response instantiates a new FleetAgentsV2Response object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetAgentsV2Response(data []FleetAgentV2) *FleetAgentsV2Response {
	this := FleetAgentsV2Response{}
	this.Data = data
	return &this
}

// NewFleetAgentsV2ResponseWithDefaults instantiates a new FleetAgentsV2Response object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetAgentsV2ResponseWithDefaults() *FleetAgentsV2Response {
	this := FleetAgentsV2Response{}
	return &this
}

// GetData returns the Data field value.
func (o *FleetAgentsV2Response) GetData() []FleetAgentV2 {
	if o == nil {
		var ret []FleetAgentV2
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FleetAgentsV2Response) GetDataOk() (*[]FleetAgentV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *FleetAgentsV2Response) SetData(v []FleetAgentV2) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *FleetAgentsV2Response) GetMeta() FleetAgentsV2ResponseMeta {
	if o == nil || o.Meta == nil {
		var ret FleetAgentsV2ResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentsV2Response) GetMetaOk() (*FleetAgentsV2ResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *FleetAgentsV2Response) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given FleetAgentsV2ResponseMeta and assigns it to the Meta field.
func (o *FleetAgentsV2Response) SetMeta(v FleetAgentsV2ResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetAgentsV2Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetAgentsV2Response) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]FleetAgentV2            `json:"data"`
		Meta *FleetAgentsV2ResponseMeta `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
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
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
