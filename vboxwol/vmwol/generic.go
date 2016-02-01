package vmwol

import (
	"fmt"
	"log"
	"os/exec"
)

type GenericVM struct {
	name string
	mac  string
}

type GenericVMProvider struct {
	inventory  map[string]GenericVM
	clearCache bool
	binary     string
}

type VMProvider interface {
	GetInventory() []GenericVM
	// wakeVMByName(string) (bool, error)
	// WakeVMByMac(string) (bool, error)
	// WakeVMById(string) (bool, error)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func DiscoverProviders() {
	// providers =
	// for _, provider := range vmproviders {
	//  New
	// }
}

func binaryIsUsable(binaryPath string) bool {
	err := exec.Command(binaryPath).Run()
	if err != nil {
		return false
	}

	return true
}

func (provider *GenericVMProvider) SetBinaryPath(overridePath string) {
	if binaryIsUsable(overridePath) {
		provider.binary = overridePath
	} else {
		errMsg := fmt.Sprintf("Unable to use binary '%s'. Ensure this exists and is executable.", overridePath)
		log.Fatal(errMsg)
	}
}

// func (p *GenericVMProvider) GetInventory() (res []string) {
//  // will return cached inventory
//  if len(p.inventory) == 0 {
//    p.BuildInventory()
//  }
//  return p.inventory
// }

// func (p *GenericVMProvider) ResetInventory() {
// 	p.clearCache = true
// 	p.inventory = make([]string, 0)
// }

func (p GenericVMProvider) RunCommand(args ...string) []byte {
	out, err := exec.Command(p.binary, args...).Output()
	checkError(err)

	return out
}
