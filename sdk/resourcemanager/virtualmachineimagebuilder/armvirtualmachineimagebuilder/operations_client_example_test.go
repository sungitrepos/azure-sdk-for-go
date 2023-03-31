//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armvirtualmachineimagebuilder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/078b90617e5e08137d0395963bd4119f4561a910/specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2022-02-14/examples/OperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvirtualmachineimagebuilder.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationListResult = armvirtualmachineimagebuilder.OperationListResult{
		// 	Value: []*armvirtualmachineimagebuilder.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.VirtualMachineImages/register/action"),
		// 			Display: &armvirtualmachineimagebuilder.OperationDisplay{
		// 				Description: to.Ptr("Register Virtual Machine Image Builder RP"),
		// 				Operation: to.Ptr("Register Virtual Machine Image Builder RP"),
		// 				Provider: to.Ptr("Virtual Machine Image Builder"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.VirtualMachineImages/unregister/action"),
		// 			Display: &armvirtualmachineimagebuilder.OperationDisplay{
		// 				Description: to.Ptr("Unregister Virtual Machine Image Builder RP"),
		// 				Operation: to.Ptr("Unregister Virtual Machine Image Builder RP"),
		// 				Provider: to.Ptr("Virtual Machine Image Builder"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.VirtualMachineImages/imageTemplates/read"),
		// 			Display: &armvirtualmachineimagebuilder.OperationDisplay{
		// 				Description: to.Ptr("Get a VM image template instance resource"),
		// 				Operation: to.Ptr("Get a VM image template instance resource"),
		// 				Provider: to.Ptr("Virtual Machine Image Builder"),
		// 				Resource: to.Ptr("VM Image template"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.VirtualMachineImages/imageTemplates/write"),
		// 			Display: &armvirtualmachineimagebuilder.OperationDisplay{
		// 				Description: to.Ptr("Create or update a VM image template instance resource"),
		// 				Operation: to.Ptr("Create or update a VM image template instance resource"),
		// 				Provider: to.Ptr("Virtual Machine Image Builder"),
		// 				Resource: to.Ptr("VM Image template"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.VirtualMachineImages/imageTemplates/delete"),
		// 			Display: &armvirtualmachineimagebuilder.OperationDisplay{
		// 				Description: to.Ptr("Delete a VM image template instance resource"),
		// 				Operation: to.Ptr("Delete a VM image template instance resource"),
		// 				Provider: to.Ptr("Virtual Machine Image Builder"),
		// 				Resource: to.Ptr("VM Image template"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.VirtualMachineImages/imageTemplates/run/action"),
		// 			Display: &armvirtualmachineimagebuilder.OperationDisplay{
		// 				Description: to.Ptr("Execute a VM image template to produce its outputs"),
		// 				Operation: to.Ptr("Execute a VM image template to produce its outputs"),
		// 				Provider: to.Ptr("Virtual Machine Image Builder"),
		// 				Resource: to.Ptr("VM Image template"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.VirtualMachineImages/imageTemplates/runOutputs/read"),
		// 			Display: &armvirtualmachineimagebuilder.OperationDisplay{
		// 				Description: to.Ptr("Get a VM image template run output resource"),
		// 				Operation: to.Ptr("Get a VM image template run output resource"),
		// 				Provider: to.Ptr("Virtual Machine Image Builder"),
		// 				Resource: to.Ptr("VM Image template run output"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.VirtualMachineImages/operations/read"),
		// 			Display: &armvirtualmachineimagebuilder.OperationDisplay{
		// 				Description: to.Ptr("List available Virtual Machine Image Builder Operations"),
		// 				Operation: to.Ptr("List available Virtual Machine Image Builder Operations"),
		// 				Provider: to.Ptr("Virtual Machine Image Builder"),
		// 				Resource: to.Ptr("Operation"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.VirtualMachineImages/locations/operations/read"),
		// 			Display: &armvirtualmachineimagebuilder.OperationDisplay{
		// 				Description: to.Ptr("Get the status of an asynchronous operation"),
		// 				Operation: to.Ptr("Get the status of an asynchronous operation"),
		// 				Provider: to.Ptr("Virtual Machine Image Builder"),
		// 				Resource: to.Ptr("Asynchronous Operation"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.VirtualMachineImages/imageTemplates/cancel/action"),
		// 			Display: &armvirtualmachineimagebuilder.OperationDisplay{
		// 				Description: to.Ptr("Cancel a running image build"),
		// 				Operation: to.Ptr("Cancel a running image build"),
		// 				Provider: to.Ptr("Virtual Machine Image Builder"),
		// 				Resource: to.Ptr("VM Image template"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
