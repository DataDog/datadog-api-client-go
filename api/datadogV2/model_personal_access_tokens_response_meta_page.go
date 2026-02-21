// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PersonalAccessTokensResponseMetaPage Pagination information for personal access tokens response.
type PersonalAccessTokensResponseMetaPage struct {
	// Total filtered personal access token count.
	TotalFilteredCount *int64 `json:"total_filtered_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPersonalAccessTokensResponseMetaPage instantiates a new PersonalAccessTokensResponseMetaPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPersonalAccessTokensResponseMetaPage() *PersonalAccessTokensResponseMetaPage {
	this := PersonalAccessTokensResponseMetaPage{}
	return &this
}

// NewPersonalAccessTokensResponseMetaPageWithDefaults instantiates a new PersonalAccessTokensResponseMetaPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPersonalAccessTokensResponseMetaPageWithDefaults() *PersonalAccessTokensResponseMetaPage {
	this := PersonalAccessTokensResponseMetaPage{}
	return &this
}

// GetTotalFilteredCount returns the TotalFilteredCount field value if set, zero value otherwise.
func (o *PersonalAccessTokensResponseMetaPage) GetTotalFilteredCount() int64 {
	if o == nil || o.TotalFilteredCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalFilteredCount
}

// GetTotalFilteredCountOk returns a tuple with the TotalFilteredCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonalAccessTokensResponseMetaPage) GetTotalFilteredCountOk() (*int64, bool) {
	if o == nil || o.TotalFilteredCount == nil {
		return nil, false
	}
	return o.TotalFilteredCount, true
}

// HasTotalFilteredCount returns a boolean if a field has been set.
func (o *PersonalAccessTokensResponseMetaPage) HasTotalFilteredCount() bool {
	return o != nil && o.TotalFilteredCount != nil
}

// SetTotalFilteredCount gets a reference to the given int64 and assigns it to the TotalFilteredCount field.
func (o *PersonalAccessTokensResponseMetaPage) SetTotalFilteredCount(v int64) {
	o.TotalFilteredCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PersonalAccessTokensResponseMetaPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TotalFilteredCount != nil {
		toSerialize["total_filtered_count"] = o.TotalFilteredCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PersonalAccessTokensResponseMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TotalFilteredCount *int64 `json:"total_filtered_count,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"total_filtered_count"})
	} else {
		return err
	}
	o.TotalFilteredCount = all.TotalFilteredCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
