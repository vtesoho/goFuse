package main

import (
	"fmt"
	"testrd/src/other"
	"time"
)

func main() {
	fmt.Println("ss")
	// go itemaShow()
	// go itembShow()
	go itema()
	go itemb()
	for {
		time.Sleep(time.Hour * 10)
	}
}

func itemaShow() {
	supplier := other.ThirdSupplierSwitch["a"].(other.ThirdSupplier)
	for {
		supplier.Show()
		time.Sleep(time.Second)
	}

}
func itembShow() {
	supplier := other.ThirdSupplierSwitch["b"].(other.ThirdSupplier)
	for {
		supplier.Show()
		time.Sleep(time.Second)
	}

}

func itema() {
	for {
		supplier := other.ThirdSupplierSwitch["a"].(other.ThirdSupplier)
		supplier.Check()
		time.Sleep(time.Second * 2)
	}

}
func itemb() {
	for {
		supplier := other.ThirdSupplierSwitch["b"].(other.ThirdSupplier)
		supplier.Check()
		time.Sleep(time.Second)
	}

}
