// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsSankeyRequestAttributes Attributes for a Sankey request.
type ProductAnalyticsSankeyRequestAttributes struct {
	// The data source for the Sankey query.
	DataSource string `json:"data_source"`
	// Sankey visualization definition. Set either `source` or `target`, not both.
	// Use `source` for forward flow (where do users go after this page?) or
	// `target` for backward flow (where did users come from?).
	Definition ProductAnalyticsSankeyDefinition `json:"definition"`
	// Override the query execution strategy.
	EnforcedExecutionType *ProductAnalyticsExecutionType `json:"enforced_execution_type,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
	// Sampling configuration.
	Sampling *ProductAnalyticsSampling `json:"sampling,omitempty"`
	// Search parameters for a Sankey query.
	Search ProductAnalyticsSankeySearch `json:"search"`
	// Time range for the Sankey query.
	Time ProductAnalyticsSankeyTime `json:"time"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsSankeyRequestAttributes instantiates a new ProductAnalyticsSankeyRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsSankeyRequestAttributes(dataSource string, definition ProductAnalyticsSankeyDefinition, search ProductAnalyticsSankeySearch, time ProductAnalyticsSankeyTime) *ProductAnalyticsSankeyRequestAttributes {
	this := ProductAnalyticsSankeyRequestAttributes{}
	this.DataSource = dataSource
	this.Definition = definition
	this.Search = search
	this.Time = time
	return &this
}

// NewProductAnalyticsSankeyRequestAttributesWithDefaults instantiates a new ProductAnalyticsSankeyRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsSankeyRequestAttributesWithDefaults() *ProductAnalyticsSankeyRequestAttributes {
	this := ProductAnalyticsSankeyRequestAttributes{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *ProductAnalyticsSankeyRequestAttributes) GetDataSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyRequestAttributes) GetDataSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *ProductAnalyticsSankeyRequestAttributes) SetDataSource(v string) {
	o.DataSource = v
}

// GetDefinition returns the Definition field value.
func (o *ProductAnalyticsSankeyRequestAttributes) GetDefinition() ProductAnalyticsSankeyDefinition {
	if o == nil {
		var ret ProductAnalyticsSankeyDefinition
		return ret
	}
	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyRequestAttributes) GetDefinitionOk() (*ProductAnalyticsSankeyDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value.
func (o *ProductAnalyticsSankeyRequestAttributes) SetDefinition(v ProductAnalyticsSankeyDefinition) {
	o.Definition = v
}

// GetEnforcedExecutionType returns the EnforcedExecutionType field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyRequestAttributes) GetEnforcedExecutionType() ProductAnalyticsExecutionType {
	if o == nil || o.EnforcedExecutionType == nil {
		var ret ProductAnalyticsExecutionType
		return ret
	}
	return *o.EnforcedExecutionType
}

// GetEnforcedExecutionTypeOk returns a tuple with the EnforcedExecutionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyRequestAttributes) GetEnforcedExecutionTypeOk() (*ProductAnalyticsExecutionType, bool) {
	if o == nil || o.EnforcedExecutionType == nil {
		return nil, false
	}
	return o.EnforcedExecutionType, true
}

// HasEnforcedExecutionType returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyRequestAttributes) HasEnforcedExecutionType() bool {
	return o != nil && o.EnforcedExecutionType != nil
}

// SetEnforcedExecutionType gets a reference to the given ProductAnalyticsExecutionType and assigns it to the EnforcedExecutionType field.
func (o *ProductAnalyticsSankeyRequestAttributes) SetEnforcedExecutionType(v ProductAnalyticsExecutionType) {
	o.EnforcedExecutionType = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyRequestAttributes) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyRequestAttributes) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyRequestAttributes) HasRequestId() bool {
	return o != nil && o.RequestId != nil
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ProductAnalyticsSankeyRequestAttributes) SetRequestId(v string) {
	o.RequestId = &v
}

// GetSampling returns the Sampling field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyRequestAttributes) GetSampling() ProductAnalyticsSampling {
	if o == nil || o.Sampling == nil {
		var ret ProductAnalyticsSampling
		return ret
	}
	return *o.Sampling
}

// GetSamplingOk returns a tuple with the Sampling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyRequestAttributes) GetSamplingOk() (*ProductAnalyticsSampling, bool) {
	if o == nil || o.Sampling == nil {
		return nil, false
	}
	return o.Sampling, true
}

// HasSampling returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyRequestAttributes) HasSampling() bool {
	return o != nil && o.Sampling != nil
}

// SetSampling gets a reference to the given ProductAnalyticsSampling and assigns it to the Sampling field.
func (o *ProductAnalyticsSankeyRequestAttributes) SetSampling(v ProductAnalyticsSampling) {
	o.Sampling = &v
}

// GetSearch returns the Search field value.
func (o *ProductAnalyticsSankeyRequestAttributes) GetSearch() ProductAnalyticsSankeySearch {
	if o == nil {
		var ret ProductAnalyticsSankeySearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyRequestAttributes) GetSearchOk() (*ProductAnalyticsSankeySearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *ProductAnalyticsSankeyRequestAttributes) SetSearch(v ProductAnalyticsSankeySearch) {
	o.Search = v
}

// GetTime returns the Time field value.
func (o *ProductAnalyticsSankeyRequestAttributes) GetTime() ProductAnalyticsSankeyTime {
	if o == nil {
		var ret ProductAnalyticsSankeyTime
		return ret
	}
	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyRequestAttributes) GetTimeOk() (*ProductAnalyticsSankeyTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value.
func (o *ProductAnalyticsSankeyRequestAttributes) SetTime(v ProductAnalyticsSankeyTime) {
	o.Time = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsSankeyRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["definition"] = o.Definition
	if o.EnforcedExecutionType != nil {
		toSerialize["enforced_execution_type"] = o.EnforcedExecutionType
	}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	if o.Sampling != nil {
		toSerialize["sampling"] = o.Sampling
	}
	toSerialize["search"] = o.Search
	toSerialize["time"] = o.Time

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsSankeyRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource            *string                           `json:"data_source"`
		Definition            *ProductAnalyticsSankeyDefinition `json:"definition"`
		EnforcedExecutionType *ProductAnalyticsExecutionType    `json:"enforced_execution_type,omitempty"`
		RequestId             *string                           `json:"request_id,omitempty"`
		Sampling              *ProductAnalyticsSampling         `json:"sampling,omitempty"`
		Search                *ProductAnalyticsSankeySearch     `json:"search"`
		Time                  *ProductAnalyticsSankeyTime       `json:"time"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Definition == nil {
		return fmt.Errorf("required field definition missing")
	}
	if all.Search == nil {
		return fmt.Errorf("required field search missing")
	}
	if all.Time == nil {
		return fmt.Errorf("required field time missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_source", "definition", "enforced_execution_type", "request_id", "sampling", "search", "time"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DataSource = *all.DataSource
	if all.Definition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Definition = *all.Definition
	if all.EnforcedExecutionType != nil && !all.EnforcedExecutionType.IsValid() {
		hasInvalidField = true
	} else {
		o.EnforcedExecutionType = all.EnforcedExecutionType
	}
	o.RequestId = all.RequestId
	if all.Sampling != nil && all.Sampling.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sampling = all.Sampling
	if all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = *all.Search
	if all.Time.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Time = *all.Time

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
