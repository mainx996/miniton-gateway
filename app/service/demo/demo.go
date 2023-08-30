package demo

import (
	"context"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	modelDemo "www.miniton-gateway.com/app/model/demo"
	schemaDemo "www.miniton-gateway.com/app/schema/demo"
	"www.miniton-gateway.com/pkg/log"
)

func Detail(ctx context.Context, req *schemaDemo.DetailReq) (resp *schemaDemo.DetailResp, err error) {
	msg := "demo.Detail"
	resp = new(schemaDemo.DetailResp)
	entity, err := modelDemo.DetailByID(ctx, req.ID)
	if err != nil {
		log.Error(ctx, msg+".DetailByID.Err", zap.Reflect("entity", entity), zap.Error(err))
		return
	}
	err = copier.Copy(&resp, &entity)
	if err != nil {
		log.Error(ctx, msg+".Copy.Err", zap.Error(err))
		return
	}
	return
}

func List(ctx context.Context) (resp *schemaDemo.ListResp, err error) {
	msg := "demo.List"
	resp = new(schemaDemo.ListResp)
	resp.Items = make([]*schemaDemo.Item, 0)
	entities, err := modelDemo.List(ctx)
	if err != nil {
		log.Error(ctx, msg+".List.Err", zap.Reflect("entities", entities), zap.Error(err))
		return
	}

	total, err := modelDemo.Count(ctx)
	if err != nil {
		log.Error(ctx, msg+".Count.Err", zap.Reflect("total", entities), zap.Error(err))
		return
	}
	err = copier.Copy(&resp.Items, &entities)
	if err != nil {
		log.Error(ctx, msg+".Copy.Err", zap.Error(err))
		return
	}

	resp.Total = total
	return
}

func Create(ctx context.Context, req *schemaDemo.CreateReq) (resp *schemaDemo.CreateResp, err error) {
	msg := "demo.Create"
	resp = new(schemaDemo.CreateResp)
	entity := new(modelDemo.Entity)

	err = copier.Copy(&entity, &req)
	if err != nil {
		log.Error(ctx, msg+".Copy.Err", zap.Error(err))
		return
	}
	err = modelDemo.Create(ctx, entity)
	if err != nil {
		log.Error(ctx, msg+".Create.Err", zap.Error(err))
		return
	}

	resp.ID = entity.ID
	return
}
