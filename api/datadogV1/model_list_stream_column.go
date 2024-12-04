// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListStreamColumn Widget column.
type ListStreamColumn struct {
	// Widget column field.
	Field string `json:"field"`
	// Identifies the clustering pattern field column, usable only with logs_pattern_stream.
	IsClusteringPatternFieldPath *bool `json:"is_clustering_pattern_field_path,omitempty"`
	// Widget column width.
	Width ListStreamColumnWidth `json:"width"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListStreamColumn instantiates a new ListStreamColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListStreamColumn(field string, width ListStreamColumnWidth) *ListStreamColumn {
	this := ListStreamColumn{}
	this.Field = field
	this.Width = width
	return &this
}

// NewListStreamColumnWithDefaults instantiates a new ListStreamColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListStreamColumnWithDefaults() *ListStreamColumn {
	this := ListStreamColumn{}
	return &this
}

// GetField returns the Field field value.
func (o *ListStreamColumn) GetField() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *ListStreamColumn) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value.
func (o *ListStreamColumn) SetField(v string) {
	o.Field = v
}

// GetIsClusteringPatternFieldPath returns the IsClusteringPatternFieldPath field value if set, zero value otherwise.
func (o *ListStreamColumn) GetIsClusteringPatternFieldPath() bool {
	if o == nil || o.IsClusteringPatternFieldPath == nil {
		var ret bool
		return ret
	}
	return *o.IsClusteringPatternFieldPath
}

// GetIsClusteringPatternFieldPathOk returns a tuple with the IsClusteringPatternFieldPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamColumn) GetIsClusteringPatternFieldPathOk() (*bool, bool) {
	if o == nil || o.IsClusteringPatternFieldPath == nil {
		return nil, false
	}
	return o.IsClusteringPatternFieldPath, true
}

// HasIsClusteringPatternFieldPath returns a boolean if a field has been set.
func (o *ListStreamColumn) HasIsClusteringPatternFieldPath() bool {
	return o != nil && o.IsClusteringPatternFieldPath != nil
}

// SetIsClusteringPatternFieldPath gets a reference to the given bool and assigns it to the IsClusteringPatternFieldPath field.
func (o *ListStreamColumn) SetIsClusteringPatternFieldPath(v bool) {
	o.IsClusteringPatternFieldPath = &v
}

// GetWidth returns the Width field value.
func (o *ListStreamColumn) GetWidth() ListStreamColumnWidth {
	if o == nil {
		var ret ListStreamColumnWidth
		return ret
	}
	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *ListStreamColumn) GetWidthOk() (*ListStreamColumnWidth, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value.
func (o *ListStreamColumn) SetWidth(v ListStreamColumnWidth) {
	o.Width = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListStreamColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["field"] = o.Field
	if o.IsClusteringPatternFieldPath != nil {
		toSerialize["is_clustering_pattern_field_path"] = o.IsClusteringPatternFieldPath
	}
	toSerialize["width"] = o.Width

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListStreamColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Field                        *string                `json:"field"`
		IsClusteringPatternFieldPath *bool                  `json:"is_clustering_pattern_field_path,omitempty"`
		Width                        *ListStreamColumnWidth `json:"width"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Field == nil {
		return fmt.Errorf("required field field missing")
	}
	if all.Width == nil {
		return fmt.Errorf("required field width missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"field", "is_clustering_pattern_field_path", "width"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Field = *all.Field
	o.IsClusteringPatternFieldPath = all.IsClusteringPatternFieldPath
	if !all.Width.IsValid() {
		hasInvalidField = true
	} else {
		o.Width = *all.Width
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
