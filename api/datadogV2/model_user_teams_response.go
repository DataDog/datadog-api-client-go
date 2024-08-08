// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserTeamsResponse Team memberships response
type UserTeamsResponse struct {
	// Team memberships response data
	Data []UserTeam `json:"data,omitempty"`
	// Resources related to the team memberships
	Included []UserTeamIncluded `json:"included,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserTeamsResponse instantiates a new UserTeamsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserTeamsResponse() *UserTeamsResponse {
	this := UserTeamsResponse{}
	return &this
}

// NewUserTeamsResponseWithDefaults instantiates a new UserTeamsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserTeamsResponseWithDefaults() *UserTeamsResponse {
	this := UserTeamsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UserTeamsResponse) GetData() []UserTeam {
	if o == nil || o.Data == nil {
		var ret []UserTeam
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTeamsResponse) GetDataOk() (*[]UserTeam, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UserTeamsResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []UserTeam and assigns it to the Data field.
func (o *UserTeamsResponse) SetData(v []UserTeam) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *UserTeamsResponse) GetIncluded() []UserTeamIncluded {
	if o == nil || o.Included == nil {
		var ret []UserTeamIncluded
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTeamsResponse) GetIncludedOk() (*[]UserTeamIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *UserTeamsResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []UserTeamIncluded and assigns it to the Included field.
func (o *UserTeamsResponse) SetIncluded(v []UserTeamIncluded) {
	o.Included = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserTeamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserTeamsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     []UserTeam         `json:"data,omitempty"`
		Included []UserTeamIncluded `json:"included,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included"})
	} else {
		return err
	}
	o.Data = all.Data
	o.Included = all.Included

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
