// Code generated by "stringer -trimprefix ImageDebugType -type ImageDebugType"; DO NOT EDIT.

package pe

import "strconv"

const (
	_ImageDebugType_name_0 = "UnknownCOFFCodeViewFPOMiscExceptionFixupOMapToSrcOMapFromSrcBorlandReserved10CLSID"
	_ImageDebugType_name_1 = "Repro"
)

var (
	_ImageDebugType_index_0 = [...]uint8{0, 7, 11, 19, 22, 26, 35, 40, 49, 60, 67, 77, 82}
)

func (i ImageDebugType) String() string {
	switch {
	case 0 <= i && i <= 11:
		return _ImageDebugType_name_0[_ImageDebugType_index_0[i]:_ImageDebugType_index_0[i+1]]
	case i == 16:
		return _ImageDebugType_name_1
	default:
		return "ImageDebugType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
