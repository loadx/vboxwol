package vmwol

type VmwareVM struct {
	GenericVM
}

type VmwareVMProvider struct {
	GenericVMProvider
}

// VIRTUALBOX
func (provider *VmwareVMProvider) BuildInventory() error {
	provider.inventory = []string{"foo2", "bar2", "baz2"}
	return nil
}
