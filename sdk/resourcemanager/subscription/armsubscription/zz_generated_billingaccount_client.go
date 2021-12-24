//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsubscription

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

// BillingAccountClient contains the methods for the BillingAccount group.
// Don't use this type directly, use NewBillingAccountClient() instead.
type BillingAccountClient struct {
	ep string
	pl runtime.Pipeline
}

// NewBillingAccountClient creates a new instance of BillingAccountClient with the specified values.
func NewBillingAccountClient(credential azcore.TokenCredential, options *arm.ClientOptions) *BillingAccountClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &BillingAccountClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// GetPolicy - Get Billing Account Policy.
// If the operation fails it returns the *ErrorResponseBody error type.
func (client *BillingAccountClient) GetPolicy(ctx context.Context, billingAccountID string, options *BillingAccountGetPolicyOptions) (BillingAccountGetPolicyResponse, error) {
	req, err := client.getPolicyCreateRequest(ctx, billingAccountID, options)
	if err != nil {
		return BillingAccountGetPolicyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BillingAccountGetPolicyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BillingAccountGetPolicyResponse{}, client.getPolicyHandleError(resp)
	}
	return client.getPolicyHandleResponse(resp)
}

// getPolicyCreateRequest creates the GetPolicy request.
func (client *BillingAccountClient) getPolicyCreateRequest(ctx context.Context, billingAccountID string, options *BillingAccountGetPolicyOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.Subscription/policies/default"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getPolicyHandleResponse handles the GetPolicy response.
func (client *BillingAccountClient) getPolicyHandleResponse(resp *http.Response) (BillingAccountGetPolicyResponse, error) {
	result := BillingAccountGetPolicyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BillingAccountPoliciesResponse); err != nil {
		return BillingAccountGetPolicyResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getPolicyHandleError handles the GetPolicy error response.
func (client *BillingAccountClient) getPolicyHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponseBody{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}