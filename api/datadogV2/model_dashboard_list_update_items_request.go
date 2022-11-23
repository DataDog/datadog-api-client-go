// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV2

import (
	"encoding/json"
	"fmt"

)


// DashboardListUpdateItemsRequest Request containing the list of dashboards to update to.
type DashboardListUpdateItemsRequest struct {
	// List of dashboards to update the dashboard list to.
	Dashboards []DashboardListItemRequest `json:"dashboards,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}



// NewDashboardListUpdateItemsRequest instantiates a new DashboardListUpdateItemsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardListUpdateItemsRequest() *DashboardListUpdateItemsRequest {
	this := DashboardListUpdateItemsRequest{}
	return &this
}

// NewDashboardListUpdateItemsRequestWithDefaults instantiates a new DashboardListUpdateItemsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardListUpdateItemsRequestWithDefaults() *DashboardListUpdateItemsRequest {
	this := DashboardListUpdateItemsRequest{}
	return &this
}
// GetDashboards returns the Dashboards field value if set, zero value otherwise.
func (o *DashboardListUpdateItemsRequest) GetDashboards() []DashboardListItemRequest {
	if o == nil || o.Dashboards == nil {
		var ret []DashboardListItemRequest
		return ret
	}
	return o.Dashboards
}

// GetDashboardsOk returns a tuple with the Dashboards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardListUpdateItemsRequest) GetDashboardsOk() (*[]DashboardListItemRequest, bool) {
	if o == nil || o.Dashboards == nil {
		return nil, false
	}
	return &o.Dashboards, true
}

// HasDashboards returns a boolean if a field has been set.
func (o *DashboardListUpdateItemsRequest) HasDashboards() bool {
	return o != nil && o.Dashboards != nil
}

// SetDashboards gets a reference to the given []DashboardListItemRequest and assigns it to the Dashboards field.
func (o *DashboardListUpdateItemsRequest) SetDashboards(v []DashboardListItemRequest) {
	o.Dashboards = v
}



// MarshalJSON serializes the struct using spec logic.
func (o DashboardListUpdateItemsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Dashboards != nil {
		toSerialize["dashboards"] = o.Dashboards
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}


// UnmarshalJSON deserializes the given payload.
func (o *DashboardListUpdateItemsRequest) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Dashboards []DashboardListItemRequest `json:"dashboards,omitempty"`
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
	o.Dashboards = all.Dashboards
	return nil
}
