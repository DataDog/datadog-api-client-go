package main

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	ddv1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	ddv2 "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

type deleteListResponse string
type deleteOneResponse string

const (
	deleteAll      deleteListResponse = "all"
	deleteOneByOne deleteListResponse = "oneByOne"
	deleteNone     deleteListResponse = "none"

	deleteYes deleteOneResponse = "yes"
	deleteNo  deleteOneResponse = "no"
)

// unique test name will be e.g. go-testCaseName-12345-123456789 or go-testCaseName-local-123456789
var uniqueTestNameRe = regexp.MustCompile(`^[^-]+-[^-]+-(?:[\d]+|local)-[\d]+.*$`)
var batch = false
var dryRun = false

func isUniqueTestName(name string) bool {
	return uniqueTestNameRe.MatchString(name)
}

func deleteListPrompt() deleteListResponse {
	if dryRun {
		return deleteNone
	}
	if batch {
		return deleteAll
	}
	for {
		fmt.Printf("Delete [a]ll/[o]ne-by-one/[n]one? ")
		var text string
		fmt.Scanln(&text)
		switch text {
		case "a":
			return deleteAll
		case "o":
			return deleteOneByOne
		case "n":
			return deleteNone
		default:
			fmt.Printf("'%s' is not an allowed option, please provide one of 'a', 'o', 'n'\n", text)
		}
	}
}

func deleteOnePrompt(what string) deleteOneResponse {
	for {
		fmt.Printf("Delete %s? [y]es/[n]o ", what)
		var text string
		fmt.Scanln(&text)
		switch text {
		case "y":
			return deleteYes
		case "n":
			return deleteNo
		default:
			fmt.Printf("'%s' is not an allowed option, please provide one of 'y', 'n'\n", text)
		}
	}
}

// genericCleanupResource is a generic function to list and delete resources
// if accepts:
// * resource name (for user output)
// * a function that lists all resources of a specific type
// * a function that can format a single resource of the type for user output
// * a function that can delete a single resource of the type
func genericCleanupResource(
	resourceName string,
	listFunc func() ([]interface{}, error),
	toStrFunc func(interface{}) string,
	deleteFunc func(interface{}) (int, error)) {

	fmt.Printf("Getting %s list ...\n", resourceName)
	all, err := listFunc()
	if err != nil {
		fmt.Printf("Failed getting %s list: %s\n", resourceName, err.Error())
	}

	fmt.Printf("Found %d %ss created by tests\n", len(all), resourceName)
	if len(all) == 0 {
		return
	}
	for _, res := range all {
		fmt.Println(toStrFunc(res))
	}

	delete := deleteListPrompt()
	if delete == deleteNone {
		fmt.Printf("Not deleting any %ss\n", resourceName)
	} else {
		deleted := 0
		for _, d := range all {
			deleteOne := deleteYes
			if delete == deleteOneByOne {
				deleteOne = deleteOnePrompt(fmt.Sprintf("%s %s", resourceName, toStrFunc(d)))
			}
			if deleteOne == deleteYes {
				fmt.Printf("Deleting %s %s ... ", resourceName, toStrFunc(d))
				statusCode, err := deleteFunc(d)
				if err != nil {
					fmt.Printf("failed: %s\n", err.Error())
				} else {
					deleted++
					fmt.Printf("success (code %d)\n", statusCode)
				}
			}
		}
		fmt.Printf("Successfully deleted %d %ss\n", deleted, resourceName)
	}
}

