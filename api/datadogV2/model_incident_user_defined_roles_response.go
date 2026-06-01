// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedRolesResponse Response with a list of incident user-defined roles.
type IncidentUserDefinedRolesResponse struct {
	// List of incident user-defined role data objects.
	Data []IncidentUserDefinedRoleDataResponse `json:"data"`
	// Included resources for an incident user-defined role response.
	Included []IncidentUserDefinedRoleIncludedItem `json:"included,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentUserDefinedRolesResponse instantiates a new IncidentUserDefinedRolesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentUserDefinedRolesResponse(data []IncidentUserDefinedRoleDataResponse) *IncidentUserDefinedRolesResponse {
	this := IncidentUserDefinedRolesResponse{}
	this.Data = data
	return &this
}

// NewIncidentUserDefinedRolesResponseWithDefaults instantiates a new IncidentUserDefinedRolesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentUserDefinedRolesResponseWithDefaults() *IncidentUserDefinedRolesResponse {
	this := IncidentUserDefinedRolesResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *IncidentUserDefinedRolesResponse) GetData() []IncidentUserDefinedRoleDataResponse {
	if o == nil {
		var ret []IncidentUserDefinedRoleDataResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedRolesResponse) GetDataOk() (*[]IncidentUserDefinedRoleDataResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *IncidentUserDefinedRolesResponse) SetData(v []IncidentUserDefinedRoleDataResponse) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *IncidentUserDefinedRolesResponse) GetIncluded() []IncidentUserDefinedRoleIncludedItem {
	if o == nil || o.Included == nil {
		var ret []IncidentUserDefinedRoleIncludedItem
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedRolesResponse) GetIncludedOk() (*[]IncidentUserDefinedRoleIncludedItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *IncidentUserDefinedRolesResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []IncidentUserDefinedRoleIncludedItem and assigns it to the Included field.
func (o *IncidentUserDefinedRolesResponse) SetIncluded(v []IncidentUserDefinedRoleIncludedItem) {
	o.Included = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentUserDefinedRolesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentUserDefinedRolesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *[]IncidentUserDefinedRoleDataResponse `json:"data"`
		Included []IncidentUserDefinedRoleIncludedItem  `json:"included,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included"})
	} else {
		return err
	}
	o.Data = *all.Data
	o.Included = all.Included

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
