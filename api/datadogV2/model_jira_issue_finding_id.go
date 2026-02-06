// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JiraIssueFindingId
type JiraIssueFindingId struct {
	// Timestamp when the finding was discovered.
	Discovered int64 `json:"discovered"`
	// Identifier of the finding.
	Id string `json:"id"`
	// Resource associated with the finding.
	Resource string `json:"resource"`
	// Tags associated with the finding.
	Tags string `json:"tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJiraIssueFindingId instantiates a new JiraIssueFindingId object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJiraIssueFindingId(discovered int64, id string, resource string, tags string) *JiraIssueFindingId {
	this := JiraIssueFindingId{}
	this.Discovered = discovered
	this.Id = id
	this.Resource = resource
	this.Tags = tags
	return &this
}

// NewJiraIssueFindingIdWithDefaults instantiates a new JiraIssueFindingId object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJiraIssueFindingIdWithDefaults() *JiraIssueFindingId {
	this := JiraIssueFindingId{}
	return &this
}

// GetDiscovered returns the Discovered field value.
func (o *JiraIssueFindingId) GetDiscovered() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Discovered
}

// GetDiscoveredOk returns a tuple with the Discovered field value
// and a boolean to check if the value has been set.
func (o *JiraIssueFindingId) GetDiscoveredOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Discovered, true
}

// SetDiscovered sets field value.
func (o *JiraIssueFindingId) SetDiscovered(v int64) {
	o.Discovered = v
}

// GetId returns the Id field value.
func (o *JiraIssueFindingId) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JiraIssueFindingId) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *JiraIssueFindingId) SetId(v string) {
	o.Id = v
}

// GetResource returns the Resource field value.
func (o *JiraIssueFindingId) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *JiraIssueFindingId) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value.
func (o *JiraIssueFindingId) SetResource(v string) {
	o.Resource = v
}

// GetTags returns the Tags field value.
func (o *JiraIssueFindingId) GetTags() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *JiraIssueFindingId) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *JiraIssueFindingId) SetTags(v string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o JiraIssueFindingId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["discovered"] = o.Discovered
	toSerialize["id"] = o.Id
	toSerialize["resource"] = o.Resource
	toSerialize["tags"] = o.Tags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JiraIssueFindingId) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Discovered *int64  `json:"discovered"`
		Id         *string `json:"id"`
		Resource   *string `json:"resource"`
		Tags       *string `json:"tags"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Discovered == nil {
		return fmt.Errorf("required field discovered missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Resource == nil {
		return fmt.Errorf("required field resource missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"discovered", "id", "resource", "tags"})
	} else {
		return err
	}
	o.Discovered = *all.Discovered
	o.Id = *all.Id
	o.Resource = *all.Resource
	o.Tags = *all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
