// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentGoogleMeetConfigurationRequest Request payload for creating a Google Meet configuration.
type IncidentGoogleMeetConfigurationRequest struct {
	// Google Meet configuration data in a create request.
	Data IncidentGoogleMeetConfigurationDataRequest `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentGoogleMeetConfigurationRequest instantiates a new IncidentGoogleMeetConfigurationRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentGoogleMeetConfigurationRequest(data IncidentGoogleMeetConfigurationDataRequest) *IncidentGoogleMeetConfigurationRequest {
	this := IncidentGoogleMeetConfigurationRequest{}
	this.Data = data
	return &this
}

// NewIncidentGoogleMeetConfigurationRequestWithDefaults instantiates a new IncidentGoogleMeetConfigurationRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentGoogleMeetConfigurationRequestWithDefaults() *IncidentGoogleMeetConfigurationRequest {
	this := IncidentGoogleMeetConfigurationRequest{}
	return &this
}

// GetData returns the Data field value.
func (o *IncidentGoogleMeetConfigurationRequest) GetData() IncidentGoogleMeetConfigurationDataRequest {
	if o == nil {
		var ret IncidentGoogleMeetConfigurationDataRequest
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentGoogleMeetConfigurationRequest) GetDataOk() (*IncidentGoogleMeetConfigurationDataRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *IncidentGoogleMeetConfigurationRequest) SetData(v IncidentGoogleMeetConfigurationDataRequest) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentGoogleMeetConfigurationRequest) MarshalJSON() ([]byte, error) {
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
func (o *IncidentGoogleMeetConfigurationRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *IncidentGoogleMeetConfigurationDataRequest `json:"data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
