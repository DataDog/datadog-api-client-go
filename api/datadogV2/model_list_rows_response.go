// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListRowsResponse Paginated list of reference table rows.
type ListRowsResponse struct {
	// The rows.
	Data []TableRowResourceData `json:"data"`
	// Pagination links for the list rows response.
	Links ListRowsResponseLinks `json:"links"`
	// Contains pagination details, including the continuation token for fetching additional rows.
	Meta *ListRowsResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListRowsResponse instantiates a new ListRowsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListRowsResponse(data []TableRowResourceData, links ListRowsResponseLinks) *ListRowsResponse {
	this := ListRowsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewListRowsResponseWithDefaults instantiates a new ListRowsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListRowsResponseWithDefaults() *ListRowsResponse {
	this := ListRowsResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *ListRowsResponse) GetData() []TableRowResourceData {
	if o == nil {
		var ret []TableRowResourceData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListRowsResponse) GetDataOk() (*[]TableRowResourceData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ListRowsResponse) SetData(v []TableRowResourceData) {
	o.Data = v
}

// GetLinks returns the Links field value.
func (o *ListRowsResponse) GetLinks() ListRowsResponseLinks {
	if o == nil {
		var ret ListRowsResponseLinks
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ListRowsResponse) GetLinksOk() (*ListRowsResponseLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value.
func (o *ListRowsResponse) SetLinks(v ListRowsResponseLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListRowsResponse) GetMeta() ListRowsResponseMeta {
	if o == nil || o.Meta == nil {
		var ret ListRowsResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRowsResponse) GetMetaOk() (*ListRowsResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListRowsResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given ListRowsResponseMeta and assigns it to the Meta field.
func (o *ListRowsResponse) SetMeta(v ListRowsResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListRowsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListRowsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data  *[]TableRowResourceData `json:"data"`
		Links *ListRowsResponseLinks  `json:"links"`
		Meta  *ListRowsResponseMeta   `json:"meta,omitempty"`
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
