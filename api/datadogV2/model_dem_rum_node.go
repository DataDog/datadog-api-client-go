// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemRumNode A RUM node within a journey step.
type DemRumNode struct {
	// The RUM application ID whose events this node query matches. This value is required for every node when creating or updating a DEM feature or journey, including variants, and is used to discover the resource in application-scoped searches. Use `GET /api/v2/rum/applications` to find RUM application IDs.
	AppId string `json:"app_id"`
	// The ID of the RUM node element.
	Id *string `json:"id,omitempty"`
	// The RUM query for matching this node.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDemRumNode instantiates a new DemRumNode object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDemRumNode(appId string, query string) *DemRumNode {
	this := DemRumNode{}
	this.AppId = appId
	this.Query = query
	return &this
}

// NewDemRumNodeWithDefaults instantiates a new DemRumNode object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDemRumNodeWithDefaults() *DemRumNode {
	this := DemRumNode{}
	return &this
}

// GetAppId returns the AppId field value.
func (o *DemRumNode) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *DemRumNode) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value.
func (o *DemRumNode) SetAppId(v string) {
	o.AppId = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DemRumNode) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemRumNode) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DemRumNode) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DemRumNode) SetId(v string) {
	o.Id = &v
}

// GetQuery returns the Query field value.
func (o *DemRumNode) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *DemRumNode) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *DemRumNode) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DemRumNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["app_id"] = o.AppId
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["query"] = o.Query

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DemRumNode) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AppId *string `json:"app_id"`
		Id    *string `json:"id,omitempty"`
		Query *string `json:"query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AppId == nil {
		return fmt.Errorf("required field app_id missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"app_id", "id", "query"})
	} else {
		return err
	}
	o.AppId = *all.AppId
	o.Id = all.Id
	o.Query = *all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
