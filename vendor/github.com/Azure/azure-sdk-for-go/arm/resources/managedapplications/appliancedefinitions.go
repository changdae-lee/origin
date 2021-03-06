package managedapplications

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ApplianceDefinitionsClient is the ARM managed applications (appliances)
type ApplianceDefinitionsClient struct {
	ManagementClient
}

// NewApplianceDefinitionsClient creates an instance of the ApplianceDefinitionsClient client.
func NewApplianceDefinitionsClient(subscriptionID string) ApplianceDefinitionsClient {
	return NewApplianceDefinitionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplianceDefinitionsClientWithBaseURI creates an instance of the ApplianceDefinitionsClient client.
func NewApplianceDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) ApplianceDefinitionsClient {
	return ApplianceDefinitionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a new appliance definition. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. The name is case insensitive. applianceDefinitionName is the
// name of the appliance definition. parameters is parameters supplied to the create or update an appliance definition.
func (client ApplianceDefinitionsClient) CreateOrUpdate(resourceGroupName string, applianceDefinitionName string, parameters ApplianceDefinition, cancel <-chan struct{}) (<-chan ApplianceDefinition, <-chan error) {
	resultChan := make(chan ApplianceDefinition, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: applianceDefinitionName,
			Constraints: []validation.Constraint{{Target: "applianceDefinitionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applianceDefinitionName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ApplianceDefinitionProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.ApplianceDefinitionProperties.Authorizations", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.ApplianceDefinitionProperties.PackageFileURI", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "managedapplications.ApplianceDefinitionsClient", "CreateOrUpdate")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result ApplianceDefinition
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, applianceDefinitionName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ApplianceDefinitionsClient) CreateOrUpdatePreparer(resourceGroupName string, applianceDefinitionName string, parameters ApplianceDefinition, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applianceDefinitionName": autorest.Encode("path", applianceDefinitionName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applianceDefinitions/{applianceDefinitionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ApplianceDefinitionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ApplianceDefinitionsClient) CreateOrUpdateResponder(resp *http.Response) (result ApplianceDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateByID creates a new appliance definition. This method may poll for completion. Polling can be canceled
// by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// applianceDefinitionID is the fully qualified ID of the appliance definition, including the appliance name and the
// appliance definition resource type. Use the format,
// /subscriptions/{guid}/resourceGroups/{resource-group-name}/Microsoft.Solutions/applianceDefinitions/{applianceDefinition-name}
// parameters is parameters supplied to the create or update an appliance definition.
func (client ApplianceDefinitionsClient) CreateOrUpdateByID(applianceDefinitionID string, parameters ApplianceDefinition, cancel <-chan struct{}) (<-chan ApplianceDefinition, <-chan error) {
	resultChan := make(chan ApplianceDefinition, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ApplianceDefinitionProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.ApplianceDefinitionProperties.Authorizations", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.ApplianceDefinitionProperties.PackageFileURI", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "managedapplications.ApplianceDefinitionsClient", "CreateOrUpdateByID")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result ApplianceDefinition
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdateByIDPreparer(applianceDefinitionID, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "CreateOrUpdateByID", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateByIDSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "CreateOrUpdateByID", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateByIDResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "CreateOrUpdateByID", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdateByIDPreparer prepares the CreateOrUpdateByID request.
func (client ApplianceDefinitionsClient) CreateOrUpdateByIDPreparer(applianceDefinitionID string, parameters ApplianceDefinition, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applianceDefinitionId": applianceDefinitionID,
	}

	const APIVersion = "2016-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{applianceDefinitionId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateByIDSender sends the CreateOrUpdateByID request. The method will close the
// http.Response Body if it receives an error.
func (client ApplianceDefinitionsClient) CreateOrUpdateByIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateByIDResponder handles the response to the CreateOrUpdateByID request. The method always
// closes the http.Response Body.
func (client ApplianceDefinitionsClient) CreateOrUpdateByIDResponder(resp *http.Response) (result ApplianceDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the appliance definition. This method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. The name is case insensitive. applianceDefinitionName is the
// name of the appliance definition to delete.
func (client ApplianceDefinitionsClient) Delete(resourceGroupName string, applianceDefinitionName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: applianceDefinitionName,
			Constraints: []validation.Constraint{{Target: "applianceDefinitionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applianceDefinitionName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "managedapplications.ApplianceDefinitionsClient", "Delete")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeletePreparer(resourceGroupName, applianceDefinitionName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client ApplianceDefinitionsClient) DeletePreparer(resourceGroupName string, applianceDefinitionName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applianceDefinitionName": autorest.Encode("path", applianceDefinitionName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applianceDefinitions/{applianceDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplianceDefinitionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ApplianceDefinitionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteByID deletes the appliance definition. This method may poll for completion. Polling can be canceled by passing
// the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// applianceDefinitionID is the fully qualified ID of the appliance definition, including the appliance name and the
// appliance definition resource type. Use the format,
// /subscriptions/{guid}/resourceGroups/{resource-group-name}/Microsoft.Solutions/applianceDefinitions/{applianceDefinition-name}
func (client ApplianceDefinitionsClient) DeleteByID(applianceDefinitionID string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeleteByIDPreparer(applianceDefinitionID, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "DeleteByID", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteByIDSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "DeleteByID", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteByIDResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "DeleteByID", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeleteByIDPreparer prepares the DeleteByID request.
func (client ApplianceDefinitionsClient) DeleteByIDPreparer(applianceDefinitionID string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applianceDefinitionId": applianceDefinitionID,
	}

	const APIVersion = "2016-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{applianceDefinitionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteByIDSender sends the DeleteByID request. The method will close the
// http.Response Body if it receives an error.
func (client ApplianceDefinitionsClient) DeleteByIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteByIDResponder handles the response to the DeleteByID request. The method always
// closes the http.Response Body.
func (client ApplianceDefinitionsClient) DeleteByIDResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the appliance definition.
//
// resourceGroupName is the name of the resource group. The name is case insensitive. applianceDefinitionName is the
// name of the appliance definition.
func (client ApplianceDefinitionsClient) Get(resourceGroupName string, applianceDefinitionName string) (result ApplianceDefinition, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: applianceDefinitionName,
			Constraints: []validation.Constraint{{Target: "applianceDefinitionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applianceDefinitionName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "managedapplications.ApplianceDefinitionsClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, applianceDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplianceDefinitionsClient) GetPreparer(resourceGroupName string, applianceDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applianceDefinitionName": autorest.Encode("path", applianceDefinitionName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applianceDefinitions/{applianceDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplianceDefinitionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplianceDefinitionsClient) GetResponder(resp *http.Response) (result ApplianceDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByID gets the appliance definition.
//
// applianceDefinitionID is the fully qualified ID of the appliance definition, including the appliance name and the
// appliance definition resource type. Use the format,
// /subscriptions/{guid}/resourceGroups/{resource-group-name}/Microsoft.Solutions/applianceDefinitions/{applianceDefinition-name}
func (client ApplianceDefinitionsClient) GetByID(applianceDefinitionID string) (result ApplianceDefinition, err error) {
	req, err := client.GetByIDPreparer(applianceDefinitionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "GetByID", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "GetByID", resp, "Failure sending request")
		return
	}

	result, err = client.GetByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "GetByID", resp, "Failure responding to request")
	}

	return
}

// GetByIDPreparer prepares the GetByID request.
func (client ApplianceDefinitionsClient) GetByIDPreparer(applianceDefinitionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applianceDefinitionId": applianceDefinitionID,
	}

	const APIVersion = "2016-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{applianceDefinitionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetByIDSender sends the GetByID request. The method will close the
// http.Response Body if it receives an error.
func (client ApplianceDefinitionsClient) GetByIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByIDResponder handles the response to the GetByID request. The method always
// closes the http.Response Body.
func (client ApplianceDefinitionsClient) GetByIDResponder(resp *http.Response) (result ApplianceDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup lists the appliance definitions in a resource group.
//
// resourceGroupName is the name of the resource group. The name is case insensitive.
func (client ApplianceDefinitionsClient) ListByResourceGroup(resourceGroupName string) (result ApplianceDefinitionListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "managedapplications.ApplianceDefinitionsClient", "ListByResourceGroup")
	}

	req, err := client.ListByResourceGroupPreparer(resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ApplianceDefinitionsClient) ListByResourceGroupPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applianceDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ApplianceDefinitionsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ApplianceDefinitionsClient) ListByResourceGroupResponder(resp *http.Response) (result ApplianceDefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroupNextResults retrieves the next set of results, if any.
func (client ApplianceDefinitionsClient) ListByResourceGroupNextResults(lastResults ApplianceDefinitionListResult) (result ApplianceDefinitionListResult, err error) {
	req, err := lastResults.ApplianceDefinitionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "ListByResourceGroup", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "ListByResourceGroup", resp, "Failure sending next results request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "ListByResourceGroup", resp, "Failure responding to next results request")
	}

	return
}

// ListByResourceGroupComplete gets all elements from the list without paging.
func (client ApplianceDefinitionsClient) ListByResourceGroupComplete(resourceGroupName string, cancel <-chan struct{}) (<-chan ApplianceDefinition, <-chan error) {
	resultChan := make(chan ApplianceDefinition)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByResourceGroup(resourceGroupName)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListByResourceGroupNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}
