package model

import "github.com/dargasht/gocrud/database/repo"

type PersonCReq struct { //CReq stands for Create Request
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func (p PersonCReq) ToRepo() repo.CreatePersonParams {
	return repo.CreatePersonParams(p)
}

type PersonUReq struct { //UReq stands for Update Request
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	ID    int32  `json:"-"` // "-" means this field is not required alway leave it like this
}

func (p PersonUReq) ToRepo() repo.UpdatePersonParams {
	return repo.UpdatePersonParams(p)
}
func (p PersonUReq) SetID(id int32) PersonUReq {
	p.ID = id
	return p
}
