package otherB

import (
	"fmt"
	"testrd/src/fuse"
)

func init() {
	fmt.Println("执行了 otherB init")
	// fuse.CreateFuse("otherB")
}

type Other struct {
}

func (c Other) Check() (bool, error) {

	isRequest := fuse.FuseStatus("otherB")
	return isRequest, nil

}
func (c Other) Show() {
	fuse := fuse.CreateFuse("otherB")
	fmt.Println("show otherB", fuse)

}
func (c Other) Err() {
	fuse := fuse.CreateFuse("otherB")
	fuse.ErrNumber++
	// fmt.Println("show otherB", fuse)

}
