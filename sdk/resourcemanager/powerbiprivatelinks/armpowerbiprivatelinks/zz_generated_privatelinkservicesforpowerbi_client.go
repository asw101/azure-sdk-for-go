//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerbiprivatelinks

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

// PrivateLinkServicesForPowerBIClient contains the methods for the PrivateLinkServicesForPowerBI group.
// Don't use this type directly, use NewPrivateLinkServicesForPowerBIClient() instead.
type PrivateLinkServicesForPowerBIClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPrivateLinkServicesForPowerBIClient creates a new instance of PrivateLinkServicesForPowerBIClient with the specified values.
func NewPrivateLinkServicesForPowerBIClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrivateLinkServicesForPowerBIClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &PrivateLinkServicesForPowerBIClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// ListBySubscriptionID - Gets all the private link resources for the given subscription id.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PrivateLinkServicesForPowerBIClient) ListBySubscriptionID(ctx context.Context, options *PrivateLinkServicesForPowerBIListBySubscriptionIDOptions) (PrivateLinkServicesForPowerBIListBySubscriptionIDResponse, error) {
	req, err := client.listBySubscriptionIDCreateRequest(ctx, options)
	if err != nil {
		return PrivateLinkServicesForPowerBIListBySubscriptionIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkServicesForPowerBIListBySubscriptionIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkServicesForPowerBIListBySubscriptionIDResponse{}, client.listBySubscriptionIDHandleError(resp)
	}
	return client.listBySubscriptionIDHandleResponse(resp)
}

// listBySubscriptionIDCreateRequest creates the ListBySubscriptionID request.
func (client *PrivateLinkServicesForPowerBIClient) listBySubscriptionIDCreateRequest(ctx context.Context, options *PrivateLinkServicesForPowerBIListBySubscriptionIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.PowerBI/privateLinkServicesForPowerBI"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionIDHandleResponse handles the ListBySubscriptionID response.
func (client *PrivateLinkServicesForPowerBIClient) listBySubscriptionIDHandleResponse(resp *http.Response) (PrivateLinkServicesForPowerBIListBySubscriptionIDResponse, error) {
	result := PrivateLinkServicesForPowerBIListBySubscriptionIDResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TenantResourceArray); err != nil {
		return PrivateLinkServicesForPowerBIListBySubscriptionIDResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listBySubscriptionIDHandleError handles the ListBySubscriptionID error response.
func (client *PrivateLinkServicesForPowerBIClient) listBySubscriptionIDHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}