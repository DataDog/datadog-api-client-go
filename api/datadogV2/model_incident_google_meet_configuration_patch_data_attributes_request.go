// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentGoogleMeetConfigurationPatchDataAttributesRequest Attributes for patching a Google Meet configuration. All fields are optional.
type IncidentGoogleMeetConfigurationPatchDataAttributesRequest struct {
	// Whether to allow manual meeting creation.
	AllowManualMeetingCreation *bool `json:"allow_manual_meeting_creation,omitempty"`
	// Whether to auto-summarize meetings.
	AutoSummarize *bool `json:"auto_summarize,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentGoogleMeetConfigurationPatchDataAttributesRequest instantiates a new IncidentGoogleMeetConfigurationPatchDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentGoogleMeetConfigurationPatchDataAttributesRequest() *IncidentGoogleMeetConfigurationPatchDataAttributesRequest {
	this := IncidentGoogleMeetConfigurationPatchDataAttributesRequest{}
	return &this
}

// NewIncidentGoogleMeetConfigurationPatchDataAttributesRequestWithDefaults instantiates a new IncidentGoogleMeetConfigurationPatchDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentGoogleMeetConfigurationPatchDataAttributesRequestWithDefaults() *IncidentGoogleMeetConfigurationPatchDataAttributesRequest {
	this := IncidentGoogleMeetConfigurationPatchDataAttributesRequest{}
	return &this
}

// GetAllowManualMeetingCreation returns the AllowManualMeetingCreation field value if set, zero value otherwise.
func (o *IncidentGoogleMeetConfigurationPatchDataAttributesRequest) GetAllowManualMeetingCreation() bool {
	if o == nil || o.AllowManualMeetingCreation == nil {
		var ret bool
		return ret
	}
	return *o.AllowManualMeetingCreation
}

// GetAllowManualMeetingCreationOk returns a tuple with the AllowManualMeetingCreation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentGoogleMeetConfigurationPatchDataAttributesRequest) GetAllowManualMeetingCreationOk() (*bool, bool) {
	if o == nil || o.AllowManualMeetingCreation == nil {
		return nil, false
	}
	return o.AllowManualMeetingCreation, true
}

// HasAllowManualMeetingCreation returns a boolean if a field has been set.
func (o *IncidentGoogleMeetConfigurationPatchDataAttributesRequest) HasAllowManualMeetingCreation() bool {
	return o != nil && o.AllowManualMeetingCreation != nil
}

// SetAllowManualMeetingCreation gets a reference to the given bool and assigns it to the AllowManualMeetingCreation field.
func (o *IncidentGoogleMeetConfigurationPatchDataAttributesRequest) SetAllowManualMeetingCreation(v bool) {
	o.AllowManualMeetingCreation = &v
}

// GetAutoSummarize returns the AutoSummarize field value if set, zero value otherwise.
func (o *IncidentGoogleMeetConfigurationPatchDataAttributesRequest) GetAutoSummarize() bool {
	if o == nil || o.AutoSummarize == nil {
		var ret bool
		return ret
	}
	return *o.AutoSummarize
}

// GetAutoSummarizeOk returns a tuple with the AutoSummarize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentGoogleMeetConfigurationPatchDataAttributesRequest) GetAutoSummarizeOk() (*bool, bool) {
	if o == nil || o.AutoSummarize == nil {
		return nil, false
	}
	return o.AutoSummarize, true
}

// HasAutoSummarize returns a boolean if a field has been set.
func (o *IncidentGoogleMeetConfigurationPatchDataAttributesRequest) HasAutoSummarize() bool {
	return o != nil && o.AutoSummarize != nil
}

// SetAutoSummarize gets a reference to the given bool and assigns it to the AutoSummarize field.
func (o *IncidentGoogleMeetConfigurationPatchDataAttributesRequest) SetAutoSummarize(v bool) {
	o.AutoSummarize = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentGoogleMeetConfigurationPatchDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowManualMeetingCreation != nil {
		toSerialize["allow_manual_meeting_creation"] = o.AllowManualMeetingCreation
	}
	if o.AutoSummarize != nil {
		toSerialize["auto_summarize"] = o.AutoSummarize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentGoogleMeetConfigurationPatchDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowManualMeetingCreation *bool `json:"allow_manual_meeting_creation,omitempty"`
		AutoSummarize              *bool `json:"auto_summarize,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allow_manual_meeting_creation", "auto_summarize"})
	} else {
		return err
	}
	o.AllowManualMeetingCreation = all.AllowManualMeetingCreation
	o.AutoSummarize = all.AutoSummarize

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
