// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsCoverageScalarResponseData The scalar coverage resource.
type CommitmentsCoverageScalarResponseData struct {
	// Attributes of a scalar coverage response.
	Attributes CommitmentsCoverageScalarResponseAttributes `json:"attributes"`
	// The identifier of the scalar coverage response.
	Id string `json:"id"`
	// The type of the resource. The value should always be scalar_response.
	Type ScalarFormulaResponseType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsCoverageScalarResponseData instantiates a new CommitmentsCoverageScalarResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsCoverageScalarResponseData(attributes CommitmentsCoverageScalarResponseAttributes, id string, typeVar ScalarFormulaResponseType) *CommitmentsCoverageScalarResponseData {
	this := CommitmentsCoverageScalarResponseData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewCommitmentsCoverageScalarResponseDataWithDefaults instantiates a new CommitmentsCoverageScalarResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsCoverageScalarResponseDataWithDefaults() *CommitmentsCoverageScalarResponseData {
	this := CommitmentsCoverageScalarResponseData{}
	var typeVar ScalarFormulaResponseType = SCALARFORMULARESPONSETYPE_SCALAR_RESPONSE
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *CommitmentsCoverageScalarResponseData) GetAttributes() CommitmentsCoverageScalarResponseAttributes {
	if o == nil {
		var ret CommitmentsCoverageScalarResponseAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CommitmentsCoverageScalarResponseData) GetAttributesOk() (*CommitmentsCoverageScalarResponseAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *CommitmentsCoverageScalarResponseData) SetAttributes(v CommitmentsCoverageScalarResponseAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *CommitmentsCoverageScalarResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CommitmentsCoverageScalarResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CommitmentsCoverageScalarResponseData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *CommitmentsCoverageScalarResponseData) GetType() ScalarFormulaResponseType {
	if o == nil {
		var ret ScalarFormulaResponseType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CommitmentsCoverageScalarResponseData) GetTypeOk() (*ScalarFormulaResponseType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CommitmentsCoverageScalarResponseData) SetType(v ScalarFormulaResponseType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsCoverageScalarResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsCoverageScalarResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CommitmentsCoverageScalarResponseAttributes `json:"attributes"`
		Id         *string                                      `json:"id"`
		Type       *ScalarFormulaResponseType                   `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
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
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
