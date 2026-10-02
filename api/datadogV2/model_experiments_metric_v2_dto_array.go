// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsMetricV2DTOArray List of metric resources with pagination information.
type ExperimentsMetricV2DTOArray struct {
	// Resources returned in this response.
	Data []ExperimentsMetricV2DTOData `json:"data"`
	// Links for navigating a paginated result set.
	Links *ExperimentsOffsetLinks `json:"links,omitempty"`
	// Pagination information for a list response.
	Meta *ExperimentsOffsetMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsMetricV2DTOArray instantiates a new ExperimentsMetricV2DTOArray object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsMetricV2DTOArray(data []ExperimentsMetricV2DTOData) *ExperimentsMetricV2DTOArray {
	this := ExperimentsMetricV2DTOArray{}
	this.Data = data
	return &this
}

// NewExperimentsMetricV2DTOArrayWithDefaults instantiates a new ExperimentsMetricV2DTOArray object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsMetricV2DTOArrayWithDefaults() *ExperimentsMetricV2DTOArray {
	this := ExperimentsMetricV2DTOArray{}
	return &this
}

// GetData returns the Data field value.
func (o *ExperimentsMetricV2DTOArray) GetData() []ExperimentsMetricV2DTOData {
	if o == nil {
		var ret []ExperimentsMetricV2DTOData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTOArray) GetDataOk() (*[]ExperimentsMetricV2DTOData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ExperimentsMetricV2DTOArray) SetData(v []ExperimentsMetricV2DTOData) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTOArray) GetLinks() ExperimentsOffsetLinks {
	if o == nil || o.Links == nil {
		var ret ExperimentsOffsetLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTOArray) GetLinksOk() (*ExperimentsOffsetLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTOArray) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given ExperimentsOffsetLinks and assigns it to the Links field.
func (o *ExperimentsMetricV2DTOArray) SetLinks(v ExperimentsOffsetLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ExperimentsMetricV2DTOArray) GetMeta() ExperimentsOffsetMeta {
	if o == nil || o.Meta == nil {
		var ret ExperimentsOffsetMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricV2DTOArray) GetMetaOk() (*ExperimentsOffsetMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ExperimentsMetricV2DTOArray) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given ExperimentsOffsetMeta and assigns it to the Meta field.
func (o *ExperimentsMetricV2DTOArray) SetMeta(v ExperimentsOffsetMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsMetricV2DTOArray) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsMetricV2DTOArray) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data  *[]ExperimentsMetricV2DTOData `json:"data"`
		Links *ExperimentsOffsetLinks       `json:"links,omitempty"`
		Meta  *ExperimentsOffsetMeta        `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "links", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	if all.Links != nil && all.Links.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Links = all.Links
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
