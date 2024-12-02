package model

import "github.com/dargasht/gocrud/database/repo"

type UserCReq struct { //CReq stands for Create Request
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func (p UserCReq) ToRepo() repo.CreateUserParams {
	return repo.CreateUserParams(p)
}

type UserUReq struct { //UReq stands for Update Request
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	ID    int32  `json:"-"` // "-" means this field is not required alway leave it like this
}

func (p UserUReq) ToRepo() repo.UpdateUserParams {
	return repo.UpdateUserParams(p)
}
func (p UserUReq) SetID(id int32) UserUReq {
	p.ID = id
	return p
}
