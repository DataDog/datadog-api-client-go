// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentPreviewAttributes Segment definition to preview and the columns to return.
// Without columns, the response returns the user ID, or the account ID for account segments.
type SegmentPreviewAttributes struct {
	// Account profile columns to return.
	AccountStoreColumns []string `json:"account_store_columns,omitempty"`
	// Defines who belongs to the segment. It contains one or more blocks, up to 20 in total.
	// Use `combination` to combine named blocks with a boolean formula such as `A AND NOT B`.
	DataQuery SegmentDataQuery `json:"data_query"`
	// Maximum number of rows to return, up to 1000. Defaults to 1000.
	Limit *int64 `json:"limit,omitempty"`
	// Reference table columns to return.
	ReferenceTableColumns []string `json:"reference_table_columns,omitempty"`
	// User profile columns to return.
	UserStoreColumns []string `json:"user_store_columns,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSegmentPreviewAttributes instantiates a new SegmentPreviewAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSegmentPreviewAttributes(dataQuery SegmentDataQuery) *SegmentPreviewAttributes {
	this := SegmentPreviewAttributes{}
	this.DataQuery = dataQuery
	return &this
}

// NewSegmentPreviewAttributesWithDefaults instantiates a new SegmentPreviewAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSegmentPreviewAttributesWithDefaults() *SegmentPreviewAttributes {
	this := SegmentPreviewAttributes{}
	return &this
}

// GetAccountStoreColumns returns the AccountStoreColumns field value if set, zero value otherwise.
func (o *SegmentPreviewAttributes) GetAccountStoreColumns() []string {
	if o == nil || o.AccountStoreColumns == nil {
		var ret []string
		return ret
	}
	return o.AccountStoreColumns
}

// GetAccountStoreColumnsOk returns a tuple with the AccountStoreColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentPreviewAttributes) GetAccountStoreColumnsOk() (*[]string, bool) {
	if o == nil || o.AccountStoreColumns == nil {
		return nil, false
	}
	return &o.AccountStoreColumns, true
}

// HasAccountStoreColumns returns a boolean if a field has been set.
func (o *SegmentPreviewAttributes) HasAccountStoreColumns() bool {
	return o != nil && o.AccountStoreColumns != nil
}

// SetAccountStoreColumns gets a reference to the given []string and assigns it to the AccountStoreColumns field.
func (o *SegmentPreviewAttributes) SetAccountStoreColumns(v []string) {
	o.AccountStoreColumns = v
}

// GetDataQuery returns the DataQuery field value.
func (o *SegmentPreviewAttributes) GetDataQuery() SegmentDataQuery {
	if o == nil {
		var ret SegmentDataQuery
		return ret
	}
	return o.DataQuery
}

// GetDataQueryOk returns a tuple with the DataQuery field value
// and a boolean to check if the value has been set.
func (o *SegmentPreviewAttributes) GetDataQueryOk() (*SegmentDataQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataQuery, true
}

// SetDataQuery sets field value.
func (o *SegmentPreviewAttributes) SetDataQuery(v SegmentDataQuery) {
	o.DataQuery = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *SegmentPreviewAttributes) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentPreviewAttributes) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *SegmentPreviewAttributes) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *SegmentPreviewAttributes) SetLimit(v int64) {
	o.Limit = &v
}

// GetReferenceTableColumns returns the ReferenceTableColumns field value if set, zero value otherwise.
func (o *SegmentPreviewAttributes) GetReferenceTableColumns() []string {
	if o == nil || o.ReferenceTableColumns == nil {
		var ret []string
		return ret
	}
	return o.ReferenceTableColumns
}

// GetReferenceTableColumnsOk returns a tuple with the ReferenceTableColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentPreviewAttributes) GetReferenceTableColumnsOk() (*[]string, bool) {
	if o == nil || o.ReferenceTableColumns == nil {
		return nil, false
	}
	return &o.ReferenceTableColumns, true
}

// HasReferenceTableColumns returns a boolean if a field has been set.
func (o *SegmentPreviewAttributes) HasReferenceTableColumns() bool {
	return o != nil && o.ReferenceTableColumns != nil
}

// SetReferenceTableColumns gets a reference to the given []string and assigns it to the ReferenceTableColumns field.
func (o *SegmentPreviewAttributes) SetReferenceTableColumns(v []string) {
	o.ReferenceTableColumns = v
}

// GetUserStoreColumns returns the UserStoreColumns field value if set, zero value otherwise.
func (o *SegmentPreviewAttributes) GetUserStoreColumns() []string {
	if o == nil || o.UserStoreColumns == nil {
		var ret []string
		return ret
	}
	return o.UserStoreColumns
}

// GetUserStoreColumnsOk returns a tuple with the UserStoreColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentPreviewAttributes) GetUserStoreColumnsOk() (*[]string, bool) {
	if o == nil || o.UserStoreColumns == nil {
		return nil, false
	}
	return &o.UserStoreColumns, true
}

// HasUserStoreColumns returns a boolean if a field has been set.
func (o *SegmentPreviewAttributes) HasUserStoreColumns() bool {
	return o != nil && o.UserStoreColumns != nil
}

// SetUserStoreColumns gets a reference to the given []string and assigns it to the UserStoreColumns field.
func (o *SegmentPreviewAttributes) SetUserStoreColumns(v []string) {
	o.UserStoreColumns = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SegmentPreviewAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountStoreColumns != nil {
		toSerialize["account_store_columns"] = o.AccountStoreColumns
	}
	toSerialize["data_query"] = o.DataQuery
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.ReferenceTableColumns != nil {
		toSerialize["reference_table_columns"] = o.ReferenceTableColumns
	}
	if o.UserStoreColumns != nil {
		toSerialize["user_store_columns"] = o.UserStoreColumns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SegmentPreviewAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountStoreColumns   []string          `json:"account_store_columns,omitempty"`
		DataQuery             *SegmentDataQuery `json:"data_query"`
		Limit                 *int64            `json:"limit,omitempty"`
		ReferenceTableColumns []string          `json:"reference_table_columns,omitempty"`
		UserStoreColumns      []string          `json:"user_store_columns,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataQuery == nil {
		return fmt.Errorf("required field data_query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_store_columns", "data_query", "limit", "reference_table_columns", "user_store_columns"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccountStoreColumns = all.AccountStoreColumns
	if all.DataQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DataQuery = *all.DataQuery
	o.Limit = all.Limit
	o.ReferenceTableColumns = all.ReferenceTableColumns
	o.UserStoreColumns = all.UserStoreColumns

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
