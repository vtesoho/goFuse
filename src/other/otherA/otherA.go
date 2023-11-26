package otherA

import (
	"fmt"
	"testrd/src/fuse"
	"time"
)

func init() {
	fmt.Println("执行了 otherA init")
	fuse := fuse.CreateFuse("otherA")

	go func() {
		for {
			fmt.Println("FuseStruct otherA", fuse.ErrNumber)

			time.Sleep(time.Second)

		}

	}()
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
