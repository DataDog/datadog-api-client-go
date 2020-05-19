package test

import (
	"context"
	"fmt"
	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	_nethttp "net/http"
	"testing"
)

func createRule(ctx context.Context, api *datadog.SecurityMonitoringApiService, ruleName string) (datadog.SecurityMonitoringRuleResponse, *_nethttp.Response, error) {

	query := datadog.NewSecurityMonitoringRuleQuery()
	query.SetName("nevermatch")
	query.SetQuery("thiswillnevernevermatch")
	query.SetGroupByFields([]string{})
	queries := []datadog.SecurityMonitoringRuleQuery{*query}

	options := datadog.NewSecurityMonitoringRuleOptions()
	options.SetEvaluationWindow(datadog.SECURITYMONITORINGRULEEVALUATIONWINDOW_FIVE_MINUTES)
	options.SetKeepAlive(datadog.SECURITYMONITORINGRULEKEEPALIVE_FIVE_MINUTES)
	options.SetMaxSignalDuration(datadog.SECURITYMONITORINGRULEMAXSIGNALDURATION_FIVE_MINUTES)

	ruleCase := datadog.NewSecurityMonitoringRuleCase()
	ruleCase.SetName("rule-case")
	ruleCase.SetCondition("nevermatch > 1000")
	ruleCase.SetStatus(datadog.SECURITYMONITORINGRULESEVERITY_INFO)
	ruleCases := []datadog.SecurityMonitoringRuleCase{*ruleCase}

	createPayload := datadog.NewSecurityMonitoringRuleCreatePayload(
		ruleCases,
		true,
		"rule message",
		ruleName,
		*options,
		queries,
		[]string{"datadog-api-client-test-go"},
		)

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

	ruleResponses := []datadog.SecurityMonitoringRuleResponse{}
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
	updatePayload.SetName(ruleResponse.GetName())
	updatePayload.SetEnabled(false)
	updatePayload.SetMessage(ruleResponse.GetMessage())
	updatePayload.SetTags(ruleResponse.GetTags())
	updatePayload.SetQueries(ruleResponse.GetQueries())
	updatePayload.SetOptions(ruleResponse.GetOptions())
	updatePayload.SetCases(ruleResponse.GetCases())
	updateResponse, httpResponse, err := api.UpdateSecurityMonitoringRule(ctx, ruleResponse.GetId()).
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