func cleanupAPIKeys(ctx context.Context, client *ddv1.APIClient, clientV2 *ddv2.APIClient) {
	genericCleanupResource(
		"API key",
		func() ([]interface{}, error) {
			matched := []ddv1.ApiKey{}
			resp, _, err := client.KeyManagementApi.ListAPIKeys(ctx).Execute()
			if err != nil {
				return []interface{}{}, err
			}
			for _, i := range resp.GetApiKeys() {
				if isUniqueTestName(i.GetName()) {
					matched = append(matched, i)
				}
			}
			matchedI := make([]interface{}, len(matched))
			for i, d := range matched {
				matchedI[i] = d
			}
			return matchedI, nil
		},
		func(intf interface{}) string {
			d := intf.(ddv1.ApiKey)
			return fmt.Sprintf("name: %s", d.GetName())
		},
		func(intf interface{}) (int, error) {
			d := intf.(ddv1.ApiKey)
			_, r, err := client.KeyManagementApi.DeleteAPIKey(ctx, d.GetKey()).Execute()
			return r.StatusCode, err
		},
	)
}

func cleanupApplicationKeys(ctx context.Context, client *ddv1.APIClient, clientV2 *ddv2.APIClient) {
	genericCleanupResource(
		"application key",
		func() ([]interface{}, error) {
			matched := []ddv1.ApplicationKey{}
			resp, _, err := client.KeyManagementApi.ListApplicationKeys(ctx).Execute()
			if err != nil {
				return []interface{}{}, err
			}
			for _, i := range resp.GetApplicationKeys() {
				if isUniqueTestName(i.GetName()) {
					matched = append(matched, i)
				}
			}
			matchedI := make([]interface{}, len(matched))
			for i, d := range matched {
				matchedI[i] = d
			}
			return matchedI, nil
		},
		func(intf interface{}) string {
			d := intf.(ddv1.ApplicationKey)
			return fmt.Sprintf("name: %s", d.GetName())
		},
		func(intf interface{}) (int, error) {
			d := intf.(ddv1.ApplicationKey)
			_, r, err := client.KeyManagementApi.DeleteApplicationKey(ctx, d.GetHash()).Execute()
			return r.StatusCode, err
		},
	)
}

func cleanupAWSIntegration(ctx context.Context, client *ddv1.APIClient, clientV2 *ddv2.APIClient) {
	genericCleanupResource(
		"aws integration",
		func() ([]interface{}, error) {
			matched := []ddv1.AWSAccount{}
			resp, _, err := client.AWSIntegrationApi.ListAWSAccounts(ctx).Execute()
			if err != nil {
				return []interface{}{}, err
			}
			for _, i := range resp.GetAccounts() {
				if isUniqueTestName(i.GetRoleName()) {
					matched = append(matched, i)
				}
			}
			matchedI := make([]interface{}, len(matched))
			for i, d := range matched {
				matchedI[i] = d
			}
			return matchedI, nil
		},
		func(intf interface{}) string {
			d := intf.(ddv1.AWSAccount)
			return fmt.Sprintf("account ID: %s, role name: %s", d.GetAccountId(), d.GetRoleName())
		},
		func(intf interface{}) (int, error) {
			d := intf.(ddv1.AWSAccount)
			_, r, err := client.AWSIntegrationApi.DeleteAWSAccount(ctx).Body(d).Execute()
			return r.StatusCode, err
		},
	)
}

func cleanupAzureIntegration(ctx context.Context, client *ddv1.APIClient, clientV2 *ddv2.APIClient) {
	genericCleanupResource(
		"azure integration",
		func() ([]interface{}, error) {
			matched := []ddv1.AzureAccount{}
			resp, _, err := client.AzureIntegrationApi.ListAzureIntegration(ctx).Execute()
			if err != nil {
				return []interface{}{}, err
			}
			for _, i := range resp {
				if isUniqueTestName(i.GetTenantName()) {
					matched = append(matched, i)
				}
			}
			matchedI := make([]interface{}, len(matched))
			for i, d := range matched {
				matchedI[i] = d
			}
			return matchedI, nil
		},
		func(intf interface{}) string {
			d := intf.(ddv1.AzureAccount)
			return fmt.Sprintf("client ID: %s, tenant name: %s", d.GetClientId(), d.GetTenantName())
		},
		func(intf interface{}) (int, error) {
			d := intf.(ddv1.AzureAccount)
			_, r, err := client.AzureIntegrationApi.DeleteAzureIntegration(ctx).Body(d).Execute()
			return r.StatusCode, err
		},
	)
}

