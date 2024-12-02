package gocrud

type CReq[T CRepo] interface {
	ToRepo() T
}

type CRepo interface {
}

type CRes interface {
}
