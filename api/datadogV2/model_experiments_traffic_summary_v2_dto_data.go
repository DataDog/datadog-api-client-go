// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsTrafficSummaryV2DTOData JSON:API resource containing the experiment traffic summary identity and fields.
type ExperimentsTrafficSummaryV2DTOData struct {
	// Details of the experiment traffic summary.
	Attributes *ExperimentsTrafficSummaryV2DTODataAttributes `json:"attributes,omitempty"`
	// Identifier of this traffic summary.
	Id string `json:"id"`
	// Traffic summary resource type.
	Type ExperimentsTrafficSummaryV2DTODataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsTrafficSummaryV2DTOData instantiates a new ExperimentsTrafficSummaryV2DTOData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsTrafficSummaryV2DTOData(id string, typeVar ExperimentsTrafficSummaryV2DTODataType) *ExperimentsTrafficSummaryV2DTOData {
	this := ExperimentsTrafficSummaryV2DTOData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewExperimentsTrafficSummaryV2DTODataWithDefaults instantiates a new ExperimentsTrafficSummaryV2DTOData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsTrafficSummaryV2DTODataWithDefaults() *ExperimentsTrafficSummaryV2DTOData {
	this := ExperimentsTrafficSummaryV2DTOData{}
	var typeVar ExperimentsTrafficSummaryV2DTODataType = EXPERIMENTSTRAFFICSUMMARYV2DTODATATYPE_TRAFFIC_SUMMARY
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ExperimentsTrafficSummaryV2DTOData) GetAttributes() ExperimentsTrafficSummaryV2DTODataAttributes {
	if o == nil || o.Attributes == nil {
		var ret ExperimentsTrafficSummaryV2DTODataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsTrafficSummaryV2DTOData) GetAttributesOk() (*ExperimentsTrafficSummaryV2DTODataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ExperimentsTrafficSummaryV2DTOData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given ExperimentsTrafficSummaryV2DTODataAttributes and assigns it to the Attributes field.
func (o *ExperimentsTrafficSummaryV2DTOData) SetAttributes(v ExperimentsTrafficSummaryV2DTODataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *ExperimentsTrafficSummaryV2DTOData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExperimentsTrafficSummaryV2DTOData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ExperimentsTrafficSummaryV2DTOData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *ExperimentsTrafficSummaryV2DTOData) GetType() ExperimentsTrafficSummaryV2DTODataType {
	if o == nil {
		var ret ExperimentsTrafficSummaryV2DTODataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExperimentsTrafficSummaryV2DTOData) GetTypeOk() (*ExperimentsTrafficSummaryV2DTODataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ExperimentsTrafficSummaryV2DTOData) SetType(v ExperimentsTrafficSummaryV2DTODataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsTrafficSummaryV2DTOData) MarshalJSON() ([]byte, error) {
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
func (o *ExperimentsTrafficSummaryV2DTOData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ExperimentsTrafficSummaryV2DTODataAttributes `json:"attributes,omitempty"`
		Id         *string                                       `json:"id"`
		Type       *ExperimentsTrafficSummaryV2DTODataType       `json:"type"`
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
