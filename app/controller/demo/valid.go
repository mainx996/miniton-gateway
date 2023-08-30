package demo

import (
	"errors"
	schemDemo "www.miniton-gateway.com/app/schema/demo"
)

func validateDetail(req *schemDemo.DetailReq) error {
	if req.ID <= 0 {
		return errors.New("id err")
	}
	return nil
}

func validateCreate(req *schemDemo.CreateReq) error {
	if req.Name == "" {
		return errors.New("name err")
	}
	if req.Age <= 0 {
		return errors.New("age err")
	}
	return nil
}
