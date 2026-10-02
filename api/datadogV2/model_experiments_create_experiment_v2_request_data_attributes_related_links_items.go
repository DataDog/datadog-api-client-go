// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems External link associated with an experiment.
type ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems struct {
	// Link ID. Omit it when adding a link.
	Id datadog.NullableString `json:"id,omitempty"`
	// Optional display title.
	Title datadog.NullableString `json:"title,omitempty"`
	// Absolute URL.
	Url string `json:"url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems instantiates a new ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems(url string) *ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems {
	this := ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems{}
	this.Url = url
	return &this
}

// NewExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItemsWithDefaults instantiates a new ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItemsWithDefaults() *ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems {
	this := ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems) HasId() bool {
	return o != nil && o.Id.IsSet()
}

// SetId gets a reference to the given datadog.NullableString and assigns it to the Id field.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems) UnsetId() {
	o.Id.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems) HasTitle() bool {
	return o != nil && o.Title.IsSet()
}

// SetTitle gets a reference to the given datadog.NullableString and assigns it to the Title field.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems) SetTitle(v string) {
	o.Title.Set(&v)
}

// SetTitleNil sets the value for Title to be an explicit nil.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems) UnsetTitle() {
	o.Title.Unset()
}

// GetUrl returns the Url field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems) SetUrl(v string) {
	o.Url = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesRelatedLinksItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id    datadog.NullableString `json:"id,omitempty"`
		Title datadog.NullableString `json:"title,omitempty"`
		Url   *string                `json:"url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Url == nil {
		return fmt.Errorf("required field url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "title", "url"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Title = all.Title
	o.Url = *all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
