package miscellaneous

import (
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"reflect"
	"unsafe"

	"github.com/ethereum/go-ethereum/common"
)

func Max(x, y int) int {
	switch {
	case x < y:
		return y
	default:
		return x
	}
}

func Min(x, y int) int {
	switch {
	case x < y:
		return x
	default:
		return y
	}
}

func Dup(a []byte) []byte {
	if a == nil {
		return nil
	}
	b := []byte{}
	return append(b, a...)
}

func E8func(a uint8) []byte {
	buf := make([]byte, 1)
	buf[0] = a & 0xFF
	return buf
}

func D8func(a []byte) (uint8, error) {
	if len(a) != 1 {
		return 0, errors.New("D8func: Illegal slice length")
	}
	return uint8(a[0]), nil
}

func E16func(a uint16) []byte {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, a)
	return buf
}

func D16func(a []byte) (uint16, error) {
	if len(a) != 2 {
		return 0, errors.New("D16func: Illegal slice length")
	}
	return binary.LittleEndian.Uint16(a), nil
}

func E32func(a uint32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, a)
	return buf
}

func D32func(a []byte) (uint32, error) {
	if len(a) != 4 {
		return 0, errors.New("D32func: Illegal slice length")
	}
	return binary.LittleEndian.Uint32(a), nil
}

func E64func(a uint64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, a)
	return buf
}

func D64func(a []byte) (uint64, error) {
	if len(a) != 8 {
		return 0, errors.New("D64func: Illegal slice length")
	}
	return binary.LittleEndian.Uint64(a), nil
}

func EB16func(a uint16) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, a)
	return buf
}

func DB16func(a []byte) (uint16, error) {
	if len(a) != 2 {
		return 0, errors.New("DB16func: Illegal slice length")
	}
	return binary.BigEndian.Uint16(a), nil
}

func EB32func(a uint32) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, a)
	return buf
}

func DB32func(a []byte) (uint32, error) {
	if len(a) != 4 {
		return 0, errors.New("DB32func: Illegal slice length")
	}
	return binary.BigEndian.Uint32(a), nil
}

func EB64func(a uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, a)
	return buf
}

func DB64func(a []byte) (uint64, error) {
	if len(a) != 8 {
		return 0, errors.New("DB64func: Illegal slice length")
	}
	return binary.BigEndian.Uint64(a), nil
}

// slice length || slice
func Eslice(a []byte) []byte {
	buf := []byte{}
	buf = append(buf, E32func(uint32(len(a)))...)
	buf = append(buf, a...)
	return buf
}

func Dslice(data []byte) ([]byte, []byte, error) {
	if len(data) < 4 {
		return []byte{}, []byte{}, errors.New("DecodeSlice: Illegal slice length")
	}
	n, _ := D32func(data[:4])
	if data = data[4:]; uint32(len(data)) < n {
		return []byte{}, []byte{}, errors.New("DecodeSlice: Illegal slice length")
	}
	return Dup(data[:n]), data[n:], nil
}

// string transformation bytes
func Str2Bytes(s string) []byte {
	hp := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	hq := reflect.SliceHeader{Data: hp.Data, Len: hp.Len, Cap: hp.Len}
	return *(*[]byte)(unsafe.Pointer(&hq))
}

// []byte transformation string
func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// []byte transformation Address
func BytesSha1Address(addr []byte) common.Address {
	addrString := Bytes2Str(addr)
	addrHAS1 := SHA1(addrString)
	a := common.HexToAddress(addrHAS1)
	return a
}

// SHA1 string shortening
func SHA1(s string) string {
	o := sha1.New()
	o.Write(Str2Bytes(s))
	//o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// EMapKey merge m and K into a []byte
func EMapKey(m, k []byte) []byte {
	buf := []byte{}
	buf = append([]byte{'m'}, E32func(uint32(len(m)))...)
	buf = append(buf, m...)
	buf = append(buf, byte('+'))
	buf = append(buf, k...)
	return buf
}
