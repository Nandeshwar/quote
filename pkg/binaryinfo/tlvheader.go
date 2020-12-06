package binaryinfo

type HeaderType uint8
type HeaderLength uint8

type Header struct {
	Type   HeaderType
	Length HeaderLength
}
