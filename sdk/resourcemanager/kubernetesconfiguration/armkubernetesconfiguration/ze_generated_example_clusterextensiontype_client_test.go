//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration"
)

// x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/preview/2022-01-01-preview/examples/ClusterExtensionType_Get.json
func ExampleClusterExtensionTypeClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkubernetesconfiguration.NewClusterExtensionTypeClient("<subscription-id>", cred, nil)
	_, err = client.Get(ctx,
		"<resource-group-name>",
		armkubernetesconfiguration.Enum0MicrosoftContainerService,
		armkubernetesconfiguration.Enum1ManagedClusters,
		"<cluster-name>",
		"<extension-type-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}