func cleanupDashboards(ctx context.Context, client *ddv1.APIClient, clientV2 *ddv2.APIClient) {
	genericCleanupResource(
		"dashboard",
		func() ([]interface{}, error) {
			matched := []ddv1.DashboardSummaryDashboards{}
			resp, _, err := client.DashboardsApi.ListDashboards(ctx).Execute()
			if err != nil {
				return []interface{}{}, err
			}
			for _, i := range resp.GetDashboards() {
				if isUniqueTestName(i.GetTitle()) {
					matched = append(matched, i)
				}
			}
			matchedI := make([]interface{}, len(matched))
			for i, d := range matched {
				matchedI[i] = d
			}
			return matchedI, nil
		},
		func(intf interface{}) string {
			d := intf.(ddv1.DashboardSummaryDashboards)
			return fmt.Sprintf("ID: %s, title: %s", d.GetId(), d.GetTitle())
		},
		func(intf interface{}) (int, error) {
			d := intf.(ddv1.DashboardSummaryDashboards)
			_, r, err := client.DashboardsApi.DeleteDashboard(ctx, d.GetId()).Execute()
			return r.StatusCode, err
		},
	)
}

func cleanupDashboardLists(ctx context.Context, client *ddv1.APIClient, clientV2 *ddv2.APIClient) {
	genericCleanupResource(
		"dashboard list",
		func() ([]interface{}, error) {
			matched := []ddv1.DashboardList{}
			resp, _, err := client.DashboardListsApi.ListDashboardLists(ctx).Execute()
			if err != nil {
				return []interface{}{}, err
			}
			for _, i := range resp.GetDashboardLists() {
				if isUniqueTestName(i.GetName()) {
					matched = append(matched, i)
				}
			}
			matchedI := make([]interface{}, len(matched))
			for i, d := range matched {
				matchedI[i] = d
			}
			return matchedI, nil
		},
		func(intf interface{}) string {
			d := intf.(ddv1.DashboardList)
			return fmt.Sprintf("ID: %d, name: %s", d.GetId(), d.GetName())
		},
		func(intf interface{}) (int, error) {
			d := intf.(ddv1.DashboardList)
			_, r, err := client.DashboardListsApi.DeleteDashboardList(ctx, d.GetId()).Execute()
			return r.StatusCode, err
		},
	)
}

func cleanupDowntimes(ctx context.Context, client *ddv1.APIClient, clientV2 *ddv2.APIClient) {
	genericCleanupResource(
		"downtime",
		func() ([]interface{}, error) {
			matched := []ddv1.Downtime{}
			resp, _, err := client.DowntimesApi.ListDowntimes(ctx).Execute()
			if err != nil {
				return []interface{}{}, err
			}
			for _, i := range resp {
				if isUniqueTestName(i.GetMessage()) && i.GetCanceled() == 0 {
					matched = append(matched, i)
				}
			}
			matchedI := make([]interface{}, len(matched))
			for i, d := range matched {
				matchedI[i] = d
			}
			return matchedI, nil
		},
		func(intf interface{}) string {
			d := intf.(ddv1.Downtime)
			return fmt.Sprintf("ID: %d, message: %s", d.GetId(), d.GetMessage())
		},
		func(intf interface{}) (int, error) {
			d := intf.(ddv1.Downtime)
			r, err := client.DowntimesApi.CancelDowntime(ctx, d.GetId()).Execute()
			return r.StatusCode, err
		},
	)
}

