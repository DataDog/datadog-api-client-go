// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudMetricsIntegrationDataflowResponse The Elastic Cloud metrics dataflow.
type ElasticCloudMetricsIntegrationDataflowResponse struct {
	// Whether the Elastic Cloud dataflow is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Read-only collection status of a dataflow.
	Status *IntegrationAccountDataflowStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticCloudMetricsIntegrationDataflowResponse instantiates a new ElasticCloudMetricsIntegrationDataflowResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticCloudMetricsIntegrationDataflowResponse() *ElasticCloudMetricsIntegrationDataflowResponse {
	this := ElasticCloudMetricsIntegrationDataflowResponse{}
	return &this
}

// NewElasticCloudMetricsIntegrationDataflowResponseWithDefaults instantiates a new ElasticCloudMetricsIntegrationDataflowResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticCloudMetricsIntegrationDataflowResponseWithDefaults() *ElasticCloudMetricsIntegrationDataflowResponse {
	this := ElasticCloudMetricsIntegrationDataflowResponse{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ElasticCloudMetricsIntegrationDataflowResponse) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudMetricsIntegrationDataflowResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ElasticCloudMetricsIntegrationDataflowResponse) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ElasticCloudMetricsIntegrationDataflowResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ElasticCloudMetricsIntegrationDataflowResponse) GetStatus() IntegrationAccountDataflowStatus {
	if o == nil || o.Status == nil {
		var ret IntegrationAccountDataflowStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudMetricsIntegrationDataflowResponse) GetStatusOk() (*IntegrationAccountDataflowStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ElasticCloudMetricsIntegrationDataflowResponse) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given IntegrationAccountDataflowStatus and assigns it to the Status field.
func (o *ElasticCloudMetricsIntegrationDataflowResponse) SetStatus(v IntegrationAccountDataflowStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticCloudMetricsIntegrationDataflowResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticCloudMetricsIntegrationDataflowResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled *bool                             `json:"enabled,omitempty"`
		Status  *IntegrationAccountDataflowStatus `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
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
