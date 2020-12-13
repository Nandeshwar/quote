package binaryinfo

type ByteConverter interface {
	ToJSON() (string, error)
}
