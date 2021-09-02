// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// MetricAlertsStatusClient contains the methods for the MetricAlertsStatus group.
// Don't use this type directly, use NewMetricAlertsStatusClient() instead.
type MetricAlertsStatusClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewMetricAlertsStatusClient creates a new instance of MetricAlertsStatusClient with the specified values.
func NewMetricAlertsStatusClient(con *armcore.Connection, subscriptionID string) *MetricAlertsStatusClient {
	return &MetricAlertsStatusClient{con: con, subscriptionID: subscriptionID}
}

// List - Retrieve an alert rule status.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MetricAlertsStatusClient) List(ctx context.Context, resourceGroupName string, ruleName string, options *MetricAlertsStatusListOptions) (MetricAlertStatusCollectionResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, ruleName, options)
	if err != nil {
		return MetricAlertStatusCollectionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return MetricAlertStatusCollectionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return MetricAlertStatusCollectionResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *MetricAlertsStatusClient) listCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, options *MetricAlertsStatusListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/metricAlerts/{ruleName}/status"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MetricAlertsStatusClient) listHandleResponse(resp *azcore.Response) (MetricAlertStatusCollectionResponse, error) {
	var val *MetricAlertStatusCollection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return MetricAlertStatusCollectionResponse{}, err
	}
	return MetricAlertStatusCollectionResponse{RawResponse: resp.Response, MetricAlertStatusCollection: val}, nil
}

// listHandleError handles the List error response.
func (client *MetricAlertsStatusClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListByName - Retrieve an alert rule status.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MetricAlertsStatusClient) ListByName(ctx context.Context, resourceGroupName string, ruleName string, statusName string, options *MetricAlertsStatusListByNameOptions) (MetricAlertStatusCollectionResponse, error) {
	req, err := client.listByNameCreateRequest(ctx, resourceGroupName, ruleName, statusName, options)
	if err != nil {
		return MetricAlertStatusCollectionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return MetricAlertStatusCollectionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return MetricAlertStatusCollectionResponse{}, client.listByNameHandleError(resp)
	}
	return client.listByNameHandleResponse(resp)
}

// listByNameCreateRequest creates the ListByName request.
func (client *MetricAlertsStatusClient) listByNameCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, statusName string, options *MetricAlertsStatusListByNameOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/metricAlerts/{ruleName}/status/{statusName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	if statusName == "" {
		return nil, errors.New("parameter statusName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{statusName}", url.PathEscape(statusName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByNameHandleResponse handles the ListByName response.
func (client *MetricAlertsStatusClient) listByNameHandleResponse(resp *azcore.Response) (MetricAlertStatusCollectionResponse, error) {
	var val *MetricAlertStatusCollection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return MetricAlertStatusCollectionResponse{}, err
	}
	return MetricAlertStatusCollectionResponse{RawResponse: resp.Response, MetricAlertStatusCollection: val}, nil
}

// listByNameHandleError handles the ListByName error response.
func (client *MetricAlertsStatusClient) listByNameHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}