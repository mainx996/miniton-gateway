package demo

import (
	"errors"
	"github.com/gin-gonic/gin"
	schemDemo "www.miniton-gateway.com/app/schema/demo"
	servDemo "www.miniton-gateway.com/app/service/demo"
	"www.miniton-gateway.com/utils/response"
)

func Detail(c *gin.Context) {
	var (
		err error
		req schemDemo.DetailReq
	)
	if err = c.ShouldBind(&req); err != nil {
		response.BindErr(c)
		return
	}
	if err = validateDetail(&req); err != nil {
		response.ValidateErr(c, err)
		return
	}
	resp, err := servDemo.Detail(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, response.NotFoundErr) {
			response.FoundErr(c)
		} else {
			response.SysErr(c, err)
		}
		return
	}
	response.Success(c, gin.H{"data": resp})
}

func Create(c *gin.Context) {
	var (
		err error
		req schemDemo.CreateReq
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		response.BindErr(c)
		return
	}
	if err = validateCreate(&req); err != nil {
		response.ValidateErr(c, err)
		return
	}
	resp, err := servDemo.Create(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, response.NotFoundErr) {
			response.FoundErr(c)
		} else {
			response.SysErr(c, err)
		}
		return
	}
	response.Success(c, gin.H{"data": resp.ID})
}

func List(c *gin.Context) {
	var (
		err error
	)
	resp, err := servDemo.List(c.Request.Context())
	if err != nil {
		if errors.Is(err, response.NotFoundErr) {
			response.FoundErr(c)
		} else {
			response.SysErr(c, err)
		}
		return
	}
	response.Success(c, gin.H{"data": resp})
}
