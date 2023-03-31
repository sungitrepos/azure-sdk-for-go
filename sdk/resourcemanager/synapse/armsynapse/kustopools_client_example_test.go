//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsListSkus.json
func ExampleKustoPoolsClient_NewListSKUsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewKustoPoolsClient().NewListSKUsPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SKUDescriptionList = armsynapse.SKUDescriptionList{
		// 	Value: []*armsynapse.SKUDescription{
		// 		{
		// 			Name: to.Ptr("Compute optimized"),
		// 			LocationInfo: []*armsynapse.SKULocationInfoItem{
		// 				{
		// 					Location: to.Ptr("West US"),
		// 					Zones: []*string{
		// 						to.Ptr("1"),
		// 						to.Ptr("2"),
		// 						to.Ptr("3")},
		// 					},
		// 					{
		// 						Location: to.Ptr("West Europe"),
		// 						Zones: []*string{
		// 						},
		// 				}},
		// 				Locations: []*string{
		// 					to.Ptr("West US"),
		// 					to.Ptr("West Europe")},
		// 					Size: to.Ptr("Medium"),
		// 				},
		// 				{
		// 					Name: to.Ptr("Compute optimized"),
		// 					LocationInfo: []*armsynapse.SKULocationInfoItem{
		// 						{
		// 							Location: to.Ptr("West US"),
		// 							Zones: []*string{
		// 								to.Ptr("1"),
		// 								to.Ptr("2"),
		// 								to.Ptr("3")},
		// 							},
		// 							{
		// 								Location: to.Ptr("West Europe"),
		// 								Zones: []*string{
		// 								},
		// 						}},
		// 						Locations: []*string{
		// 							to.Ptr("West US"),
		// 							to.Ptr("West Europe")},
		// 							Size: to.Ptr("Large"),
		// 						},
		// 						{
		// 							Name: to.Ptr("Storage optimized"),
		// 							LocationInfo: []*armsynapse.SKULocationInfoItem{
		// 								{
		// 									Location: to.Ptr("West US"),
		// 									Zones: []*string{
		// 										to.Ptr("1"),
		// 										to.Ptr("2"),
		// 										to.Ptr("3")},
		// 									},
		// 									{
		// 										Location: to.Ptr("West Europe"),
		// 										Zones: []*string{
		// 										},
		// 								}},
		// 								Locations: []*string{
		// 									to.Ptr("West US"),
		// 									to.Ptr("West Europe")},
		// 									Size: to.Ptr("Medium"),
		// 								},
		// 								{
		// 									Name: to.Ptr("Storage optimized"),
		// 									LocationInfo: []*armsynapse.SKULocationInfoItem{
		// 										{
		// 											Location: to.Ptr("West US"),
		// 											Zones: []*string{
		// 												to.Ptr("1"),
		// 												to.Ptr("2"),
		// 												to.Ptr("3")},
		// 											},
		// 											{
		// 												Location: to.Ptr("West Europe"),
		// 												Zones: []*string{
		// 												},
		// 										}},
		// 										Locations: []*string{
		// 											to.Ptr("West US"),
		// 											to.Ptr("West Europe")},
		// 											Size: to.Ptr("Large"),
		// 									}},
		// 								}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsCheckNameAvailability.json
func ExampleKustoPoolsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewKustoPoolsClient().CheckNameAvailability(ctx, "westus", armsynapse.KustoPoolCheckNameRequest{
		Name: to.Ptr("kustoclusterrptest4"),
		Type: to.Ptr("Microsoft.Synapse/workspaces/kustoPools"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameResult = armsynapse.CheckNameResult{
	// 	Name: to.Ptr("kustoclusterrptest4"),
	// 	Message: to.Ptr("Name 'kustoclusterrptest4' is already taken. Please specify a different name"),
	// 	NameAvailable: to.Ptr(false),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsListByWorkspace.json
func ExampleKustoPoolsClient_ListByWorkspace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewKustoPoolsClient().ListByWorkspace(ctx, "kustorptest", "kustorptest", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.KustoPoolListResult = armsynapse.KustoPoolListResult{
	// 	Value: []*armsynapse.KustoPool{
	// 		{
	// 			Name: to.Ptr("KustoClusterRPTest4"),
	// 			Type: to.Ptr("Microsoft.Synapse/workspaces/kustopools"),
	// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/kustorptest/kustopools/KustoClusterRPTest4"),
	// 			Location: to.Ptr("westus"),
	// 			Etag: to.Ptr("abcd123"),
	// 			Properties: &armsynapse.KustoPoolProperties{
	// 				ProvisioningState: to.Ptr(armsynapse.ResourceProvisioningStateSucceeded),
	// 			},
	// 			SKU: &armsynapse.AzureSKU{
	// 				Name: to.Ptr(armsynapse.SKUNameStorageOptimized),
	// 				Capacity: to.Ptr[int32](2),
	// 				Size: to.Ptr(armsynapse.SKUSizeMedium),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("KustoClusterRPTest3"),
	// 			Type: to.Ptr("Microsoft.Synapse/workspaces/kustopools"),
	// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/kustorptest/kustopools/KustoClusterRPTest3"),
	// 			Location: to.Ptr("westus"),
	// 			Etag: to.Ptr("abcd123"),
	// 			Properties: &armsynapse.KustoPoolProperties{
	// 				ProvisioningState: to.Ptr(armsynapse.ResourceProvisioningStateSucceeded),
	// 			},
	// 			SKU: &armsynapse.AzureSKU{
	// 				Name: to.Ptr(armsynapse.SKUNameComputeOptimized),
	// 				Capacity: to.Ptr[int32](2),
	// 				Size: to.Ptr(armsynapse.SKUSizeSmall),
	// 			},
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsGet.json
func ExampleKustoPoolsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewKustoPoolsClient().Get(ctx, "synapseWorkspaceName", "kustoclusterrptest4", "kustorptest", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.KustoPool = armsynapse.KustoPool{
	// 	Name: to.Ptr("KustoClusterRPTest5"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/kustopools"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/synapseWorkspaceName/kustopools/KustoClusterRPTest5"),
	// 	Location: to.Ptr("westus"),
	// 	Etag: to.Ptr("abcd123"),
	// 	Properties: &armsynapse.KustoPoolProperties{
	// 		ProvisioningState: to.Ptr(armsynapse.ResourceProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armsynapse.AzureSKU{
	// 		Name: to.Ptr(armsynapse.SKUNameStorageOptimized),
	// 		Capacity: to.Ptr[int32](2),
	// 		Size: to.Ptr(armsynapse.SKUSizeMedium),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsCreateOrUpdate.json
func ExampleKustoPoolsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewKustoPoolsClient().BeginCreateOrUpdate(ctx, "synapseWorkspaceName", "kustorptest", "kustoclusterrptest4", armsynapse.KustoPool{
		Location: to.Ptr("westus"),
		Properties: &armsynapse.KustoPoolProperties{
			EnablePurge:           to.Ptr(true),
			EnableStreamingIngest: to.Ptr(true),
			WorkspaceUID:          to.Ptr("11111111-2222-3333-444444444444"),
		},
		SKU: &armsynapse.AzureSKU{
			Name:     to.Ptr(armsynapse.SKUNameStorageOptimized),
			Capacity: to.Ptr[int32](2),
			Size:     to.Ptr(armsynapse.SKUSizeMedium),
		},
	}, &armsynapse.KustoPoolsClientBeginCreateOrUpdateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.KustoPool = armsynapse.KustoPool{
	// 	Name: to.Ptr("KustoClusterRPTest4"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/kustopools"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/synapseWorkspaceName/kustopools/KustoClusterRPTest4"),
	// 	Location: to.Ptr("westus"),
	// 	Etag: to.Ptr("abcd"),
	// 	Properties: &armsynapse.KustoPoolProperties{
	// 		EnablePurge: to.Ptr(true),
	// 		EnableStreamingIngest: to.Ptr(true),
	// 		ProvisioningState: to.Ptr(armsynapse.ResourceProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armsynapse.AzureSKU{
	// 		Name: to.Ptr(armsynapse.SKUNameStorageOptimized),
	// 		Capacity: to.Ptr[int32](2),
	// 		Size: to.Ptr(armsynapse.SKUSizeMedium),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsUpdate.json
func ExampleKustoPoolsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewKustoPoolsClient().BeginUpdate(ctx, "synapseWorkspaceName", "kustorptest", "kustoclusterrptest4", armsynapse.KustoPoolUpdate{
		Properties: &armsynapse.KustoPoolProperties{
			EnablePurge:           to.Ptr(true),
			EnableStreamingIngest: to.Ptr(true),
			WorkspaceUID:          to.Ptr("11111111-2222-3333-444444444444"),
		},
		SKU: &armsynapse.AzureSKU{
			Name:     to.Ptr(armsynapse.SKUNameStorageOptimized),
			Capacity: to.Ptr[int32](2),
			Size:     to.Ptr(armsynapse.SKUSizeMedium),
		},
	}, &armsynapse.KustoPoolsClientBeginUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.KustoPool = armsynapse.KustoPool{
	// 	Name: to.Ptr("KustoClusterRPTest4"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/kustopools"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/synapseWorkspaceName/kustopools/KustoClusterRPTest4"),
	// 	Location: to.Ptr("westus"),
	// 	Etag: to.Ptr("abcd"),
	// 	Properties: &armsynapse.KustoPoolProperties{
	// 		EnablePurge: to.Ptr(true),
	// 		EnableStreamingIngest: to.Ptr(true),
	// 		ProvisioningState: to.Ptr(armsynapse.ResourceProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armsynapse.AzureSKU{
	// 		Name: to.Ptr(armsynapse.SKUNameStorageOptimized),
	// 		Capacity: to.Ptr[int32](2),
	// 		Size: to.Ptr(armsynapse.SKUSizeMedium),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsDelete.json
func ExampleKustoPoolsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewKustoPoolsClient().BeginDelete(ctx, "kustorptest", "kustorptest", "kustoclusterrptest4", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsStop.json
func ExampleKustoPoolsClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewKustoPoolsClient().BeginStop(ctx, "kustorptest", "kustoclusterrptest4", "kustorptest", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsStart.json
func ExampleKustoPoolsClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewKustoPoolsClient().BeginStart(ctx, "kustorptest", "kustoclusterrptest4", "kustorptest", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsListResourceSkus.json
func ExampleKustoPoolsClient_NewListSKUsByResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewKustoPoolsClient().NewListSKUsByResourcePager("synapseWorkspaceName", "kustoclusterrptest4", "kustorptest", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ListResourceSKUsResult = armsynapse.ListResourceSKUsResult{
		// 	Value: []*armsynapse.AzureResourceSKU{
		// 		{
		// 			Capacity: &armsynapse.AzureCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](100),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr(armsynapse.AzureScaleTypeAutomatic),
		// 			},
		// 			ResourceType: to.Ptr("workspaces/kustoPools"),
		// 			SKU: &armsynapse.AzureSKU{
		// 				Name: to.Ptr(armsynapse.SKUNameComputeOptimized),
		// 				Size: to.Ptr(armsynapse.SKUSizeMedium),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armsynapse.AzureCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](100),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr(armsynapse.AzureScaleTypeAutomatic),
		// 			},
		// 			ResourceType: to.Ptr("workspaces/kustoPools"),
		// 			SKU: &armsynapse.AzureSKU{
		// 				Name: to.Ptr(armsynapse.SKUNameComputeOptimized),
		// 				Size: to.Ptr(armsynapse.SKUSizeLarge),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armsynapse.AzureCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](100),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr(armsynapse.AzureScaleTypeAutomatic),
		// 			},
		// 			ResourceType: to.Ptr("workspaces/kustoPools"),
		// 			SKU: &armsynapse.AzureSKU{
		// 				Name: to.Ptr(armsynapse.SKUNameStorageOptimized),
		// 				Size: to.Ptr(armsynapse.SKUSizeMedium),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armsynapse.AzureCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](100),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr(armsynapse.AzureScaleTypeAutomatic),
		// 			},
		// 			ResourceType: to.Ptr("workspaces/kustoPools"),
		// 			SKU: &armsynapse.AzureSKU{
		// 				Name: to.Ptr(armsynapse.SKUNameStorageOptimized),
		// 				Size: to.Ptr(armsynapse.SKUSizeLarge),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolLanguageExtensionsList.json
func ExampleKustoPoolsClient_NewListLanguageExtensionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewKustoPoolsClient().NewListLanguageExtensionsPager("kustorptest", "kustoclusterrptest4", "kustorptest", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.LanguageExtensionsList = armsynapse.LanguageExtensionsList{
		// 	Value: []*armsynapse.LanguageExtension{
		// 		{
		// 			LanguageExtensionName: to.Ptr(armsynapse.LanguageExtensionNamePYTHON),
		// 		},
		// 		{
		// 			LanguageExtensionName: to.Ptr(armsynapse.LanguageExtensionNameR),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolLanguageExtensionsAdd.json
func ExampleKustoPoolsClient_BeginAddLanguageExtensions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewKustoPoolsClient().BeginAddLanguageExtensions(ctx, "kustorptest", "kustoclusterrptest4", "kustorptest", armsynapse.LanguageExtensionsList{
		Value: []*armsynapse.LanguageExtension{
			{
				LanguageExtensionName: to.Ptr(armsynapse.LanguageExtensionNamePYTHON),
			},
			{
				LanguageExtensionName: to.Ptr(armsynapse.LanguageExtensionNameR),
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolLanguageExtensionsRemove.json
func ExampleKustoPoolsClient_BeginRemoveLanguageExtensions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewKustoPoolsClient().BeginRemoveLanguageExtensions(ctx, "kustorptest", "kustoclusterrptest4", "kustorptest", armsynapse.LanguageExtensionsList{
		Value: []*armsynapse.LanguageExtension{
			{
				LanguageExtensionName: to.Ptr(armsynapse.LanguageExtensionNamePYTHON),
			},
			{
				LanguageExtensionName: to.Ptr(armsynapse.LanguageExtensionNameR),
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolFollowerDatabasesList.json
func ExampleKustoPoolsClient_NewListFollowerDatabasesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewKustoPoolsClient().NewListFollowerDatabasesPager("kustorptest", "kustoclusterrptest4", "kustorptest", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.FollowerDatabaseListResult = armsynapse.FollowerDatabaseListResult{
		// 	Value: []*armsynapse.FollowerDatabaseDefinition{
		// 		{
		// 			AttachedDatabaseConfigurationName: to.Ptr("attachedDbConfiguration"),
		// 			KustoPoolResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/kustorptest/kustoPools/follower1"),
		// 			DatabaseName: to.Ptr("*"),
		// 		},
		// 		{
		// 			AttachedDatabaseConfigurationName: to.Ptr("attachedDbConfiguration2"),
		// 			KustoPoolResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/kustorptest/kustoPools/follower4"),
		// 			DatabaseName: to.Ptr("db1"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolFollowerDatabasesDetach.json
func ExampleKustoPoolsClient_BeginDetachFollowerDatabases() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewKustoPoolsClient().BeginDetachFollowerDatabases(ctx, "kustorptest", "kustoclusterrptest4", "kustorptest", armsynapse.FollowerDatabaseDefinition{
		AttachedDatabaseConfigurationName: to.Ptr("myAttachedDatabaseConfiguration"),
		KustoPoolResourceID:               to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/kustorptest/kustoPools/leader4"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
