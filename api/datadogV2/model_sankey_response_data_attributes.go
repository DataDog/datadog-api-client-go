// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyResponseDataAttributes
type SankeyResponseDataAttributes struct {
	//
	Links []SankeyResponseDataAttributesLinksItems `json:"links,omitempty"`
	//
	Nodes []SankeyResponseDataAttributesNodesItems `json:"nodes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSankeyResponseDataAttributes instantiates a new SankeyResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSankeyResponseDataAttributes() *SankeyResponseDataAttributes {
	this := SankeyResponseDataAttributes{}
	return &this
}

// NewSankeyResponseDataAttributesWithDefaults instantiates a new SankeyResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSankeyResponseDataAttributesWithDefaults() *SankeyResponseDataAttributes {
	this := SankeyResponseDataAttributes{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributes) GetLinks() []SankeyResponseDataAttributesLinksItems {
	if o == nil || o.Links == nil {
		var ret []SankeyResponseDataAttributesLinksItems
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributes) GetLinksOk() (*[]SankeyResponseDataAttributesLinksItems, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return &o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributes) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given []SankeyResponseDataAttributesLinksItems and assigns it to the Links field.
func (o *SankeyResponseDataAttributes) SetLinks(v []SankeyResponseDataAttributesLinksItems) {
	o.Links = v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *SankeyResponseDataAttributes) GetNodes() []SankeyResponseDataAttributesNodesItems {
	if o == nil || o.Nodes == nil {
		var ret []SankeyResponseDataAttributesNodesItems
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyResponseDataAttributes) GetNodesOk() (*[]SankeyResponseDataAttributesNodesItems, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *SankeyResponseDataAttributes) HasNodes() bool {
	return o != nil && o.Nodes != nil
}

// SetNodes gets a reference to the given []SankeyResponseDataAttributesNodesItems and assigns it to the Nodes field.
func (o *SankeyResponseDataAttributes) SetNodes(v []SankeyResponseDataAttributesNodesItems) {
	o.Nodes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SankeyResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Nodes != nil {
		toSerialize["nodes"] = o.Nodes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SankeyResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Links []SankeyResponseDataAttributesLinksItems `json:"links,omitempty"`
		Nodes []SankeyResponseDataAttributesNodesItems `json:"nodes,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"links", "nodes"})
	} else {
		return err
	}
	o.Links = all.Links
	o.Nodes = all.Nodes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
