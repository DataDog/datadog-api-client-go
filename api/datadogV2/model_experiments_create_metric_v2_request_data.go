// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateMetricV2RequestData Metric resource to create.
type ExperimentsCreateMetricV2RequestData struct {
	// Configuration for the new metric. Supply either numerator_aggregation or percentile_aggregation. A denominator_aggregation requires numerator_aggregation. Omit unused aggregation fields; do not send them as null.
	Attributes ExperimentsCreateMetricV2RequestDataAttributes `json:"attributes"`
	// The metric resource type.
	Type MetricType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsCreateMetricV2RequestData instantiates a new ExperimentsCreateMetricV2RequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCreateMetricV2RequestData(attributes ExperimentsCreateMetricV2RequestDataAttributes, typeVar MetricType) *ExperimentsCreateMetricV2RequestData {
	this := ExperimentsCreateMetricV2RequestData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewExperimentsCreateMetricV2RequestDataWithDefaults instantiates a new ExperimentsCreateMetricV2RequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCreateMetricV2RequestDataWithDefaults() *ExperimentsCreateMetricV2RequestData {
	this := ExperimentsCreateMetricV2RequestData{}
	var typeVar MetricType = METRICTYPE_METRICS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *ExperimentsCreateMetricV2RequestData) GetAttributes() ExperimentsCreateMetricV2RequestDataAttributes {
	if o == nil {
		var ret ExperimentsCreateMetricV2RequestDataAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateMetricV2RequestData) GetAttributesOk() (*ExperimentsCreateMetricV2RequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ExperimentsCreateMetricV2RequestData) SetAttributes(v ExperimentsCreateMetricV2RequestDataAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *ExperimentsCreateMetricV2RequestData) GetType() MetricType {
	if o == nil {
		var ret MetricType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateMetricV2RequestData) GetTypeOk() (*MetricType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ExperimentsCreateMetricV2RequestData) SetType(v MetricType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCreateMetricV2RequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCreateMetricV2RequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ExperimentsCreateMetricV2RequestDataAttributes `json:"attributes"`
		Type       *MetricType                                     `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Attributes = *all.Attributes
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
