// Create a pipeline with Schema Processor and preserve_source false returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.LogsPipeline{
		Filter: &datadogV1.LogsFilter{
			Query: datadog.PtrString("source:python"),
		},
		Name: "testSchemaProcessor",
		Processors: []datadogV1.LogsProcessor{
			datadogV1.LogsProcessor{
				LogsSchemaProcessor: &datadogV1.LogsSchemaProcessor{
					Type:      datadogV1.LOGSSCHEMAPROCESSORTYPE_SCHEMA_PROCESSOR,
					IsEnabled: datadog.PtrBool(true),
					Name:      "Apply OCSF schema for 3001",
					Schema: datadogV1.LogsSchemaData{
						SchemaType: "ocsf",
						Version:    "1.5.0",
						ClassUid:   3001,
						ClassName:  "Account Change",
						Profiles: []string{
							"cloud",
							"datetime",
						},
					},
					Mappers: []datadogV1.LogsSchemaMapper{
						datadogV1.LogsSchemaMapper{
							LogsSchemaCategoryMapper: &datadogV1.LogsSchemaCategoryMapper{
								Type: datadogV1.LOGSSCHEMACATEGORYMAPPERTYPE_SCHEMA_CATEGORY_MAPPER,
								Name: "activity_id and activity_name",
								Categories: []datadogV1.LogsSchemaCategoryMapperCategory{
									{
										Filter: datadogV1.LogsFilter{
											Query: datadog.PtrString("@eventName:(*Create*)"),
										},
										Name: "Create",
										Id:   1,
									},
									{
										Filter: datadogV1.LogsFilter{
											Query: datadog.PtrString("@eventName:(ChangePassword OR PasswordUpdated)"),
										},
										Name: "Password Change",
										Id:   3,
									},
									{
										Filter: datadogV1.LogsFilter{
											Query: datadog.PtrString("@eventName:(*Attach*)"),
										},
										Name: "Attach Policy",
										Id:   7,
									},
									{
										Filter: datadogV1.LogsFilter{
											Query: datadog.PtrString("@eventName:(*Detach* OR *Remove*)"),
										},
										Name: "Detach Policy",
										Id:   8,
									},
									{
										Filter: datadogV1.LogsFilter{
											Query: datadog.PtrString("@eventName:(*Delete*)"),
										},
										Name: "Delete",
										Id:   6,
									},
									{
										Filter: datadogV1.LogsFilter{
											Query: datadog.PtrString("@eventName:*"),
										},
										Name: "Other",
										Id:   99,
									},
								},
								Targets: datadogV1.LogsSchemaCategoryMapperTargets{
									Name: datadog.PtrString("ocsf.activity_name"),
									Id:   datadog.PtrString("ocsf.activity_id"),
								},
								Fallback: &datadogV1.LogsSchemaCategoryMapperFallback{
									Values: map[string]string{
										"ocsf.activity_id":   "99",
										"ocsf.activity_name": "Other",
									},
									Sources: map[string][]string{
										"ocsf.activity_name": []string{
											"eventName",
										},
									},
								},
							}},
						datadogV1.LogsSchemaMapper{
							LogsSchemaCategoryMapper: &datadogV1.LogsSchemaCategoryMapper{
								Type: datadogV1.LOGSSCHEMACATEGORYMAPPERTYPE_SCHEMA_CATEGORY_MAPPER,
								Name: "status",
								Categories: []datadogV1.LogsSchemaCategoryMapperCategory{
									{
										Filter: datadogV1.LogsFilter{
											Query: datadog.PtrString("-@errorCode:*"),
										},
										Id:   1,
										Name: "Success",
									},
									{
										Filter: datadogV1.LogsFilter{
											Query: datadog.PtrString("@errorCode:*"),
										},
										Id:   2,
										Name: "Failure",
									},
								},
								Targets: datadogV1.LogsSchemaCategoryMapperTargets{
									Id:   datadog.PtrString("ocsf.status_id"),
									Name: datadog.PtrString("ocsf.status"),
								},
							}},
						datadogV1.LogsSchemaMapper{
							LogsSchemaCategoryMapper: &datadogV1.LogsSchemaCategoryMapper{
								Type: datadogV1.LOGSSCHEMACATEGORYMAPPERTYPE_SCHEMA_CATEGORY_MAPPER,
								Name: "Set default severity",
								Categories: []datadogV1.LogsSchemaCategoryMapperCategory{
									{
										Filter: datadogV1.LogsFilter{
											Query: datadog.PtrString("@eventName:*"),
										},
										Name: "Informational",
										Id:   1,
									},
								},
								Targets: datadogV1.LogsSchemaCategoryMapperTargets{
									Name: datadog.PtrString("ocsf.severity"),
									Id:   datadog.PtrString("ocsf.severity_id"),
								},
							}},
						datadogV1.LogsSchemaMapper{
							LogsSchemaRemapper: &datadogV1.LogsSchemaRemapper{
								Type: datadogV1.LOGSSCHEMAREMAPPERTYPE_SCHEMA_REMAPPER,
								Name: "Map userIdentity to ocsf.user.uid",
								Sources: []string{
									"userIdentity.principalId",
									"responseElements.role.roleId",
									"responseElements.user.userId",
								},
								Target:         "ocsf.user.uid",
								PreserveSource: datadog.PtrBool(false),
							}},
						datadogV1.LogsSchemaMapper{
							LogsSchemaRemapper: &datadogV1.LogsSchemaRemapper{
								Type: datadogV1.LOGSSCHEMAREMAPPERTYPE_SCHEMA_REMAPPER,
								Name: "Map userName to ocsf.user.name",
								Sources: []string{
									"requestParameters.userName",
									"responseElements.role.roleName",
									"requestParameters.roleName",
									"responseElements.user.userName",
								},
								Target:         "ocsf.user.name",
								PreserveSource: datadog.PtrBool(false),
							}},
						datadogV1.LogsSchemaMapper{
							LogsSchemaRemapper: &datadogV1.LogsSchemaRemapper{
								Type: datadogV1.LOGSSCHEMAREMAPPERTYPE_SCHEMA_REMAPPER,
								Name: "Map api to ocsf.api",
								Sources: []string{
									"api",
								},
								Target:         "ocsf.api",
								PreserveSource: datadog.PtrBool(false),
							}},
						datadogV1.LogsSchemaMapper{
							LogsSchemaRemapper: &datadogV1.LogsSchemaRemapper{
								Type: datadogV1.LOGSSCHEMAREMAPPERTYPE_SCHEMA_REMAPPER,
								Name: "Map user to ocsf.user",
								Sources: []string{
									"user",
								},
								Target:         "ocsf.user",
								PreserveSource: datadog.PtrBool(false),
							}},
						datadogV1.LogsSchemaMapper{
							LogsSchemaRemapper: &datadogV1.LogsSchemaRemapper{
								Type: datadogV1.LOGSSCHEMAREMAPPERTYPE_SCHEMA_REMAPPER,
								Name: "Map actor to ocsf.actor",
								Sources: []string{
									"actor",
								},
								Target:         "ocsf.actor",
								PreserveSource: datadog.PtrBool(false),
							}},
						datadogV1.LogsSchemaMapper{
							LogsSchemaRemapper: &datadogV1.LogsSchemaRemapper{
								Type: datadogV1.LOGSSCHEMAREMAPPERTYPE_SCHEMA_REMAPPER,
								Name: "Map cloud to ocsf.cloud",
								Sources: []string{
									"cloud",
								},
								Target:         "ocsf.cloud",
								PreserveSource: datadog.PtrBool(false),
							}},
						datadogV1.LogsSchemaMapper{
							LogsSchemaRemapper: &datadogV1.LogsSchemaRemapper{
								Type: datadogV1.LOGSSCHEMAREMAPPERTYPE_SCHEMA_REMAPPER,
								Name: "Map http_request to ocsf.http_request",
								Sources: []string{
									"http_request",
								},
								Target:         "ocsf.http_request",
								PreserveSource: datadog.PtrBool(false),
							}},
						datadogV1.LogsSchemaMapper{
							LogsSchemaRemapper: &datadogV1.LogsSchemaRemapper{
								Type: datadogV1.LOGSSCHEMAREMAPPERTYPE_SCHEMA_REMAPPER,
								Name: "Map metadata to ocsf.metadata",
								Sources: []string{
									"metadata",
								},
								Target:         "ocsf.metadata",
								PreserveSource: datadog.PtrBool(false),
							}},
						datadogV1.LogsSchemaMapper{
							LogsSchemaRemapper: &datadogV1.LogsSchemaRemapper{
								Type: datadogV1.LOGSSCHEMAREMAPPERTYPE_SCHEMA_REMAPPER,
								Name: "Map time to ocsf.time",
								Sources: []string{
									"time",
								},
								Target:         "ocsf.time",
								PreserveSource: datadog.PtrBool(false),
							}},
						datadogV1.LogsSchemaMapper{
							LogsSchemaRemapper: &datadogV1.LogsSchemaRemapper{
								Type: datadogV1.LOGSSCHEMAREMAPPERTYPE_SCHEMA_REMAPPER,
								Name: "Map src_endpoint to ocsf.src_endpoint",
								Sources: []string{
									"src_endpoint",
								},
								Target:         "ocsf.src_endpoint",
								PreserveSource: datadog.PtrBool(false),
							}},
						datadogV1.LogsSchemaMapper{
							LogsSchemaRemapper: &datadogV1.LogsSchemaRemapper{
								Type: datadogV1.LOGSSCHEMAREMAPPERTYPE_SCHEMA_REMAPPER,
								Name: "Map severity to ocsf.severity",
								Sources: []string{
									"severity",
								},
								Target:         "ocsf.severity",
								PreserveSource: datadog.PtrBool(false),
							}},
						datadogV1.LogsSchemaMapper{
							LogsSchemaRemapper: &datadogV1.LogsSchemaRemapper{
								Type: datadogV1.LOGSSCHEMAREMAPPERTYPE_SCHEMA_REMAPPER,
								Name: "Map severity_id to ocsf.severity_id",
								Sources: []string{
									"severity_id",
								},
								Target:         "ocsf.severity_id",
								PreserveSource: datadog.PtrBool(false),
							}},
					},
				}},
		},
		Tags: []string{},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewLogsPipelinesApi(apiClient)
	resp, r, err := api.CreateLogsPipeline(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsPipelinesApi.CreateLogsPipeline`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsPipelinesApi.CreateLogsPipeline`:\n%s\n", responseContent)
}