func cleanupGCPIntegration(ctx context.Context, client *ddv1.APIClient, clientV2 *ddv2.APIClient) {
	genericCleanupResource(
		"gcp integration",
		func() ([]interface{}, error) {
			matched := []ddv1.GCPAccount{}
			resp, _, err := client.GCPIntegrationApi.ListGCPIntegration(ctx).Execute()
			if err != nil {
				return []interface{}{}, err
			}
			for _, i := range resp {
				if isUniqueTestName(i.GetProjectId()) {
					matched = append(matched, i)
				}
			}
			matchedI := make([]interface{}, len(matched))
			for i, d := range matched {
				matchedI[i] = d
			}
			return matchedI, nil
		},
		func(intf interface{}) string {
			d := intf.(ddv1.GCPAccount)
			return fmt.Sprintf("client ID: %s, project ID: %s", d.GetClientId(), d.GetProjectId())
		},
		func(intf interface{}) (int, error) {
			d := intf.(ddv1.GCPAccount)
			_, r, err := client.GCPIntegrationApi.DeleteGCPIntegration(ctx).Body(d).Execute()
			return r.StatusCode, err
		},
	)
}

func cleanupLogsPipelines(ctx context.Context, client *ddv1.APIClient, clientV2 *ddv2.APIClient) {
	genericCleanupResource(
		"logs pipeline",
		func() ([]interface{}, error) {
			matched := []ddv1.LogsPipeline{}
			resp, _, err := client.LogsPipelinesApi.ListLogsPipelines(ctx).Execute()
			if err != nil {
				return []interface{}{}, err
			}
			for _, i := range resp {
				if isUniqueTestName(i.GetName()) {
					matched = append(matched, i)
				}
			}
			matchedI := make([]interface{}, len(matched))
			for i, d := range matched {
				matchedI[i] = d
			}
			return matchedI, nil
		},
		func(intf interface{}) string {
			d := intf.(ddv1.LogsPipeline)
			return fmt.Sprintf("ID: %s, name: %s", d.GetId(), d.GetName())
		},
		func(intf interface{}) (int, error) {
			d := intf.(ddv1.LogsPipeline)
			r, err := client.LogsPipelinesApi.DeleteLogsPipeline(ctx, d.GetId()).Execute()
			return r.StatusCode, err
		},
	)
}

func cleanupMonitors(ctx context.Context, client *ddv1.APIClient, clientV2 *ddv2.APIClient) {
	genericCleanupResource(
		"monitor",
		func() ([]interface{}, error) {
			matched := []ddv1.Monitor{}
			resp, _, err := client.MonitorsApi.ListMonitors(ctx).Execute()
			if err != nil {
				return []interface{}{}, err
			}
			for _, i := range resp {
				if isUniqueTestName(i.GetName()) {
					matched = append(matched, i)
				}
			}
			matchedI := make([]interface{}, len(matched))
			for i, d := range matched {
				matchedI[i] = d
			}
			return matchedI, nil
		},
		func(intf interface{}) string {
			d := intf.(ddv1.Monitor)
			return fmt.Sprintf("ID: %d, name: %s", d.GetId(), d.GetName())
		},
		func(intf interface{}) (int, error) {
			d := intf.(ddv1.Monitor)
			_, r, err := client.MonitorsApi.DeleteMonitor(ctx, d.GetId()).Execute()
			return r.StatusCode, err
		},
	)
}

func cleanupRoles(ctx context.Context, client *ddv1.APIClient, clientV2 *ddv2.APIClient) {
	genericCleanupResource(
		"role",
		func() ([]interface{}, error) {
			matched := []ddv2.Role{}
			resp, _, err := clientV2.RolesApi.ListRoles(ctx).Execute()
			if err != nil {
				return []interface{}{}, err
			}
			for _, i := range resp.GetData() {
				attrs := i.GetAttributes()
				if isUniqueTestName(attrs.GetName()) {
					matched = append(matched, i)
				}
			}
			matchedI := make([]interface{}, len(matched))
			for i, d := range matched {
				matchedI[i] = d
			}
			return matchedI, nil
		},
		func(intf interface{}) string {
			d := intf.(ddv2.Role)
			attrs := d.GetAttributes()
			return fmt.Sprintf("ID: %s, name: %s", d.GetId(), attrs.GetName())
		},
		func(intf interface{}) (int, error) {
			d := intf.(ddv2.Role)
			r, err := clientV2.RolesApi.DeleteRole(ctx, d.GetId()).Execute()
			return r.StatusCode, err
		},
	)
}

