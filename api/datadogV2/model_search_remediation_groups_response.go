// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SearchRemediationGroupsResponse A page of remediation groups.
type SearchRemediationGroupsResponse struct {
	// Remediation groups of the page.
	Data []RemediationGroupData `json:"data"`
	// Cases of the page's groups, when the request has `include=case`.
	Included []FindingCaseResponseData `json:"included,omitempty"`
	// Metadata of the search.
	Meta RemediationGroupsMeta `json:"meta"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSearchRemediationGroupsResponse instantiates a new SearchRemediationGroupsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSearchRemediationGroupsResponse(data []RemediationGroupData, meta RemediationGroupsMeta) *SearchRemediationGroupsResponse {
	this := SearchRemediationGroupsResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewSearchRemediationGroupsResponseWithDefaults instantiates a new SearchRemediationGroupsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSearchRemediationGroupsResponseWithDefaults() *SearchRemediationGroupsResponse {
	this := SearchRemediationGroupsResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *SearchRemediationGroupsResponse) GetData() []RemediationGroupData {
	if o == nil {
		var ret []RemediationGroupData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SearchRemediationGroupsResponse) GetDataOk() (*[]RemediationGroupData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *SearchRemediationGroupsResponse) SetData(v []RemediationGroupData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SearchRemediationGroupsResponse) GetIncluded() []FindingCaseResponseData {
	if o == nil || o.Included == nil {
		var ret []FindingCaseResponseData
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRemediationGroupsResponse) GetIncludedOk() (*[]FindingCaseResponseData, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SearchRemediationGroupsResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []FindingCaseResponseData and assigns it to the Included field.
func (o *SearchRemediationGroupsResponse) SetIncluded(v []FindingCaseResponseData) {
	o.Included = v
}

// GetMeta returns the Meta field value.
func (o *SearchRemediationGroupsResponse) GetMeta() RemediationGroupsMeta {
	if o == nil {
		var ret RemediationGroupsMeta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *SearchRemediationGroupsResponse) GetMetaOk() (*RemediationGroupsMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *SearchRemediationGroupsResponse) SetMeta(v RemediationGroupsMeta) {
	o.Meta = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SearchRemediationGroupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SearchRemediationGroupsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *[]RemediationGroupData   `json:"data"`
		Included []FindingCaseResponseData `json:"included,omitempty"`
		Meta     *RemediationGroupsMeta    `json:"meta"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	if all.Meta == nil {
		return fmt.Errorf("required field meta missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	o.Included = all.Included
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
