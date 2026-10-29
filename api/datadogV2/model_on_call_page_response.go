// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageResponse A single On-Call Page with its related included resources.
type OnCallPageResponse struct {
	// An On-Call Page.
	Data OnCallPageData `json:"data"`
	// Resources related to the Page, such as users, teams, scheduled actions, and links.
	Included []OnCallPagesIncluded `json:"included,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnCallPageResponse instantiates a new OnCallPageResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnCallPageResponse(data OnCallPageData) *OnCallPageResponse {
	this := OnCallPageResponse{}
	this.Data = data
	return &this
}

// NewOnCallPageResponseWithDefaults instantiates a new OnCallPageResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnCallPageResponseWithDefaults() *OnCallPageResponse {
	this := OnCallPageResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *OnCallPageResponse) GetData() OnCallPageData {
	if o == nil {
		var ret OnCallPageData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OnCallPageResponse) GetDataOk() (*OnCallPageData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *OnCallPageResponse) SetData(v OnCallPageData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *OnCallPageResponse) GetIncluded() []OnCallPagesIncluded {
	if o == nil || o.Included == nil {
		var ret []OnCallPagesIncluded
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageResponse) GetIncludedOk() (*[]OnCallPagesIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *OnCallPageResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []OnCallPagesIncluded and assigns it to the Included field.
func (o *OnCallPageResponse) SetIncluded(v []OnCallPagesIncluded) {
	o.Included = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnCallPageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnCallPageResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *OnCallPageData       `json:"data"`
		Included []OnCallPagesIncluded `json:"included,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data
	o.Included = all.Included

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
