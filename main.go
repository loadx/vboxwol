package main

import "fmt"

type VM struct {
	name string
	id   string
	mac  string
}

type GenericVMProvider struct {
	inventory []VM
}

type VboxVMProvider struct {
	GenericVMProvider
}

type VmwareVMProvider struct {
	GenericVMProvider
}

type VMProvider interface {
	getInventory() error
	// wakeVMByName(string) (bool, error)
	// WakeVMByMac(string) (bool, error)
	// WakeVMById(string) (bool, error)
}

func main() {
	hypervisors := []VMProvider{
		VboxVMProvider{},
		VmwareVMProvider{},
	}

	for _, hypervisor := range hypervisors {
		hypervisor.buildInventory()
		fmt.Println(hypervisor)
	}
}

// // VIRTUALBOX
// func (provider *VboxVMProvider) buildInventory() error {
// 	boxes := make([]VM, 0)
// 	out, err := exec.Command("VBoxManage", "list", "vms").Output()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	machines := strings.Split(string(out), "\n")
// 	boxRegex := regexp.MustCompile("^\"([a-zA-Z]+)\"")
// 	idRegex := regexp.MustCompile("\\{(\\w+-\\w+-+\\w+-\\w+-\\w+)\\}")
// 	for _, data := range machines {
// 		if data != "" {
// 			boxes = append(boxes, VM{
// 				name: boxRegex.FindStringSubmatch(data)[0],
// 				id:   idRegex.FindStringSubmatch(data)[0],
// 			})
// 		}
// 	}

// 	for index, box := range boxes {
// 		out, err := exec.Command("VBoxManage", "showvminfo", box.id).Output()

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		scanner := bufio.NewScanner(bytes.NewReader(out))
// 		regexMac := regexp.MustCompile("MAC: (0-9A-F-a-f{12})")
// 		for scanner.Scan() {
// 			found := regexMac.FindStringSubmatch(scanner.Text())
// 			if len(found) > 0 {
// 				boxes[index].mac = found[1]
// 			}
// 		}
// 	}

// 	provider.inventory = boxes
// 	return nil
// }
