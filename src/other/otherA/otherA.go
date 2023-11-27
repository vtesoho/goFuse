package otherA

import (
	"fmt"
	"testrd/src/fuse"
)

func init() {
	fmt.Println("执行了 otherA init")

}

type Other struct {
}

func (c Other) Check() (bool, error) {
	fuse.CreateFuse("otherA").ErrNumber++
	return false, nil

}

func (c Other) Show() {
	fuse := fuse.CreateFuse("otherA")
	fmt.Println("show otherA", fuse)

}
