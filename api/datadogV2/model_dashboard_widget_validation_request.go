// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardWidgetValidationRequest Request containing dashboard widgets and their layout context.
type DashboardWidgetValidationRequest struct {
	// Layout type used to apply dashboard-specific widget layout validation.
	LayoutType DashboardWidgetValidationLayoutType `json:"layout_type"`
	// Reflow behavior used for an ordered dashboard.
	ReflowType *DashboardWidgetValidationReflowType `json:"reflow_type,omitempty"`
	// Dashboard widgets to validate.
	Widgets []map[string]interface{} `json:"widgets"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDashboardWidgetValidationRequest instantiates a new DashboardWidgetValidationRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardWidgetValidationRequest(layoutType DashboardWidgetValidationLayoutType, widgets []map[string]interface{}) *DashboardWidgetValidationRequest {
	this := DashboardWidgetValidationRequest{}
	this.LayoutType = layoutType
	this.Widgets = widgets
	return &this
}

// NewDashboardWidgetValidationRequestWithDefaults instantiates a new DashboardWidgetValidationRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardWidgetValidationRequestWithDefaults() *DashboardWidgetValidationRequest {
	this := DashboardWidgetValidationRequest{}
	return &this
}

// GetLayoutType returns the LayoutType field value.
func (o *DashboardWidgetValidationRequest) GetLayoutType() DashboardWidgetValidationLayoutType {
	if o == nil {
		var ret DashboardWidgetValidationLayoutType
		return ret
	}
	return o.LayoutType
}

// GetLayoutTypeOk returns a tuple with the LayoutType field value
// and a boolean to check if the value has been set.
func (o *DashboardWidgetValidationRequest) GetLayoutTypeOk() (*DashboardWidgetValidationLayoutType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LayoutType, true
}

// SetLayoutType sets field value.
func (o *DashboardWidgetValidationRequest) SetLayoutType(v DashboardWidgetValidationLayoutType) {
	o.LayoutType = v
}

// GetReflowType returns the ReflowType field value if set, zero value otherwise.
func (o *DashboardWidgetValidationRequest) GetReflowType() DashboardWidgetValidationReflowType {
	if o == nil || o.ReflowType == nil {
		var ret DashboardWidgetValidationReflowType
		return ret
	}
	return *o.ReflowType
}

// GetReflowTypeOk returns a tuple with the ReflowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardWidgetValidationRequest) GetReflowTypeOk() (*DashboardWidgetValidationReflowType, bool) {
	if o == nil || o.ReflowType == nil {
		return nil, false
	}
	return o.ReflowType, true
}

// HasReflowType returns a boolean if a field has been set.
func (o *DashboardWidgetValidationRequest) HasReflowType() bool {
	return o != nil && o.ReflowType != nil
}

// SetReflowType gets a reference to the given DashboardWidgetValidationReflowType and assigns it to the ReflowType field.
func (o *DashboardWidgetValidationRequest) SetReflowType(v DashboardWidgetValidationReflowType) {
	o.ReflowType = &v
}

// GetWidgets returns the Widgets field value.
func (o *DashboardWidgetValidationRequest) GetWidgets() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Widgets
}

// GetWidgetsOk returns a tuple with the Widgets field value
// and a boolean to check if the value has been set.
func (o *DashboardWidgetValidationRequest) GetWidgetsOk() (*[]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Widgets, true
}

// SetWidgets sets field value.
func (o *DashboardWidgetValidationRequest) SetWidgets(v []map[string]interface{}) {
	o.Widgets = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardWidgetValidationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["layout_type"] = o.LayoutType
	if o.ReflowType != nil {
		toSerialize["reflow_type"] = o.ReflowType
	}
	toSerialize["widgets"] = o.Widgets

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardWidgetValidationRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LayoutType *DashboardWidgetValidationLayoutType `json:"layout_type"`
		ReflowType *DashboardWidgetValidationReflowType `json:"reflow_type,omitempty"`
		Widgets    *[]map[string]interface{}            `json:"widgets"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.LayoutType == nil {
		return fmt.Errorf("required field layout_type missing")
	}
	if all.Widgets == nil {
		return fmt.Errorf("required field widgets missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"layout_type", "reflow_type", "widgets"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.LayoutType.IsValid() {
		hasInvalidField = true
	} else {
		o.LayoutType = *all.LayoutType
	}
	if all.ReflowType != nil && !all.ReflowType.IsValid() {
		hasInvalidField = true
	} else {
		o.ReflowType = all.ReflowType
	}
	o.Widgets = *all.Widgets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
