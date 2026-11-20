// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricHostmapRequests Test-only generated schema for the dashboard versioning proof of concept.
type MetricHostmapRequests struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Fill *HostmapTileDefFillRequest `json:"fill,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Size *HostmapTileDefSizeRequest `json:"size,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMetricHostmapRequests instantiates a new MetricHostmapRequests object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricHostmapRequests() *MetricHostmapRequests {
	this := MetricHostmapRequests{}
	return &this
}

// NewMetricHostmapRequestsWithDefaults instantiates a new MetricHostmapRequests object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricHostmapRequestsWithDefaults() *MetricHostmapRequests {
	this := MetricHostmapRequests{}
	return &this
}

// GetFill returns the Fill field value if set, zero value otherwise.
func (o *MetricHostmapRequests) GetFill() HostmapTileDefFillRequest {
	if o == nil || o.Fill == nil {
		var ret HostmapTileDefFillRequest
		return ret
	}
	return *o.Fill
}

// GetFillOk returns a tuple with the Fill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricHostmapRequests) GetFillOk() (*HostmapTileDefFillRequest, bool) {
	if o == nil || o.Fill == nil {
		return nil, false
	}
	return o.Fill, true
}

// HasFill returns a boolean if a field has been set.
func (o *MetricHostmapRequests) HasFill() bool {
	return o != nil && o.Fill != nil
}

// SetFill gets a reference to the given HostmapTileDefFillRequest and assigns it to the Fill field.
func (o *MetricHostmapRequests) SetFill(v HostmapTileDefFillRequest) {
	o.Fill = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *MetricHostmapRequests) GetSize() HostmapTileDefSizeRequest {
	if o == nil || o.Size == nil {
		var ret HostmapTileDefSizeRequest
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricHostmapRequests) GetSizeOk() (*HostmapTileDefSizeRequest, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *MetricHostmapRequests) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given HostmapTileDefSizeRequest and assigns it to the Size field.
func (o *MetricHostmapRequests) SetSize(v HostmapTileDefSizeRequest) {
	o.Size = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricHostmapRequests) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Fill != nil {
		toSerialize["fill"] = o.Fill
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricHostmapRequests) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fill *HostmapTileDefFillRequest `json:"fill,omitempty"`
		Size *HostmapTileDefSizeRequest `json:"size,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.Fill != nil && all.Fill.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Fill = all.Fill
	if all.Size != nil && all.Size.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Size = all.Size

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
