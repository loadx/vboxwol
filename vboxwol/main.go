package main

import (
	"fmt"
	"log"
)

import "./vmwol"

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	vbox := vmwol.NewVirtualBoxProvider()
	fmt.Println(vbox.GetInventory())
	fmt.Println(vbox.GetInventory())
	vbox.ResetInventory()
	fmt.Println(vbox.GetInventory())

	// p.BuildInventory()
	// fmt.Println(p.GetInventory())
	// providers := map[string]vmwol.VMProvider{
	// 	"vbox":   &vmwol.VirtualBoxVMProvider{},
	// 	"vmware": &vmwol.VmwareVMProvider{},
	// }

	// for _, provider := range providers {
	// 	// InitProvider(provider)
	// 	provider.BuildInventory()
	// 	provider.BuildInventory()
	// 	fmt.Println(provider.GetInventory())
	// }

	// vmwol.DiscoverProviders()
}
