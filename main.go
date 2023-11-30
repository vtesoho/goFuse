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
	// go itema()
	go func() {
		itemb(5000000)
	}()

	// go func() {
	// 	itembError(5)
	// 	time.Sleep(time.Second * 1)
	// 	itembError(5)
	// 	time.Sleep(time.Second * 1)
	// 	itembError(5)
	// 	time.Sleep(time.Second * 1)
	// 	itembError(5)
	// 	time.Sleep(time.Second * 1)
	// 	itembError(5)
	// 	time.Sleep(time.Second * 1)
	// 	itembError(30)
	// 	time.Sleep(time.Second * 5)
	// 	itembError(5)
	// }()

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

func itemb(number int) {
	supplier := other.ThirdSupplierSwitch["b"].(other.ThirdSupplier)
	for i := 0; i < number; i++ {
		if i%10000 == 0 {
			fmt.Println("i", i)
		}
		supplier.Check()
		// fmt.Println("status", status, "----", i)
		time.Sleep(time.Microsecond * 500)
	}

}

func itembError(number int) {
	supplier := other.ThirdSupplierSwitch["b"].(other.ThirdSupplier)
	for i := 0; i < number; i++ {
		supplier.Err()
		time.Sleep(time.Millisecond * 20)

	}
}
