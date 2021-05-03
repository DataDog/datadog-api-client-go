# RelationshipToIncidentIntegrationMetadataData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Id** | **string** | A unique identifier that represents the integration metadata. | 
**Type** | [**IncidentIntegrationMetadataType**](IncidentIntegrationMetadataType.md) |  | [default to INCIDENTINTEGRATIONMETADATATYPE_INCIDENT_INTEGRATION_METADATA]

## Methods

### NewRelationshipToIncidentIntegrationMetadataData

`func NewRelationshipToIncidentIntegrationMetadataData(id string, type_ IncidentIntegrationMetadataType, ) *RelationshipToIncidentIntegrationMetadataData`

NewRelationshipToIncidentIntegrationMetadataData instantiates a new RelationshipToIncidentIntegrationMetadataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipToIncidentIntegrationMetadataDataWithDefaults

`func NewRelationshipToIncidentIntegrationMetadataDataWithDefaults() *RelationshipToIncidentIntegrationMetadataData`

NewRelationshipToIncidentIntegrationMetadataDataWithDefaults instantiates a new RelationshipToIncidentIntegrationMetadataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RelationshipToIncidentIntegrationMetadataData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationshipToIncidentIntegrationMetadataData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationshipToIncidentIntegrationMetadataData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *RelationshipToIncidentIntegrationMetadataData) GetType() IncidentIntegrationMetadataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipToIncidentIntegrationMetadataData) GetTypeOk() (*IncidentIntegrationMetadataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipToIncidentIntegrationMetadataData) SetType(v IncidentIntegrationMetadataType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


