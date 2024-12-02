package gocrud

type CReq[T any] interface {
	ToRepo() T
}

type UReq[T any] interface {
	ToRepo() T
	SetID(id int32) UReq[T]
}
