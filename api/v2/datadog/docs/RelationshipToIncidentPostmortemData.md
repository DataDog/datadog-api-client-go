# RelationshipToIncidentPostmortemData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Id** | **string** | A unique identifier that represents the postmortem. | 
**Type** | [**IncidentPostmortemType**](IncidentPostmortemType.md) |  | [default to INCIDENTPOSTMORTEMTYPE_INCIDENT_POSTMORTEMS]

## Methods

### NewRelationshipToIncidentPostmortemData

`func NewRelationshipToIncidentPostmortemData(id string, type_ IncidentPostmortemType, ) *RelationshipToIncidentPostmortemData`

NewRelationshipToIncidentPostmortemData instantiates a new RelationshipToIncidentPostmortemData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipToIncidentPostmortemDataWithDefaults

`func NewRelationshipToIncidentPostmortemDataWithDefaults() *RelationshipToIncidentPostmortemData`

NewRelationshipToIncidentPostmortemDataWithDefaults instantiates a new RelationshipToIncidentPostmortemData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RelationshipToIncidentPostmortemData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationshipToIncidentPostmortemData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationshipToIncidentPostmortemData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *RelationshipToIncidentPostmortemData) GetType() IncidentPostmortemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipToIncidentPostmortemData) GetTypeOk() (*IncidentPostmortemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipToIncidentPostmortemData) SetType(v IncidentPostmortemType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


