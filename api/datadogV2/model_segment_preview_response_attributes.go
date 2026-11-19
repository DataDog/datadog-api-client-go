// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentPreviewResponseAttributes Members of the previewed segment.
type SegmentPreviewResponseAttributes struct {
	// Rows of the preview, each as a map of column names to values.
	Hits []map[string]interface{} `json:"hits"`
	// Total number of entities in the segment.
	TotalRows int64 `json:"total_rows"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSegmentPreviewResponseAttributes instantiates a new SegmentPreviewResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSegmentPreviewResponseAttributes(hits []map[string]interface{}, totalRows int64) *SegmentPreviewResponseAttributes {
	this := SegmentPreviewResponseAttributes{}
	this.Hits = hits
	this.TotalRows = totalRows
	return &this
}

// NewSegmentPreviewResponseAttributesWithDefaults instantiates a new SegmentPreviewResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSegmentPreviewResponseAttributesWithDefaults() *SegmentPreviewResponseAttributes {
	this := SegmentPreviewResponseAttributes{}
	return &this
}

// GetHits returns the Hits field value.
func (o *SegmentPreviewResponseAttributes) GetHits() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Hits
}

// GetHitsOk returns a tuple with the Hits field value
// and a boolean to check if the value has been set.
func (o *SegmentPreviewResponseAttributes) GetHitsOk() (*[]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hits, true
}

// SetHits sets field value.
func (o *SegmentPreviewResponseAttributes) SetHits(v []map[string]interface{}) {
	o.Hits = v
}

// GetTotalRows returns the TotalRows field value.
func (o *SegmentPreviewResponseAttributes) GetTotalRows() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalRows
}

// GetTotalRowsOk returns a tuple with the TotalRows field value
// and a boolean to check if the value has been set.
func (o *SegmentPreviewResponseAttributes) GetTotalRowsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalRows, true
}

// SetTotalRows sets field value.
func (o *SegmentPreviewResponseAttributes) SetTotalRows(v int64) {
	o.TotalRows = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SegmentPreviewResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["hits"] = o.Hits
	toSerialize["total_rows"] = o.TotalRows

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SegmentPreviewResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Hits      *[]map[string]interface{} `json:"hits"`
		TotalRows *int64                    `json:"total_rows"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Hits == nil {
		return fmt.Errorf("required field hits missing")
	}
	if all.TotalRows == nil {
		return fmt.Errorf("required field total_rows missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"hits", "total_rows"})
	} else {
		return err
	}
	o.Hits = *all.Hits
	o.TotalRows = *all.TotalRows

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
