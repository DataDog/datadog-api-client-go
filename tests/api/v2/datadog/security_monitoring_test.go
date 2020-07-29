package test

import (
	"context"
	"fmt"
	_nethttp "net/http"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func createRule(ctx context.Context, api *datadog.SecurityMonitoringApiService, ruleName string) (datadog.SecurityMonitoringRuleResponseCreate, *_nethttp.Response, error) {

	query := datadog.NewSecurityMonitoringRuleQueryCreate("thiswillnevernevermatch")
	query.SetName("nevermatch")
	query.SetGroupByFields([]string{})
	query.SetDistinctFields([]string{})
	queries := []datadog.SecurityMonitoringRuleQueryCreate{*query}

	options := datadog.NewSecurityMonitoringRuleOptions()
	options.SetEvaluationWindow(datadog.SECURITYMONITORINGRULEEVALUATIONWINDOW_FIVE_MINUTES)
	options.SetKeepAlive(datadog.SECURITYMONITORINGRULEKEEPALIVE_FIVE_MINUTES)
	options.SetMaxSignalDuration(datadog.SECURITYMONITORINGRULEMAXSIGNALDURATION_FIVE_MINUTES)

	ruleCase := datadog.NewSecurityMonitoringRuleCaseCreate(datadog.SECURITYMONITORINGRULESEVERITY_INFO)
	ruleCase.SetName("rule-case")
	ruleCase.SetCondition("nevermatch > 1000")
	ruleCases := []datadog.SecurityMonitoringRuleCaseCreate{*ruleCase}

	createPayload := datadog.NewSecurityMonitoringRuleCreatePayload(
		ruleCases,
		true,
		"rule message",
		ruleName,
		*options,
		queries,
	)
	createPayload.SetTags([]string{"datadog-api-client-test-go"})

	return api.CreateSecurityMonitoringRule(ctx).Body(*createPayload).Execute()
}

func deleteRule(t *testing.T, ctx context.Context, api *datadog.SecurityMonitoringApiService, id string) {
	_, err := api.DeleteSecurityMonitoringRule(ctx, id).Execute()
	if err != nil {
		t.Logf("Error delete rule: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
}

func TestSecMonRulesCRUD(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := Client(ctx).SecurityMonitoringApi

	ruleResponses := []datadog.SecurityMonitoringRuleResponseCreate{}
	uniqueName := *tests.UniqueEntityName(ctx, t)
	// create rules
	for i := 0; i < 5; i++ {
		ruleName := fmt.Sprintf("%s-%d", uniqueName, i)
		ruleResponse, httpResponse, err := createRule(ctx, api, ruleName)
		if err != nil {
			t.Fatalf("Error creating rule %d: Response: %v", i, err)
		}
		assert.Equal(200, httpResponse.StatusCode)
		ruleResponses = append(ruleResponses, ruleResponse)
		defer deleteRule(t, ctx, api, ruleResponse.GetId())
	}

	// get single rule
	ruleResponse := ruleResponses[0]
	getResponse, httpResponse, err := api.GetSecurityMonitoringRule(ctx, ruleResponse.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error getting rule: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpResponse.StatusCode)
	assert.Equal(ruleResponse.GetName(), getResponse.GetName())

	//// get rule list
	// get filter count
	listResponse, httpResponse, err := api.
		ListSecurityMonitoringRules(ctx).
		PageSize(1).
		PageNumber(0).
		Execute()
	if err != nil {
		t.Fatalf("Error listing rules: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpResponse.StatusCode)
	assert.Equal(1, len(listResponse.GetData()))

	meta := listResponse.GetMeta()
	page := meta.GetPage()
	ruleCount := page.GetTotalCount()
	assert.GreaterOrEqual(ruleCount, int64(5))

	// check that all known rules are present in the response
	// we are not asserting the size of getData as this could be flaky
	listResponse, httpResponse, err = api.ListSecurityMonitoringRules(ctx).PageSize(ruleCount).PageNumber(0).Execute()
	if err != nil {
		t.Fatalf("Error listing rules: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpResponse.StatusCode)
	ruleIds := getRuleIds(listResponse.GetData())
	for _, rule := range ruleResponses {
		assert.Contains(ruleIds, rule.GetId())
	}

	// paging
	firstPageResponse, httpResponse, err := api.
		ListSecurityMonitoringRules(ctx).
		PageSize(2).
		PageNumber(0).
		Execute()
	if err != nil {
		t.Fatalf("Error listing rules: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpResponse.StatusCode)
	assert.Equal(2, len(firstPageResponse.GetData()))
	secondPageResponse, httpResponse, err := api.
		ListSecurityMonitoringRules(ctx).
		PageSize(2).
		PageNumber(0).
		Execute()
	if err != nil {
		t.Fatalf("Error listing rules: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpResponse.StatusCode)
	assert.Equal(2, len(secondPageResponse.GetData()))

	firstPageIds := getRuleIds(firstPageResponse.GetData())
	secondPageIds := getRuleIds(secondPageResponse.GetData())
	for _, id := range firstPageIds {
		assert.NotContains(secondPageIds, id)
	}

	// update rule
	updatePayload := datadog.NewSecurityMonitoringRuleUpdatePayload()
	updatePayload.SetName(getResponse.GetName())
	updatePayload.SetIsEnabled(false)
	updatePayload.SetMessage(getResponse.GetMessage())
	updatePayload.SetTags(getResponse.GetTags())
	updatePayload.SetQueries(getResponse.GetQueries())
	updatePayload.SetOptions(getResponse.GetOptions())
	updatePayload.SetCases(getResponse.GetCases())
	updateResponse, httpResponse, err := api.UpdateSecurityMonitoringRule(ctx, getResponse.GetId()).
		Body(*updatePayload).
		Execute()
	if err != nil {
		t.Fatalf("Error updating rule: Response %v", err)
	}
	assert.Equal(200, httpResponse.StatusCode)
	assert.Equal(false, updateResponse.GetIsEnabled())

	getUpdatedResponse, httpResponse, err := api.GetSecurityMonitoringRule(ctx, ruleResponse.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error updating rule: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpResponse.StatusCode)
	assert.Equal(false, getUpdatedResponse.GetIsEnabled())

	// Delete rule
	httpResponse, err = api.DeleteSecurityMonitoringRule(ctx, ruleResponse.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error delete rule: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(204, httpResponse.StatusCode)

	_, httpResponse, _ = api.GetSecurityMonitoringRule(ctx, ruleResponse.GetId()).Execute()
	assert.Equal(404, httpResponse.StatusCode)
}

func getRuleIds(rules []datadog.SecurityMonitoringRuleResponse) map[string]bool {
	ruleIds := map[string]bool{}
	for _, rule := range rules {
		ruleIds[rule.GetId()] = true
	}
	return ruleIds
}
