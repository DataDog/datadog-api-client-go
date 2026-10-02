// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateExposureSQLModelV2RequestData Exposure SQL model resource to create.
type ExperimentsCreateExposureSQLModelV2RequestData struct {
	// Complete column mappings and query used to replace the exposure SQL model.
	Attributes ExperimentsUpdateExposureSQLModelV2RequestDataAttributes `json:"attributes"`
	// Optional JSON:API resource identifier field.
	Id *string `json:"id,omitempty"`
	// Exposure SQL models resource type.
	Type ExperimentsUpdateExposureSQLModelV2RequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsCreateExposureSQLModelV2RequestData instantiates a new ExperimentsCreateExposureSQLModelV2RequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCreateExposureSQLModelV2RequestData(attributes ExperimentsUpdateExposureSQLModelV2RequestDataAttributes, typeVar ExperimentsUpdateExposureSQLModelV2RequestDataType) *ExperimentsCreateExposureSQLModelV2RequestData {
	this := ExperimentsCreateExposureSQLModelV2RequestData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewExperimentsCreateExposureSQLModelV2RequestDataWithDefaults instantiates a new ExperimentsCreateExposureSQLModelV2RequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCreateExposureSQLModelV2RequestDataWithDefaults() *ExperimentsCreateExposureSQLModelV2RequestData {
	this := ExperimentsCreateExposureSQLModelV2RequestData{}
	var typeVar ExperimentsUpdateExposureSQLModelV2RequestDataType = EXPERIMENTSUPDATEEXPOSURESQLMODELV2REQUESTDATATYPE_EXPOSURE_SQL_MODELS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *ExperimentsCreateExposureSQLModelV2RequestData) GetAttributes() ExperimentsUpdateExposureSQLModelV2RequestDataAttributes {
	if o == nil {
		var ret ExperimentsUpdateExposureSQLModelV2RequestDataAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExposureSQLModelV2RequestData) GetAttributesOk() (*ExperimentsUpdateExposureSQLModelV2RequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ExperimentsCreateExposureSQLModelV2RequestData) SetAttributes(v ExperimentsUpdateExposureSQLModelV2RequestDataAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExperimentsCreateExposureSQLModelV2RequestData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExposureSQLModelV2RequestData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExperimentsCreateExposureSQLModelV2RequestData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExperimentsCreateExposureSQLModelV2RequestData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *ExperimentsCreateExposureSQLModelV2RequestData) GetType() ExperimentsUpdateExposureSQLModelV2RequestDataType {
	if o == nil {
		var ret ExperimentsUpdateExposureSQLModelV2RequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExposureSQLModelV2RequestData) GetTypeOk() (*ExperimentsUpdateExposureSQLModelV2RequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ExperimentsCreateExposureSQLModelV2RequestData) SetType(v ExperimentsUpdateExposureSQLModelV2RequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCreateExposureSQLModelV2RequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCreateExposureSQLModelV2RequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ExperimentsUpdateExposureSQLModelV2RequestDataAttributes `json:"attributes"`
		Id         *string                                                   `json:"id,omitempty"`
		Type       *ExperimentsUpdateExposureSQLModelV2RequestDataType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = all.Id
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
