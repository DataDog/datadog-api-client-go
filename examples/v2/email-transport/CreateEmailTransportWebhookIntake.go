// Ingest email transport webhook events returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := []datadogV2.TransportWebhookLog{
		{
			Attributes: datadogV2.TransportWebhookLogAttributes{
				Category: []string{
					"transactional",
				},
				Email: &datadogV2.TransportWebhookLogEmail{
					Address: datadog.PtrString("user@example.com"),
					Domain:  datadog.PtrString("example.com"),
					Subject: datadog.PtrString("[Monitor Alert] CPU usage is high"),
					Type: []string{
						"transactional",
					},
				},
				EmailId:              datadog.PtrString("abc123-def456"),
				EmailTypeDisplayName: datadog.PtrString("Monitor Alert"),
				Message: &datadogV2.TransportWebhookLogMessage{
					Auth: &datadogV2.TransportWebhookLogMessageAuth{
						DeliveredWithTls: datadog.PtrString("TLSv1.2"),
					},
					CustomArgs: &datadogV2.TransportWebhookLogMessageCustomArgs{
						EmailId:              datadog.PtrString("abc123-def456"),
						EmailTypeDisplayName: datadog.PtrString("Monitor Alert"),
						OrgUuid:              datadog.PtrString("8dee7c38-00cb-11ea-a77b-8b5a08d3b091"),
						QueueTime:            datadog.PtrString("2024-01-15T10:29:00Z"),
						Subject:              datadog.PtrString("[Monitor Alert] CPU usage is high"),
					},
					Id: &datadogV2.TransportWebhookLogMessageId{
						MessageId:        datadog.PtrString("<message-id@example.com>"),
						SmtpId:           datadog.PtrString("<abc123@mail.example.com>"),
						TransportEventId: datadog.PtrString("evt_abc123"),
					},
					Name: datadog.PtrString("delivered"),
					Response: &datadogV2.TransportWebhookLogMessageResponse{
						EnhancedSmtpCode: datadog.PtrString("2.0.0"),
						Reason:           datadog.PtrString("250 2.0.0 OK"),
						SmtpCode:         datadog.PtrString("250"),
					},
					SenderIp: datadog.PtrString("192.168.1.1"),
					Timestamp: &datadogV2.TransportWebhookLogMessageTimestamp{
						EventTimestamp: datadog.PtrFloat64(1705312200.0),
						Lifetime:       datadog.PtrFloat64(3.2),
						QueueTime:      datadog.PtrFloat64(1.5),
						ScheduledTime:  datadog.PtrFloat64(1705312190.0),
					},
				},
				Network: &datadogV2.TransportWebhookLogNetwork{
					Ip: &datadogV2.TransportWebhookLogNetworkIp{
						Attributes: []datadogV2.TransportWebhookLogIpAttribute{
							{
								Ip: datadog.PtrString("192.168.1.1"),
								Source: []string{
									"sendgrid",
								},
							},
						},
						List: []string{
							"192.168.1.1",
						},
					},
				},
				Org:         datadog.PtrInt64(1234),
				OrgMetadata: &datadogV2.TransportWebhookLogOrgMetadata{},
				OrgUuid:     datadog.PtrString("8dee7c38-00cb-11ea-a77b-8b5a08d3b091"),
				QueueTime:   datadog.PtrString("2024-01-15T10:29:00Z"),
				Subject:     datadog.PtrString("[Monitor Alert] CPU usage is high"),
				Useragent:   datadog.PtrString("Mozilla/5.0"),
			},
			Date:   time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC),
			LogId:  "AQAAAZPHnBT0TwJAdgAAAABBWlBIblVlNEFBQ0dFMmVkYTFDSnRR",
			Source: "sendgrid",
			Status: "info",
			Tags: []string{
				"env:production",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateEmailTransportWebhookIntake", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewEmailTransportApi(apiClient)
	r, err := api.CreateEmailTransportWebhookIntake(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailTransportApi.CreateEmailTransportWebhookIntake`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
