// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookSearchMetadata Metadata about the notebook.
type NotebookSearchMetadata struct {
	// User information.
	Author NotebookSearchUser `json:"author"`
	// Number of cells in the notebook.
	CellCount int64 `json:"cell_count"`
	// Time at which the notebook was created.
	CreatedAt time.Time `json:"created_at"`
	// Time at which the notebook was deleted, or null if not deleted.
	DeletedAt datadog.NullableTime `json:"deleted_at"`
	// Experience type of the notebook.
	ExperienceType datadog.NullableString `json:"experience_type"`
	// Whether the notebook has computational cells.
	HasComputationalCells bool `json:"has_computational_cells"`
	// Whether the notebook is favorited by the user.
	IsFavorited bool `json:"is_favorited"`
	// Whether the notebook is a template.
	IsTemplate bool `json:"is_template"`
	// Time at which the notebook was last updated.
	ModifiedAt time.Time `json:"modified_at"`
	// Status of the notebook.
	Status string `json:"status"`
	// Whether the notebook can take a snapshot.
	TakeSnapshots bool `json:"take_snapshots"`
	// Notebook type.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNotebookSearchMetadata instantiates a new NotebookSearchMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotebookSearchMetadata(author NotebookSearchUser, cellCount int64, createdAt time.Time, deletedAt datadog.NullableTime, experienceType datadog.NullableString, hasComputationalCells bool, isFavorited bool, isTemplate bool, modifiedAt time.Time, status string, takeSnapshots bool, typeVar string) *NotebookSearchMetadata {
	this := NotebookSearchMetadata{}
	this.Author = author
	this.CellCount = cellCount
	this.CreatedAt = createdAt
	this.DeletedAt = deletedAt
	this.ExperienceType = experienceType
	this.HasComputationalCells = hasComputationalCells
	this.IsFavorited = isFavorited
	this.IsTemplate = isTemplate
	this.ModifiedAt = modifiedAt
	this.Status = status
	this.TakeSnapshots = takeSnapshots
	this.Type = typeVar
	return &this
}

// NewNotebookSearchMetadataWithDefaults instantiates a new NotebookSearchMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotebookSearchMetadataWithDefaults() *NotebookSearchMetadata {
	this := NotebookSearchMetadata{}
	return &this
}

// GetAuthor returns the Author field value.
func (o *NotebookSearchMetadata) GetAuthor() NotebookSearchUser {
	if o == nil {
		var ret NotebookSearchUser
		return ret
	}
	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *NotebookSearchMetadata) GetAuthorOk() (*NotebookSearchUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value.
func (o *NotebookSearchMetadata) SetAuthor(v NotebookSearchUser) {
	o.Author = v
}

// GetCellCount returns the CellCount field value.
func (o *NotebookSearchMetadata) GetCellCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CellCount
}

// GetCellCountOk returns a tuple with the CellCount field value
// and a boolean to check if the value has been set.
func (o *NotebookSearchMetadata) GetCellCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CellCount, true
}

// SetCellCount sets field value.
func (o *NotebookSearchMetadata) SetCellCount(v int64) {
	o.CellCount = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *NotebookSearchMetadata) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NotebookSearchMetadata) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *NotebookSearchMetadata) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDeletedAt returns the DeletedAt field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *NotebookSearchMetadata) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *NotebookSearchMetadata) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// SetDeletedAt sets field value.
func (o *NotebookSearchMetadata) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// GetExperienceType returns the ExperienceType field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *NotebookSearchMetadata) GetExperienceType() string {
	if o == nil || o.ExperienceType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExperienceType.Get()
}

// GetExperienceTypeOk returns a tuple with the ExperienceType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *NotebookSearchMetadata) GetExperienceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExperienceType.Get(), o.ExperienceType.IsSet()
}

// SetExperienceType sets field value.
func (o *NotebookSearchMetadata) SetExperienceType(v string) {
	o.ExperienceType.Set(&v)
}

// GetHasComputationalCells returns the HasComputationalCells field value.
func (o *NotebookSearchMetadata) GetHasComputationalCells() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HasComputationalCells
}

// GetHasComputationalCellsOk returns a tuple with the HasComputationalCells field value
// and a boolean to check if the value has been set.
func (o *NotebookSearchMetadata) GetHasComputationalCellsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasComputationalCells, true
}

// SetHasComputationalCells sets field value.
func (o *NotebookSearchMetadata) SetHasComputationalCells(v bool) {
	o.HasComputationalCells = v
}

// GetIsFavorited returns the IsFavorited field value.
func (o *NotebookSearchMetadata) GetIsFavorited() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsFavorited
}

