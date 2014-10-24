package ethpipe

import (
	"github.com/eris-ltd/thelonious/ethstate"
	"github.com/eris-ltd/thelonious/ethutil"
)

type Object struct {
	*ethstate.StateObject
}

func (self *Object) StorageString(str string) *ethutil.Value {
	if ethutil.IsHex(str) {
		return self.Storage(ethutil.Hex2Bytes(str[2:]))
	} else {
		return self.Storage(ethutil.RightPadBytes([]byte(str), 32))
	}
}

func (self *Object) StorageValue(addr *ethutil.Value) *ethutil.Value {
	return self.Storage(addr.Bytes())
}

func (self *Object) Storage(addr []byte) *ethutil.Value {
	return self.StateObject.GetStorage(ethutil.BigD(addr))
}
