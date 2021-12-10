//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armm365securityandcompliance

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PrivateEndpointConnectionsSecClient contains the methods for the PrivateEndpointConnectionsSec group.
// Don't use this type directly, use NewPrivateEndpointConnectionsSecClient() instead.
type PrivateEndpointConnectionsSecClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPrivateEndpointConnectionsSecClient creates a new instance of PrivateEndpointConnectionsSecClient with the specified values.
func NewPrivateEndpointConnectionsSecClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrivateEndpointConnectionsSecClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &PrivateEndpointConnectionsSecClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Update the state of the specified private endpoint connection associated with the service.
// If the operation fails it returns the *ErrorDetails error type.
func (client *PrivateEndpointConnectionsSecClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, properties PrivateEndpointConnection, options *PrivateEndpointConnectionsSecBeginCreateOrUpdateOptions) (PrivateEndpointConnectionsSecCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, privateEndpointConnectionName, properties, options)
	if err != nil {
		return PrivateEndpointConnectionsSecCreateOrUpdatePollerResponse{}, err
	}
	result := PrivateEndpointConnectionsSecCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PrivateEndpointConnectionsSecClient.CreateOrUpdate", "location", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return PrivateEndpointConnectionsSecCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &PrivateEndpointConnectionsSecCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Update the state of the specified private endpoint connection associated with the service.
// If the operation fails it returns the *ErrorDetails error type.
func (client *PrivateEndpointConnectionsSecClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, properties PrivateEndpointConnection, options *PrivateEndpointConnectionsSecBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, privateEndpointConnectionName, properties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PrivateEndpointConnectionsSecClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, properties PrivateEndpointConnection, options *PrivateEndpointConnectionsSecBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/{resourceName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, properties)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *PrivateEndpointConnectionsSecClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes a private endpoint connection.
// If the operation fails it returns the *ErrorDetails error type.
func (client *PrivateEndpointConnectionsSecClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsSecBeginDeleteOptions) (PrivateEndpointConnectionsSecDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionsSecDeletePollerResponse{}, err
	}
	result := PrivateEndpointConnectionsSecDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PrivateEndpointConnectionsSecClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return PrivateEndpointConnectionsSecDeletePollerResponse{}, err
	}
	result.Poller = &PrivateEndpointConnectionsSecDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a private endpoint connection.
// If the operation fails it returns the *ErrorDetails error type.
func (client *PrivateEndpointConnectionsSecClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsSecBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, privateEndpointConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateEndpointConnectionsSecClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsSecBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/{resourceName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *PrivateEndpointConnectionsSecClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the specified private endpoint connection associated with the service.
// If the operation fails it returns the *ErrorDetails error type.
func (client *PrivateEndpointConnectionsSecClient) Get(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsSecGetOptions) (PrivateEndpointConnectionsSecGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionsSecGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionsSecGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateEndpointConnectionsSecGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateEndpointConnectionsSecClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsSecGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/{resourceName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateEndpointConnectionsSecClient) getHandleResponse(resp *http.Response) (PrivateEndpointConnectionsSecGetResponse, error) {
	result := PrivateEndpointConnectionsSecGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionsSecGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *PrivateEndpointConnectionsSecClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByService - Lists all private endpoint connections for a service.
// If the operation fails it returns the *ErrorDetails error type.
func (client *PrivateEndpointConnectionsSecClient) ListByService(resourceGroupName string, resourceName string, options *PrivateEndpointConnectionsSecListByServiceOptions) *PrivateEndpointConnectionsSecListByServicePager {
	return &PrivateEndpointConnectionsSecListByServicePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServiceCreateRequest(ctx, resourceGroupName, resourceName, options)
		},
		advancer: func(ctx context.Context, resp PrivateEndpointConnectionsSecListByServiceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PrivateEndpointConnectionListResult.NextLink)
		},
	}
}

// listByServiceCreateRequest creates the ListByService request.
func (client *PrivateEndpointConnectionsSecClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateEndpointConnectionsSecListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/{resourceName}/privateEndpointConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *PrivateEndpointConnectionsSecClient) listByServiceHandleResponse(resp *http.Response) (PrivateEndpointConnectionsSecListByServiceResponse, error) {
	result := PrivateEndpointConnectionsSecListByServiceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionListResult); err != nil {
		return PrivateEndpointConnectionsSecListByServiceResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByServiceHandleError handles the ListByService error response.
func (client *PrivateEndpointConnectionsSecClient) listByServiceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}