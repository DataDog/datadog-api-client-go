// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DdsqlTabularQueryRequestAttributes Attributes describing the DDSQL query to execute.
type DdsqlTabularQueryRequestAttributes struct {
	// The DDSQL statement to execute. DDSQL is Datadog's SQL dialect, which is a subset
	// of PostgreSQL, scoped to Datadog data sources.
	Query string `json:"query"`
	// Cap on the number of rows returned. Defaults to 5,000 when omitted. Must be
	// between 1 and 10,000 inclusive; values outside this range are rejected with 400.
	RowLimit *int64 `json:"row_limit,omitempty"`
	// Time window scoping the underlying data sources, expressed in Unix milliseconds
	// since the epoch. Inclusive on `from_timestamp`, exclusive on `to_timestamp`.
	// Results from static tables (for example, `dd.hosts`) are not affected by the
	// time window, but the field must still be provided.
	Time DdsqlTabularQueryTimeWindow `json:"time"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDdsqlTabularQueryRequestAttributes instantiates a new DdsqlTabularQueryRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDdsqlTabularQueryRequestAttributes(query string, time DdsqlTabularQueryTimeWindow) *DdsqlTabularQueryRequestAttributes {
	this := DdsqlTabularQueryRequestAttributes{}
	this.Query = query
	this.Time = time
	return &this
}

// NewDdsqlTabularQueryRequestAttributesWithDefaults instantiates a new DdsqlTabularQueryRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDdsqlTabularQueryRequestAttributesWithDefaults() *DdsqlTabularQueryRequestAttributes {
	this := DdsqlTabularQueryRequestAttributes{}
	return &this
}

// GetQuery returns the Query field value.
func (o *DdsqlTabularQueryRequestAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *DdsqlTabularQueryRequestAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *DdsqlTabularQueryRequestAttributes) SetQuery(v string) {
	o.Query = v
}

// GetRowLimit returns the RowLimit field value if set, zero value otherwise.
func (o *DdsqlTabularQueryRequestAttributes) GetRowLimit() int64 {
	if o == nil || o.RowLimit == nil {
		var ret int64
		return ret
	}
	return *o.RowLimit
}

// GetRowLimitOk returns a tuple with the RowLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdsqlTabularQueryRequestAttributes) GetRowLimitOk() (*int64, bool) {
	if o == nil || o.RowLimit == nil {
		return nil, false
	}
	return o.RowLimit, true
}

// HasRowLimit returns a boolean if a field has been set.
func (o *DdsqlTabularQueryRequestAttributes) HasRowLimit() bool {
	return o != nil && o.RowLimit != nil
}

// SetRowLimit gets a reference to the given int64 and assigns it to the RowLimit field.
func (o *DdsqlTabularQueryRequestAttributes) SetRowLimit(v int64) {
	o.RowLimit = &v
}

// GetTime returns the Time field value.
func (o *DdsqlTabularQueryRequestAttributes) GetTime() DdsqlTabularQueryTimeWindow {
	if o == nil {
		var ret DdsqlTabularQueryTimeWindow
		return ret
	}
	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *DdsqlTabularQueryRequestAttributes) GetTimeOk() (*DdsqlTabularQueryTimeWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value.
func (o *DdsqlTabularQueryRequestAttributes) SetTime(v DdsqlTabularQueryTimeWindow) {
	o.Time = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DdsqlTabularQueryRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["query"] = o.Query
	if o.RowLimit != nil {
		toSerialize["row_limit"] = o.RowLimit
	}
	toSerialize["time"] = o.Time

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DdsqlTabularQueryRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Query    *string                      `json:"query"`
		RowLimit *int64                       `json:"row_limit,omitempty"`
		Time     *DdsqlTabularQueryTimeWindow `json:"time"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.Time == nil {
		return fmt.Errorf("required field time missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"query", "row_limit", "time"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Query = *all.Query
	o.RowLimit = all.RowLimit
	if all.Time.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Time = *all.Time

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
