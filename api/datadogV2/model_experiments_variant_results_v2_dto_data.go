// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsVariantResultsV2DTOData JSON:API resource containing the variant result identity and fields.
type ExperimentsVariantResultsV2DTOData struct {
	// Details of the variant result.
	Attributes *ExperimentsVariantResultsV2DTODataAttributes `json:"attributes,omitempty"`
	// ID of the variant result.
	Id string `json:"id"`
	// Experiment variant results resource type.
	Type ExperimentsVariantResultsV2DTODataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsVariantResultsV2DTOData instantiates a new ExperimentsVariantResultsV2DTOData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsVariantResultsV2DTOData(id string, typeVar ExperimentsVariantResultsV2DTODataType) *ExperimentsVariantResultsV2DTOData {
	this := ExperimentsVariantResultsV2DTOData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewExperimentsVariantResultsV2DTODataWithDefaults instantiates a new ExperimentsVariantResultsV2DTOData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsVariantResultsV2DTODataWithDefaults() *ExperimentsVariantResultsV2DTOData {
	this := ExperimentsVariantResultsV2DTOData{}
	var typeVar ExperimentsVariantResultsV2DTODataType = EXPERIMENTSVARIANTRESULTSV2DTODATATYPE_EXPERIMENT_VARIANT_RESULTS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTOData) GetAttributes() ExperimentsVariantResultsV2DTODataAttributes {
	if o == nil || o.Attributes == nil {
		var ret ExperimentsVariantResultsV2DTODataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTOData) GetAttributesOk() (*ExperimentsVariantResultsV2DTODataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTOData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given ExperimentsVariantResultsV2DTODataAttributes and assigns it to the Attributes field.
func (o *ExperimentsVariantResultsV2DTOData) SetAttributes(v ExperimentsVariantResultsV2DTODataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *ExperimentsVariantResultsV2DTOData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTOData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ExperimentsVariantResultsV2DTOData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *ExperimentsVariantResultsV2DTOData) GetType() ExperimentsVariantResultsV2DTODataType {
	if o == nil {
		var ret ExperimentsVariantResultsV2DTODataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTOData) GetTypeOk() (*ExperimentsVariantResultsV2DTODataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ExperimentsVariantResultsV2DTOData) SetType(v ExperimentsVariantResultsV2DTODataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsVariantResultsV2DTOData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsVariantResultsV2DTOData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ExperimentsVariantResultsV2DTODataAttributes `json:"attributes,omitempty"`
		Id         *string                                       `json:"id"`
		Type       *ExperimentsVariantResultsV2DTODataType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
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
	o.Id = *all.Id
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
