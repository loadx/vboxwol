package vmwol

import (
	"bufio"
	"bytes"
	"regexp"
	"strings"
)

var virtualbox_binary_path = "VBoxManage"

type VirtualBoxVM struct {
	GenericVM
	id string
}

type VirtualBoxProvider struct {
	GenericVMProvider
}

func NewVirtualBoxProvider() *VirtualBoxProvider {
	provider := &VirtualBoxProvider{}
	provider.SetBinaryPath(virtualbox_binary_path)
	provider.GetInventory()

	return provider
}

func (provider *VirtualBoxProvider) GetInventory() []VirtualBoxVM {
	boxes := make([]VirtualBoxVM, 0)

	// get a list of all virtual machines in virtualbox
	out := provider.RunCommand("list", "vms")
	machines := strings.Split(string(out), "\n")
	boxRegex := regexp.MustCompile("^\"([a-zA-Z]+)\"")
	idRegex := regexp.MustCompile("\\{(\\w+-\\w+-+\\w+-\\w+-\\w+)\\}")

	// loop through and build a box struct for each machine
	for _, data := range machines {
		if data != "" {
			// fill in the known box information
			box := VirtualBoxVM{
				GenericVM: GenericVM{
					name: boxRegex.FindStringSubmatch(data)[0],
				},
				id: idRegex.FindStringSubmatch(data)[0],
			}

			// now fetch the MAC address and update the box information
			out := provider.RunCommand("showvminfo", box.id)
			scanner := bufio.NewScanner(bytes.NewReader(out))
			regexMac := regexp.MustCompile("MAC: [0-9A-F-a-f]{12}")
			for scanner.Scan() {
				found := regexMac.FindStringSubmatch(scanner.Text())
				if len(found) > 0 {
					box.mac = found[0]
				}
			}

			boxes = append(boxes, box)
		}
	}

	return boxes
}
