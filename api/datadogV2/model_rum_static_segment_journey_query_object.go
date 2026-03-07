// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumStaticSegmentJourneyQueryObject The journey query object used to compute the static segment user list.
type RumStaticSegmentJourneyQueryObject struct {
	// The list of journey nodes defining the query.
	Nodes []RumStaticSegmentJourneyNode `json:"nodes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumStaticSegmentJourneyQueryObject instantiates a new RumStaticSegmentJourneyQueryObject object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumStaticSegmentJourneyQueryObject(nodes []RumStaticSegmentJourneyNode) *RumStaticSegmentJourneyQueryObject {
	this := RumStaticSegmentJourneyQueryObject{}
	this.Nodes = nodes
	return &this
}

// NewRumStaticSegmentJourneyQueryObjectWithDefaults instantiates a new RumStaticSegmentJourneyQueryObject object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumStaticSegmentJourneyQueryObjectWithDefaults() *RumStaticSegmentJourneyQueryObject {
	this := RumStaticSegmentJourneyQueryObject{}
	return &this
}

// GetNodes returns the Nodes field value.
func (o *RumStaticSegmentJourneyQueryObject) GetNodes() []RumStaticSegmentJourneyNode {
	if o == nil {
		var ret []RumStaticSegmentJourneyNode
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *RumStaticSegmentJourneyQueryObject) GetNodesOk() (*[]RumStaticSegmentJourneyNode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// SetNodes sets field value.
func (o *RumStaticSegmentJourneyQueryObject) SetNodes(v []RumStaticSegmentJourneyNode) {
	o.Nodes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumStaticSegmentJourneyQueryObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["nodes"] = o.Nodes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumStaticSegmentJourneyQueryObject) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Nodes *[]RumStaticSegmentJourneyNode `json:"nodes"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Nodes == nil {
		return fmt.Errorf("required field nodes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"nodes"})
	} else {
		return err
	}
	o.Nodes = *all.Nodes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
