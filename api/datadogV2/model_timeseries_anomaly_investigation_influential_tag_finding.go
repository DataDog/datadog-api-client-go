// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationInfluentialTagFinding Finding that attributes an anomaly to an influential tag.
type TimeseriesAnomalyInvestigationInfluentialTagFinding struct {
	// Deterministic explanation of the finding.
	Description string `json:"description"`
	// Concise, deterministic finding title.
	Headline string `json:"headline"`
	// Structured tag evidence for an influential-tag finding.
	Tag TimeseriesAnomalyInvestigationFindingTag `json:"tag"`
	// Finding category for an influential tag.
	Type TimeseriesAnomalyInvestigationInfluentialTagFindingType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesAnomalyInvestigationInfluentialTagFinding instantiates a new TimeseriesAnomalyInvestigationInfluentialTagFinding object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesAnomalyInvestigationInfluentialTagFinding(description string, headline string, tag TimeseriesAnomalyInvestigationFindingTag, typeVar TimeseriesAnomalyInvestigationInfluentialTagFindingType) *TimeseriesAnomalyInvestigationInfluentialTagFinding {
	this := TimeseriesAnomalyInvestigationInfluentialTagFinding{}
	this.Description = description
	this.Headline = headline
	this.Tag = tag
	this.Type = typeVar
	return &this
}

// NewTimeseriesAnomalyInvestigationInfluentialTagFindingWithDefaults instantiates a new TimeseriesAnomalyInvestigationInfluentialTagFinding object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesAnomalyInvestigationInfluentialTagFindingWithDefaults() *TimeseriesAnomalyInvestigationInfluentialTagFinding {
	this := TimeseriesAnomalyInvestigationInfluentialTagFinding{}
	return &this
}

// GetDescription returns the Description field value.
func (o *TimeseriesAnomalyInvestigationInfluentialTagFinding) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationInfluentialTagFinding) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *TimeseriesAnomalyInvestigationInfluentialTagFinding) SetDescription(v string) {
	o.Description = v
}

// GetHeadline returns the Headline field value.
func (o *TimeseriesAnomalyInvestigationInfluentialTagFinding) GetHeadline() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationInfluentialTagFinding) GetHeadlineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Headline, true
}

// SetHeadline sets field value.
func (o *TimeseriesAnomalyInvestigationInfluentialTagFinding) SetHeadline(v string) {
	o.Headline = v
}

// GetTag returns the Tag field value.
func (o *TimeseriesAnomalyInvestigationInfluentialTagFinding) GetTag() TimeseriesAnomalyInvestigationFindingTag {
	if o == nil {
		var ret TimeseriesAnomalyInvestigationFindingTag
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationInfluentialTagFinding) GetTagOk() (*TimeseriesAnomalyInvestigationFindingTag, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value.
func (o *TimeseriesAnomalyInvestigationInfluentialTagFinding) SetTag(v TimeseriesAnomalyInvestigationFindingTag) {
	o.Tag = v
}

// GetType returns the Type field value.
func (o *TimeseriesAnomalyInvestigationInfluentialTagFinding) GetType() TimeseriesAnomalyInvestigationInfluentialTagFindingType {
	if o == nil {
		var ret TimeseriesAnomalyInvestigationInfluentialTagFindingType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationInfluentialTagFinding) GetTypeOk() (*TimeseriesAnomalyInvestigationInfluentialTagFindingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TimeseriesAnomalyInvestigationInfluentialTagFinding) SetType(v TimeseriesAnomalyInvestigationInfluentialTagFindingType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesAnomalyInvestigationInfluentialTagFinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["headline"] = o.Headline
	toSerialize["tag"] = o.Tag
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesAnomalyInvestigationInfluentialTagFinding) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                                                  `json:"description"`
		Headline    *string                                                  `json:"headline"`
		Tag         *TimeseriesAnomalyInvestigationFindingTag                `json:"tag"`
		Type        *TimeseriesAnomalyInvestigationInfluentialTagFindingType `json:"type"`
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
	if all.Tag == nil {
		return fmt.Errorf("required field tag missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "headline", "tag", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = *all.Description
	o.Headline = *all.Headline
	if all.Tag.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Tag = *all.Tag
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
