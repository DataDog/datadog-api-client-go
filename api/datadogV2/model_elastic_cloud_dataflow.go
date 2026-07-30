// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudDataflow An Elastic Cloud dataflow toggle. The set of dataflow ids is fixed by the interface schema.
type ElasticCloudDataflow struct {
	// Whether the dataflow is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Identifier of an Elastic Cloud dataflow.
	Id ElasticCloudDataflowId `json:"id"`
	// Read-only, server-computed collection status of a dataflow.
	Status *IntegrationAccountDataflowStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticCloudDataflow instantiates a new ElasticCloudDataflow object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticCloudDataflow(id ElasticCloudDataflowId) *ElasticCloudDataflow {
	this := ElasticCloudDataflow{}
	this.Id = id
	return &this
}

// NewElasticCloudDataflowWithDefaults instantiates a new ElasticCloudDataflow object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticCloudDataflowWithDefaults() *ElasticCloudDataflow {
	this := ElasticCloudDataflow{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ElasticCloudDataflow) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudDataflow) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ElasticCloudDataflow) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ElasticCloudDataflow) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetId returns the Id field value.
func (o *ElasticCloudDataflow) GetId() ElasticCloudDataflowId {
	if o == nil {
		var ret ElasticCloudDataflowId
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ElasticCloudDataflow) GetIdOk() (*ElasticCloudDataflowId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ElasticCloudDataflow) SetId(v ElasticCloudDataflowId) {
	o.Id = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ElasticCloudDataflow) GetStatus() IntegrationAccountDataflowStatus {
	if o == nil || o.Status == nil {
		var ret IntegrationAccountDataflowStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudDataflow) GetStatusOk() (*IntegrationAccountDataflowStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ElasticCloudDataflow) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given IntegrationAccountDataflowStatus and assigns it to the Status field.
func (o *ElasticCloudDataflow) SetStatus(v IntegrationAccountDataflowStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticCloudDataflow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["id"] = o.Id
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticCloudDataflow) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled *bool                             `json:"enabled,omitempty"`
		Id      *ElasticCloudDataflowId           `json:"id"`
		Status  *IntegrationAccountDataflowStatus `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "id", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	if !all.Id.IsValid() {
		hasInvalidField = true
	} else {
		o.Id = *all.Id
	}
	if all.Status != nil && all.Status.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