// GetIsFavoritedOk returns a tuple with the IsFavorited field value
// and a boolean to check if the value has been set.
func (o *NotebookSearchMetadata) GetIsFavoritedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFavorited, true
}

// SetIsFavorited sets field value.
func (o *NotebookSearchMetadata) SetIsFavorited(v bool) {
	o.IsFavorited = v
}

// GetIsTemplate returns the IsTemplate field value.
func (o *NotebookSearchMetadata) GetIsTemplate() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsTemplate
}

// GetIsTemplateOk returns a tuple with the IsTemplate field value
// and a boolean to check if the value has been set.
func (o *NotebookSearchMetadata) GetIsTemplateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTemplate, true
}

// SetIsTemplate sets field value.
func (o *NotebookSearchMetadata) SetIsTemplate(v bool) {
	o.IsTemplate = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *NotebookSearchMetadata) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *NotebookSearchMetadata) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *NotebookSearchMetadata) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetStatus returns the Status field value.
func (o *NotebookSearchMetadata) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *NotebookSearchMetadata) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *NotebookSearchMetadata) SetStatus(v string) {
	o.Status = v
}

// GetTakeSnapshots returns the TakeSnapshots field value.
func (o *NotebookSearchMetadata) GetTakeSnapshots() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.TakeSnapshots
}

// GetTakeSnapshotsOk returns a tuple with the TakeSnapshots field value
// and a boolean to check if the value has been set.
func (o *NotebookSearchMetadata) GetTakeSnapshotsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TakeSnapshots, true
}

// SetTakeSnapshots sets field value.
func (o *NotebookSearchMetadata) SetTakeSnapshots(v bool) {
	o.TakeSnapshots = v
}

// GetType returns the Type field value.
func (o *NotebookSearchMetadata) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotebookSearchMetadata) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *NotebookSearchMetadata) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotebookSearchMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["author"] = o.Author
	toSerialize["cell_count"] = o.CellCount
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["deleted_at"] = o.DeletedAt.Get()
	toSerialize["experience_type"] = o.ExperienceType.Get()
	toSerialize["has_computational_cells"] = o.HasComputationalCells
	toSerialize["is_favorited"] = o.IsFavorited
	toSerialize["is_template"] = o.IsTemplate
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["status"] = o.Status
	toSerialize["take_snapshots"] = o.TakeSnapshots
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotebookSearchMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Author                *NotebookSearchUser    `json:"author"`
		CellCount             *int64                 `json:"cell_count"`
		CreatedAt             *time.Time             `json:"created_at"`
		DeletedAt             datadog.NullableTime   `json:"deleted_at"`
		ExperienceType        datadog.NullableString `json:"experience_type"`
		HasComputationalCells *bool                  `json:"has_computational_cells"`
		IsFavorited           *bool                  `json:"is_favorited"`
		IsTemplate            *bool                  `json:"is_template"`
		ModifiedAt            *time.Time             `json:"modified_at"`
		Status                *string                `json:"status"`
		TakeSnapshots         *bool                  `json:"take_snapshots"`
		Type                  *string                `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Author == nil {
		return fmt.Errorf("required field author missing")
	}
	if all.CellCount == nil {
		return fmt.Errorf("required field cell_count missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if !all.DeletedAt.IsSet() {
		return fmt.Errorf("required field deleted_at missing")
	}
	if !all.ExperienceType.IsSet() {
		return fmt.Errorf("required field experience_type missing")
	}
	if all.HasComputationalCells == nil {
		return fmt.Errorf("required field has_computational_cells missing")
	}
	if all.IsFavorited == nil {
		return fmt.Errorf("required field is_favorited missing")
	}
	if all.IsTemplate == nil {
		return fmt.Errorf("required field is_template missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.TakeSnapshots == nil {
		return fmt.Errorf("required field take_snapshots missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author", "cell_count", "created_at", "deleted_at", "experience_type", "has_computational_cells", "is_favorited", "is_template", "modified_at", "status", "take_snapshots", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Author.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Author = *all.Author
	o.CellCount = *all.CellCount
	o.CreatedAt = *all.CreatedAt
	o.DeletedAt = all.DeletedAt
	o.ExperienceType = all.ExperienceType
	o.HasComputationalCells = *all.HasComputationalCells
	o.IsFavorited = *all.IsFavorited
	o.IsTemplate = *all.IsTemplate
	o.ModifiedAt = *all.ModifiedAt
	o.Status = *all.Status
	o.TakeSnapshots = *all.TakeSnapshots
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
