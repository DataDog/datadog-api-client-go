// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListRowsResponseMetaPage Contains the continuation token for navigating to the next page of rows.
type ListRowsResponseMetaPage struct {
	// Opaque token to pass as the `page[continuation_token]` query parameter to fetch the next page of results. Only present when more rows are available.
	NextContinuationToken *string `json:"next_continuation_token,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListRowsResponseMetaPage instantiates a new ListRowsResponseMetaPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListRowsResponseMetaPage() *ListRowsResponseMetaPage {
	this := ListRowsResponseMetaPage{}
	return &this
}

// NewListRowsResponseMetaPageWithDefaults instantiates a new ListRowsResponseMetaPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListRowsResponseMetaPageWithDefaults() *ListRowsResponseMetaPage {
	this := ListRowsResponseMetaPage{}
	return &this
}

// GetNextContinuationToken returns the NextContinuationToken field value if set, zero value otherwise.
func (o *ListRowsResponseMetaPage) GetNextContinuationToken() string {
	if o == nil || o.NextContinuationToken == nil {
		var ret string
		return ret
	}
	return *o.NextContinuationToken
}

// GetNextContinuationTokenOk returns a tuple with the NextContinuationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRowsResponseMetaPage) GetNextContinuationTokenOk() (*string, bool) {
	if o == nil || o.NextContinuationToken == nil {
		return nil, false
	}
	return o.NextContinuationToken, true
}

// HasNextContinuationToken returns a boolean if a field has been set.
func (o *ListRowsResponseMetaPage) HasNextContinuationToken() bool {
	return o != nil && o.NextContinuationToken != nil
}

// SetNextContinuationToken gets a reference to the given string and assigns it to the NextContinuationToken field.
func (o *ListRowsResponseMetaPage) SetNextContinuationToken(v string) {
	o.NextContinuationToken = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListRowsResponseMetaPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.NextContinuationToken != nil {
		toSerialize["next_continuation_token"] = o.NextContinuationToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListRowsResponseMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NextContinuationToken *string `json:"next_continuation_token,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"next_continuation_token"})
	} else {
		return err
	}
	o.NextContinuationToken = all.NextContinuationToken

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
