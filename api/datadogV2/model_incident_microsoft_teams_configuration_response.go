// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentMicrosoftTeamsConfigurationResponse Response with a Microsoft Teams configuration.
type IncidentMicrosoftTeamsConfigurationResponse struct {
	// Microsoft Teams configuration data in a response.
	Data IncidentMicrosoftTeamsConfigurationDataResponse `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentMicrosoftTeamsConfigurationResponse instantiates a new IncidentMicrosoftTeamsConfigurationResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentMicrosoftTeamsConfigurationResponse(data IncidentMicrosoftTeamsConfigurationDataResponse) *IncidentMicrosoftTeamsConfigurationResponse {
	this := IncidentMicrosoftTeamsConfigurationResponse{}
	this.Data = data
	return &this
}

// NewIncidentMicrosoftTeamsConfigurationResponseWithDefaults instantiates a new IncidentMicrosoftTeamsConfigurationResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentMicrosoftTeamsConfigurationResponseWithDefaults() *IncidentMicrosoftTeamsConfigurationResponse {
	this := IncidentMicrosoftTeamsConfigurationResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *IncidentMicrosoftTeamsConfigurationResponse) GetData() IncidentMicrosoftTeamsConfigurationDataResponse {
	if o == nil {
		var ret IncidentMicrosoftTeamsConfigurationDataResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentMicrosoftTeamsConfigurationResponse) GetDataOk() (*IncidentMicrosoftTeamsConfigurationDataResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *IncidentMicrosoftTeamsConfigurationResponse) SetData(v IncidentMicrosoftTeamsConfigurationDataResponse) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentMicrosoftTeamsConfigurationResponse) MarshalJSON() ([]byte, error) {
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
func (o *IncidentMicrosoftTeamsConfigurationResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *IncidentMicrosoftTeamsConfigurationDataResponse `json:"data"`
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
