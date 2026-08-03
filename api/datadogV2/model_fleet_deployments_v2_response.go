// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetDeploymentsV2Response Response containing a paginated list of deployments.
type FleetDeploymentsV2Response struct {
	// Array of deployments matching the query criteria.
	Data []FleetDeploymentV2 `json:"data"`
	// Metadata for the v2 list of deployments, including pagination information.
	Meta *FleetDeploymentsV2ResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetDeploymentsV2Response instantiates a new FleetDeploymentsV2Response object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetDeploymentsV2Response(data []FleetDeploymentV2) *FleetDeploymentsV2Response {
	this := FleetDeploymentsV2Response{}
	this.Data = data
	return &this
}

// NewFleetDeploymentsV2ResponseWithDefaults instantiates a new FleetDeploymentsV2Response object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetDeploymentsV2ResponseWithDefaults() *FleetDeploymentsV2Response {
	this := FleetDeploymentsV2Response{}
	return &this
}

// GetData returns the Data field value.
func (o *FleetDeploymentsV2Response) GetData() []FleetDeploymentV2 {
	if o == nil {
		var ret []FleetDeploymentV2
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FleetDeploymentsV2Response) GetDataOk() (*[]FleetDeploymentV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *FleetDeploymentsV2Response) SetData(v []FleetDeploymentV2) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *FleetDeploymentsV2Response) GetMeta() FleetDeploymentsV2ResponseMeta {
	if o == nil || o.Meta == nil {
		var ret FleetDeploymentsV2ResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentsV2Response) GetMetaOk() (*FleetDeploymentsV2ResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *FleetDeploymentsV2Response) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given FleetDeploymentsV2ResponseMeta and assigns it to the Meta field.
func (o *FleetDeploymentsV2Response) SetMeta(v FleetDeploymentsV2ResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetDeploymentsV2Response) MarshalJSON() ([]byte, error) {
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
func (o *FleetDeploymentsV2Response) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]FleetDeploymentV2            `json:"data"`
		Meta *FleetDeploymentsV2ResponseMeta `json:"meta,omitempty"`
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
