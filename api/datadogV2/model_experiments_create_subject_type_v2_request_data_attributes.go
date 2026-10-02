// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateSubjectTypeV2RequestDataAttributes Name and data field mappings for the new subject type.
type ExperimentsCreateSubjectTypeV2RequestDataAttributes struct {
	// Metadata associated with migration of this resource.
	MigrationMetadata interface{} `json:"migration_metadata,omitempty"`
	// Display name of the subject type.
	Name string `json:"name"`
	// Product Analytics attribute used to identify subjects of this type.
	ProductAnalyticsAttribute *string `json:"product_analytics_attribute,omitempty"`
	// Warehouse column names associated with this subject type.
	WarehouseColumnNames []string `json:"warehouse_column_names,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsCreateSubjectTypeV2RequestDataAttributes instantiates a new ExperimentsCreateSubjectTypeV2RequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCreateSubjectTypeV2RequestDataAttributes(name string) *ExperimentsCreateSubjectTypeV2RequestDataAttributes {
	this := ExperimentsCreateSubjectTypeV2RequestDataAttributes{}
	this.Name = name
	return &this
}

// NewExperimentsCreateSubjectTypeV2RequestDataAttributesWithDefaults instantiates a new ExperimentsCreateSubjectTypeV2RequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCreateSubjectTypeV2RequestDataAttributesWithDefaults() *ExperimentsCreateSubjectTypeV2RequestDataAttributes {
	this := ExperimentsCreateSubjectTypeV2RequestDataAttributes{}
	return &this
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise.
func (o *ExperimentsCreateSubjectTypeV2RequestDataAttributes) GetMigrationMetadata() interface{} {
	if o == nil || o.MigrationMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateSubjectTypeV2RequestDataAttributes) GetMigrationMetadataOk() (*interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsCreateSubjectTypeV2RequestDataAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsCreateSubjectTypeV2RequestDataAttributes) SetMigrationMetadata(v interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value.
func (o *ExperimentsCreateSubjectTypeV2RequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateSubjectTypeV2RequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ExperimentsCreateSubjectTypeV2RequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetProductAnalyticsAttribute returns the ProductAnalyticsAttribute field value if set, zero value otherwise.
func (o *ExperimentsCreateSubjectTypeV2RequestDataAttributes) GetProductAnalyticsAttribute() string {
	if o == nil || o.ProductAnalyticsAttribute == nil {
		var ret string
		return ret
	}
	return *o.ProductAnalyticsAttribute
}

// GetProductAnalyticsAttributeOk returns a tuple with the ProductAnalyticsAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateSubjectTypeV2RequestDataAttributes) GetProductAnalyticsAttributeOk() (*string, bool) {
	if o == nil || o.ProductAnalyticsAttribute == nil {
		return nil, false
	}
	return o.ProductAnalyticsAttribute, true
}

// HasProductAnalyticsAttribute returns a boolean if a field has been set.
func (o *ExperimentsCreateSubjectTypeV2RequestDataAttributes) HasProductAnalyticsAttribute() bool {
	return o != nil && o.ProductAnalyticsAttribute != nil
}

// SetProductAnalyticsAttribute gets a reference to the given string and assigns it to the ProductAnalyticsAttribute field.
func (o *ExperimentsCreateSubjectTypeV2RequestDataAttributes) SetProductAnalyticsAttribute(v string) {
	o.ProductAnalyticsAttribute = &v
}

// GetWarehouseColumnNames returns the WarehouseColumnNames field value if set, zero value otherwise.
func (o *ExperimentsCreateSubjectTypeV2RequestDataAttributes) GetWarehouseColumnNames() []string {
	if o == nil || o.WarehouseColumnNames == nil {
		var ret []string
		return ret
	}
	return o.WarehouseColumnNames
}

// GetWarehouseColumnNamesOk returns a tuple with the WarehouseColumnNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateSubjectTypeV2RequestDataAttributes) GetWarehouseColumnNamesOk() (*[]string, bool) {
	if o == nil || o.WarehouseColumnNames == nil {
		return nil, false
	}
	return &o.WarehouseColumnNames, true
}

// HasWarehouseColumnNames returns a boolean if a field has been set.
func (o *ExperimentsCreateSubjectTypeV2RequestDataAttributes) HasWarehouseColumnNames() bool {
	return o != nil && o.WarehouseColumnNames != nil
}

// SetWarehouseColumnNames gets a reference to the given []string and assigns it to the WarehouseColumnNames field.
func (o *ExperimentsCreateSubjectTypeV2RequestDataAttributes) SetWarehouseColumnNames(v []string) {
	o.WarehouseColumnNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCreateSubjectTypeV2RequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.MigrationMetadata != nil {
		toSerialize["migration_metadata"] = o.MigrationMetadata
	}
	toSerialize["name"] = o.Name
	if o.ProductAnalyticsAttribute != nil {
		toSerialize["product_analytics_attribute"] = o.ProductAnalyticsAttribute
	}
	if o.WarehouseColumnNames != nil {
		toSerialize["warehouse_column_names"] = o.WarehouseColumnNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCreateSubjectTypeV2RequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MigrationMetadata         interface{} `json:"migration_metadata,omitempty"`
		Name                      *string     `json:"name"`
		ProductAnalyticsAttribute *string     `json:"product_analytics_attribute,omitempty"`
		WarehouseColumnNames      []string    `json:"warehouse_column_names,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"migration_metadata", "name", "product_analytics_attribute", "warehouse_column_names"})
	} else {
		return err
	}
	o.MigrationMetadata = all.MigrationMetadata
	o.Name = *all.Name
	o.ProductAnalyticsAttribute = all.ProductAnalyticsAttribute
	o.WarehouseColumnNames = all.WarehouseColumnNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
