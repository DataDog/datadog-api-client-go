// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemJourneyResponseData Data object in a DEM journey response.
type DemJourneyResponseData struct {
	// Attributes returned in a DEM journey response.
	Attributes DemJourneyResponseAttributes `json:"attributes"`
	// The unique identifier of the DEM journey.
	Id string `json:"id"`
	// The type identifier for DEM journeys.
	Type DemJourneyType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDemJourneyResponseData instantiates a new DemJourneyResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDemJourneyResponseData(attributes DemJourneyResponseAttributes, id string, typeVar DemJourneyType) *DemJourneyResponseData {
	this := DemJourneyResponseData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewDemJourneyResponseDataWithDefaults instantiates a new DemJourneyResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDemJourneyResponseDataWithDefaults() *DemJourneyResponseData {
	this := DemJourneyResponseData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *DemJourneyResponseData) GetAttributes() DemJourneyResponseAttributes {
	if o == nil {
		var ret DemJourneyResponseAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *DemJourneyResponseData) GetAttributesOk() (*DemJourneyResponseAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *DemJourneyResponseData) SetAttributes(v DemJourneyResponseAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *DemJourneyResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DemJourneyResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *DemJourneyResponseData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *DemJourneyResponseData) GetType() DemJourneyType {
	if o == nil {
		var ret DemJourneyType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DemJourneyResponseData) GetTypeOk() (*DemJourneyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DemJourneyResponseData) SetType(v DemJourneyType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DemJourneyResponseData) MarshalJSON() ([]byte, error) {
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
func (o *DemJourneyResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *DemJourneyResponseAttributes `json:"attributes"`
		Id         *string                       `json:"id"`
		Type       *DemJourneyType               `json:"type"`
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
