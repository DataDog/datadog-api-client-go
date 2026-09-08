// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationResponseData JSON:API resource containing anomaly investigation results.
type TimeseriesAnomalyInvestigationResponseData struct {
	// Attributes of an anomaly investigation response.
	Attributes TimeseriesAnomalyInvestigationResponseAttributes `json:"attributes"`
	// Stable identifier for an anomaly investigation response resource.
	Id TimeseriesAnomalyInvestigationResponseID `json:"id"`
	// Resource type for a timeseries anomaly investigation.
	Type TimeseriesAnomalyInvestigationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesAnomalyInvestigationResponseData instantiates a new TimeseriesAnomalyInvestigationResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesAnomalyInvestigationResponseData(attributes TimeseriesAnomalyInvestigationResponseAttributes, id TimeseriesAnomalyInvestigationResponseID, typeVar TimeseriesAnomalyInvestigationType) *TimeseriesAnomalyInvestigationResponseData {
	this := TimeseriesAnomalyInvestigationResponseData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewTimeseriesAnomalyInvestigationResponseDataWithDefaults instantiates a new TimeseriesAnomalyInvestigationResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesAnomalyInvestigationResponseDataWithDefaults() *TimeseriesAnomalyInvestigationResponseData {
	this := TimeseriesAnomalyInvestigationResponseData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *TimeseriesAnomalyInvestigationResponseData) GetAttributes() TimeseriesAnomalyInvestigationResponseAttributes {
	if o == nil {
		var ret TimeseriesAnomalyInvestigationResponseAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationResponseData) GetAttributesOk() (*TimeseriesAnomalyInvestigationResponseAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *TimeseriesAnomalyInvestigationResponseData) SetAttributes(v TimeseriesAnomalyInvestigationResponseAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *TimeseriesAnomalyInvestigationResponseData) GetId() TimeseriesAnomalyInvestigationResponseID {
	if o == nil {
		var ret TimeseriesAnomalyInvestigationResponseID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationResponseData) GetIdOk() (*TimeseriesAnomalyInvestigationResponseID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *TimeseriesAnomalyInvestigationResponseData) SetId(v TimeseriesAnomalyInvestigationResponseID) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *TimeseriesAnomalyInvestigationResponseData) GetType() TimeseriesAnomalyInvestigationType {
	if o == nil {
		var ret TimeseriesAnomalyInvestigationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationResponseData) GetTypeOk() (*TimeseriesAnomalyInvestigationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TimeseriesAnomalyInvestigationResponseData) SetType(v TimeseriesAnomalyInvestigationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesAnomalyInvestigationResponseData) MarshalJSON() ([]byte, error) {
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
func (o *TimeseriesAnomalyInvestigationResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *TimeseriesAnomalyInvestigationResponseAttributes `json:"attributes"`
		Id         *TimeseriesAnomalyInvestigationResponseID         `json:"id"`
		Type       *TimeseriesAnomalyInvestigationType               `json:"type"`
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
	if !all.Id.IsValid() {
		hasInvalidField = true
	} else {
		o.Id = *all.Id
	}
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