func cleanupSecurityMonitoringRules(ctx context.Context, client *ddv1.APIClient, clientV2 *ddv2.APIClient) {
	genericCleanupResource(
		"security monitoring rule",
		func() ([]interface{}, error) {
			matched := []ddv2.SecurityMonitoringRuleResponse{}
			all := []ddv2.SecurityMonitoringRuleResponse{}
			page := int64(0)
			for {
				resp, _, err := clientV2.SecurityMonitoringApi.ListSecurityMonitoringRules(ctx).PageSize(200).PageNumber(page).Execute()
				if err != nil {
					return []interface{}{}, err
				}
				all = append(all, resp.GetData()...)
				meta := resp.GetMeta()
				if *meta.GetPage().TotalCount == int64(len(all)) {
					break
				}
				page++
			}
			for _, i := range all {
				if isUniqueTestName(i.GetName()) {
					matched = append(matched, i)
				}
			}
			matchedI := make([]interface{}, len(matched))
			for i, d := range matched {
				matchedI[i] = d
			}
			return matchedI, nil
		},
		func(intf interface{}) string {
			d := intf.(ddv2.SecurityMonitoringRuleResponse)
			return fmt.Sprintf("ID: %s, name: %s", d.GetId(), d.GetName())
		},
		func(intf interface{}) (int, error) {
			d := intf.(ddv2.SecurityMonitoringRuleResponse)
			r, err := clientV2.SecurityMonitoringApi.DeleteSecurityMonitoringRule(ctx, d.GetId()).Execute()
			return r.StatusCode, err
		},
	)
}

func cleanupSynthetics(ctx context.Context, client *ddv1.APIClient, clientV2 *ddv2.APIClient) {
	genericCleanupResource(
		"synthetics test",
		func() ([]interface{}, error) {
			matched := []ddv1.SyntheticsTestDetails{}
			resp, _, err := client.SyntheticsApi.ListTests(ctx).Execute()
			if err != nil {
				return []interface{}{}, err
			}
			for _, i := range resp.GetTests() {
				if isUniqueTestName(i.GetName()) {
					matched = append(matched, i)
				}
			}
			matchedI := make([]interface{}, len(matched))
			for i, d := range matched {
				matchedI[i] = d
			}
			return matchedI, nil
		},
		func(intf interface{}) string {
			d := intf.(ddv1.SyntheticsTestDetails)
			return fmt.Sprintf("public ID: %s, name: %s", d.GetPublicId(), d.GetName())
		},
		func(intf interface{}) (int, error) {
			d := intf.(ddv1.SyntheticsTestDetails)
			dp := ddv1.SyntheticsDeleteTestsPayload{}
			dp.SetPublicIds([]string{d.GetPublicId()})
			_, r, err := client.SyntheticsApi.DeleteTests(ctx).Body(dp).Execute()
			return r.StatusCode, err
		},
	)
}

