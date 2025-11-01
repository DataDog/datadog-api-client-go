// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetCohortRequestDataAttributesDefinition
type GetCohortRequestDataAttributesDefinition struct {
	//
	AudienceFilters *GetCohortRequestDataAttributesDefinitionAudienceFilters `json:"audience_filters,omitempty"`
	//
	InclusionSearch *string `json:"inclusion_search,omitempty"`
	//
	ReturnSearch *string `json:"return_search,omitempty"`
	//
	SegmentId *string `json:"segment_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetCohortRequestDataAttributesDefinition instantiates a new GetCohortRequestDataAttributesDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetCohortRequestDataAttributesDefinition() *GetCohortRequestDataAttributesDefinition {
	this := GetCohortRequestDataAttributesDefinition{}
	return &this
}

// NewGetCohortRequestDataAttributesDefinitionWithDefaults instantiates a new GetCohortRequestDataAttributesDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetCohortRequestDataAttributesDefinitionWithDefaults() *GetCohortRequestDataAttributesDefinition {
	this := GetCohortRequestDataAttributesDefinition{}
	return &this
}

// GetAudienceFilters returns the AudienceFilters field value if set, zero value otherwise.
func (o *GetCohortRequestDataAttributesDefinition) GetAudienceFilters() GetCohortRequestDataAttributesDefinitionAudienceFilters {
	if o == nil || o.AudienceFilters == nil {
		var ret GetCohortRequestDataAttributesDefinitionAudienceFilters
		return ret
	}
	return *o.AudienceFilters
}

// GetAudienceFiltersOk returns a tuple with the AudienceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortRequestDataAttributesDefinition) GetAudienceFiltersOk() (*GetCohortRequestDataAttributesDefinitionAudienceFilters, bool) {
	if o == nil || o.AudienceFilters == nil {
		return nil, false
	}
	return o.AudienceFilters, true
}

// HasAudienceFilters returns a boolean if a field has been set.
func (o *GetCohortRequestDataAttributesDefinition) HasAudienceFilters() bool {
	return o != nil && o.AudienceFilters != nil
}

// SetAudienceFilters gets a reference to the given GetCohortRequestDataAttributesDefinitionAudienceFilters and assigns it to the AudienceFilters field.
func (o *GetCohortRequestDataAttributesDefinition) SetAudienceFilters(v GetCohortRequestDataAttributesDefinitionAudienceFilters) {
	o.AudienceFilters = &v
}

// GetInclusionSearch returns the InclusionSearch field value if set, zero value otherwise.
func (o *GetCohortRequestDataAttributesDefinition) GetInclusionSearch() string {
	if o == nil || o.InclusionSearch == nil {
		var ret string
		return ret
	}
	return *o.InclusionSearch
}

// GetInclusionSearchOk returns a tuple with the InclusionSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortRequestDataAttributesDefinition) GetInclusionSearchOk() (*string, bool) {
	if o == nil || o.InclusionSearch == nil {
		return nil, false
	}
	return o.InclusionSearch, true
}

// HasInclusionSearch returns a boolean if a field has been set.
func (o *GetCohortRequestDataAttributesDefinition) HasInclusionSearch() bool {
	return o != nil && o.InclusionSearch != nil
}

// SetInclusionSearch gets a reference to the given string and assigns it to the InclusionSearch field.
func (o *GetCohortRequestDataAttributesDefinition) SetInclusionSearch(v string) {
	o.InclusionSearch = &v
}

// GetReturnSearch returns the ReturnSearch field value if set, zero value otherwise.
func (o *GetCohortRequestDataAttributesDefinition) GetReturnSearch() string {
	if o == nil || o.ReturnSearch == nil {
		var ret string
		return ret
	}
	return *o.ReturnSearch
}

// GetReturnSearchOk returns a tuple with the ReturnSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortRequestDataAttributesDefinition) GetReturnSearchOk() (*string, bool) {
	if o == nil || o.ReturnSearch == nil {
		return nil, false
	}
	return o.ReturnSearch, true
}

// HasReturnSearch returns a boolean if a field has been set.
func (o *GetCohortRequestDataAttributesDefinition) HasReturnSearch() bool {
	return o != nil && o.ReturnSearch != nil
}

// SetReturnSearch gets a reference to the given string and assigns it to the ReturnSearch field.
func (o *GetCohortRequestDataAttributesDefinition) SetReturnSearch(v string) {
	o.ReturnSearch = &v
}

// GetSegmentId returns the SegmentId field value if set, zero value otherwise.
func (o *GetCohortRequestDataAttributesDefinition) GetSegmentId() string {
	if o == nil || o.SegmentId == nil {
		var ret string
		return ret
	}
	return *o.SegmentId
}

// GetSegmentIdOk returns a tuple with the SegmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortRequestDataAttributesDefinition) GetSegmentIdOk() (*string, bool) {
	if o == nil || o.SegmentId == nil {
		return nil, false
	}
	return o.SegmentId, true
}

// HasSegmentId returns a boolean if a field has been set.
func (o *GetCohortRequestDataAttributesDefinition) HasSegmentId() bool {
	return o != nil && o.SegmentId != nil
}

// SetSegmentId gets a reference to the given string and assigns it to the SegmentId field.
func (o *GetCohortRequestDataAttributesDefinition) SetSegmentId(v string) {
	o.SegmentId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetCohortRequestDataAttributesDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AudienceFilters != nil {
		toSerialize["audience_filters"] = o.AudienceFilters
	}
	if o.InclusionSearch != nil {
		toSerialize["inclusion_search"] = o.InclusionSearch
	}
	if o.ReturnSearch != nil {
		toSerialize["return_search"] = o.ReturnSearch
	}
	if o.SegmentId != nil {
		toSerialize["segment_id"] = o.SegmentId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetCohortRequestDataAttributesDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AudienceFilters *GetCohortRequestDataAttributesDefinitionAudienceFilters `json:"audience_filters,omitempty"`
		InclusionSearch *string                                                  `json:"inclusion_search,omitempty"`
		ReturnSearch    *string                                                  `json:"return_search,omitempty"`
		SegmentId       *string                                                  `json:"segment_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"audience_filters", "inclusion_search", "return_search", "segment_id"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AudienceFilters != nil && all.AudienceFilters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AudienceFilters = all.AudienceFilters
	o.InclusionSearch = all.InclusionSearch
	o.ReturnSearch = all.ReturnSearch
	o.SegmentId = all.SegmentId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
