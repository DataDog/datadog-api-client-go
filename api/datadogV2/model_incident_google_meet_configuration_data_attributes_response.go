// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentGoogleMeetConfigurationDataAttributesResponse Attributes of a Google Meet configuration.
type IncidentGoogleMeetConfigurationDataAttributesResponse struct {
	// Whether manual meeting creation is allowed.
	AllowManualMeetingCreation bool `json:"allow_manual_meeting_creation"`
	// Whether meetings are auto-summarized.
	AutoSummarize bool `json:"auto_summarize"`
	// Timestamp when the configuration was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Timestamp when the configuration was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentGoogleMeetConfigurationDataAttributesResponse instantiates a new IncidentGoogleMeetConfigurationDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentGoogleMeetConfigurationDataAttributesResponse(allowManualMeetingCreation bool, autoSummarize bool, modifiedAt time.Time) *IncidentGoogleMeetConfigurationDataAttributesResponse {
	this := IncidentGoogleMeetConfigurationDataAttributesResponse{}
	this.AllowManualMeetingCreation = allowManualMeetingCreation
	this.AutoSummarize = autoSummarize
	this.ModifiedAt = modifiedAt
	return &this
}

// NewIncidentGoogleMeetConfigurationDataAttributesResponseWithDefaults instantiates a new IncidentGoogleMeetConfigurationDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentGoogleMeetConfigurationDataAttributesResponseWithDefaults() *IncidentGoogleMeetConfigurationDataAttributesResponse {
	this := IncidentGoogleMeetConfigurationDataAttributesResponse{}
	return &this
}

// GetAllowManualMeetingCreation returns the AllowManualMeetingCreation field value.
func (o *IncidentGoogleMeetConfigurationDataAttributesResponse) GetAllowManualMeetingCreation() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.AllowManualMeetingCreation
}

// GetAllowManualMeetingCreationOk returns a tuple with the AllowManualMeetingCreation field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleMeetConfigurationDataAttributesResponse) GetAllowManualMeetingCreationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowManualMeetingCreation, true
}

// SetAllowManualMeetingCreation sets field value.
func (o *IncidentGoogleMeetConfigurationDataAttributesResponse) SetAllowManualMeetingCreation(v bool) {
	o.AllowManualMeetingCreation = v
}

// GetAutoSummarize returns the AutoSummarize field value.
func (o *IncidentGoogleMeetConfigurationDataAttributesResponse) GetAutoSummarize() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.AutoSummarize
}

// GetAutoSummarizeOk returns a tuple with the AutoSummarize field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleMeetConfigurationDataAttributesResponse) GetAutoSummarizeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoSummarize, true
}

// SetAutoSummarize sets field value.
func (o *IncidentGoogleMeetConfigurationDataAttributesResponse) SetAutoSummarize(v bool) {
	o.AutoSummarize = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IncidentGoogleMeetConfigurationDataAttributesResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentGoogleMeetConfigurationDataAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IncidentGoogleMeetConfigurationDataAttributesResponse) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *IncidentGoogleMeetConfigurationDataAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *IncidentGoogleMeetConfigurationDataAttributesResponse) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleMeetConfigurationDataAttributesResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *IncidentGoogleMeetConfigurationDataAttributesResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentGoogleMeetConfigurationDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["allow_manual_meeting_creation"] = o.AllowManualMeetingCreation
	toSerialize["auto_summarize"] = o.AutoSummarize
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentGoogleMeetConfigurationDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowManualMeetingCreation *bool      `json:"allow_manual_meeting_creation"`
		AutoSummarize              *bool      `json:"auto_summarize"`
		CreatedAt                  *time.Time `json:"created_at,omitempty"`
		ModifiedAt                 *time.Time `json:"modified_at"`
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
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allow_manual_meeting_creation", "auto_summarize", "created_at", "modified_at"})
	} else {
		return err
	}
	o.AllowManualMeetingCreation = *all.AllowManualMeetingCreation
	o.AutoSummarize = *all.AutoSummarize
	o.CreatedAt = all.CreatedAt
	o.ModifiedAt = *all.ModifiedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
