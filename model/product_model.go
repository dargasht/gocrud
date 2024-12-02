package model

import "github.com/dargasht/gocrud/database/repo"

type ProductCReq struct {
	Name  string `json:"name" validate:"required"`
	Price int64  `json:"price" validate:"gt=0"`
}

func (p ProductCReq) ToRepo() repo.CreateProductParams {
	return repo.CreateProductParams(p)
}

type ProductUReq struct {
	Name  string `json:"name" validate:"required"`
	Price int64  `json:"price" validate:"gt=0"`
	ID    int32  `json:"-"`
}

func (p ProductUReq) ToRepo() repo.UpdateProductParams {
	return repo.UpdateProductParams(p)
}

func (p ProductUReq) SetID(id int32) any {
	p.ID = id
	return p
}
