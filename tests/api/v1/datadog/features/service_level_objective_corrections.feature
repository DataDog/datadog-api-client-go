@endpoint(service-level-objective-corrections)
Feature: Service Level Objective Corrections
  SLO Status Corrections allow you to prevent specific time periods from
  negatively impacting your SLO’s status and error budget. You can use
  Status Corrections for various purposes, such as removing planned
  maintenance windows, non-business hours, or other time periods that do not
  correspond to genuine issues.

  Background:
    Given a valid "apiKeyAuth" key in the system
    And a valid "appKeyAuth" key in the system
    And an instance of "ServiceLevelObjectiveCorrections" API

  @generated @skip
  Scenario: Create an SLO correction returns "Bad Request" response
    Given operation "CreateSLOCorrection" enabled
    And new "CreateSLOCorrection" request
    And body {}
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip
  Scenario: Create an SLO correction returns "OK" response
    Given operation "CreateSLOCorrection" enabled
    And new "CreateSLOCorrection" request
    And body {}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Delete an SLO Correction returns "Not found" response
    Given operation "DeleteSLOCorrection" enabled
    And new "DeleteSLOCorrection" request
    And request contains "slo_correction_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 404 Not found

  @generated @skip
  Scenario: Delete an SLO Correction returns "OK" response
    Given operation "DeleteSLOCorrection" enabled
    And new "DeleteSLOCorrection" request
    And request contains "slo_correction_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 204 OK

  @generated @skip
  Scenario: Get all SLO corrections returns "OK" response
    Given operation "ListSLOCorrection" enabled
    And new "ListSLOCorrection" request
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Get an SLO correction for an SLO returns "Bad Request" response
    Given operation "GetSLOCorrection" enabled
    And new "GetSLOCorrection" request
    And request contains "slo_correction_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 400 Bad Request

  Scenario: Get an SLO correction for an SLO returns "OK" response
    Given operation "GetSLOCorrection" enabled
    And there is a valid "slo" in the system
    And there is a valid "correction" for "slo"
    And new "GetSLOCorrection" request
    And request contains "slo_correction_id" parameter from "correction.data.id"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Update an SLO Correction returns "Bad Request" response
    Given operation "UpdateSLOCorrection" enabled
    And new "UpdateSLOCorrection" request
    And request contains "slo_correction_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip
  Scenario: Update an SLO Correction returns "Not Found" response
    Given operation "UpdateSLOCorrection" enabled
    And new "UpdateSLOCorrection" request
    And request contains "slo_correction_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 404 Not Found

  @generated @skip
  Scenario: Update an SLO Correction returns "OK" response
    Given operation "UpdateSLOCorrection" enabled
    And new "UpdateSLOCorrection" request
    And request contains "slo_correction_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 200 OK