func cleanupUsers(ctx context.Context, client *ddv1.APIClient, clientV2 *ddv2.APIClient) {
	genericCleanupResource(
		"user",
		func() ([]interface{}, error) {
			matched := []ddv1.User{}
			resp, _, err := client.UsersApi.ListUsers(ctx).Execute()
			if err != nil {
				return []interface{}{}, err
			}
			for _, i := range resp.GetUsers() {
				if isUniqueTestName(i.GetName()) && !i.GetDisabled() {
					matched = append(matched, i)
				}
			}
			matchedI := make([]interface{}, len(matched))
			for i, d := range matched {
				matchedI[i] = d
			}
			return matchedI, nil
		},
		func(intf interface{}) string {
			d := intf.(ddv1.User)
			return fmt.Sprintf("handle: %s, name: %s", d.GetHandle(), d.GetName())
		},
		func(intf interface{}) (int, error) {
			d := intf.(ddv1.User)
			_, r, err := client.UsersApi.DisableUser(ctx, d.GetHandle()).Execute()
			return r.StatusCode, err
		},
	)
}

var cleanupFuncs = map[string]func(context.Context, *ddv1.APIClient, *ddv2.APIClient){
	"api-keys": cleanupAPIKeys,
	"app-keys": cleanupApplicationKeys,
	"dashboards": cleanupDashboards,
	"dashboard-lists": cleanupDashboardLists,
	"downtimes": cleanupDowntimes,
	"integration-aws": cleanupAWSIntegration,
	"integration-azure": cleanupAzureIntegration,
	"integration-gcp": cleanupGCPIntegration,
	"logs-pipelines": cleanupLogsPipelines,
	"monitors": cleanupMonitors,
	// NOTE: pagerduty has no `List` methods for all services, so we can't do a cleanup for it
	"roles": cleanupRoles,
	"security-monitoring-rules": cleanupSecurityMonitoringRules,
	// NOTE: SLO `List` method requires list of IDs, which we don't know, so we can't do a cleanup for it
	"synthetics": cleanupSynthetics,
	"users": cleanupUsers,
}

func printHelp(keys []string) {
	usage := `Usage: go run cleanup.go [--dry-run] [--batch] [OBJECT_TYPE] ...

Description:
  List and delete leftover test objects of given OBJECT_TYPE(s) (if not given,
  list and delete all leftover test objects).

  --batch   Prompt-less run (delete all objects unless used with --dry-run)
  --dry-run Don't delete any objects

  Recognized OBJECT_TYPEs:
  - ` + strings.Join(keys, "\n  - ")
	fmt.Println(usage)
}

func main() {
	ctx := context.WithValue(
		context.Background(),
		ddv1.ContextAPIKeys,
		map[string]ddv1.APIKey{
			"apiKeyAuth": {
				Key: os.Getenv("DD_CLIENT_API_KEY"),
			},
			"appKeyAuth": {
				Key: os.Getenv("DD_CLIENT_APP_KEY"),
			},
		},
	)
	ctx = context.WithValue(
		ctx,
		ddv2.ContextAPIKeys,
		map[string]ddv2.APIKey{
			"apiKeyAuth": {
				Key: os.Getenv("DD_CLIENT_API_KEY"),
			},
			"appKeyAuth": {
				Key: os.Getenv("DD_CLIENT_APP_KEY"),
			},
		},
	)

	configurationV1 := ddv1.NewConfiguration()
	configurationV2 := ddv2.NewConfiguration()
	clientV1 := ddv1.NewAPIClient(configurationV1)
	clientV2 := ddv2.NewAPIClient(configurationV2)
	// make sure we traverse endpoints in sorted order
	keys := []string{}
	for k, _ := range cleanupFuncs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// read args
	args := os.Args[1:]
	argKeys := []string{}
	for _, arg := range args {
		switch arg {
		case "--help":
			printHelp(keys)
			return
		case "--dry-run":
			dryRun = true
		case "--batch":
			batch = true
		default:
			if strings.HasPrefix(arg, "--") {
				fmt.Printf("Unknown argument %s\n", arg)
				printHelp(keys)
				return
			}
			if _, ok := cleanupFuncs[arg]; ok {
				argKeys = append(argKeys, arg)
			}
		}
	}
	if len(argKeys) > 0 {
		keys = argKeys
	}

	for _, key := range keys {
		cleanupFuncs[key](ctx, clientV1, clientV2)
	}
}
