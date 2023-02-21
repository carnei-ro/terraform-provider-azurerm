package compute

//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=DiskEncryptionSet -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/diskEncryptionSets/set1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=GalleryApplication -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/galleries/gallery1/applications/galleryApplication1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=GalleryApplicationVersion -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/galleries/gallery1/applications/galleryApplication1/versions/galleryApplicationVersion1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Image -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/images/image1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SharedImage -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/galleries/gallery1/images/image1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SharedImageGallery -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/galleries/gallery1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SharedImageVersion -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/galleries/gallery1/images/image1/versions/version1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=VirtualMachine -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/virtualMachines/machine1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=VirtualMachineExtension -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/virtualMachines/machine1/extensions/extension1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=VirtualMachineScaleSet -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/virtualMachineScaleSets/scaleSet1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=VirtualMachineScaleSetExtension -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/virtualMachineScaleSets/scaleSet1/extensions/extension1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SSHPublicKey -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/sshPublicKeys/sshpublickey1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=HybridMachine -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.HybridCompute/machines/machine1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=DataDisk -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Compute/virtualMachines/machine1/dataDisks/disk1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Plan -id=/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MarketplaceOrdering/agreements/agreement1/offers/offer1/plans/hourly
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=HostGroup -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Compute/hostGroups/hostgroup1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=VMSSInstance -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/virtualMachines/vm1
