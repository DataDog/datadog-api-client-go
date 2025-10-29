// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetCohortUsersRequestDataAttributes
type GetCohortUsersRequestDataAttributes struct {
	//
	DataSource *string `json:"data_source,omitempty"`
	//
	Definition *GetCohortUsersRequestDataAttributesDefinition `json:"definition,omitempty"`
	//
	Execution *int64 `json:"execution,omitempty"`
	//
	Time *GetCohortUsersRequestDataAttributesTime `json:"time,omitempty"`
	//
	UserSelection *string `json:"user_selection,omitempty"`
	//
	WindowSize *string `json:"window_size,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetCohortUsersRequestDataAttributes instantiates a new GetCohortUsersRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetCohortUsersRequestDataAttributes() *GetCohortUsersRequestDataAttributes {
	this := GetCohortUsersRequestDataAttributes{}
	return &this
}

// NewGetCohortUsersRequestDataAttributesWithDefaults instantiates a new GetCohortUsersRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetCohortUsersRequestDataAttributesWithDefaults() *GetCohortUsersRequestDataAttributes {
	this := GetCohortUsersRequestDataAttributes{}
	return &this
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *GetCohortUsersRequestDataAttributes) GetDataSource() string {
	if o == nil || o.DataSource == nil {
		var ret string
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortUsersRequestDataAttributes) GetDataSourceOk() (*string, bool) {
	if o == nil || o.DataSource == nil {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *GetCohortUsersRequestDataAttributes) HasDataSource() bool {
	return o != nil && o.DataSource != nil
}

// SetDataSource gets a reference to the given string and assigns it to the DataSource field.
func (o *GetCohortUsersRequestDataAttributes) SetDataSource(v string) {
	o.DataSource = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *GetCohortUsersRequestDataAttributes) GetDefinition() GetCohortUsersRequestDataAttributesDefinition {
	if o == nil || o.Definition == nil {
		var ret GetCohortUsersRequestDataAttributesDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortUsersRequestDataAttributes) GetDefinitionOk() (*GetCohortUsersRequestDataAttributesDefinition, bool) {
	if o == nil || o.Definition == nil {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *GetCohortUsersRequestDataAttributes) HasDefinition() bool {
	return o != nil && o.Definition != nil
}

// SetDefinition gets a reference to the given GetCohortUsersRequestDataAttributesDefinition and assigns it to the Definition field.
func (o *GetCohortUsersRequestDataAttributes) SetDefinition(v GetCohortUsersRequestDataAttributesDefinition) {
	o.Definition = &v
}

// GetExecution returns the Execution field value if set, zero value otherwise.
func (o *GetCohortUsersRequestDataAttributes) GetExecution() int64 {
	if o == nil || o.Execution == nil {
		var ret int64
		return ret
	}
	return *o.Execution
}

// GetExecutionOk returns a tuple with the Execution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortUsersRequestDataAttributes) GetExecutionOk() (*int64, bool) {
	if o == nil || o.Execution == nil {
		return nil, false
	}
	return o.Execution, true
}

// HasExecution returns a boolean if a field has been set.
func (o *GetCohortUsersRequestDataAttributes) HasExecution() bool {
	return o != nil && o.Execution != nil
}

// SetExecution gets a reference to the given int64 and assigns it to the Execution field.
func (o *GetCohortUsersRequestDataAttributes) SetExecution(v int64) {
	o.Execution = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetCohortUsersRequestDataAttributes) GetTime() GetCohortUsersRequestDataAttributesTime {
	if o == nil || o.Time == nil {
		var ret GetCohortUsersRequestDataAttributesTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortUsersRequestDataAttributes) GetTimeOk() (*GetCohortUsersRequestDataAttributesTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetCohortUsersRequestDataAttributes) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given GetCohortUsersRequestDataAttributesTime and assigns it to the Time field.
func (o *GetCohortUsersRequestDataAttributes) SetTime(v GetCohortUsersRequestDataAttributesTime) {
	o.Time = &v
}

// GetUserSelection returns the UserSelection field value if set, zero value otherwise.
func (o *GetCohortUsersRequestDataAttributes) GetUserSelection() string {
	if o == nil || o.UserSelection == nil {
		var ret string
		return ret
	}
	return *o.UserSelection
}

// GetUserSelectionOk returns a tuple with the UserSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortUsersRequestDataAttributes) GetUserSelectionOk() (*string, bool) {
	if o == nil || o.UserSelection == nil {
		return nil, false
	}
	return o.UserSelection, true
}

// HasUserSelection returns a boolean if a field has been set.
func (o *GetCohortUsersRequestDataAttributes) HasUserSelection() bool {
	return o != nil && o.UserSelection != nil
}

// SetUserSelection gets a reference to the given string and assigns it to the UserSelection field.
func (o *GetCohortUsersRequestDataAttributes) SetUserSelection(v string) {
	o.UserSelection = &v
}

// GetWindowSize returns the WindowSize field value if set, zero value otherwise.
func (o *GetCohortUsersRequestDataAttributes) GetWindowSize() string {
	if o == nil || o.WindowSize == nil {
		var ret string
		return ret
	}
	return *o.WindowSize
}

// GetWindowSizeOk returns a tuple with the WindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortUsersRequestDataAttributes) GetWindowSizeOk() (*string, bool) {
	if o == nil || o.WindowSize == nil {
		return nil, false
	}
	return o.WindowSize, true
}

// HasWindowSize returns a boolean if a field has been set.
func (o *GetCohortUsersRequestDataAttributes) HasWindowSize() bool {
	return o != nil && o.WindowSize != nil
}

// SetWindowSize gets a reference to the given string and assigns it to the WindowSize field.
func (o *GetCohortUsersRequestDataAttributes) SetWindowSize(v string) {
	o.WindowSize = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetCohortUsersRequestDataAttributes) MarshalJSON() ([]byte, error) {
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
	if o.Execution != nil {
		toSerialize["execution"] = o.Execution
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.UserSelection != nil {
		toSerialize["user_selection"] = o.UserSelection
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
func (o *GetCohortUsersRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource    *string                                        `json:"data_source,omitempty"`
		Definition    *GetCohortUsersRequestDataAttributesDefinition `json:"definition,omitempty"`
		Execution     *int64                                         `json:"execution,omitempty"`
		Time          *GetCohortUsersRequestDataAttributesTime       `json:"time,omitempty"`
		UserSelection *string                                        `json:"user_selection,omitempty"`
		WindowSize    *string                                        `json:"window_size,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_source", "definition", "execution", "time", "user_selection", "window_size"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DataSource = all.DataSource
	if all.Definition != nil && all.Definition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Definition = all.Definition
	o.Execution = all.Execution
	if all.Time != nil && all.Time.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Time = all.Time
	o.UserSelection = all.UserSelection
	o.WindowSize = all.WindowSize

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
