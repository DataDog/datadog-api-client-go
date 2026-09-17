// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationFindingTag Structured tag evidence for an influential-tag finding.
type TimeseriesAnomalyInvestigationFindingTag struct {
	// Kind of influence a tag has on a series.
	InfluenceType TimeseriesAnomalyInvestigationInfluenceType `json:"influence_type"`
	// Influential tag key.
	Key string `json:"key"`
	// Influence rating from 1 through 5.
	Rating float64 `json:"rating"`
	// Tags grouped with this tag by Variation of Influence synonym analysis.
	Synonyms []TimeseriesAnomalyInvestigationFindingSynonym `json:"synonyms"`
	// Influential values for the tag key.
	Values []string `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesAnomalyInvestigationFindingTag instantiates a new TimeseriesAnomalyInvestigationFindingTag object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesAnomalyInvestigationFindingTag(influenceType TimeseriesAnomalyInvestigationInfluenceType, key string, rating float64, synonyms []TimeseriesAnomalyInvestigationFindingSynonym, values []string) *TimeseriesAnomalyInvestigationFindingTag {
	this := TimeseriesAnomalyInvestigationFindingTag{}
	this.InfluenceType = influenceType
	this.Key = key
	this.Rating = rating
	this.Synonyms = synonyms
	this.Values = values
	return &this
}

// NewTimeseriesAnomalyInvestigationFindingTagWithDefaults instantiates a new TimeseriesAnomalyInvestigationFindingTag object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesAnomalyInvestigationFindingTagWithDefaults() *TimeseriesAnomalyInvestigationFindingTag {
	this := TimeseriesAnomalyInvestigationFindingTag{}
	return &this
}

// GetInfluenceType returns the InfluenceType field value.
func (o *TimeseriesAnomalyInvestigationFindingTag) GetInfluenceType() TimeseriesAnomalyInvestigationInfluenceType {
	if o == nil {
		var ret TimeseriesAnomalyInvestigationInfluenceType
		return ret
	}
	return o.InfluenceType
}

// GetInfluenceTypeOk returns a tuple with the InfluenceType field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationFindingTag) GetInfluenceTypeOk() (*TimeseriesAnomalyInvestigationInfluenceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InfluenceType, true
}

// SetInfluenceType sets field value.
func (o *TimeseriesAnomalyInvestigationFindingTag) SetInfluenceType(v TimeseriesAnomalyInvestigationInfluenceType) {
	o.InfluenceType = v
}

// GetKey returns the Key field value.
func (o *TimeseriesAnomalyInvestigationFindingTag) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationFindingTag) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *TimeseriesAnomalyInvestigationFindingTag) SetKey(v string) {
	o.Key = v
}

// GetRating returns the Rating field value.
func (o *TimeseriesAnomalyInvestigationFindingTag) GetRating() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationFindingTag) GetRatingOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rating, true
}

// SetRating sets field value.
func (o *TimeseriesAnomalyInvestigationFindingTag) SetRating(v float64) {
	o.Rating = v
}

// GetSynonyms returns the Synonyms field value.
func (o *TimeseriesAnomalyInvestigationFindingTag) GetSynonyms() []TimeseriesAnomalyInvestigationFindingSynonym {
	if o == nil {
		var ret []TimeseriesAnomalyInvestigationFindingSynonym
		return ret
	}
	return o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationFindingTag) GetSynonymsOk() (*[]TimeseriesAnomalyInvestigationFindingSynonym, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Synonyms, true
}

// SetSynonyms sets field value.
func (o *TimeseriesAnomalyInvestigationFindingTag) SetSynonyms(v []TimeseriesAnomalyInvestigationFindingSynonym) {
	o.Synonyms = v
}

// GetValues returns the Values field value.
func (o *TimeseriesAnomalyInvestigationFindingTag) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationFindingTag) GetValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *TimeseriesAnomalyInvestigationFindingTag) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesAnomalyInvestigationFindingTag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["influence_type"] = o.InfluenceType
	toSerialize["key"] = o.Key
	toSerialize["rating"] = o.Rating
	toSerialize["synonyms"] = o.Synonyms
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesAnomalyInvestigationFindingTag) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InfluenceType *TimeseriesAnomalyInvestigationInfluenceType    `json:"influence_type"`
		Key           *string                                         `json:"key"`
		Rating        *float64                                        `json:"rating"`
		Synonyms      *[]TimeseriesAnomalyInvestigationFindingSynonym `json:"synonyms"`
		Values        *[]string                                       `json:"values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.InfluenceType == nil {
		return fmt.Errorf("required field influence_type missing")
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Rating == nil {
		return fmt.Errorf("required field rating missing")
	}
	if all.Synonyms == nil {
		return fmt.Errorf("required field synonyms missing")
	}
	if all.Values == nil {
		return fmt.Errorf("required field values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"influence_type", "key", "rating", "synonyms", "values"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.InfluenceType.IsValid() {
		hasInvalidField = true
	} else {
		o.InfluenceType = *all.InfluenceType
	}
	o.Key = *all.Key
	o.Rating = *all.Rating
	o.Synonyms = *all.Synonyms
	o.Values = *all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
