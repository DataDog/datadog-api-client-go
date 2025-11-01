// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyRequestDataAttributes
type SankeyRequestDataAttributes struct {
	//
	DataSource *string `json:"data_source,omitempty"`
	//
	Definition *SankeyRequestDataAttributesDefinition `json:"definition,omitempty"`
	//
	EnforcedExecutionType *string `json:"enforced_execution_type,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
	//
	Sampling *SankeyRequestDataAttributesSampling `json:"sampling,omitempty"`
	//
	Search *SankeyRequestDataAttributesSearch `json:"search,omitempty"`
	//
	Time *SankeyRequestDataAttributesTime `json:"time,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSankeyRequestDataAttributes instantiates a new SankeyRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSankeyRequestDataAttributes() *SankeyRequestDataAttributes {
	this := SankeyRequestDataAttributes{}
	return &this
}

// NewSankeyRequestDataAttributesWithDefaults instantiates a new SankeyRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSankeyRequestDataAttributesWithDefaults() *SankeyRequestDataAttributes {
	this := SankeyRequestDataAttributes{}
	return &this
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *SankeyRequestDataAttributes) GetDataSource() string {
	if o == nil || o.DataSource == nil {
		var ret string
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRequestDataAttributes) GetDataSourceOk() (*string, bool) {
	if o == nil || o.DataSource == nil {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *SankeyRequestDataAttributes) HasDataSource() bool {
	return o != nil && o.DataSource != nil
}

// SetDataSource gets a reference to the given string and assigns it to the DataSource field.
func (o *SankeyRequestDataAttributes) SetDataSource(v string) {
	o.DataSource = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *SankeyRequestDataAttributes) GetDefinition() SankeyRequestDataAttributesDefinition {
	if o == nil || o.Definition == nil {
		var ret SankeyRequestDataAttributesDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRequestDataAttributes) GetDefinitionOk() (*SankeyRequestDataAttributesDefinition, bool) {
	if o == nil || o.Definition == nil {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *SankeyRequestDataAttributes) HasDefinition() bool {
	return o != nil && o.Definition != nil
}

// SetDefinition gets a reference to the given SankeyRequestDataAttributesDefinition and assigns it to the Definition field.
func (o *SankeyRequestDataAttributes) SetDefinition(v SankeyRequestDataAttributesDefinition) {
	o.Definition = &v
}

// GetEnforcedExecutionType returns the EnforcedExecutionType field value if set, zero value otherwise.
func (o *SankeyRequestDataAttributes) GetEnforcedExecutionType() string {
	if o == nil || o.EnforcedExecutionType == nil {
		var ret string
		return ret
	}
	return *o.EnforcedExecutionType
}

// GetEnforcedExecutionTypeOk returns a tuple with the EnforcedExecutionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRequestDataAttributes) GetEnforcedExecutionTypeOk() (*string, bool) {
	if o == nil || o.EnforcedExecutionType == nil {
		return nil, false
	}
	return o.EnforcedExecutionType, true
}

// HasEnforcedExecutionType returns a boolean if a field has been set.
func (o *SankeyRequestDataAttributes) HasEnforcedExecutionType() bool {
	return o != nil && o.EnforcedExecutionType != nil
}

// SetEnforcedExecutionType gets a reference to the given string and assigns it to the EnforcedExecutionType field.
func (o *SankeyRequestDataAttributes) SetEnforcedExecutionType(v string) {
	o.EnforcedExecutionType = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *SankeyRequestDataAttributes) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRequestDataAttributes) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *SankeyRequestDataAttributes) HasRequestId() bool {
	return o != nil && o.RequestId != nil
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *SankeyRequestDataAttributes) SetRequestId(v string) {
	o.RequestId = &v
}

// GetSampling returns the Sampling field value if set, zero value otherwise.
func (o *SankeyRequestDataAttributes) GetSampling() SankeyRequestDataAttributesSampling {
	if o == nil || o.Sampling == nil {
		var ret SankeyRequestDataAttributesSampling
		return ret
	}
	return *o.Sampling
}

// GetSamplingOk returns a tuple with the Sampling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRequestDataAttributes) GetSamplingOk() (*SankeyRequestDataAttributesSampling, bool) {
	if o == nil || o.Sampling == nil {
		return nil, false
	}
	return o.Sampling, true
}

// HasSampling returns a boolean if a field has been set.
func (o *SankeyRequestDataAttributes) HasSampling() bool {
	return o != nil && o.Sampling != nil
}

// SetSampling gets a reference to the given SankeyRequestDataAttributesSampling and assigns it to the Sampling field.
func (o *SankeyRequestDataAttributes) SetSampling(v SankeyRequestDataAttributesSampling) {
	o.Sampling = &v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *SankeyRequestDataAttributes) GetSearch() SankeyRequestDataAttributesSearch {
	if o == nil || o.Search == nil {
		var ret SankeyRequestDataAttributesSearch
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRequestDataAttributes) GetSearchOk() (*SankeyRequestDataAttributesSearch, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *SankeyRequestDataAttributes) HasSearch() bool {
	return o != nil && o.Search != nil
}

// SetSearch gets a reference to the given SankeyRequestDataAttributesSearch and assigns it to the Search field.
func (o *SankeyRequestDataAttributes) SetSearch(v SankeyRequestDataAttributesSearch) {
	o.Search = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SankeyRequestDataAttributes) GetTime() SankeyRequestDataAttributesTime {
	if o == nil || o.Time == nil {
		var ret SankeyRequestDataAttributesTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRequestDataAttributes) GetTimeOk() (*SankeyRequestDataAttributesTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SankeyRequestDataAttributes) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given SankeyRequestDataAttributesTime and assigns it to the Time field.
func (o *SankeyRequestDataAttributes) SetTime(v SankeyRequestDataAttributesTime) {
	o.Time = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SankeyRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DataSource != nil {
		toSerialize["data_source"] = o.DataSource
	}
	if o.Definition != nil {
		toSerialize["definition"] = o.Definition
	}
	if o.EnforcedExecutionType != nil {
		toSerialize["enforced_execution_type"] = o.EnforcedExecutionType
	}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	if o.Sampling != nil {
		toSerialize["sampling"] = o.Sampling
	}
	if o.Search != nil {
		toSerialize["search"] = o.Search
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SankeyRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource            *string                                `json:"data_source,omitempty"`
		Definition            *SankeyRequestDataAttributesDefinition `json:"definition,omitempty"`
		EnforcedExecutionType *string                                `json:"enforced_execution_type,omitempty"`
		RequestId             *string                                `json:"request_id,omitempty"`
		Sampling              *SankeyRequestDataAttributesSampling   `json:"sampling,omitempty"`
		Search                *SankeyRequestDataAttributesSearch     `json:"search,omitempty"`
		Time                  *SankeyRequestDataAttributesTime       `json:"time,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_source", "definition", "enforced_execution_type", "request_id", "sampling", "search", "time"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DataSource = all.DataSource
	if all.Definition != nil && all.Definition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Definition = all.Definition
	o.EnforcedExecutionType = all.EnforcedExecutionType
	o.RequestId = all.RequestId
	if all.Sampling != nil && all.Sampling.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sampling = all.Sampling
	if all.Search != nil && all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = all.Search
	if all.Time != nil && all.Time.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Time = all.Time

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
