// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// ProximityPlacementGroupsOperations contains the methods for the ProximityPlacementGroups group.
type ProximityPlacementGroupsOperations interface {
	// CreateOrUpdate - Create or update a proximity placement group.
	CreateOrUpdate(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters ProximityPlacementGroup) (*ProximityPlacementGroupResponse, error)
	// Delete - Delete a proximity placement group.
	Delete(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string) (*http.Response, error)
	// Get - Retrieves information about a proximity placement group .
	Get(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, proximityPlacementGroupsGetOptions *ProximityPlacementGroupsGetOptions) (*ProximityPlacementGroupResponse, error)
	// ListByResourceGroup - Lists all proximity placement groups in a resource group.
	ListByResourceGroup(resourceGroupName string) (ProximityPlacementGroupListResultPager, error)
	// ListBySubscription - Lists all proximity placement groups in a subscription.
	ListBySubscription() (ProximityPlacementGroupListResultPager, error)
	// Update - Update a proximity placement group.
	Update(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters UpdateResource) (*ProximityPlacementGroupResponse, error)
}

// proximityPlacementGroupsOperations implements the ProximityPlacementGroupsOperations interface.
type proximityPlacementGroupsOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Create or update a proximity placement group.
func (client *proximityPlacementGroupsOperations) CreateOrUpdate(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters ProximityPlacementGroup) (*ProximityPlacementGroupResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, proximityPlacementGroupName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *proximityPlacementGroupsOperations) createOrUpdateCreateRequest(resourceGroupName string, proximityPlacementGroupName string, parameters ProximityPlacementGroup) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{proximityPlacementGroupName}", url.PathEscape(proximityPlacementGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *proximityPlacementGroupsOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*ProximityPlacementGroupResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	result := ProximityPlacementGroupResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ProximityPlacementGroup)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *proximityPlacementGroupsOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}

// Delete - Delete a proximity placement group.
func (client *proximityPlacementGroupsOperations) Delete(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string) (*http.Response, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, proximityPlacementGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.deleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// deleteCreateRequest creates the Delete request.
func (client *proximityPlacementGroupsOperations) deleteCreateRequest(resourceGroupName string, proximityPlacementGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{proximityPlacementGroupName}", url.PathEscape(proximityPlacementGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *proximityPlacementGroupsOperations) deleteHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.deleteHandleError(resp)
	}
	return resp.Response, nil
}

// deleteHandleError handles the Delete error response.
func (client *proximityPlacementGroupsOperations) deleteHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}

// Get - Retrieves information about a proximity placement group .
func (client *proximityPlacementGroupsOperations) Get(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, proximityPlacementGroupsGetOptions *ProximityPlacementGroupsGetOptions) (*ProximityPlacementGroupResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, proximityPlacementGroupName, proximityPlacementGroupsGetOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client *proximityPlacementGroupsOperations) getCreateRequest(resourceGroupName string, proximityPlacementGroupName string, proximityPlacementGroupsGetOptions *ProximityPlacementGroupsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{proximityPlacementGroupName}", url.PathEscape(proximityPlacementGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	if proximityPlacementGroupsGetOptions != nil && proximityPlacementGroupsGetOptions.IncludeColocationStatus != nil {
		query.Set("includeColocationStatus", *proximityPlacementGroupsGetOptions.IncludeColocationStatus)
	}
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *proximityPlacementGroupsOperations) getHandleResponse(resp *azcore.Response) (*ProximityPlacementGroupResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := ProximityPlacementGroupResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ProximityPlacementGroup)
}

// getHandleError handles the Get error response.
func (client *proximityPlacementGroupsOperations) getHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}

// ListByResourceGroup - Lists all proximity placement groups in a resource group.
func (client *proximityPlacementGroupsOperations) ListByResourceGroup(resourceGroupName string) (ProximityPlacementGroupListResultPager, error) {
	req, err := client.listByResourceGroupCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &proximityPlacementGroupListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listByResourceGroupHandleResponse,
		advancer: func(resp *ProximityPlacementGroupListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.ProximityPlacementGroupListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.ProximityPlacementGroupListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *proximityPlacementGroupsOperations) listByResourceGroupCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *proximityPlacementGroupsOperations) listByResourceGroupHandleResponse(resp *azcore.Response) (*ProximityPlacementGroupListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listByResourceGroupHandleError(resp)
	}
	result := ProximityPlacementGroupListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ProximityPlacementGroupListResult)
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *proximityPlacementGroupsOperations) listByResourceGroupHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}

// ListBySubscription - Lists all proximity placement groups in a subscription.
func (client *proximityPlacementGroupsOperations) ListBySubscription() (ProximityPlacementGroupListResultPager, error) {
	req, err := client.listBySubscriptionCreateRequest()
	if err != nil {
		return nil, err
	}
	return &proximityPlacementGroupListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listBySubscriptionHandleResponse,
		advancer: func(resp *ProximityPlacementGroupListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.ProximityPlacementGroupListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.ProximityPlacementGroupListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *proximityPlacementGroupsOperations) listBySubscriptionCreateRequest() (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/proximityPlacementGroups"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *proximityPlacementGroupsOperations) listBySubscriptionHandleResponse(resp *azcore.Response) (*ProximityPlacementGroupListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listBySubscriptionHandleError(resp)
	}
	result := ProximityPlacementGroupListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ProximityPlacementGroupListResult)
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *proximityPlacementGroupsOperations) listBySubscriptionHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}

// Update - Update a proximity placement group.
func (client *proximityPlacementGroupsOperations) Update(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters UpdateResource) (*ProximityPlacementGroupResponse, error) {
	req, err := client.updateCreateRequest(resourceGroupName, proximityPlacementGroupName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.updateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// updateCreateRequest creates the Update request.
func (client *proximityPlacementGroupsOperations) updateCreateRequest(resourceGroupName string, proximityPlacementGroupName string, parameters UpdateResource) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{proximityPlacementGroupName}", url.PathEscape(proximityPlacementGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(parameters)
}

// updateHandleResponse handles the Update response.
func (client *proximityPlacementGroupsOperations) updateHandleResponse(resp *azcore.Response) (*ProximityPlacementGroupResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateHandleError(resp)
	}
	result := ProximityPlacementGroupResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ProximityPlacementGroup)
}

// updateHandleError handles the Update error response.
func (client *proximityPlacementGroupsOperations) updateHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}