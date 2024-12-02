package gocrud

type UReq[T URepo] interface {
	SetID(id int32) UReq[T]
	ToRepo() T
}

type URepo interface{}
