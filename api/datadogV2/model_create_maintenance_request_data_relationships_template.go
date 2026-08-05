// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateMaintenanceRequestDataRelationshipsTemplate The template used to create the maintenance.
type CreateMaintenanceRequestDataRelationshipsTemplate struct {
	// The data object identifying the template used to create the maintenance.
	Data CreateMaintenanceRequestDataRelationshipsTemplateData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateMaintenanceRequestDataRelationshipsTemplate instantiates a new CreateMaintenanceRequestDataRelationshipsTemplate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateMaintenanceRequestDataRelationshipsTemplate(data CreateMaintenanceRequestDataRelationshipsTemplateData) *CreateMaintenanceRequestDataRelationshipsTemplate {
	this := CreateMaintenanceRequestDataRelationshipsTemplate{}
	this.Data = data
	return &this
}

// NewCreateMaintenanceRequestDataRelationshipsTemplateWithDefaults instantiates a new CreateMaintenanceRequestDataRelationshipsTemplate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateMaintenanceRequestDataRelationshipsTemplateWithDefaults() *CreateMaintenanceRequestDataRelationshipsTemplate {
	this := CreateMaintenanceRequestDataRelationshipsTemplate{}
	return &this
}

// GetData returns the Data field value.
func (o *CreateMaintenanceRequestDataRelationshipsTemplate) GetData() CreateMaintenanceRequestDataRelationshipsTemplateData {
	if o == nil {
		var ret CreateMaintenanceRequestDataRelationshipsTemplateData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateMaintenanceRequestDataRelationshipsTemplate) GetDataOk() (*CreateMaintenanceRequestDataRelationshipsTemplateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *CreateMaintenanceRequestDataRelationshipsTemplate) SetData(v CreateMaintenanceRequestDataRelationshipsTemplateData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateMaintenanceRequestDataRelationshipsTemplate) MarshalJSON() ([]byte, error) {
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
func (o *CreateMaintenanceRequestDataRelationshipsTemplate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *CreateMaintenanceRequestDataRelationshipsTemplateData `json:"data"`
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
