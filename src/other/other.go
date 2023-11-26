package other

import (
	"testrd/src/other/otherA"
	"testrd/src/other/otherB"
)

type ThirdSupplier interface {
	Check() (bool, error)
	Show()
}

var ThirdSupplierSwitch = map[string]interface{}{
	"a": otherA.Other{},
	"b": otherB.Other{},
}
