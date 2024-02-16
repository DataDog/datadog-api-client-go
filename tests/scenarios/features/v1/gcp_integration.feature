@endpoint(gcp-integration) @endpoint(gcp-integration-v1)
Feature: GCP Integration
  Configure your Datadog-Google Cloud Platform (GCP) integration directly
  through the Datadog API. Read more about the [Datadog-Google Cloud
  Platform integration](https://docs.datadoghq.com/integrations/google_cloud
  _platform).

  Background:
    Given a valid "apiKeyAuth" key in the system
    And a valid "appKeyAuth" key in the system
    And an instance of "GCPIntegration" API

  @generated @skip @team:DataDog/gcp-integrations
  Scenario: Create a GCP integration returns "Bad Request" response
    Given new "CreateGCPIntegration" request
    And body with value {"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs", "auth_uri": "https://accounts.google.com/o/oauth2/auth", "client_email": "api-dev@datadog-sandbox.iam.gserviceaccount.com", "client_id": "123456712345671234567", "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/$CLIENT_EMAIL", "cloud_run_revision_filters": ["$KEY:$VALUE"], "errors": ["*"], "host_filters": "key:value,filter:example", "is_cspm_enabled": true, "is_security_command_center_enabled": true, "private_key": "private_key", "private_key_id": "123456789abcdefghi123456789abcdefghijklm", "project_id": "datadog-apitest", "resource_collection_enabled": true, "token_uri": "https://accounts.google.com/o/oauth2/token", "type": "service_account"}
    When the request is sent
    Then the response status is 400 Bad Request

  @team:DataDog/gcp-integrations
  Scenario: Create a GCP integration returns "OK" response
    Given new "CreateGCPIntegration" request
    And body with value {"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs", "auth_uri": "https://accounts.google.com/o/oauth2/auth", "client_email": "{{unique_hash}}@example.com", "client_id": "{{ timestamp("now") }}{{ timestamp("now") }}0", "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/$CLIENT_EMAIL", "host_filters": "key:value,filter:example", "cloud_run_revision_filters": ["dr:dre"], "is_cspm_enabled": true, "is_security_command_center_enabled": true, "private_key": "private_key", "private_key_id": "123456789abcdefghi123456789abcdefghijklm", "project_id": "datadog-apitest", "resource_collection_enabled": true, "token_uri": "https://accounts.google.com/o/oauth2/token", "type": "service_account"}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip @team:DataDog/gcp-integrations
  Scenario: Delete a GCP integration returns "Bad Request" response
    Given new "DeleteGCPIntegration" request
    And body with value {"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs", "auth_uri": "https://accounts.google.com/o/oauth2/auth", "client_email": "api-dev@datadog-sandbox.iam.gserviceaccount.com", "client_id": "123456712345671234567", "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/$CLIENT_EMAIL", "cloud_run_revision_filters": ["$KEY:$VALUE"], "errors": ["*"], "host_filters": "key:value,filter:example", "is_cspm_enabled": true, "is_security_command_center_enabled": true, "private_key": "private_key", "private_key_id": "123456789abcdefghi123456789abcdefghijklm", "project_id": "datadog-apitest", "resource_collection_enabled": true, "token_uri": "https://accounts.google.com/o/oauth2/token", "type": "service_account"}
    When the request is sent
    Then the response status is 400 Bad Request

  @team:DataDog/gcp-integrations
  Scenario: Delete a GCP integration returns "OK" response
    Given there is a valid "gcp_account" in the system
    And new "DeleteGCPIntegration" request
    And body with value {"client_email": "{{unique_hash}}@example.com", "client_id": "{{ timestamp("now") }}{{ timestamp("now") }}0", "project_id": "datadog-apitest"}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip @team:DataDog/gcp-integrations
  Scenario: List all GCP integrations returns "Bad Request" response
    Given new "ListGCPIntegration" request
    When the request is sent
    Then the response status is 400 Bad Request

  @team:DataDog/gcp-integrations
  Scenario: List all GCP integrations returns "OK" response
    Given new "ListGCPIntegration" request
    When the request is sent
    Then the response status is 200 OK

  @team:DataDog/gcp-integrations
  Scenario: Update a GCP integration cloud run revision filters returns "OK" response
    Given there is a valid "gcp_account" in the system
    And new "UpdateGCPIntegration" request
    And body with value {"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs", "auth_uri": "https://accounts.google.com/o/oauth2/auth", "client_email": "{{unique_hash}}@example.com", "client_id": "{{ timestamp("now") }}{{ timestamp("now") }}0", "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/$CLIENT_EMAIL", "host_filters": "key:value,filter:example", "cloud_run_revision_filters": ["merp:derp"], "is_cspm_enabled": true, "is_security_command_center_enabled": true, "private_key": "private_key", "private_key_id": "123456789abcdefghi123456789abcdefghijklm", "project_id": "datadog-apitest", "resource_collection_enabled": true, "token_uri": "https://accounts.google.com/o/oauth2/token", "type": "service_account"}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip @team:DataDog/gcp-integrations
  Scenario: Update a GCP integration returns "Bad Request" response
    Given new "UpdateGCPIntegration" request
    And body with value {"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs", "auth_uri": "https://accounts.google.com/o/oauth2/auth", "client_email": "api-dev@datadog-sandbox.iam.gserviceaccount.com", "client_id": "123456712345671234567", "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/$CLIENT_EMAIL", "cloud_run_revision_filters": ["$KEY:$VALUE"], "errors": ["*"], "host_filters": "key:value,filter:example", "is_cspm_enabled": true, "is_security_command_center_enabled": true, "private_key": "private_key", "private_key_id": "123456789abcdefghi123456789abcdefghijklm", "project_id": "datadog-apitest", "resource_collection_enabled": true, "token_uri": "https://accounts.google.com/o/oauth2/token", "type": "service_account"}
    When the request is sent
    Then the response status is 400 Bad Request

  @team:DataDog/gcp-integrations
  Scenario: Update a GCP integration returns "OK" response
    Given there is a valid "gcp_account" in the system
    And new "UpdateGCPIntegration" request
    And body with value {"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs", "auth_uri": "https://accounts.google.com/o/oauth2/auth", "client_email": "{{unique_hash}}@example.com", "client_id": "{{ timestamp("now") }}{{ timestamp("now") }}0", "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/$CLIENT_EMAIL", "host_filters": "key:value,filter:example", "is_cspm_enabled": true, "is_security_command_center_enabled": true, "private_key": "private_key", "private_key_id": "123456789abcdefghi123456789abcdefghijklm", "project_id": "datadog-apitest", "resource_collection_enabled": true, "token_uri": "https://accounts.google.com/o/oauth2/token", "type": "service_account"}
    When the request is sent
    Then the response status is 200 OK
