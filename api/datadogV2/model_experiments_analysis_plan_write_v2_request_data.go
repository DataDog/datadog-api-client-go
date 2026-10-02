// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsAnalysisPlanWriteV2RequestData Analysis plan resource to update.
type ExperimentsAnalysisPlanWriteV2RequestData struct {
	// Statistical settings and duration targets to apply to the experiment.
	Attributes *ExperimentsAnalysisPlanWriteV2RequestDataAttributes `json:"attributes,omitempty"`
	// Identifier of the experiment. If supplied, it must match experiment_id in the path.
	Id *uuid.UUID `json:"id,omitempty"`
	// Analysis plans resource type.
	Type ExperimentsAnalysisPlanWriteV2RequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsAnalysisPlanWriteV2RequestData instantiates a new ExperimentsAnalysisPlanWriteV2RequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsAnalysisPlanWriteV2RequestData(typeVar ExperimentsAnalysisPlanWriteV2RequestDataType) *ExperimentsAnalysisPlanWriteV2RequestData {
	this := ExperimentsAnalysisPlanWriteV2RequestData{}
	this.Type = typeVar
	return &this
}

// NewExperimentsAnalysisPlanWriteV2RequestDataWithDefaults instantiates a new ExperimentsAnalysisPlanWriteV2RequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsAnalysisPlanWriteV2RequestDataWithDefaults() *ExperimentsAnalysisPlanWriteV2RequestData {
	this := ExperimentsAnalysisPlanWriteV2RequestData{}
	var typeVar ExperimentsAnalysisPlanWriteV2RequestDataType = EXPERIMENTSANALYSISPLANWRITEV2REQUESTDATATYPE_ANALYSIS_PLANS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanWriteV2RequestData) GetAttributes() ExperimentsAnalysisPlanWriteV2RequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret ExperimentsAnalysisPlanWriteV2RequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestData) GetAttributesOk() (*ExperimentsAnalysisPlanWriteV2RequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given ExperimentsAnalysisPlanWriteV2RequestDataAttributes and assigns it to the Attributes field.
func (o *ExperimentsAnalysisPlanWriteV2RequestData) SetAttributes(v ExperimentsAnalysisPlanWriteV2RequestDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanWriteV2RequestData) GetId() uuid.UUID {
	if o == nil || o.Id == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *ExperimentsAnalysisPlanWriteV2RequestData) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *ExperimentsAnalysisPlanWriteV2RequestData) GetType() ExperimentsAnalysisPlanWriteV2RequestDataType {
	if o == nil {
		var ret ExperimentsAnalysisPlanWriteV2RequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanWriteV2RequestData) GetTypeOk() (*ExperimentsAnalysisPlanWriteV2RequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ExperimentsAnalysisPlanWriteV2RequestData) SetType(v ExperimentsAnalysisPlanWriteV2RequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsAnalysisPlanWriteV2RequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
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
func (o *ExperimentsAnalysisPlanWriteV2RequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ExperimentsAnalysisPlanWriteV2RequestDataAttributes `json:"attributes,omitempty"`
		Id         *uuid.UUID                                           `json:"id,omitempty"`
		Type       *ExperimentsAnalysisPlanWriteV2RequestDataType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
