package otherB

import (
	"fmt"
	"testrd/src/fuse"
)

func init() {
	fmt.Println("执行了 otherB init")
	fuse.CreateFuse("otherB")
}

type Other struct {
}

func (c Other) Check() (bool, error) {
	if fuse.FushStatus("otherB") {
		fmt.Println("能请求")
	}
	return false, nil

}
func (c Other) Show() {
	fuse := fuse.CreateFuse("otherB")
	fmt.Println("show otherB", fuse)

}
