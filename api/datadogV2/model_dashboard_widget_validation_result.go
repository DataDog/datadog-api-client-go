// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardWidgetValidationResult Validation result for one dashboard widget.
type DashboardWidgetValidationResult struct {
	// Validation error message, when the widget is invalid.
	ErrorMessage datadog.NullableString `json:"error_message"`
	// Path to the invalid value, when available.
	ErrorPath datadog.NullableString `json:"error_path"`
	// Whether the widget passed validation.
	IsValid bool `json:"is_valid"`
	// Type of the validated widget, when available.
	WidgetType datadog.NullableString `json:"widget_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDashboardWidgetValidationResult instantiates a new DashboardWidgetValidationResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardWidgetValidationResult(errorMessage datadog.NullableString, errorPath datadog.NullableString, isValid bool, widgetType datadog.NullableString) *DashboardWidgetValidationResult {
	this := DashboardWidgetValidationResult{}
	this.ErrorMessage = errorMessage
	this.ErrorPath = errorPath
	this.IsValid = isValid
	this.WidgetType = widgetType
	return &this
}

// NewDashboardWidgetValidationResultWithDefaults instantiates a new DashboardWidgetValidationResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardWidgetValidationResultWithDefaults() *DashboardWidgetValidationResult {
	this := DashboardWidgetValidationResult{}
	return &this
}

// GetErrorMessage returns the ErrorMessage field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *DashboardWidgetValidationResult) GetErrorMessage() string {
	if o == nil || o.ErrorMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardWidgetValidationResult) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// SetErrorMessage sets field value.
func (o *DashboardWidgetValidationResult) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}

// GetErrorPath returns the ErrorPath field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *DashboardWidgetValidationResult) GetErrorPath() string {
	if o == nil || o.ErrorPath.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorPath.Get()
}

// GetErrorPathOk returns a tuple with the ErrorPath field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardWidgetValidationResult) GetErrorPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorPath.Get(), o.ErrorPath.IsSet()
}

// SetErrorPath sets field value.
func (o *DashboardWidgetValidationResult) SetErrorPath(v string) {
	o.ErrorPath.Set(&v)
}

// GetIsValid returns the IsValid field value.
func (o *DashboardWidgetValidationResult) GetIsValid() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value
// and a boolean to check if the value has been set.
func (o *DashboardWidgetValidationResult) GetIsValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsValid, true
}

// SetIsValid sets field value.
func (o *DashboardWidgetValidationResult) SetIsValid(v bool) {
	o.IsValid = v
}

// GetWidgetType returns the WidgetType field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *DashboardWidgetValidationResult) GetWidgetType() string {
	if o == nil || o.WidgetType.Get() == nil {
		var ret string
		return ret
	}
	return *o.WidgetType.Get()
}

// GetWidgetTypeOk returns a tuple with the WidgetType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardWidgetValidationResult) GetWidgetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WidgetType.Get(), o.WidgetType.IsSet()
}

// SetWidgetType sets field value.
func (o *DashboardWidgetValidationResult) SetWidgetType(v string) {
	o.WidgetType.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardWidgetValidationResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["error_message"] = o.ErrorMessage.Get()
	toSerialize["error_path"] = o.ErrorPath.Get()
	toSerialize["is_valid"] = o.IsValid
	toSerialize["widget_type"] = o.WidgetType.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardWidgetValidationResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ErrorMessage datadog.NullableString `json:"error_message"`
		ErrorPath    datadog.NullableString `json:"error_path"`
		IsValid      *bool                  `json:"is_valid"`
		WidgetType   datadog.NullableString `json:"widget_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.ErrorMessage.IsSet() {
		return fmt.Errorf("required field error_message missing")
	}
	if !all.ErrorPath.IsSet() {
		return fmt.Errorf("required field error_path missing")
	}
	if all.IsValid == nil {
		return fmt.Errorf("required field is_valid missing")
	}
	if !all.WidgetType.IsSet() {
		return fmt.Errorf("required field widget_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error_message", "error_path", "is_valid", "widget_type"})
	} else {
		return err
	}
	o.ErrorMessage = all.ErrorMessage
	o.ErrorPath = all.ErrorPath
	o.IsValid = *all.IsValid
	o.WidgetType = all.WidgetType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
