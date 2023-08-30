package demo

import p "www.miniton-gateway.com/pkg/schema"

type (
	DetailReq struct {
		ID int64 `json:"id" binding:"required" form:"id"`
	}

	DetailResp struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	CreateReq struct {
		Name string `json:"name" binding:"required"`
		Age  int    `json:"age" binding:"required"`
	}

	CreateResp struct {
		ID int64 `json:"id"`
	}

	Item struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	ListResp struct {
		Items []*Item `json:"items,omitempty"`
		p.Page
	}
)
