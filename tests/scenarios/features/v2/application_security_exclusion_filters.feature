@endpoint(application-security-exclusion-filters) @endpoint(application-security-exclusion-filters-v2)
Feature: Application Security Exclusion Filters
  Exclusion filters in Application Security libraries are used to circumvent
  false positives in protection.

  Background:
    Given a valid "apiKeyAuth" key in the system
    And a valid "appKeyAuth" key in the system
    And an instance of "ApplicationSecurityExclusionFilters" API

  @generated @skip @team:DataDog/asm-respond-backend
  Scenario: Create an ASM Exclusion filter returns "Bad Request" response
    Given new "CreateApplicationSecurityExclusionFilter" request
    And body with value {"data": {"type": "exclusion_filter"}}
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/asm-respond-backend
  Scenario: Create an ASM Exclusion filter returns "Conflict" response
    Given new "CreateApplicationSecurityExclusionFilter" request
    And body with value {"data": {"type": "exclusion_filter"}}
    When the request is sent
    Then the response status is 409 Conflict

  @skip @team:DataDog/asm-respond-backend
  Scenario: Create an ASM Exclusion filter returns "OK" response
    Given new "CreateApplicationSecurityExclusionFilter" request
    And body with value {"data": {"type": "exclusion_filter","attributes":{"description":"my description","enabled":true,"path_glob":"*","rules_target":[{"tags":{"category":"attack_attempt","type":"sql_injection"}}],"scope":[{"env":"staging","service":"container-resolver"}]}}}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip @team:DataDog/asm-respond-backend
  Scenario: Delete an ASM Exclusion Filter returns "Not Found" response
    Given new "DeleteApplicationSecurityExclusionFilter" request
    And request contains "exclusion_filter_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 404 Not Found

  @generated @skip @team:DataDog/asm-respond-backend
  Scenario: Delete an ASM Exclusion Filter returns "OK" response
    Given new "DeleteApplicationSecurityExclusionFilter" request
    And request contains "exclusion_filter_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 204 OK

  @generated @skip @team:DataDog/asm-respond-backend
  Scenario: Get a specific ASM Exclusion Filter returns "OK" response
    Given new "GetApplicationSecurityExclusionFilters" request
    And request contains "exclusion_filter_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip @team:DataDog/asm-respond-backend
  Scenario: List all ASM Exclusion Filters returns "OK" response
    Given new "ListApplicationSecurityExclusionFilters" request
    When the request is sent
    Then the response status is 200 OK

  @generated @skip @team:DataDog/asm-respond-backend
  Scenario: Update an ASM Exclusion filter returns "Bad Request" response
    Given new "UpdateApplicationSecurityExclusionFilter" request
    And request contains "exclusion_filter_id" parameter from "REPLACE.ME"
    And body with value {"data": {"attributes": {"description": "My Exclusion filter", "enabled": true, "ip_list": ["127.0.0.1"], "parameters": ["list.search.query"], "path_glob": "/lfi_include/*", "rules_target": [{"rule_id": "dog-913-009"}], "scope": [{"env": "dd-appsec-php-support", "service": "anil-php-weblog"}]}, "id": "3dd-0uc-h1s", "type": "exclusion_filter"}}
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/asm-respond-backend
  Scenario: Update an ASM Exclusion filter returns "Concurrent Modification" response
    Given new "UpdateApplicationSecurityExclusionFilter" request
    And request contains "exclusion_filter_id" parameter from "REPLACE.ME"
    And body with value {"data": {"attributes": {"description": "My Exclusion filter", "enabled": true, "ip_list": ["127.0.0.1"], "parameters": ["list.search.query"], "path_glob": "/lfi_include/*", "rules_target": [{"rule_id": "dog-913-009"}], "scope": [{"env": "dd-appsec-php-support", "service": "anil-php-weblog"}]}, "id": "3dd-0uc-h1s", "type": "exclusion_filter"}}
    When the request is sent
    Then the response status is 409 Concurrent Modification

  @generated @skip @team:DataDog/asm-respond-backend
  Scenario: Update an ASM Exclusion filter returns "Not Found" response
    Given new "UpdateApplicationSecurityExclusionFilter" request
    And request contains "exclusion_filter_id" parameter from "REPLACE.ME"
    And body with value {"data": {"attributes": {"description": "My Exclusion filter", "enabled": true, "ip_list": ["127.0.0.1"], "parameters": ["list.search.query"], "path_glob": "/lfi_include/*", "rules_target": [{"rule_id": "dog-913-009"}], "scope": [{"env": "dd-appsec-php-support", "service": "anil-php-weblog"}]}, "id": "3dd-0uc-h1s", "type": "exclusion_filter"}}
    When the request is sent
    Then the response status is 404 Not Found

  @generated @skip @team:DataDog/asm-respond-backend
  Scenario: Update an ASM Exclusion filter returns "OK" response
    Given new "UpdateApplicationSecurityExclusionFilter" request
    And request contains "exclusion_filter_id" parameter from "REPLACE.ME"
    And body with value {"data": {"attributes": {"description": "My Exclusion filter", "enabled": true, "ip_list": ["127.0.0.1"], "parameters": ["list.search.query"], "path_glob": "/lfi_include/*", "rules_target": [{"rule_id": "dog-913-009"}], "scope": [{"env": "dd-appsec-php-support", "service": "anil-php-weblog"}]}, "id": "3dd-0uc-h1s", "type": "exclusion_filter"}}
    When the request is sent
    Then the response status is 200 OK
