package resources

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.11.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// Resources Client
type ResourcesClient struct {
	ResourceManagementClient
}

func NewResourcesClient(subscriptionId string) ResourcesClient {
	return NewResourcesClientWithBaseUri(DefaultBaseUri, subscriptionId)
}

func NewResourcesClientWithBaseUri(baseUri string, subscriptionId string) ResourcesClient {
	return ResourcesClient{NewWithBaseUri(baseUri, subscriptionId)}
}

// MoveResources move resources within or across subscriptions.
//
// sourceResourceGroupName is source resource group name. parameters is move
// resources' parameters.
func (client ResourcesClient) MoveResources(sourceResourceGroupName string, parameters ResourcesMoveInfo) (result autorest.Response, ae autorest.Error) {
	req, err := client.NewMoveResourcesRequest(sourceResourceGroupName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ResourcesClient", "MoveResources", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ResourcesClient", "MoveResources", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(client, req)
	if err == nil {
		err = client.ShouldPoll(resp)
		if err == nil {
			resp, err = client.PollAsNeeded(resp)
		}
	}

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK())
		if err != nil {
			ae = autorest.NewErrorWithError(err, "resources.ResourcesClient", "MoveResources", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "resources.ResourcesClient", "MoveResources", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = resp

	return
}

// Create the MoveResources request.
func (client ResourcesClient) NewMoveResourcesRequest(sourceResourceGroupName string, parameters ResourcesMoveInfo) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"sourceResourceGroupName": url.QueryEscape(sourceResourceGroupName),
		"subscriptionId":          url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.MoveResourcesRequestPreparer(),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the MoveResources request.
func (client ResourcesClient) MoveResourcesRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{sourceResourceGroupName}/moveResources"))
}

// List get all of the resources under a subscription.
//
// filter is the filter to apply on the operation. top is query parameters. If
// null is passed returns all resource groups.
func (client ResourcesClient) List(filter string, top int) (result ResourceListResult, ae autorest.Error) {
	req, err := client.NewListRequest(filter, top)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ResourcesClient", "List", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ResourcesClient", "List", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(client, req)

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "resources.ResourcesClient", "List", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "resources.ResourcesClient", "List", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the List request.
func (client ResourcesClient) NewListRequest(filter string, top int) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"$filter":     filter,
		"$top":        top,
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.ListRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the List request.
func (client ResourcesClient) ListRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resources"))
}

// CheckExistence checks whether resource exists.
//
// resourceGroupName is the name of the resource group. The name is case
// insensitive. resourceProviderNamespace is resource identity.
// parentResourcePath is resource identity. resourceType is resource
// identity. resourceName is resource identity.
func (client ResourcesClient) CheckExistence(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, apiVersion string) (result autorest.Response, ae autorest.Error) {
	req, err := client.NewCheckExistenceRequest(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, apiVersion)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ResourcesClient", "CheckExistence", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ResourcesClient", "CheckExistence", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(client, req)

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK())
		if err != nil {
			ae = autorest.NewErrorWithError(err, "resources.ResourcesClient", "CheckExistence", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "resources.ResourcesClient", "CheckExistence", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = resp

	return
}

// Create the CheckExistence request.
func (client ResourcesClient) NewCheckExistenceRequest(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, apiVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         url.QueryEscape(resourceGroupName),
		"resourceName":              url.QueryEscape(resourceName),
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.CheckExistenceRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the CheckExistence request.
func (client ResourcesClient) CheckExistenceRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}"))
}

// Delete delete resource and all of its resources.
//
// resourceGroupName is the name of the resource group. The name is case
// insensitive. resourceProviderNamespace is resource identity.
// parentResourcePath is resource identity. resourceType is resource
// identity. resourceName is resource identity.
func (client ResourcesClient) Delete(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, apiVersion string) (result autorest.Response, ae autorest.Error) {
	req, err := client.NewDeleteRequest(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, apiVersion)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ResourcesClient", "Delete", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ResourcesClient", "Delete", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(client, req)
	if err == nil {
		err = client.ShouldPoll(resp)
		if err == nil {
			resp, err = client.PollAsNeeded(resp)
		}
	}

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK())
		if err != nil {
			ae = autorest.NewErrorWithError(err, "resources.ResourcesClient", "Delete", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "resources.ResourcesClient", "Delete", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = resp

	return
}

// Create the Delete request.
func (client ResourcesClient) NewDeleteRequest(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, apiVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         url.QueryEscape(resourceGroupName),
		"resourceName":              url.QueryEscape(resourceName),
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.DeleteRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the Delete request.
func (client ResourcesClient) DeleteRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}"))
}

// CreateOrUpdate create a resource.
//
// resourceGroupName is the name of the resource group. The name is case
// insensitive. resourceProviderNamespace is resource identity.
// parentResourcePath is resource identity. resourceType is resource
// identity. resourceName is resource identity. parameters is create or
// update resource parameters.
func (client ResourcesClient) CreateOrUpdate(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, apiVersion string, parameters GenericResource) (result GenericResource, ae autorest.Error) {
	req, err := client.NewCreateOrUpdateRequest(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, apiVersion, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ResourcesClient", "CreateOrUpdate", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ResourcesClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(client, req)
	if err == nil {
		err = client.ShouldPoll(resp)
		if err == nil {
			resp, err = client.PollAsNeeded(resp)
		}
	}

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "resources.ResourcesClient", "CreateOrUpdate", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "resources.ResourcesClient", "CreateOrUpdate", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the CreateOrUpdate request.
func (client ResourcesClient) NewCreateOrUpdateRequest(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, apiVersion string, parameters GenericResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         url.QueryEscape(resourceGroupName),
		"resourceName":              url.QueryEscape(resourceName),
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.CreateOrUpdateRequestPreparer(),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the CreateOrUpdate request.
func (client ResourcesClient) CreateOrUpdateRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}"))
}

// Get returns a resource belonging to a resource group.
//
// resourceGroupName is the name of the resource group. The name is case
// insensitive. resourceProviderNamespace is resource identity.
// parentResourcePath is resource identity. resourceType is resource
// identity. resourceName is resource identity.
func (client ResourcesClient) Get(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, apiVersion string) (result GenericResource, ae autorest.Error) {
	req, err := client.NewGetRequest(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, apiVersion)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ResourcesClient", "Get", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ResourcesClient", "Get", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(client, req)

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "resources.ResourcesClient", "Get", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "resources.ResourcesClient", "Get", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the Get request.
func (client ResourcesClient) NewGetRequest(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, apiVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         url.QueryEscape(resourceGroupName),
		"resourceName":              url.QueryEscape(resourceName),
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.GetRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the Get request.
func (client ResourcesClient) GetRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}"))
}
