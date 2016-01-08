package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type ThingBase struct {
	cached string
}

type ThingOne struct {
	ThingBase
	one string
}

type ThingTwo struct {
	two    string
	three  string
	cached []string
}

type Thing interface {
	get() string
	set(...string)
	// wakeVMByName(string) (bool, error)
	// WakeVMByMac(string) (bool, error)
	// WakeVMById(string) (bool, error)
}

func main() {
	things := []Thing{
		&ThingOne{
			ThingBase{"none"},
			"1",
		},
		&ThingTwo{
			"a",
			"b",
			[]string{"1", "2"},
		},
	}

	things[0].set("fred")
	things[1].set("bob", "julius")

	for _, thing := range things {
		spew.Dump(thing)
	}
}

// VMWARE
func (this ThingOne) get() string {
	return this.one
}

func (this *ThingOne) set(newvals ...string) {
	this.cached = this.one
	this.one = newvals[0]
}

func (this ThingTwo) get() string {
	return fmt.Sprint(this.two, this.three)
}

func (this *ThingTwo) set(newvals ...string) {
	this.cached = []string{this.two, this.three}
	this.two = newvals[0]
	this.three = newvals[1]
}
