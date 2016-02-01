package vmwol

var vmware_binary_path = "VmwareManage"

type VmwareVM struct {
	GenericVM
}

type VmwareVMProvider struct {
	GenericVMProvider
}

func NewVmwareProvider() *VmwareVMProvider {
	provider := &VmwareVMProvider{}
	// provider.SetBinaryPath(vmware_binary_path)
	provider.GetInventory()

	return provider
}

// VIRTUALBOX
// func (provider *VmwareVMProvider) BuildInventory() error {
// 	provider.inventory = []string{"foo2", "bar2", "baz2"}
// 	return nil
// }

func (provider *VmwareVMProvider) GetInventory() []VmwareVM {
	return []VmwareVM{
		VmwareVM{
			GenericVM{
				name: "test",
				mac:  "123",
			},
		},
		VmwareVM{
			GenericVM{
				name: "test 2",
				mac:  "123 5",
			},
		},
	}
}
