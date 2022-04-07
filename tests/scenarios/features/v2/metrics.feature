@endpoint(metrics) @endpoint(metrics-v2)
Feature: Metrics
  The metrics endpoint allows you to:  - Post metrics data so it can be
  graphed on Datadog’s dashboards - Query metrics from any time period -
  Modify tag configurations for metrics - View tags and volumes for metrics
  **Note**: A graph can only contain a set number of points and as the
  timeframe over which a metric is viewed increases, aggregation between
  points occurs to stay below that set number.  The Post, Patch, and Delete
  `manage_tags` API methods can only be performed by a user who has the
  `Manage Tags for Metrics` permission.

  Background:
    Given a valid "apiKeyAuth" key in the system
    And a valid "appKeyAuth" key in the system
    And an instance of "Metrics" API

  @team:DataDog/points-aggregation
  Scenario: Configure tags for multiple metrics returns "Accepted" response
    Given there is a valid "user" in the system
    And new "CreateBulkTagsMetricsConfiguration" request
    And body with value {"data": {"attributes": {"emails": ["{{ user.data.attributes.email }}"], "tags": ["test", "{{ unique_lower_alnum }}"]}, "id": "system.load.1", "type": "metric_bulk_configure_tags"}}
    When the request is sent
    Then the response status is 202 Accepted

  @generated @skip @team:DataDog/points-aggregation
  Scenario: Configure tags for multiple metrics returns "Bad Request" response
    Given new "CreateBulkTagsMetricsConfiguration" request
    And body with value {"data": {"attributes": {"emails": ["sue@example.com", "bob@example.com"], "tags": ["host", "pod_name", "is_shadow"]}, "id": "kafka.lag", "type": "metric_bulk_configure_tags"}}
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/points-aggregation
  Scenario: Configure tags for multiple metrics returns "Not Found" response
    Given new "CreateBulkTagsMetricsConfiguration" request
    And body with value {"data": {"attributes": {"emails": ["sue@example.com", "bob@example.com"], "tags": ["host", "pod_name", "is_shadow"]}, "id": "kafka.lag", "type": "metric_bulk_configure_tags"}}
    When the request is sent
    Then the response status is 404 Not Found

  @generated @skip @team:DataDog/points-aggregation
  Scenario: Create a tag configuration returns "Bad Request" response
    Given operation "CreateTagConfiguration" enabled
    And new "CreateTagConfiguration" request
    And request contains "metric_name" parameter from "REPLACE.ME"
    And body with value {"data": {"attributes": {"include_percentiles": false, "metric_type": "distribution", "tags": ["app", "datacenter"]}, "id": "http.endpoint.request", "type": "manage_tags"}}
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/points-aggregation
  Scenario: Create a tag configuration returns "Conflict" response
    Given operation "CreateTagConfiguration" enabled
    And new "CreateTagConfiguration" request
    And request contains "metric_name" parameter from "REPLACE.ME"
    And body with value {"data": {"attributes": {"include_percentiles": false, "metric_type": "distribution", "tags": ["app", "datacenter"]}, "id": "http.endpoint.request", "type": "manage_tags"}}
    When the request is sent
    Then the response status is 409 Conflict

  @skip @team:DataDog/points-aggregation
  Scenario: Create a tag configuration returns "Created" response
    Given operation "CreateTagConfiguration" enabled
    And new "CreateTagConfiguration" request
    And request contains "metric_name" parameter with value "{{ unique_alnum }}"
    And body with value {"data": {"type": "manage_tags", "id": "{{ unique_alnum }}", "attributes": {"tags": ["app","datacenter"], "metric_type": "distribution"}}}
    When the request is sent
    Then the response status is 201 Created

  @skip @team:DataDog/points-aggregation
  Scenario: Delete a tag configuration returns "No Content" response
    Given there is a valid "metric_tag_configuration" in the system
    And operation "DeleteTagConfiguration" enabled
    And new "DeleteTagConfiguration" request
    And request contains "metric_name" parameter from "metric_tag_configuration.data.id"
    When the request is sent
    Then the response status is 204 No Content

  @generated @skip @team:DataDog/points-aggregation
  Scenario: Delete a tag configuration returns "Not found" response
    Given operation "DeleteTagConfiguration" enabled
    And new "DeleteTagConfiguration" request
    And request contains "metric_name" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 404 Not found

  @generated @skip @team:DataDog/points-aggregation
  Scenario: List distinct metric volumes by metric name returns "Bad Request" response
    Given new "ListVolumesByMetricName" request
    And request contains "metric_name" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/points-aggregation
  Scenario: List distinct metric volumes by metric name returns "Not Found" response
    Given new "ListVolumesByMetricName" request
    And request contains "metric_name" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 404 Not Found

  @skip @team:DataDog/points-aggregation
  Scenario: List distinct metric volumes by metric name returns "Success" response
    Given there is a valid "metric_tag_configuration" in the system
    And new "ListVolumesByMetricName" request
    And request contains "metric_name" parameter from "metric_tag_configuration.data.id"
    When the request is sent
    Then the response status is 200 Success
    And the response "data.id" has the same value as "metric_tag_configuration.data.id"

  @generated @skip @team:DataDog/points-aggregation
  Scenario: List tag configuration by name returns "Not Found" response
    Given operation "ListTagConfigurationByName" enabled
    And new "ListTagConfigurationByName" request
    And request contains "metric_name" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 404 Not Found

  @skip @team:DataDog/points-aggregation
  Scenario: List tag configuration by name returns "Success" response
    Given there is a valid "metric_tag_configuration" in the system
    And operation "ListTagConfigurationByName" enabled
    And new "ListTagConfigurationByName" request
    And request contains "metric_name" parameter from "metric_tag_configuration.data.id"
    When the request is sent
    Then the response status is 200 Success
    And the response "data.id" has the same value as "metric_tag_configuration.data.id"

  @generated @skip @team:DataDog/points-aggregation
  Scenario: List tag configurations returns "Bad Request" response
    Given operation "ListTagConfigurations" enabled
    And new "ListTagConfigurations" request
    When the request is sent
    Then the response status is 400 Bad Request

  @skip @team:DataDog/points-aggregation
  Scenario: List tag configurations returns "Success" response
    Given there is a valid "metric_tag_configuration" in the system
    And operation "ListTagConfigurations" enabled
    And new "ListTagConfigurations" request
    And request contains "filter[tags_configured]" parameter with value "{{ unique_alnum }}"
    When the request is sent
    Then the response status is 200 Success
    And the response "data[0].id" has the same value as "metric_tag_configuration.data.id"

  @team:DataDog/points-aggregation
  Scenario: List tag configurations with a tag filter returns "Success" response
    Given operation "ListTagConfigurations" enabled
    And new "ListTagConfigurations" request
    And request contains "filter[tags]" parameter with value "{{ unique_alnum }}"
    When the request is sent
    Then the response status is 200 Success

  @generated @skip @team:DataDog/points-aggregation
  Scenario: List tags by metric name returns "Bad Request" response
    Given new "ListTagsByMetricName" request
    And request contains "metric_name" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/points-aggregation
  Scenario: List tags by metric name returns "Not Found" response
    Given new "ListTagsByMetricName" request
    And request contains "metric_name" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 404 Not Found

  @skip @team:DataDog/points-aggregation
  Scenario: List tags by metric name returns "Success" response
    Given there is a valid "metric_tag_configuration" in the system
    And new "ListTagsByMetricName" request
    And request contains "metric_name" parameter from "metric_tag_configuration.data.id"
    When the request is sent
    Then the response status is 200 Success
    And the response "data.id" has the same value as "metric_tag_configuration.data.id"

  @generated @skip @team:DataDog/points-aggregation
  Scenario: Update a tag configuration returns "Bad Request" response
    Given operation "UpdateTagConfiguration" enabled
    And new "UpdateTagConfiguration" request
    And request contains "metric_name" parameter from "REPLACE.ME"
    And body with value {"data": {"attributes": {"group_by": ["app", "datacenter"], "include_percentiles": false}, "id": "http.endpoint.request", "type": "manage_tags"}}
    When the request is sent
    Then the response status is 400 Bad Request

  @skip @team:DataDog/points-aggregation
  Scenario: Update a tag configuration returns "OK" response
    Given operation "UpdateTagConfiguration" enabled
    And there is a valid "metric_tag_configuration" in the system
    And new "UpdateTagConfiguration" request
    And request contains "metric_name" parameter from "metric_tag_configuration.data.id"
    And body with value {"data": {"type": "manage_tags", "id": "{{ metric_tag_configuration.data.id }}", "attributes": {"tags": ["app"]}}}
    When the request is sent
    Then the response status is 200 OK
    And the response "data.attributes.tags[0]" is equal to "app"

  @generated @skip @team:DataDog/points-aggregation
  Scenario: Update a tag configuration returns "Unprocessable Entity" response
    Given operation "UpdateTagConfiguration" enabled
    And new "UpdateTagConfiguration" request
    And request contains "metric_name" parameter from "REPLACE.ME"
    And body with value {"data": {"attributes": {"group_by": ["app", "datacenter"], "include_percentiles": false}, "id": "http.endpoint.request", "type": "manage_tags"}}
    When the request is sent
    Then the response status is 422 Unprocessable Entity
