package aadapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/aad/mgmt/2020-03-01/aad"
	"github.com/Azure/go-autorest/autorest"
)

// PrivateLinkForAzureAdClientAPI contains the set of methods on the PrivateLinkForAzureAdClient type.
type PrivateLinkForAzureAdClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, policyName string, privateLinkPolicy aad.PrivateLinkPolicy) (result aad.PrivateLinkForAzureAdCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, policyName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, policyName string) (result aad.PrivateLinkPolicy, err error)
	List(ctx context.Context, resourceGroupName string) (result aad.PrivateLinkPolicyListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result aad.PrivateLinkPolicyListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result aad.PrivateLinkPolicyListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result aad.PrivateLinkPolicyListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, policyName string, privateLinkPolicy *aad.PrivateLinkPolicyUpdateParameter) (result aad.PrivateLinkPolicy, err error)
}

var _ PrivateLinkForAzureAdClientAPI = (*aad.PrivateLinkForAzureAdClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, policyName string, groupName string) (result aad.PrivateLinkResource, err error)
	ListByPrivateLinkPolicy(ctx context.Context, resourceGroupName string, policyName string) (result aad.PrivateLinkResourceListResultPage, err error)
	ListByPrivateLinkPolicyComplete(ctx context.Context, resourceGroupName string, policyName string) (result aad.PrivateLinkResourceListResultIterator, err error)
}

var _ PrivateLinkResourcesClientAPI = (*aad.PrivateLinkResourcesClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, policyName string, privateEndpointConnectionName string, parameters aad.PrivateEndpointConnection) (result aad.PrivateEndpointConnectionsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, policyName string, privateEndpointConnectionName string) (result aad.PrivateEndpointConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, policyName string, privateEndpointConnectionName string) (result aad.PrivateEndpointConnection, err error)
	ListByPolicyName(ctx context.Context, resourceGroupName string, policyName string) (result aad.PrivateEndpointConnectionListResultPage, err error)
	ListByPolicyNameComplete(ctx context.Context, resourceGroupName string, policyName string) (result aad.PrivateEndpointConnectionListResultIterator, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*aad.PrivateEndpointConnectionsClient)(nil)
