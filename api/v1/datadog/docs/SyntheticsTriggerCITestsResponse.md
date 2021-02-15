# SyntheticsTriggerCITestsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locations** | Pointer to [**[]SyntheticsTriggerCITestLocation**](SyntheticsTriggerCITestLocation.md) | List of Synthetics locations. | [optional] 
**Results** | Pointer to [**[]SyntheticsTriggerCITestRunResult**](SyntheticsTriggerCITestRunResult.md) | Information about the tests runs. | [optional] 
**TriggeredCheckIds** | Pointer to **[]string** | The public IDs of the Synthetics test triggered. | [optional] 

## Methods

### NewSyntheticsTriggerCITestsResponse

`func NewSyntheticsTriggerCITestsResponse() *SyntheticsTriggerCITestsResponse`

NewSyntheticsTriggerCITestsResponse instantiates a new SyntheticsTriggerCITestsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsTriggerCITestsResponseWithDefaults

`func NewSyntheticsTriggerCITestsResponseWithDefaults() *SyntheticsTriggerCITestsResponse`

NewSyntheticsTriggerCITestsResponseWithDefaults instantiates a new SyntheticsTriggerCITestsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocations

`func (o *SyntheticsTriggerCITestsResponse) GetLocations() []SyntheticsTriggerCITestLocation`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *SyntheticsTriggerCITestsResponse) GetLocationsOk() (*[]SyntheticsTriggerCITestLocation, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *SyntheticsTriggerCITestsResponse) SetLocations(v []SyntheticsTriggerCITestLocation)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *SyntheticsTriggerCITestsResponse) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetResults

`func (o *SyntheticsTriggerCITestsResponse) GetResults() []SyntheticsTriggerCITestRunResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SyntheticsTriggerCITestsResponse) GetResultsOk() (*[]SyntheticsTriggerCITestRunResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SyntheticsTriggerCITestsResponse) SetResults(v []SyntheticsTriggerCITestRunResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *SyntheticsTriggerCITestsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTriggeredCheckIds

`func (o *SyntheticsTriggerCITestsResponse) GetTriggeredCheckIds() []string`

GetTriggeredCheckIds returns the TriggeredCheckIds field if non-nil, zero value otherwise.

### GetTriggeredCheckIdsOk

`func (o *SyntheticsTriggerCITestsResponse) GetTriggeredCheckIdsOk() (*[]string, bool)`

GetTriggeredCheckIdsOk returns a tuple with the TriggeredCheckIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredCheckIds

`func (o *SyntheticsTriggerCITestsResponse) SetTriggeredCheckIds(v []string)`

SetTriggeredCheckIds sets TriggeredCheckIds field to given value.

### HasTriggeredCheckIds

`func (o *SyntheticsTriggerCITestsResponse) HasTriggeredCheckIds() bool`

HasTriggeredCheckIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


