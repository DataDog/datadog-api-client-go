// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditV2ListDTOArray Response containing a page of bandit summaries.
type ExperimentsBanditV2ListDTOArray struct {
	// Bandits on this page.
	Data []ExperimentsBanditV2ListDTOData `json:"data"`
	// Links for navigating a paginated result set.
	Links ExperimentsOffsetLinks `json:"links"`
	// Pagination information for a list response.
	Meta ExperimentsOffsetMeta `json:"meta"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditV2ListDTOArray instantiates a new ExperimentsBanditV2ListDTOArray object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditV2ListDTOArray(data []ExperimentsBanditV2ListDTOData, links ExperimentsOffsetLinks, meta ExperimentsOffsetMeta) *ExperimentsBanditV2ListDTOArray {
	this := ExperimentsBanditV2ListDTOArray{}
	this.Data = data
	this.Links = links
	this.Meta = meta
	return &this
}

// NewExperimentsBanditV2ListDTOArrayWithDefaults instantiates a new ExperimentsBanditV2ListDTOArray object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditV2ListDTOArrayWithDefaults() *ExperimentsBanditV2ListDTOArray {
	this := ExperimentsBanditV2ListDTOArray{}
	return &this
}

// GetData returns the Data field value.
func (o *ExperimentsBanditV2ListDTOArray) GetData() []ExperimentsBanditV2ListDTOData {
	if o == nil {
		var ret []ExperimentsBanditV2ListDTOData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2ListDTOArray) GetDataOk() (*[]ExperimentsBanditV2ListDTOData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ExperimentsBanditV2ListDTOArray) SetData(v []ExperimentsBanditV2ListDTOData) {
	o.Data = v
}

// GetLinks returns the Links field value.
func (o *ExperimentsBanditV2ListDTOArray) GetLinks() ExperimentsOffsetLinks {
	if o == nil {
		var ret ExperimentsOffsetLinks
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2ListDTOArray) GetLinksOk() (*ExperimentsOffsetLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value.
func (o *ExperimentsBanditV2ListDTOArray) SetLinks(v ExperimentsOffsetLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value.
func (o *ExperimentsBanditV2ListDTOArray) GetMeta() ExperimentsOffsetMeta {
	if o == nil {
		var ret ExperimentsOffsetMeta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2ListDTOArray) GetMetaOk() (*ExperimentsOffsetMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *ExperimentsBanditV2ListDTOArray) SetMeta(v ExperimentsOffsetMeta) {
	o.Meta = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditV2ListDTOArray) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditV2ListDTOArray) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data  *[]ExperimentsBanditV2ListDTOData `json:"data"`
		Links *ExperimentsOffsetLinks           `json:"links"`
		Meta  *ExperimentsOffsetMeta            `json:"meta"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	if all.Links == nil {
		return fmt.Errorf("required field links missing")
	}
	if all.Meta == nil {
		return fmt.Errorf("required field meta missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "links", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	if all.Links.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Links = *all.Links
	if all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = *all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
