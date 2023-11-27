package other

import (
	"testrd/src/other/otherB"
)

type ThirdSupplier interface {
	Check() (bool, error)
	Show()
	Err()
}

var ThirdSupplierSwitch = map[string]interface{}{
	// "a": otherA.Other{},
	"b": otherB.Other{},
}
