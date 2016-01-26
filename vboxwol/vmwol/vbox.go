package vmwol

import "fmt"

var binaryPath = "VBoxManage"

type VirtualBoxVM struct {
	GenericVM
	id string
}

type VirtualBoxProvider struct {
	GenericVMProvider
}

func NewVirtualBoxProvider() *VirtualBoxProvider {
	provider := &VirtualBoxProvider{}
	provider.SetBinaryPath(binaryPath)
	provider.BuildInventory()

	return provider
}

func (provider *VirtualBoxProvider) GetInventory() []string {
	if len(provider.inventory) == 0 && provider.clearCache == true {
		provider.BuildInventory()
	}

	return provider.inventory
}

// VIRTUALBOX
func (provider *VirtualBoxProvider) BuildInventory() error {
	// boxes := make([]VirtualBoxVM, 0)
	// out, err := exec.Command("VBoxManage", "list", "vms").Output()
	// if err != nil {
	//   log.Fatal(err)
	// }

	// machines := strings.Split(string(out), "\n")
	// boxRegex := regexp.MustCompile("^\"([a-zA-Z]+)\"")
	// idRegex := regexp.MustCompile("\\{(\\w+-\\w+-+\\w+-\\w+-\\w+)\\}")
	// for index, data := range machines {
	//   if data != "" {
	//     boxes = append(boxes, GenericVM{
	//       name: boxRegex.FindStringSubmatch(data)[0],
	//     })

	//     boxes[index].id = idRegex.FindStringSubmatch(data)[0]
	//   }
	// }

	// for index, box := range boxes {
	//   out, err := exec.Command("VBoxManage", "showvminfo", box.id).Output()

	//   if err != nil {
	//     log.Fatal(err)
	//   }

	//   scanner := bufio.NewScanner(bytes.NewReader(out))
	//   regexMac := regexp.MustCompile("MAC: (0-9A-F-a-f{12})")
	//   for scanner.Scan() {
	//     found := regexMac.FindStringSubmatch(scanner.Text())
	//     if len(found) > 0 {
	//       boxes[index].mac = found[1]
	//     }
	//   }
	// }

	fmt.Println("len", len(provider.inventory))
	provider.inventory = []string{"not", "cached", "now"}

	return nil
}
