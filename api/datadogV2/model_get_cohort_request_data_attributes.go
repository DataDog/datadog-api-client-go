// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetCohortRequestDataAttributes
type GetCohortRequestDataAttributes struct {
	//
	DataSource *string `json:"data_source,omitempty"`
	//
	Definition *GetCohortRequestDataAttributesDefinition `json:"definition,omitempty"`
	//
	EnforcedExecutionType *string `json:"enforced_execution_type,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
	//
	Time *GetCohortRequestDataAttributesTime `json:"time,omitempty"`
	//
	WindowSize *string `json:"window_size,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetCohortRequestDataAttributes instantiates a new GetCohortRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetCohortRequestDataAttributes() *GetCohortRequestDataAttributes {
	this := GetCohortRequestDataAttributes{}
	return &this
}

// NewGetCohortRequestDataAttributesWithDefaults instantiates a new GetCohortRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetCohortRequestDataAttributesWithDefaults() *GetCohortRequestDataAttributes {
	this := GetCohortRequestDataAttributes{}
	return &this
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *GetCohortRequestDataAttributes) GetDataSource() string {
	if o == nil || o.DataSource == nil {
		var ret string
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortRequestDataAttributes) GetDataSourceOk() (*string, bool) {
	if o == nil || o.DataSource == nil {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *GetCohortRequestDataAttributes) HasDataSource() bool {
	return o != nil && o.DataSource != nil
}

// SetDataSource gets a reference to the given string and assigns it to the DataSource field.
func (o *GetCohortRequestDataAttributes) SetDataSource(v string) {
	o.DataSource = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *GetCohortRequestDataAttributes) GetDefinition() GetCohortRequestDataAttributesDefinition {
	if o == nil || o.Definition == nil {
		var ret GetCohortRequestDataAttributesDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortRequestDataAttributes) GetDefinitionOk() (*GetCohortRequestDataAttributesDefinition, bool) {
	if o == nil || o.Definition == nil {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *GetCohortRequestDataAttributes) HasDefinition() bool {
	return o != nil && o.Definition != nil
}

// SetDefinition gets a reference to the given GetCohortRequestDataAttributesDefinition and assigns it to the Definition field.
func (o *GetCohortRequestDataAttributes) SetDefinition(v GetCohortRequestDataAttributesDefinition) {
	o.Definition = &v
}

// GetEnforcedExecutionType returns the EnforcedExecutionType field value if set, zero value otherwise.
func (o *GetCohortRequestDataAttributes) GetEnforcedExecutionType() string {
	if o == nil || o.EnforcedExecutionType == nil {
		var ret string
		return ret
	}
	return *o.EnforcedExecutionType
}

// GetEnforcedExecutionTypeOk returns a tuple with the EnforcedExecutionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortRequestDataAttributes) GetEnforcedExecutionTypeOk() (*string, bool) {
	if o == nil || o.EnforcedExecutionType == nil {
		return nil, false
	}
	return o.EnforcedExecutionType, true
}

// HasEnforcedExecutionType returns a boolean if a field has been set.
func (o *GetCohortRequestDataAttributes) HasEnforcedExecutionType() bool {
	return o != nil && o.EnforcedExecutionType != nil
}

// SetEnforcedExecutionType gets a reference to the given string and assigns it to the EnforcedExecutionType field.
func (o *GetCohortRequestDataAttributes) SetEnforcedExecutionType(v string) {
	o.EnforcedExecutionType = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *GetCohortRequestDataAttributes) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortRequestDataAttributes) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *GetCohortRequestDataAttributes) HasRequestId() bool {
	return o != nil && o.RequestId != nil
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *GetCohortRequestDataAttributes) SetRequestId(v string) {
	o.RequestId = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetCohortRequestDataAttributes) GetTime() GetCohortRequestDataAttributesTime {
	if o == nil || o.Time == nil {
		var ret GetCohortRequestDataAttributesTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortRequestDataAttributes) GetTimeOk() (*GetCohortRequestDataAttributesTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetCohortRequestDataAttributes) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given GetCohortRequestDataAttributesTime and assigns it to the Time field.
func (o *GetCohortRequestDataAttributes) SetTime(v GetCohortRequestDataAttributesTime) {
	o.Time = &v
}

// GetWindowSize returns the WindowSize field value if set, zero value otherwise.
func (o *GetCohortRequestDataAttributes) GetWindowSize() string {
	if o == nil || o.WindowSize == nil {
		var ret string
		return ret
	}
	return *o.WindowSize
}

// GetWindowSizeOk returns a tuple with the WindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortRequestDataAttributes) GetWindowSizeOk() (*string, bool) {
	if o == nil || o.WindowSize == nil {
		return nil, false
	}
	return o.WindowSize, true
}

// HasWindowSize returns a boolean if a field has been set.
func (o *GetCohortRequestDataAttributes) HasWindowSize() bool {
	return o != nil && o.WindowSize != nil
}

// SetWindowSize gets a reference to the given string and assigns it to the WindowSize field.
func (o *GetCohortRequestDataAttributes) SetWindowSize(v string) {
	o.WindowSize = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetCohortRequestDataAttributes) MarshalJSON() ([]byte, error) {
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
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.WindowSize != nil {
		toSerialize["window_size"] = o.WindowSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetCohortRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource            *string                                   `json:"data_source,omitempty"`
		Definition            *GetCohortRequestDataAttributesDefinition `json:"definition,omitempty"`
		EnforcedExecutionType *string                                   `json:"enforced_execution_type,omitempty"`
		RequestId             *string                                   `json:"request_id,omitempty"`
		Time                  *GetCohortRequestDataAttributesTime       `json:"time,omitempty"`
		WindowSize            *string                                   `json:"window_size,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_source", "definition", "enforced_execution_type", "request_id", "time", "window_size"})
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
	if all.Time != nil && all.Time.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Time = all.Time
	o.WindowSize = all.WindowSize

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
