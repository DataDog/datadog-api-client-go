// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationAnomalyFinding Finding that describes the anomaly when completed analysis produces no displayable influential tags.
type TimeseriesAnomalyInvestigationAnomalyFinding struct {
	// Deterministic explanation of the finding.
	Description string `json:"description"`
	// Concise, deterministic finding title.
	Headline string `json:"headline"`
	// Finding category for an anomaly without a displayable influential tag.
	Type TimeseriesAnomalyInvestigationAnomalyFindingType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesAnomalyInvestigationAnomalyFinding instantiates a new TimeseriesAnomalyInvestigationAnomalyFinding object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesAnomalyInvestigationAnomalyFinding(description string, headline string, typeVar TimeseriesAnomalyInvestigationAnomalyFindingType) *TimeseriesAnomalyInvestigationAnomalyFinding {
	this := TimeseriesAnomalyInvestigationAnomalyFinding{}
	this.Description = description
	this.Headline = headline
	this.Type = typeVar
	return &this
}

// NewTimeseriesAnomalyInvestigationAnomalyFindingWithDefaults instantiates a new TimeseriesAnomalyInvestigationAnomalyFinding object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesAnomalyInvestigationAnomalyFindingWithDefaults() *TimeseriesAnomalyInvestigationAnomalyFinding {
	this := TimeseriesAnomalyInvestigationAnomalyFinding{}
	return &this
}

// GetDescription returns the Description field value.
func (o *TimeseriesAnomalyInvestigationAnomalyFinding) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationAnomalyFinding) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *TimeseriesAnomalyInvestigationAnomalyFinding) SetDescription(v string) {
	o.Description = v
}

// GetHeadline returns the Headline field value.
func (o *TimeseriesAnomalyInvestigationAnomalyFinding) GetHeadline() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationAnomalyFinding) GetHeadlineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Headline, true
}

// SetHeadline sets field value.
func (o *TimeseriesAnomalyInvestigationAnomalyFinding) SetHeadline(v string) {
	o.Headline = v
}

// GetType returns the Type field value.
func (o *TimeseriesAnomalyInvestigationAnomalyFinding) GetType() TimeseriesAnomalyInvestigationAnomalyFindingType {
	if o == nil {
		var ret TimeseriesAnomalyInvestigationAnomalyFindingType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationAnomalyFinding) GetTypeOk() (*TimeseriesAnomalyInvestigationAnomalyFindingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TimeseriesAnomalyInvestigationAnomalyFinding) SetType(v TimeseriesAnomalyInvestigationAnomalyFindingType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesAnomalyInvestigationAnomalyFinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["headline"] = o.Headline
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesAnomalyInvestigationAnomalyFinding) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                                           `json:"description"`
		Headline    *string                                           `json:"headline"`
		Type        *TimeseriesAnomalyInvestigationAnomalyFindingType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Headline == nil {
		return fmt.Errorf("required field headline missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "headline", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = *all.Description
	o.Headline = *all.Headline
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
