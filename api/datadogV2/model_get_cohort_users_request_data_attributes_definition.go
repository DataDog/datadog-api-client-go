// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetCohortUsersRequestDataAttributesDefinition
type GetCohortUsersRequestDataAttributesDefinition struct {
	//
	AudienceFilters *GetCohortUsersRequestDataAttributesDefinitionAudienceFilters `json:"audience_filters,omitempty"`
	//
	Cohort *string `json:"cohort,omitempty"`
	//
	InclusionSearch *string `json:"inclusion_search,omitempty"`
	//
	ReturnSearch *string `json:"return_search,omitempty"`
	//
	SegmentId *string `json:"segment_id,omitempty"`
	//
	Window *int64 `json:"window,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetCohortUsersRequestDataAttributesDefinition instantiates a new GetCohortUsersRequestDataAttributesDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetCohortUsersRequestDataAttributesDefinition() *GetCohortUsersRequestDataAttributesDefinition {
	this := GetCohortUsersRequestDataAttributesDefinition{}
	return &this
}

// NewGetCohortUsersRequestDataAttributesDefinitionWithDefaults instantiates a new GetCohortUsersRequestDataAttributesDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetCohortUsersRequestDataAttributesDefinitionWithDefaults() *GetCohortUsersRequestDataAttributesDefinition {
	this := GetCohortUsersRequestDataAttributesDefinition{}
	return &this
}

// GetAudienceFilters returns the AudienceFilters field value if set, zero value otherwise.
func (o *GetCohortUsersRequestDataAttributesDefinition) GetAudienceFilters() GetCohortUsersRequestDataAttributesDefinitionAudienceFilters {
	if o == nil || o.AudienceFilters == nil {
		var ret GetCohortUsersRequestDataAttributesDefinitionAudienceFilters
		return ret
	}
	return *o.AudienceFilters
}

// GetAudienceFiltersOk returns a tuple with the AudienceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortUsersRequestDataAttributesDefinition) GetAudienceFiltersOk() (*GetCohortUsersRequestDataAttributesDefinitionAudienceFilters, bool) {
	if o == nil || o.AudienceFilters == nil {
		return nil, false
	}
	return o.AudienceFilters, true
}

// HasAudienceFilters returns a boolean if a field has been set.
func (o *GetCohortUsersRequestDataAttributesDefinition) HasAudienceFilters() bool {
	return o != nil && o.AudienceFilters != nil
}

// SetAudienceFilters gets a reference to the given GetCohortUsersRequestDataAttributesDefinitionAudienceFilters and assigns it to the AudienceFilters field.
func (o *GetCohortUsersRequestDataAttributesDefinition) SetAudienceFilters(v GetCohortUsersRequestDataAttributesDefinitionAudienceFilters) {
	o.AudienceFilters = &v
}

// GetCohort returns the Cohort field value if set, zero value otherwise.
func (o *GetCohortUsersRequestDataAttributesDefinition) GetCohort() string {
	if o == nil || o.Cohort == nil {
		var ret string
		return ret
	}
	return *o.Cohort
}

// GetCohortOk returns a tuple with the Cohort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortUsersRequestDataAttributesDefinition) GetCohortOk() (*string, bool) {
	if o == nil || o.Cohort == nil {
		return nil, false
	}
	return o.Cohort, true
}

// HasCohort returns a boolean if a field has been set.
func (o *GetCohortUsersRequestDataAttributesDefinition) HasCohort() bool {
	return o != nil && o.Cohort != nil
}

// SetCohort gets a reference to the given string and assigns it to the Cohort field.
func (o *GetCohortUsersRequestDataAttributesDefinition) SetCohort(v string) {
	o.Cohort = &v
}

// GetInclusionSearch returns the InclusionSearch field value if set, zero value otherwise.
func (o *GetCohortUsersRequestDataAttributesDefinition) GetInclusionSearch() string {
	if o == nil || o.InclusionSearch == nil {
		var ret string
		return ret
	}
	return *o.InclusionSearch
}

// GetInclusionSearchOk returns a tuple with the InclusionSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortUsersRequestDataAttributesDefinition) GetInclusionSearchOk() (*string, bool) {
	if o == nil || o.InclusionSearch == nil {
		return nil, false
	}
	return o.InclusionSearch, true
}

// HasInclusionSearch returns a boolean if a field has been set.
func (o *GetCohortUsersRequestDataAttributesDefinition) HasInclusionSearch() bool {
	return o != nil && o.InclusionSearch != nil
}

// SetInclusionSearch gets a reference to the given string and assigns it to the InclusionSearch field.
func (o *GetCohortUsersRequestDataAttributesDefinition) SetInclusionSearch(v string) {
	o.InclusionSearch = &v
}

// GetReturnSearch returns the ReturnSearch field value if set, zero value otherwise.
func (o *GetCohortUsersRequestDataAttributesDefinition) GetReturnSearch() string {
	if o == nil || o.ReturnSearch == nil {
		var ret string
		return ret
	}
	return *o.ReturnSearch
}

// GetReturnSearchOk returns a tuple with the ReturnSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortUsersRequestDataAttributesDefinition) GetReturnSearchOk() (*string, bool) {
	if o == nil || o.ReturnSearch == nil {
		return nil, false
	}
	return o.ReturnSearch, true
}

// HasReturnSearch returns a boolean if a field has been set.
func (o *GetCohortUsersRequestDataAttributesDefinition) HasReturnSearch() bool {
	return o != nil && o.ReturnSearch != nil
}

// SetReturnSearch gets a reference to the given string and assigns it to the ReturnSearch field.
func (o *GetCohortUsersRequestDataAttributesDefinition) SetReturnSearch(v string) {
	o.ReturnSearch = &v
}

// GetSegmentId returns the SegmentId field value if set, zero value otherwise.
func (o *GetCohortUsersRequestDataAttributesDefinition) GetSegmentId() string {
	if o == nil || o.SegmentId == nil {
		var ret string
		return ret
	}
	return *o.SegmentId
}

// GetSegmentIdOk returns a tuple with the SegmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortUsersRequestDataAttributesDefinition) GetSegmentIdOk() (*string, bool) {
	if o == nil || o.SegmentId == nil {
		return nil, false
	}
	return o.SegmentId, true
}

// HasSegmentId returns a boolean if a field has been set.
func (o *GetCohortUsersRequestDataAttributesDefinition) HasSegmentId() bool {
	return o != nil && o.SegmentId != nil
}

// SetSegmentId gets a reference to the given string and assigns it to the SegmentId field.
func (o *GetCohortUsersRequestDataAttributesDefinition) SetSegmentId(v string) {
	o.SegmentId = &v
}

// GetWindow returns the Window field value if set, zero value otherwise.
func (o *GetCohortUsersRequestDataAttributesDefinition) GetWindow() int64 {
	if o == nil || o.Window == nil {
		var ret int64
		return ret
	}
	return *o.Window
}

// GetWindowOk returns a tuple with the Window field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortUsersRequestDataAttributesDefinition) GetWindowOk() (*int64, bool) {
	if o == nil || o.Window == nil {
		return nil, false
	}
	return o.Window, true
}

// HasWindow returns a boolean if a field has been set.
func (o *GetCohortUsersRequestDataAttributesDefinition) HasWindow() bool {
	return o != nil && o.Window != nil
}

// SetWindow gets a reference to the given int64 and assigns it to the Window field.
func (o *GetCohortUsersRequestDataAttributesDefinition) SetWindow(v int64) {
	o.Window = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetCohortUsersRequestDataAttributesDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AudienceFilters != nil {
		toSerialize["audience_filters"] = o.AudienceFilters
	}
	if o.Cohort != nil {
		toSerialize["cohort"] = o.Cohort
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
	if o.Window != nil {
		toSerialize["window"] = o.Window
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetCohortUsersRequestDataAttributesDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AudienceFilters *GetCohortUsersRequestDataAttributesDefinitionAudienceFilters `json:"audience_filters,omitempty"`
		Cohort          *string                                                       `json:"cohort,omitempty"`
		InclusionSearch *string                                                       `json:"inclusion_search,omitempty"`
		ReturnSearch    *string                                                       `json:"return_search,omitempty"`
		SegmentId       *string                                                       `json:"segment_id,omitempty"`
		Window          *int64                                                        `json:"window,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"audience_filters", "cohort", "inclusion_search", "return_search", "segment_id", "window"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AudienceFilters != nil && all.AudienceFilters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AudienceFilters = all.AudienceFilters
	o.Cohort = all.Cohort
	o.InclusionSearch = all.InclusionSearch
	o.ReturnSearch = all.ReturnSearch
	o.SegmentId = all.SegmentId
	o.Window = all.Window

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
