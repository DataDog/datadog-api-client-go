package test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func enableSecMonUnstableOperations(ctx context.Context) func() {
	Client(ctx).GetConfig().SetUnstableOperationEnabled("SearchSecurityMonitoringSignals", true)
	Client(ctx).GetConfig().SetUnstableOperationEnabled("ListSecurityMonitoringSignals", true)
	return func() { disableSecMonUnstableOperations(ctx) }
}

func disableSecMonUnstableOperations(ctx context.Context) {
	Client(ctx).GetConfig().SetUnstableOperationEnabled("SearchSecurityMonitoringSignals", false)
	Client(ctx).GetConfig().SetUnstableOperationEnabled("ListSecurityMonitoringSignals", false)
}

func createRule(ctx context.Context, api *datadog.SecurityMonitoringApiService, ruleName string) (datadog.SecurityMonitoringRuleResponse, *http.Response, error) {

	query := datadog.NewSecurityMonitoringRuleQueryCreate(ruleName)
	query.SetName("rule")
	query.SetGroupByFields([]string{})
	queries := []datadog.SecurityMonitoringRuleQueryCreate{*query}

	options := datadog.NewSecurityMonitoringRuleOptions()
	options.SetEvaluationWindow(datadog.SECURITYMONITORINGRULEEVALUATIONWINDOW_FIVE_MINUTES)
	options.SetKeepAlive(datadog.SECURITYMONITORINGRULEKEEPALIVE_FIVE_MINUTES)
	options.SetMaxSignalDuration(datadog.SECURITYMONITORINGRULEMAXSIGNALDURATION_FIVE_MINUTES)

	ruleCase := datadog.NewSecurityMonitoringRuleCaseCreate(datadog.SECURITYMONITORINGRULESEVERITY_INFO)
	ruleCase.SetName("rule-case")
	ruleCase.SetCondition("rule > 0")
	ruleCases := []datadog.SecurityMonitoringRuleCaseCreate{*ruleCase}

	filter := datadog.NewSecurityMonitoringFilter()
	filter.SetAction(datadog.SECURITYMONITORINGFILTERACTION_REQUIRE)
	filter.SetQuery("env:prod")
	filters := []datadog.SecurityMonitoringFilter{*filter}

	createPayload := datadog.NewSecurityMonitoringRuleCreatePayload(
		ruleCases,
		true,
		"rule message",
		ruleName,
		*options,
		queries,
	)
	createPayload.SetFilters(filters)

	return api.CreateSecurityMonitoringRule(ctx, *createPayload)
}

