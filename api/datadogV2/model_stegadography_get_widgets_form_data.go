// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	_io "io"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StegadographyGetWidgetsFormData The form data submitted to look up widgets from a watermarked image.
type StegadographyGetWidgetsFormData struct {
	// A PNG image (for example, a dashboard screenshot) containing embedded widget watermarks.
	Image *_io.Reader `json:"image,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStegadographyGetWidgetsFormData instantiates a new StegadographyGetWidgetsFormData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStegadographyGetWidgetsFormData() *StegadographyGetWidgetsFormData {
	this := StegadographyGetWidgetsFormData{}
	return &this
}

// NewStegadographyGetWidgetsFormDataWithDefaults instantiates a new StegadographyGetWidgetsFormData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStegadographyGetWidgetsFormDataWithDefaults() *StegadographyGetWidgetsFormData {
	this := StegadographyGetWidgetsFormData{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *StegadographyGetWidgetsFormData) GetImage() _io.Reader {
	if o == nil || o.Image == nil {
		var ret _io.Reader
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StegadographyGetWidgetsFormData) GetImageOk() (*_io.Reader, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *StegadographyGetWidgetsFormData) HasImage() bool {
	return o != nil && o.Image != nil
}

// SetImage gets a reference to the given _io.Reader and assigns it to the Image field.
func (o *StegadographyGetWidgetsFormData) SetImage(v _io.Reader) {
	o.Image = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o StegadographyGetWidgetsFormData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StegadographyGetWidgetsFormData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Image *_io.Reader `json:"image,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"image"})
	} else {
		return err
	}
	o.Image = all.Image

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
