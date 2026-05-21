// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTimelineEntriesResponse Response with a list of timeline entries.
type IncidentTimelineEntriesResponse struct {
	// List of timeline entries.
	Data []IncidentTimelineEntryDataResponse `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTimelineEntriesResponse instantiates a new IncidentTimelineEntriesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTimelineEntriesResponse(data []IncidentTimelineEntryDataResponse) *IncidentTimelineEntriesResponse {
	this := IncidentTimelineEntriesResponse{}
	this.Data = data
	return &this
}

// NewIncidentTimelineEntriesResponseWithDefaults instantiates a new IncidentTimelineEntriesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTimelineEntriesResponseWithDefaults() *IncidentTimelineEntriesResponse {
	this := IncidentTimelineEntriesResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *IncidentTimelineEntriesResponse) GetData() []IncidentTimelineEntryDataResponse {
	if o == nil {
		var ret []IncidentTimelineEntryDataResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentTimelineEntriesResponse) GetDataOk() (*[]IncidentTimelineEntryDataResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *IncidentTimelineEntriesResponse) SetData(v []IncidentTimelineEntryDataResponse) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTimelineEntriesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentTimelineEntriesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]IncidentTimelineEntryDataResponse `json:"data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
	o.Data = *all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
