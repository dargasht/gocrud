package gocrud

type UReq[T URepo] interface {
	SetID(id int32) any
	ToRepo() T
}

type URepo interface{}
