package gate

import (
	"context"
	"log"

	"google.golang.org/grpc"

	rpc "github.com/aiyun/gomqtt/proto"
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

func (r *rpcServie) login(acm *rpc.AccMsg) error {
	req, err := r.client.Login(context.Background(), acm)
	if err != nil {
		Logger.Error("Login", zap.Error(err))
		return err
	}
	log.Println(req)
	return nil
}

func (r *rpcServie) logout(acm *rpc.AccMsg) error {
	req, err := r.client.Logout(context.Background(), acm)
	if err != nil {
		Logger.Error("LogOut", zap.Error(err))
		return err
	}
	log.Println(req)
	return nil
}

func (r *rpcServie) subscribe(tm *rpc.TcMsg) error {
	req, err := r.client.Subscribe(context.Background(), tm)
	if err != nil {
		Logger.Error("Subscribe", zap.Error(err))
		return err
	}
	log.Println(req)
	return nil
}

func (r *rpcServie) unSubscribe(tm *rpc.TcMsg) error {
	req, err := r.client.Subscribe(context.Background(), tm)
	if err != nil {
		Logger.Error("UnSubscribe", zap.Error(err))
		return err
	}
	log.Println(req)
	return nil
}

func (r *rpcServie) pChat(ctx context.Context, pm *rpc.PChatMsg) error {
	req, err := r.client.PChat(context.Background(), pm)
	if err != nil {
		Logger.Error("PChat", zap.Error(err))
		return err
	}
	log.Println(req)
	return nil
}

func (r *rpcServie) gChat(ctx context.Context, gm *rpc.GChatMsg) error {
	req, err := r.client.GChat(context.Background(), gm)
	if err != nil {
		Logger.Error("GChat", zap.Error(err))
		return err
	}
	log.Println(req)
	return nil
}
