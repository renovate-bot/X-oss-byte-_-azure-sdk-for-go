//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/ReplicationStorageClassificationMappings_ListByReplicationStorageClassifications.json
func ExampleReplicationStorageClassificationMappingsClient_ListByReplicationStorageClassifications() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationStorageClassificationMappingsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	pager := client.ListByReplicationStorageClassifications("<fabric-name>",
		"<storage-classification-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("StorageClassificationMapping.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/ReplicationStorageClassificationMappings_Get.json
func ExampleReplicationStorageClassificationMappingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationStorageClassificationMappingsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<fabric-name>",
		"<storage-classification-name>",
		"<storage-classification-mapping-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("StorageClassificationMapping.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/ReplicationStorageClassificationMappings_Create.json
func ExampleReplicationStorageClassificationMappingsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationStorageClassificationMappingsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<fabric-name>",
		"<storage-classification-name>",
		"<storage-classification-mapping-name>",
		armrecoveryservicessiterecovery.StorageClassificationMappingInput{
			Properties: &armrecoveryservicessiterecovery.StorageMappingInputProperties{
				TargetStorageClassificationID: to.StringPtr("<target-storage-classification-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("StorageClassificationMapping.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/ReplicationStorageClassificationMappings_Delete.json
func ExampleReplicationStorageClassificationMappingsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationStorageClassificationMappingsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<fabric-name>",
		"<storage-classification-name>",
		"<storage-classification-mapping-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/ReplicationStorageClassificationMappings_List.json
func ExampleReplicationStorageClassificationMappingsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationStorageClassificationMappingsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	pager := client.List(nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("StorageClassificationMapping.ID: %s\n", *v.ID)
		}
	}
}
