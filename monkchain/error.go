package monkchain

import (
	"fmt"
	"github.com/eris-ltd/thelonious/monkutil"
	"math/big"
)

// Parent error. In case a parent is unknown this error will be thrown
// by the block manager
type ParentErr struct {
	Message string
}

func (err *ParentErr) Error() string {
	return err.Message
}

func ParentError(hash []byte) error {
	return &ParentErr{Message: fmt.Sprintf("Block's parent unkown %x", hash)}
}

func IsParentErr(err error) bool {
	_, ok := err.(*ParentErr)

	return ok
}

type UncleErr struct {
	Message string
}

func (err *UncleErr) Error() string {
	return err.Message
}

func UncleError(str string) error {
	return &UncleErr{Message: str}
}

func IsUncleErr(err error) bool {
	_, ok := err.(*UncleErr)

	return ok
}

// Block validation error. If any validation fails, this error will be thrown
type ValidationErr struct {
	Message string
}

func (err *ValidationErr) Error() string {
	return err.Message
}

func ValidationError(format string, v ...interface{}) *ValidationErr {
	return &ValidationErr{Message: fmt.Sprintf(format, v...)}
}

func IsValidationErr(err error) bool {
	_, ok := err.(*ValidationErr)

	return ok
}

type GasLimitErr struct {
	Message string
	Is, Max *big.Int
}

func IsGasLimitErr(err error) bool {
	_, ok := err.(*GasLimitErr)

	return ok
}
func (err *GasLimitErr) Error() string {
	return err.Message
}
func GasLimitError(is, max *big.Int) *GasLimitErr {
	return &GasLimitErr{Message: fmt.Sprintf("GasLimit error. Max %s, transaction would take it to %s", max, is), Is: is, Max: max}
}

type GasLimitTxErr struct {
	Message string
	Is, Max *big.Int
}

func IsGasLimitTxErr(err error) bool {
	_, ok := err.(*GasLimitTxErr)

	return ok
}
func (err *GasLimitTxErr) Error() string {
	return err.Message
}
func GasLimitTxError(is, max *big.Int) *GasLimitTxErr {
	return &GasLimitTxErr{Message: fmt.Sprintf("GasLimitTx error. Max %s, transaction would take %s", max, is), Is: is, Max: max}
}

type NonceErr struct {
	Message string
	Is, Exp uint64
}

func (err *NonceErr) Error() string {
	return err.Message
}

func NonceError(is, exp uint64) *NonceErr {
	return &NonceErr{Message: fmt.Sprintf("Nonce err. Is %d, expected %d", is, exp), Is: is, Exp: exp}
}

func IsNonceErr(err error) bool {
	_, ok := err.(*NonceErr)

	return ok
}

type OutOfGasErr struct {
	Message string
}

func OutOfGasError() *OutOfGasErr {
	return &OutOfGasErr{Message: "Out of gas"}
}
func (self *OutOfGasErr) Error() string {
	return self.Message
}

func IsOutOfGasErr(err error) bool {
	_, ok := err.(*OutOfGasErr)

	return ok
}

type TxFail struct {
	Tx  *Transaction
	Err error
}

type InvalidPermErr string

func InvalidPermError(addr []byte, role string) *InvalidPermErr {
	s := InvalidPermErr(fmt.Sprintf("Invalid permissions err on role %s for adddress %s", role, monkutil.Bytes2Hex(addr)))
	return &s
}

func (self *InvalidPermErr) Error() string {
	return string(*self)
}

type InvalidSigErr string

func InvalidSigError(signer, coinbase []byte) *InvalidSigErr {
	s := InvalidSigErr(fmt.Sprintf("Invalid signature err for coinbase %s signed by %s", monkutil.Bytes2Hex(coinbase), monkutil.Bytes2Hex(signer)))
	return &s
}

func (self *InvalidSigErr) Error() string {
	return string(*self)
}

type InvalidTurnErr string

func InvalidTurnError(observed, expected []byte) *InvalidTurnErr {
	s := InvalidTurnErr(fmt.Sprintf("Invalid miner in sequence. Got %s, expected %s", monkutil.Bytes2Hex(observed), monkutil.Bytes2Hex(expected)))
	return &s
}

func (self *InvalidTurnErr) Error() string {
	return string(*self)
}

type InvalidDifficultyErr string

func InvalidDifficultyError(observed, expected *big.Int, coinbase []byte) *InvalidDifficultyErr {
	s := InvalidDifficultyErr(fmt.Sprintf("Invalid difficulty for coinbase %s. Got %s, expected %s", monkutil.Bytes2Hex(coinbase), observed.String(), expected.String()))
	return &s
}

func (self *InvalidDifficultyErr) Error() string {
	return string(*self)
}

type TDError struct {
	a, b *big.Int
}

func (self *TDError) Error() string {
	return fmt.Sprintf("incoming chain has a lower or equal TD (%v <= %v)", self.a, self.b)
}
func IsTDError(e error) bool {
	_, ok := e.(*TDError)
	return ok
}
