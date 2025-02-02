package rpc_server

import (
	"context"
	etcdInit "feedService/rpc_server/etcd"
	tokenproto "feedService/services/token"
	"fmt"

	"github.com/micro/go-micro/v2"
)

func GetIdByToken(token string) (int64, error) {
	tokenMicroService := micro.NewService(micro.Registry(etcdInit.EtcdReg))
	tokenService := tokenproto.NewTokenService("tokenService", tokenMicroService.Client())

	var req tokenproto.ParseTokenToIdRequest
	req.Token = token

	resp, err := tokenService.ParseTokenToId(context.TODO(), &req)
	if err != nil {
		fmt.Println(err)
	}

	return int64(resp.UserId), err
}
