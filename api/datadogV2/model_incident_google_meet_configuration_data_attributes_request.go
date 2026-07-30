// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentGoogleMeetConfigurationDataAttributesRequest Attributes for creating a Google Meet configuration.
type IncidentGoogleMeetConfigurationDataAttributesRequest struct {
	// Whether to allow manual meeting creation.
	AllowManualMeetingCreation bool `json:"allow_manual_meeting_creation"`
	// Whether to auto-summarize meetings.
	AutoSummarize bool `json:"auto_summarize"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentGoogleMeetConfigurationDataAttributesRequest instantiates a new IncidentGoogleMeetConfigurationDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentGoogleMeetConfigurationDataAttributesRequest(allowManualMeetingCreation bool, autoSummarize bool) *IncidentGoogleMeetConfigurationDataAttributesRequest {
	this := IncidentGoogleMeetConfigurationDataAttributesRequest{}
	this.AllowManualMeetingCreation = allowManualMeetingCreation
	this.AutoSummarize = autoSummarize
	return &this
}

// NewIncidentGoogleMeetConfigurationDataAttributesRequestWithDefaults instantiates a new IncidentGoogleMeetConfigurationDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentGoogleMeetConfigurationDataAttributesRequestWithDefaults() *IncidentGoogleMeetConfigurationDataAttributesRequest {
	this := IncidentGoogleMeetConfigurationDataAttributesRequest{}
	return &this
}

// GetAllowManualMeetingCreation returns the AllowManualMeetingCreation field value.
func (o *IncidentGoogleMeetConfigurationDataAttributesRequest) GetAllowManualMeetingCreation() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.AllowManualMeetingCreation
}

// GetAllowManualMeetingCreationOk returns a tuple with the AllowManualMeetingCreation field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleMeetConfigurationDataAttributesRequest) GetAllowManualMeetingCreationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowManualMeetingCreation, true
}

// SetAllowManualMeetingCreation sets field value.
func (o *IncidentGoogleMeetConfigurationDataAttributesRequest) SetAllowManualMeetingCreation(v bool) {
	o.AllowManualMeetingCreation = v
}

// GetAutoSummarize returns the AutoSummarize field value.
func (o *IncidentGoogleMeetConfigurationDataAttributesRequest) GetAutoSummarize() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.AutoSummarize
}

// GetAutoSummarizeOk returns a tuple with the AutoSummarize field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleMeetConfigurationDataAttributesRequest) GetAutoSummarizeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoSummarize, true
}

// SetAutoSummarize sets field value.
func (o *IncidentGoogleMeetConfigurationDataAttributesRequest) SetAutoSummarize(v bool) {
	o.AutoSummarize = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentGoogleMeetConfigurationDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["allow_manual_meeting_creation"] = o.AllowManualMeetingCreation
	toSerialize["auto_summarize"] = o.AutoSummarize

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentGoogleMeetConfigurationDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowManualMeetingCreation *bool `json:"allow_manual_meeting_creation"`
		AutoSummarize              *bool `json:"auto_summarize"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AllowManualMeetingCreation == nil {
		return fmt.Errorf("required field allow_manual_meeting_creation missing")
	}
	if all.AutoSummarize == nil {
		return fmt.Errorf("required field auto_summarize missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allow_manual_meeting_creation", "auto_summarize"})
	} else {
		return err
	}
	o.AllowManualMeetingCreation = *all.AllowManualMeetingCreation
	o.AutoSummarize = *all.AutoSummarize

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
