// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSNamespace The namespace associated with the tag filter entry.
type AWSNamespace string

// List of AWSNamespace.
const (
	AWSNAMESPACE_API_GATEWAY                    AWSNamespace = "api_gateway"
	AWSNAMESPACE_APPLICATION_ELB                AWSNamespace = "application_elb"
	AWSNAMESPACE_APPRUNNER                      AWSNamespace = "apprunner"
	AWSNAMESPACE_APPSTREAM                      AWSNamespace = "appstream"
	AWSNAMESPACE_APPSYNC                        AWSNamespace = "appsync"
	AWSNAMESPACE_ATHENA                         AWSNamespace = "athena"
	AWSNAMESPACE_AUTO_SCALING                   AWSNamespace = "auto_scaling"
	AWSNAMESPACE_BACKUP                         AWSNamespace = "backup"
	AWSNAMESPACE_BEDROCK                        AWSNamespace = "bedrock"
	AWSNAMESPACE_BILLING                        AWSNamespace = "billing"
	AWSNAMESPACE_BUDGETING                      AWSNamespace = "budgeting"
	AWSNAMESPACE_CERTIFICATEMANAGER             AWSNamespace = "certificatemanager"
	AWSNAMESPACE_CLOUDFRONT                     AWSNamespace = "cloudfront"
	AWSNAMESPACE_CLOUDHSM                       AWSNamespace = "cloudhsm"
	AWSNAMESPACE_CLOUDSEARCH                    AWSNamespace = "cloudsearch"
	AWSNAMESPACE_CLOUDWATCH_EVENTS              AWSNamespace = "cloudwatch_events"
	AWSNAMESPACE_CLOUDWATCH_LOGS                AWSNamespace = "cloudwatch_logs"
	AWSNAMESPACE_CODEBUILD                      AWSNamespace = "codebuild"
	AWSNAMESPACE_CODEWHISPERER                  AWSNamespace = "codewhisperer"
	AWSNAMESPACE_COGNITO                        AWSNamespace = "cognito"
	AWSNAMESPACE_COLLECT_CUSTOM_METRICS         AWSNamespace = "collect_custom_metrics"
	AWSNAMESPACE_CONFIG                         AWSNamespace = "config"
	AWSNAMESPACE_CONNECT                        AWSNamespace = "connect"
	AWSNAMESPACE_CRAWL_ALARMS                   AWSNamespace = "crawl_alarms"
	AWSNAMESPACE_CUSTOM                         AWSNamespace = "custom"
	AWSNAMESPACE_DIRECTCONNECT                  AWSNamespace = "directconnect"
	AWSNAMESPACE_DMS                            AWSNamespace = "dms"
	AWSNAMESPACE_DOCUMENTDB                     AWSNamespace = "documentdb"
	AWSNAMESPACE_DYNAMODB                       AWSNamespace = "dynamodb"
	AWSNAMESPACE_DYNAMODBACCELERATOR            AWSNamespace = "dynamodbaccelerator"
	AWSNAMESPACE_EBS                            AWSNamespace = "ebs"
	AWSNAMESPACE_EC2                            AWSNamespace = "ec2"
	AWSNAMESPACE_EC2API                         AWSNamespace = "ec2api"
	AWSNAMESPACE_EC2SPOT                        AWSNamespace = "ec2spot"
	AWSNAMESPACE_ECR                            AWSNamespace = "ecr"
	AWSNAMESPACE_ECS                            AWSNamespace = "ecs"
	AWSNAMESPACE_EFS                            AWSNamespace = "efs"
	AWSNAMESPACE_ELASTICACHE                    AWSNamespace = "elasticache"
	AWSNAMESPACE_ELASTICBEANSTALK               AWSNamespace = "elasticbeanstalk"
	AWSNAMESPACE_ELASTICINFERENCE               AWSNamespace = "elasticinference"
	AWSNAMESPACE_ELASTICTRANSCODER              AWSNamespace = "elastictranscoder"
	AWSNAMESPACE_ELB                            AWSNamespace = "elb"
	AWSNAMESPACE_EMR                            AWSNamespace = "emr"
	AWSNAMESPACE_ES                             AWSNamespace = "es"
	AWSNAMESPACE_FIREHOSE                       AWSNamespace = "firehose"
	AWSNAMESPACE_FSX                            AWSNamespace = "fsx"
	AWSNAMESPACE_GAMELIFT                       AWSNamespace = "gamelift"
	AWSNAMESPACE_GLOBALACCELERATOR              AWSNamespace = "globalaccelerator"
	AWSNAMESPACE_GLUE                           AWSNamespace = "glue"
	AWSNAMESPACE_INSPECTOR                      AWSNamespace = "inspector"
	AWSNAMESPACE_IOT                            AWSNamespace = "iot"
	AWSNAMESPACE_KEYSPACES                      AWSNamespace = "keyspaces"
	AWSNAMESPACE_KINESIS                        AWSNamespace = "kinesis"
	AWSNAMESPACE_KINESIS_ANALYTICS              AWSNamespace = "kinesis_analytics"
	AWSNAMESPACE_KMS                            AWSNamespace = "kms"
	AWSNAMESPACE_LAMBDA                         AWSNamespace = "lambda"
	AWSNAMESPACE_LEX                            AWSNamespace = "lex"
	AWSNAMESPACE_MEDIACONNECT                   AWSNamespace = "mediaconnect"
	AWSNAMESPACE_MEDIACONVERT                   AWSNamespace = "mediaconvert"
	AWSNAMESPACE_MEDIALIVE                      AWSNamespace = "medialive"
	AWSNAMESPACE_MEDIAPACKAGE                   AWSNamespace = "mediapackage"
	AWSNAMESPACE_MEDIASTORE                     AWSNamespace = "mediastore"
	AWSNAMESPACE_MEDIATAILOR                    AWSNamespace = "mediatailor"
	AWSNAMESPACE_MEMORYDB                       AWSNamespace = "memorydb"
	AWSNAMESPACE_ML                             AWSNamespace = "ml"
	AWSNAMESPACE_MQ                             AWSNamespace = "mq"
	AWSNAMESPACE_MSK                            AWSNamespace = "msk"
	AWSNAMESPACE_MWAA                           AWSNamespace = "mwaa"
	AWSNAMESPACE_NAT_GATEWAY                    AWSNamespace = "nat_gateway"
	AWSNAMESPACE_NEPTUNE                        AWSNamespace = "neptune"
	AWSNAMESPACE_NETWORK_ELB                    AWSNamespace = "network_elb"
	AWSNAMESPACE_NETWORKFIREWALL                AWSNamespace = "networkfirewall"
	AWSNAMESPACE_NETWORKMONITOR                 AWSNamespace = "networkmonitor"
	AWSNAMESPACE_OPENSEARCHSERVERLESS           AWSNamespace = "opensearchserverless"
	AWSNAMESPACE_OPSWORKS                       AWSNamespace = "opsworks"
	AWSNAMESPACE_POLLY                          AWSNamespace = "polly"
	AWSNAMESPACE_PRIVATELINKENDPOINTS           AWSNamespace = "privatelinkendpoints"
	AWSNAMESPACE_PRIVATELINKSERVICES            AWSNamespace = "privatelinkservices"
	AWSNAMESPACE_RDS                            AWSNamespace = "rds"
	AWSNAMESPACE_RDSPROXY                       AWSNamespace = "rdsproxy"
	AWSNAMESPACE_REDSHIFT                       AWSNamespace = "redshift"
	AWSNAMESPACE_REKOGNITION                    AWSNamespace = "rekognition"
	AWSNAMESPACE_ROUTE53                        AWSNamespace = "route53"
	AWSNAMESPACE_ROUTE53RESOLVER                AWSNamespace = "route53resolver"
	AWSNAMESPACE_S3                             AWSNamespace = "s3"
	AWSNAMESPACE_S3STORAGELENS                  AWSNamespace = "s3storagelens"
	AWSNAMESPACE_SAGEMAKER                      AWSNamespace = "sagemaker"
	AWSNAMESPACE_SAGEMAKERENDPOINTS             AWSNamespace = "sagemakerendpoints"
	AWSNAMESPACE_SAGEMAKERLABELINGJOBS          AWSNamespace = "sagemakerlabelingjobs"
	AWSNAMESPACE_SAGEMAKERMODELBUILDINGPIPELINE AWSNamespace = "sagemakermodelbuildingpipeline"
	AWSNAMESPACE_SAGEMAKERPROCESSINGJOBS        AWSNamespace = "sagemakerprocessingjobs"
	AWSNAMESPACE_SAGEMAKERTRAININGJOBS          AWSNamespace = "sagemakertrainingjobs"
	AWSNAMESPACE_SAGEMAKERTRANSFORMJOBS         AWSNamespace = "sagemakertransformjobs"
	AWSNAMESPACE_SAGEMAKERWORKTEAM              AWSNamespace = "sagemakerworkteam"
	AWSNAMESPACE_SERVICE_QUOTAS                 AWSNamespace = "service_quotas"
	AWSNAMESPACE_SES                            AWSNamespace = "ses"
	AWSNAMESPACE_SHIELD                         AWSNamespace = "shield"
	AWSNAMESPACE_SNS                            AWSNamespace = "sns"
	AWSNAMESPACE_SQS                            AWSNamespace = "sqs"
	AWSNAMESPACE_STEP_FUNCTIONS                 AWSNamespace = "step_functions"
	AWSNAMESPACE_STORAGE_GATEWAY                AWSNamespace = "storage_gateway"
	AWSNAMESPACE_SWF                            AWSNamespace = "swf"
	AWSNAMESPACE_TEXTRACT                       AWSNamespace = "textract"
	AWSNAMESPACE_TRANSITGATEWAY                 AWSNamespace = "transitgateway"
	AWSNAMESPACE_TRANSLATE                      AWSNamespace = "translate"
	AWSNAMESPACE_TRUSTED_ADVISOR                AWSNamespace = "trusted_advisor"
	AWSNAMESPACE_USAGE                          AWSNamespace = "usage"
	AWSNAMESPACE_VPN                            AWSNamespace = "vpn"
	AWSNAMESPACE_WAF                            AWSNamespace = "waf"
	AWSNAMESPACE_WAFV2                          AWSNamespace = "wafv2"
	AWSNAMESPACE_WORKSPACES                     AWSNamespace = "workspaces"
	AWSNAMESPACE_XRAY                           AWSNamespace = "xray"
)

