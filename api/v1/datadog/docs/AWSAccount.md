# AwsAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | **string** | Your AWS access key ID. Only required if your AWS account is a GovCloud or China account. | [optional] 
**AccountId** | **string** | Your AWS Account ID without dashes. | [optional] 
**AccountSpecificNamespaceRules** | **map[string]bool** | An object, (in the form &#x60;{\&quot;namespace1\&quot;:true/false, \&quot;namespace2\&quot;:true/false}&#x60;), that enables or disables metric collection for specific AWS namespaces for this AWS account only. | [optional] 
**ExcludedRegions** | **[]string** | An array of AWS regions to exclude from metrics collection. | [optional] 
**FilterTags** | **[]string** | The array of EC2 tags (in the form &#x60;key:value&#x60;) defines a filter that Datadog uses when collecting metrics from EC2. Wildcards, such as &#x60;?&#x60; (for single characters) and &#x60;*&#x60; (for multiple characters) can also be used. Only hosts that match one of the defined tags will be imported into Datadog. The rest will be ignored. Host matching a given tag can also be excluded by adding &#x60;!&#x60; before the tag. For example, &#x60;env:production,instance-type:c1.*,!region:us-east-1&#x60; | [optional] 
**HostTags** | **[]string** | Array of tags (in the form &#x60;key:value&#x60;) to add to all hosts and metrics reporting through this integration. | [optional] 
**RoleName** | **string** | Your Datadog role delegation name. | [optional] 
**SecretAccessKey** | **string** | Your AWS secret access key. Only required if your AWS account is a GovCloud or China account. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


