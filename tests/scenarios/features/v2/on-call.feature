@endpoint(on-call) @endpoint(on-call-v2)
Feature: On-Call
  Configure your [Datadog On-
  Call](https://docs.datadoghq.com/service_management/on-call/) directly
  through the Datadog API.

  Background:
    Given a valid "apiKeyAuth" key in the system
    And a valid "appKeyAuth" key in the system
    And an instance of "On-Call" API

  @generated @skip @team:DataDog/bugle
  Scenario: Create on-call schedule returns "Bad Request" response
    Given new "CreateOnCallSchedule" request
    And body with value {"data": {"attributes": {"layers": [{"effective_date": "2025-02-03T05:00:00Z", "end_date": "2025-12-31T00:00:00Z", "interval": {"days": 1}, "members": [{"user": {"id": "00000000-aba1-0000-0000-000000000000"}}], "name": "Layer 1", "restrictions": [{"end_day": "friday", "end_time": "17:00:00", "start_day": "monday", "start_time": "09:00:00"}], "rotation_start": "2025-02-01T00:00:00Z"}], "name": "On-Call Schedule", "tags": ["tag1", "tag2"], "time_zone": "America/New_York"}, "relationships": {"teams": {"data": [{"id": "00000000-da3a-0000-0000-000000000000", "type": "teams"}]}}, "type": "schedules"}}
    When the request is sent
    Then the response status is 400 Bad Request

  @skip-python @team:DataDog/bugle
  Scenario: Create on-call schedule returns "Created" response
    Given new "CreateOnCallSchedule" request
    And there is a valid "user" in the system
    And there is a valid "team" in the system
    And body with value {"data": {"attributes": {"layers": [{"effective_date": "{{ timeISO('now - 10d') }}", "end_date": "{{ timeISO('now + 10d') }}", "interval": {"days": 1}, "members": [{"user": {"id": "{{user.data.id}}"}}], "name": "Layer 1", "restrictions": [{"end_day": "friday", "end_time": "17:00:00", "start_day": "monday", "start_time": "09:00:00"}], "rotation_start": "{{ timeISO('now - 5d') }}"}], "name": "{{ unique }}", "tags": ["tag1", "tag2"], "time_zone": "America/New_York"}, "relationships": {"teams": {"data": [{"id": "{{team.data.id}}", "type": "teams"}]}}, "type": "schedules"}}
    When the request is sent
    Then the response status is 201 Created

  @team:DataDog/bugle
  Scenario: Delete on-call schedule returns "No Content" response
    Given new "DeleteOnCallSchedule" request
    And there is a valid "schedule" in the system
    And request contains "schedule_id" parameter from "schedule.data.id"
    When the request is sent
    Then the response status is 204 No Content

  @generated @skip @team:DataDog/bugle
  Scenario: Delete on-call schedule returns "Not Found" response
    Given new "DeleteOnCallSchedule" request
    And request contains "schedule_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 404 Not Found

  @generated @skip @team:DataDog/bugle
  Scenario: Get on-call schedule returns "Not Found" response
    Given new "GetOnCallSchedule" request
    And request contains "schedule_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 404 Not Found

  @team:DataDog/bugle
  Scenario: Get on-call schedule returns "OK" response
    Given new "GetOnCallSchedule" request
    And there is a valid "schedule" in the system
    And request contains "schedule_id" parameter from "schedule.data.id"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip @team:DataDog/bugle
  Scenario: Update on-call schedule returns "Bad Request" response
    Given new "UpdateOnCallSchedule" request
    And request contains "schedule_id" parameter from "REPLACE.ME"
    And body with value {"data": {"attributes": {"layers": [{"effective_date": "2025-02-03T05:00:00Z", "end_date": "2025-12-31T00:00:00Z", "interval": {"seconds": 300}, "members": [{"user": {"id": "00000000-aba1-0000-0000-000000000000"}}], "name": "Layer 1", "restrictions": [{"end_day": "friday", "end_time": "17:00:00", "start_day": "monday", "start_time": "09:00:00"}], "rotation_start": "2025-02-01T00:00:00Z"}], "name": "On-Call Schedule Updated", "tags": ["tag1", "tag2", "tag3"], "time_zone": "America/New_York"}, "id": "3653d3c6-0c75-11ea-ad28-fb5701eabc7d", "relationships": {"teams": {"data": [{"id": "00000000-da3a-0000-0000-000000000000", "type": "teams"}]}}, "type": "schedules"}}
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/bugle
  Scenario: Update on-call schedule returns "Not Found" response
    Given new "UpdateOnCallSchedule" request
    And request contains "schedule_id" parameter from "REPLACE.ME"
    And body with value {"data": {"attributes": {"layers": [{"effective_date": "2025-02-03T05:00:00Z", "end_date": "2025-12-31T00:00:00Z", "interval": {"seconds": 300}, "members": [{"user": {"id": "00000000-aba1-0000-0000-000000000000"}}], "name": "Layer 1", "restrictions": [{"end_day": "friday", "end_time": "17:00:00", "start_day": "monday", "start_time": "09:00:00"}], "rotation_start": "2025-02-01T00:00:00Z"}], "name": "On-Call Schedule Updated", "tags": ["tag1", "tag2", "tag3"], "time_zone": "America/New_York"}, "id": "3653d3c6-0c75-11ea-ad28-fb5701eabc7d", "relationships": {"teams": {"data": [{"id": "00000000-da3a-0000-0000-000000000000", "type": "teams"}]}}, "type": "schedules"}}
    When the request is sent
    Then the response status is 404 Not Found

  @skip-python @team:DataDog/bugle
  Scenario: Update on-call schedule returns "OK" response
    Given new "UpdateOnCallSchedule" request
    And there is a valid "schedule" in the system
    And there is a valid "team" in the system
    And there is a valid "user" in the system
    And request contains "schedule_id" parameter from "schedule.data.id"
    And body with value {"data": { "id": "{{ schedule.data.id }}", "attributes": {"layers": [{"id": "{{ schedule.data.relationships.layers.data[0].id }}" , "effective_date": "{{ timeISO('now - 10d') }}", "end_date": "{{ timeISO('now + 10d') }}", "interval": {"seconds": 300}, "members": [{"user": {"id": "{{user.data.id}}"}}], "name": "Layer 1", "restrictions": [{"end_day": "friday", "end_time": "17:00:00", "start_day": "monday", "start_time": "09:00:00"}], "rotation_start": "{{ timeISO('now - 5d') }}"}], "name": "{{ unique }}", "tags": ["tag1", "tag2", "tag3"], "time_zone": "America/New_York"}, "relationships": {"teams": {"data": [{"id": "{{team.data.id}}", "type": "teams"}]}}, "type": "schedules"}}
    When the request is sent
    Then the response status is 200 OK