func deleteRule(ctx context.Context, t *testing.T, api *datadog.SecurityMonitoringApiService, id string) {
	_, err := api.DeleteSecurityMonitoringRule(ctx, id)
	if err != nil {
		t.Logf("Error delete rule: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
}

func TestSecMonRulesCRUD(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
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
		defer deleteRule(ctx, t, api, ruleResponse.GetId())
	}

	// get single rule
	ruleResponse := ruleResponses[0]
	getResponse, httpResponse, err := api.GetSecurityMonitoringRule(ctx, ruleResponse.GetId())
	if err != nil {
		t.Fatalf("Error getting rule: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpResponse.StatusCode)
	assert.Equal(ruleResponse.GetName(), getResponse.GetName())

	//// get rule list
	// get filter count
	listResponse, httpResponse, err := api.ListSecurityMonitoringRules(ctx, *datadog.NewListSecurityMonitoringRulesOptionalParameters().
		WithPageSize(1).
		WithPageNumber(0))
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
	listResponse, httpResponse, err = api.ListSecurityMonitoringRules(ctx, *datadog.NewListSecurityMonitoringRulesOptionalParameters().
		WithPageSize(ruleCount).
		WithPageNumber(0))
	if err != nil {
		t.Fatalf("Error listing rules: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpResponse.StatusCode)
	ruleIds := getRuleIds(listResponse.GetData())
	for _, rule := range ruleResponses {
		assert.Contains(ruleIds, rule.GetId())
	}

	// paging
	firstPageResponse, httpResponse, err := api.ListSecurityMonitoringRules(ctx, *datadog.NewListSecurityMonitoringRulesOptionalParameters().
		WithPageSize(2).
		WithPageNumber(0))
	if err != nil {
		t.Fatalf("Error listing rules: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpResponse.StatusCode)
	assert.Equal(2, len(firstPageResponse.GetData()))
	secondPageResponse, httpResponse, err := api.ListSecurityMonitoringRules(ctx, *datadog.NewListSecurityMonitoringRulesOptionalParameters().
		WithPageSize(2).
		WithPageNumber(0))
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
	updatePayload.SetIsEnabled(false)
	updatePayload.SetMessage(ruleResponse.GetMessage())
	updatePayload.SetTags(ruleResponse.GetTags())
	updatePayload.SetQueries(ruleResponse.GetQueries())
	updatePayload.SetOptions(ruleResponse.GetOptions())
	updatePayload.SetCases(ruleResponse.GetCases())
	updateResponse, httpResponse, err := api.UpdateSecurityMonitoringRule(ctx, ruleResponse.GetId(), *updatePayload)
	if err != nil {
		t.Fatalf("Error updating rule: Response %v", err)
	}
	assert.Equal(200, httpResponse.StatusCode)
	assert.Equal(false, updateResponse.GetIsEnabled())

	getUpdatedResponse, httpResponse, err := api.GetSecurityMonitoringRule(ctx, ruleResponse.GetId())
	if err != nil {
		t.Fatalf("Error updating rule: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpResponse.StatusCode)
	assert.Equal(false, getUpdatedResponse.GetIsEnabled())

	// Delete rule
	httpResponse, err = api.DeleteSecurityMonitoringRule(ctx, ruleResponse.GetId())
	if err != nil {
		t.Fatalf("Error delete rule: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(204, httpResponse.StatusCode)

	_, httpResponse, _ = api.GetSecurityMonitoringRule(ctx, ruleResponse.GetId())
	assert.Equal(404, httpResponse.StatusCode)
}

func TestSearchSecurityMonitoringSignals(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	defer enableSecMonUnstableOperations(ctx)()
	client := Client(ctx)
	api := client.SecurityMonitoringApi

	if tests.GetRecording() == tests.ModeIgnore {
		t.Skipf("Unreliable test")
	}

	now := tests.ClockFromContext(ctx).Now()
	uniqueName := tests.UniqueEntityName(ctx, t)

	_, _, err := createRule(ctx, api, *uniqueName)
	if err != nil {
		t.Fatalf("Error creating rule: Response: %v", err)
	}

	var response datadog.SecurityMonitoringSignalsListResponse
	var httpResp *http.Response

	filter := datadog.NewSecurityMonitoringSignalListRequestFilter()
	filter.SetQuery(*uniqueName)
	filter.SetFrom(now.Add(time.Duration(-2) * time.Hour))
	filter.SetTo(now.Add(time.Duration(2) * time.Hour))

	request := datadog.NewSecurityMonitoringSignalListRequestWithDefaults()
	request.SetFilter(*filter)

	// Make sure both signals are generated
	err = tests.Retry(time.Duration(5)*time.Second, 30, func() bool {
		_ = sendLogsSignals(ctx, client, *uniqueName)
		response, httpResp, err = api.SearchSecurityMonitoringSignals(ctx, *datadog.NewSearchSecurityMonitoringSignalsOptionalParameters().
			WithBody(*request))
		return err == nil && 200 == httpResp.StatusCode && 2 == len(response.GetData())
	})

	if err != nil {
		t.Fatalf("Could not list sent signals: %v", err)
	}

	// Sort works correctly
	request.SetSort(datadog.SECURITYMONITORINGSIGNALSSORT_TIMESTAMP_ASCENDING)

	err = tests.Retry(time.Duration(5)*time.Second, 30, func() bool {
		response, httpResp, err = api.SearchSecurityMonitoringSignals(ctx, *datadog.NewSearchSecurityMonitoringSignalsOptionalParameters().
			WithBody(*request))
		return err == nil && 200 == httpResp.StatusCode && 2 == len(response.GetData())
	})
	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(2, len(response.GetData()))
	firstTimestamp := response.GetData()[0].GetAttributes().Timestamp
	secondTimestamp := response.GetData()[1].GetAttributes().Timestamp
	assert.True(firstTimestamp.Before(*secondTimestamp))

	request.SetSort(datadog.SECURITYMONITORINGSIGNALSSORT_TIMESTAMP_DESCENDING)

	err = tests.Retry(time.Duration(5)*time.Second, 30, func() bool {
		response, httpResp, err = api.SearchSecurityMonitoringSignals(ctx, *datadog.NewSearchSecurityMonitoringSignalsOptionalParameters().
			WithBody(*request))
		return err == nil && 200 == httpResp.StatusCode && 2 == len(response.GetData())
	})
	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(2, len(response.GetData()))
	firstTimestamp = response.GetData()[0].GetAttributes().Timestamp
	secondTimestamp = response.GetData()[1].GetAttributes().Timestamp
	assert.True(firstTimestamp.After(*secondTimestamp))

	// Paging
	page := datadog.NewSecurityMonitoringSignalListRequestPage()
	page.SetLimit(1)
	request.SetPage(*page)
	response, httpResp, err = api.SearchSecurityMonitoringSignals(ctx, *datadog.NewSearchSecurityMonitoringSignalsOptionalParameters().
		WithBody(*request))
	if err != nil {
		t.Fatalf("Could not list signals: %v", err)
	}

	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(1, len(response.GetData()))
	respMeta := response.GetMeta()
	respPage := respMeta.GetPage()
	cursor := respPage.GetAfter()
	firstID := response.GetData()[0].GetId()

	request.Page.SetCursor(cursor)
	response, httpResp, err = api.SearchSecurityMonitoringSignals(ctx, *datadog.NewSearchSecurityMonitoringSignalsOptionalParameters().
		WithBody(*request))
	if err != nil {
		t.Fatalf("Could not list signals: %v", err)
	}
	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(1, len(response.GetData()))
	secondID := response.GetData()[0].GetId()

	assert.NotEqual(firstID, secondID)
}

func TestListSecurityMonitoringSignals(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	defer enableSecMonUnstableOperations(ctx)()
	client := Client(ctx)
	api := client.SecurityMonitoringApi

	if tests.GetRecording() == tests.ModeIgnore {
		t.Skipf("Unreliable test")
	}

	now := tests.ClockFromContext(ctx).Now()
	uniqueName := tests.UniqueEntityName(ctx, t)

	_, _, err := createRule(ctx, api, *uniqueName)
	if err != nil {
		t.Fatalf("Error creating rule: Response: %v", err)
	}

	var response datadog.SecurityMonitoringSignalsListResponse
	var httpResp *http.Response

	from := now.Add(time.Duration(-2) * time.Hour)
	to := now.Add(time.Duration(2) * time.Hour)

	// Make sure both signals are generated
	err = tests.Retry(time.Duration(5)*time.Second, 30, func() bool {
		_ = sendLogsSignals(ctx, client, *uniqueName)

		response, httpResp, err = api.ListSecurityMonitoringSignals(ctx, *datadog.NewListSecurityMonitoringSignalsOptionalParameters().
			WithFilterQuery(*uniqueName).
			WithFilterFrom(from).
			WithFilterTo(to))
		return err == nil && 200 == httpResp.StatusCode && 2 == len(response.GetData())
	})

	if err != nil {
		t.Fatalf("Could not list signals: %v", err)
	}

	// Sort works correctly
	err = tests.Retry(time.Duration(5)*time.Second, 30, func() bool {
		response, httpResp, err = api.ListSecurityMonitoringSignals(ctx, *datadog.NewListSecurityMonitoringSignalsOptionalParameters().
			WithFilterQuery(*uniqueName).
			WithFilterFrom(from).
			WithFilterTo(to).
			WithSort(datadog.SECURITYMONITORINGSIGNALSSORT_TIMESTAMP_ASCENDING))
		return err == nil && 200 == httpResp.StatusCode && 2 == len(response.GetData())
	})
	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(2, len(response.GetData()))
	firstTimestamp := response.GetData()[0].GetAttributes().Timestamp
	secondTimestamp := response.GetData()[1].GetAttributes().Timestamp
	assert.True(firstTimestamp.Before(*secondTimestamp))

	err = tests.Retry(time.Duration(5)*time.Second, 30, func() bool {
		response, httpResp, err = api.ListSecurityMonitoringSignals(ctx, *datadog.NewListSecurityMonitoringSignalsOptionalParameters().
			WithFilterQuery(*uniqueName).
			WithFilterFrom(from).
			WithFilterTo(to).
			WithSort(datadog.SECURITYMONITORINGSIGNALSSORT_TIMESTAMP_DESCENDING))
		return err == nil && 200 == httpResp.StatusCode && 2 == len(response.GetData())
	})
	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(2, len(response.GetData()))
	firstTimestamp = response.GetData()[0].GetAttributes().Timestamp
	secondTimestamp = response.GetData()[1].GetAttributes().Timestamp
	assert.True(firstTimestamp.After(*secondTimestamp))

	// Paging
	response, httpResp, err = api.ListSecurityMonitoringSignals(ctx, *datadog.NewListSecurityMonitoringSignalsOptionalParameters().
		WithFilterQuery(*uniqueName).
		WithFilterFrom(from).
		WithFilterTo(to).
		WithPageLimit(1))
	if err != nil {
		t.Fatalf("Could not list signals: %v", err)
	}

	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(1, len(response.GetData()))
	respMeta := response.GetMeta()
	respPage := respMeta.GetPage()
	cursor := respPage.GetAfter()
	firstID := response.GetData()[0].GetId()

	response, httpResp, err = api.ListSecurityMonitoringSignals(ctx, *datadog.NewListSecurityMonitoringSignalsOptionalParameters().
		WithFilterQuery(*uniqueName).
		WithFilterFrom(from).
		WithFilterTo(to).
		WithPageLimit(1).
		WithPageCursor(cursor))
	if err != nil {
		t.Fatalf("Could not list signals: %v", err)
	}
	assert.Equal(200, httpResp.StatusCode)
	assert.Equal(1, len(response.GetData()))
	secondID := response.GetData()[0].GetId()

	assert.NotEqual(firstID, secondID)
}

func sendLogsSignals(ctx context.Context, client *datadog.APIClient, suffix string) error {
	now := tests.ClockFromContext(ctx).Now()
	source := fmt.Sprintf("go-client-test-%s", suffix)

	httpLog := fmt.Sprintf(
		`{"ddsource": "%s", "ddtags": "go,test,signal,env:prod", "message": "{\"timestamp\": %d, \"message\": \"%s\"}"}`,
		source, (now.Unix()-1000)*1000, suffix,
	)

	domain, err := getTestDomain(ctx, client)
	if err != nil {
		return fmt.Errorf("parsing domain: %v", err)
	}
	intakeURL := fmt.Sprintf("https://http-intake.logs.%s/v1/input", domain)

	httpresp, respBody, err := SendRequest(ctx, "POST", intakeURL, []byte(httpLog))
	if err != nil {
		return fmt.Errorf("response %s: %v", respBody, err)
	}
	if httpresp.StatusCode != 200 {
		return fmt.Errorf("status=%d response=%v", httpresp.StatusCode, respBody)
	}

	httpLog = fmt.Sprintf(
		`{"ddsource": "%s", "ddtags": "go,test,signal,env:prod", "message": "{\"timestamp\": %d, \"message\": \"%s\"}"}`,
		source, now.Unix()*1000, suffix,
	)

	httpresp, respBody, err = SendRequest(ctx, "POST", intakeURL, []byte(httpLog))
	if err != nil {
		return fmt.Errorf("error creating log: Response %s: %v", respBody, err)
	}
	if httpresp.StatusCode != 200 {
		return fmt.Errorf("unable to send log: Status=%d Response=%v", httpresp.StatusCode, respBody)
	}

	return nil
}

func getRuleIds(rules []datadog.SecurityMonitoringRuleResponse) map[string]bool {
	ruleIds := map[string]bool{}
	for _, rule := range rules {
		ruleIds[rule.GetId()] = true
	}
	return ruleIds
}