var allowedAWSNamespaceEnumValues = []AWSNamespace{
	AWSNAMESPACE_API_GATEWAY,
	AWSNAMESPACE_APPLICATION_ELB,
	AWSNAMESPACE_APPRUNNER,
	AWSNAMESPACE_APPSTREAM,
	AWSNAMESPACE_APPSYNC,
	AWSNAMESPACE_ATHENA,
	AWSNAMESPACE_AUTO_SCALING,
	AWSNAMESPACE_BACKUP,
	AWSNAMESPACE_BEDROCK,
	AWSNAMESPACE_BILLING,
	AWSNAMESPACE_BUDGETING,
	AWSNAMESPACE_CERTIFICATEMANAGER,
	AWSNAMESPACE_CLOUDFRONT,
	AWSNAMESPACE_CLOUDHSM,
	AWSNAMESPACE_CLOUDSEARCH,
	AWSNAMESPACE_CLOUDWATCH_EVENTS,
	AWSNAMESPACE_CLOUDWATCH_LOGS,
	AWSNAMESPACE_CODEBUILD,
	AWSNAMESPACE_CODEWHISPERER,
	AWSNAMESPACE_COGNITO,
	AWSNAMESPACE_COLLECT_CUSTOM_METRICS,
	AWSNAMESPACE_CONFIG,
	AWSNAMESPACE_CONNECT,
	AWSNAMESPACE_CRAWL_ALARMS,
	AWSNAMESPACE_CUSTOM,
	AWSNAMESPACE_DIRECTCONNECT,
	AWSNAMESPACE_DMS,
	AWSNAMESPACE_DOCUMENTDB,
	AWSNAMESPACE_DYNAMODB,
	AWSNAMESPACE_DYNAMODBACCELERATOR,
	AWSNAMESPACE_EBS,
	AWSNAMESPACE_EC2,
	AWSNAMESPACE_EC2API,
	AWSNAMESPACE_EC2SPOT,
	AWSNAMESPACE_ECR,
	AWSNAMESPACE_ECS,
	AWSNAMESPACE_EFS,
	AWSNAMESPACE_ELASTICACHE,
	AWSNAMESPACE_ELASTICBEANSTALK,
	AWSNAMESPACE_ELASTICINFERENCE,
	AWSNAMESPACE_ELASTICTRANSCODER,
	AWSNAMESPACE_ELB,
	AWSNAMESPACE_EMR,
	AWSNAMESPACE_ES,
	AWSNAMESPACE_FIREHOSE,
	AWSNAMESPACE_FSX,
	AWSNAMESPACE_GAMELIFT,
	AWSNAMESPACE_GLOBALACCELERATOR,
	AWSNAMESPACE_GLUE,
	AWSNAMESPACE_INSPECTOR,
	AWSNAMESPACE_IOT,
	AWSNAMESPACE_KEYSPACES,
	AWSNAMESPACE_KINESIS,
	AWSNAMESPACE_KINESIS_ANALYTICS,
	AWSNAMESPACE_KMS,
	AWSNAMESPACE_LAMBDA,
	AWSNAMESPACE_LEX,
	AWSNAMESPACE_MEDIACONNECT,
	AWSNAMESPACE_MEDIACONVERT,
	AWSNAMESPACE_MEDIALIVE,
	AWSNAMESPACE_MEDIAPACKAGE,
	AWSNAMESPACE_MEDIASTORE,
	AWSNAMESPACE_MEDIATAILOR,
	AWSNAMESPACE_MEMORYDB,
	AWSNAMESPACE_ML,
	AWSNAMESPACE_MQ,
	AWSNAMESPACE_MSK,
	AWSNAMESPACE_MWAA,
	AWSNAMESPACE_NAT_GATEWAY,
	AWSNAMESPACE_NEPTUNE,
	AWSNAMESPACE_NETWORK_ELB,
	AWSNAMESPACE_NETWORKFIREWALL,
	AWSNAMESPACE_NETWORKMONITOR,
	AWSNAMESPACE_OPENSEARCHSERVERLESS,
	AWSNAMESPACE_OPSWORKS,
	AWSNAMESPACE_POLLY,
	AWSNAMESPACE_PRIVATELINKENDPOINTS,
	AWSNAMESPACE_PRIVATELINKSERVICES,
	AWSNAMESPACE_RDS,
	AWSNAMESPACE_RDSPROXY,
	AWSNAMESPACE_REDSHIFT,
	AWSNAMESPACE_REKOGNITION,
	AWSNAMESPACE_ROUTE53,
	AWSNAMESPACE_ROUTE53RESOLVER,
	AWSNAMESPACE_S3,
	AWSNAMESPACE_S3STORAGELENS,
	AWSNAMESPACE_SAGEMAKER,
	AWSNAMESPACE_SAGEMAKERENDPOINTS,
	AWSNAMESPACE_SAGEMAKERLABELINGJOBS,
	AWSNAMESPACE_SAGEMAKERMODELBUILDINGPIPELINE,
	AWSNAMESPACE_SAGEMAKERPROCESSINGJOBS,
	AWSNAMESPACE_SAGEMAKERTRAININGJOBS,
	AWSNAMESPACE_SAGEMAKERTRANSFORMJOBS,
	AWSNAMESPACE_SAGEMAKERWORKTEAM,
	AWSNAMESPACE_SERVICE_QUOTAS,
	AWSNAMESPACE_SES,
	AWSNAMESPACE_SHIELD,
	AWSNAMESPACE_SNS,
	AWSNAMESPACE_SQS,
	AWSNAMESPACE_STEP_FUNCTIONS,
	AWSNAMESPACE_STORAGE_GATEWAY,
	AWSNAMESPACE_SWF,
	AWSNAMESPACE_TEXTRACT,
	AWSNAMESPACE_TRANSITGATEWAY,
	AWSNAMESPACE_TRANSLATE,
	AWSNAMESPACE_TRUSTED_ADVISOR,
	AWSNAMESPACE_USAGE,
	AWSNAMESPACE_VPN,
	AWSNAMESPACE_WAF,
	AWSNAMESPACE_WAFV2,
	AWSNAMESPACE_WORKSPACES,
	AWSNAMESPACE_XRAY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AWSNamespace) GetAllowedValues() []AWSNamespace {
	return allowedAWSNamespaceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AWSNamespace) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AWSNamespace(value)
	return nil
}

// NewAWSNamespaceFromValue returns a pointer to a valid AWSNamespace
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAWSNamespaceFromValue(v string) (*AWSNamespace, error) {
	ev := AWSNamespace(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AWSNamespace: valid values are %v", v, allowedAWSNamespaceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AWSNamespace) IsValid() bool {
	for _, existing := range allowedAWSNamespaceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AWSNamespace value.
func (v AWSNamespace) Ptr() *AWSNamespace {
	return &v
}
