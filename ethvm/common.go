package ethvm

import (
	"github.com/eris-ltd/eth-go-mods/ethlog"
	"github.com/eris-ltd/eth-go-mods/ethutil"
	"math/big"
)

var vmlogger = ethlog.NewLogger("VM")

var (
	GasStep    = big.NewInt(1)
	GasSha     = big.NewInt(20)
	GasSLoad   = big.NewInt(20)
	GasSStore  = big.NewInt(100)
	GasBalance = big.NewInt(20)
	GasCreate  = big.NewInt(100)
	GasCall    = big.NewInt(20)
	GasMemory  = big.NewInt(1)
	GasData    = big.NewInt(5)
	GasTx      = big.NewInt(500)

	Pow256 = ethutil.BigPow(2, 256)

	LogTyPretty byte = 0x1
	LogTyDiff   byte = 0x2
)
