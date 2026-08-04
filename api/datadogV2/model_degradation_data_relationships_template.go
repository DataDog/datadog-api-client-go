// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DegradationDataRelationshipsTemplate The template the degradation was created from.
type DegradationDataRelationshipsTemplate struct {
	// The data object identifying the template the degradation was created from.
	Data DegradationDataRelationshipsTemplateData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDegradationDataRelationshipsTemplate instantiates a new DegradationDataRelationshipsTemplate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDegradationDataRelationshipsTemplate(data DegradationDataRelationshipsTemplateData) *DegradationDataRelationshipsTemplate {
	this := DegradationDataRelationshipsTemplate{}
	this.Data = data
	return &this
}

// NewDegradationDataRelationshipsTemplateWithDefaults instantiates a new DegradationDataRelationshipsTemplate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDegradationDataRelationshipsTemplateWithDefaults() *DegradationDataRelationshipsTemplate {
	this := DegradationDataRelationshipsTemplate{}
	return &this
}

// GetData returns the Data field value.
func (o *DegradationDataRelationshipsTemplate) GetData() DegradationDataRelationshipsTemplateData {
	if o == nil {
		var ret DegradationDataRelationshipsTemplateData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DegradationDataRelationshipsTemplate) GetDataOk() (*DegradationDataRelationshipsTemplateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *DegradationDataRelationshipsTemplate) SetData(v DegradationDataRelationshipsTemplateData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DegradationDataRelationshipsTemplate) MarshalJSON() ([]byte, error) {
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
func (o *DegradationDataRelationshipsTemplate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *DegradationDataRelationshipsTemplateData `json:"data"`
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
