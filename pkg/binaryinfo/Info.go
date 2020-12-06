package binaryinfo

type Info struct {
	Id        uint8
	Name      string // max 10 bytes
	Fee       float64
	CreatedAt string // 2012-11-01T22:08:41+00:00 (25 char length) - RFC3339
}

const INFO_TYPE = 1
const INFO_LEN = 44
