package gate

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"

	rpc "github.com/aiyun/gomqtt/proto"
	"github.com/corego/tools"
	"github.com/uber-go/zap"
)

func initRpc() {

}

type rpcServie struct {
	conn   *grpc.ClientConn
	client rpc.RpcClient
}

func (r *rpcServie) init(addr string) error {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}

	c := rpc.NewRpcClient(conn)

	r.conn = conn
	r.client = c

	return nil
}

func (r *rpcServie) close() {
	r.conn.Close()
}

func (r *rpcServie) login(msg *rpc.LoginMsg) error {
	req, err := r.client.Login(context.Background(), msg)
	if err != nil {
		return err
	}

	if !req.R {
		return errors.New(tools.Bytes2String(req.M))
	}

	return nil
}

func (r *rpcServie) logout(msg *rpc.LogoutMsg) error {
	req, err := r.client.Logout(context.Background(), msg)
	if err != nil {
		return err
	}

	if !req.R {
		return errors.New(tools.Bytes2String(req.M))
	}

	return nil
}

func (r *rpcServie) subscribe(msg *rpc.SubMsg) error {
	req, err := r.client.Subscribe(context.Background(), msg)
	if err != nil {
		Logger.Error("Subscribe", zap.Error(err))
		return err
	}
	log.Println(req)
	return nil
}

func (r *rpcServie) unSubscribe(msg *rpc.UnSubMsg) error {
	req, err := r.client.UnSubscribe(context.Background(), msg)
	if err != nil {
		Logger.Error("UnSubscribe", zap.Error(err))
		return err
	}
	log.Println(req)
	return nil
}
