/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// AuditLogsSearchEventsRequest The request for a Audit Logs events list.
type AuditLogsSearchEventsRequest struct {
	Filter  *AuditLogsQueryFilter      `json:"filter,omitempty"`
	Options *AuditLogsQueryOptions     `json:"options,omitempty"`
	Page    *AuditLogsQueryPageOptions `json:"page,omitempty"`
	Sort    *AuditLogsSort             `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewAuditLogsSearchEventsRequest instantiates a new AuditLogsSearchEventsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogsSearchEventsRequest() *AuditLogsSearchEventsRequest {
	this := AuditLogsSearchEventsRequest{}
	return &this
}

// NewAuditLogsSearchEventsRequestWithDefaults instantiates a new AuditLogsSearchEventsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogsSearchEventsRequestWithDefaults() *AuditLogsSearchEventsRequest {
	this := AuditLogsSearchEventsRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *AuditLogsSearchEventsRequest) GetFilter() AuditLogsQueryFilter {
	if o == nil || o.Filter == nil {
		var ret AuditLogsQueryFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogsSearchEventsRequest) GetFilterOk() (*AuditLogsQueryFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *AuditLogsSearchEventsRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given AuditLogsQueryFilter and assigns it to the Filter field.
func (o *AuditLogsSearchEventsRequest) SetFilter(v AuditLogsQueryFilter) {
	o.Filter = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *AuditLogsSearchEventsRequest) GetOptions() AuditLogsQueryOptions {
	if o == nil || o.Options == nil {
		var ret AuditLogsQueryOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogsSearchEventsRequest) GetOptionsOk() (*AuditLogsQueryOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *AuditLogsSearchEventsRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given AuditLogsQueryOptions and assigns it to the Options field.
func (o *AuditLogsSearchEventsRequest) SetOptions(v AuditLogsQueryOptions) {
	o.Options = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *AuditLogsSearchEventsRequest) GetPage() AuditLogsQueryPageOptions {
	if o == nil || o.Page == nil {
		var ret AuditLogsQueryPageOptions
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogsSearchEventsRequest) GetPageOk() (*AuditLogsQueryPageOptions, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *AuditLogsSearchEventsRequest) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given AuditLogsQueryPageOptions and assigns it to the Page field.
func (o *AuditLogsSearchEventsRequest) SetPage(v AuditLogsQueryPageOptions) {
	o.Page = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *AuditLogsSearchEventsRequest) GetSort() AuditLogsSort {
	if o == nil || o.Sort == nil {
		var ret AuditLogsSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogsSearchEventsRequest) GetSortOk() (*AuditLogsSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *AuditLogsSearchEventsRequest) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given AuditLogsSort and assigns it to the Sort field.
func (o *AuditLogsSearchEventsRequest) SetSort(v AuditLogsSort) {
	o.Sort = &v
}

func (o AuditLogsSearchEventsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	return json.Marshal(toSerialize)
}

func (o *AuditLogsSearchEventsRequest) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Filter  *AuditLogsQueryFilter      `json:"filter,omitempty"`
		Options *AuditLogsQueryOptions     `json:"options,omitempty"`
		Page    *AuditLogsQueryPageOptions `json:"page,omitempty"`
		Sort    *AuditLogsSort             `json:"sort,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Sort; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Filter = all.Filter
	o.Options = all.Options
	o.Page = all.Page
	o.Sort = all.Sort
	return nil
}
