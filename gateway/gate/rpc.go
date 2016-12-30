package gate

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"

	"github.com/corego/tools"
	rpc "github.com/taitan-io/gomqtt/proto"
)

func initRpc() {

}

func getRpc(ci *connInfo) (*rpcServie, error) {
	// connect to stream
	ip, err := consist.Get(tools.Bytes2String(ci.acc))
	if err != nil {
		return nil, err
	}

	rpc, ok := rpcRoutes[ip]
	if !ok {
		return nil, fmt.Errorf("no stream rpc available: %v", ip)
	}

	return rpc, nil
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

func (r *rpcServie) pubText(msg *rpc.PubTextMsg) error {
	req, err := r.client.PubText(context.Background(), msg)
	if err != nil {
		return err
	}

	if !req.R {
		return errors.New(tools.Bytes2String(req.M))
	}

	return nil
}

func (r *rpcServie) pubJson(msg *rpc.PubJsonMsg) error {
	req, err := r.client.PubJson(context.Background(), msg)
	if err != nil {
		return err
	}

	if !req.R {
		return errors.New(tools.Bytes2String(req.M))
	}

	return nil
}

func (r *rpcServie) puback(msg *rpc.PubAckMsg) error {
	req, err := r.client.PubAck(context.Background(), msg)
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
		return err
	}

	if !req.R {
		return errors.New(tools.Bytes2String(req.M))
	}
	return nil
}

func (r *rpcServie) unSubscribe(msg *rpc.UnSubMsg) error {
	req, err := r.client.UnSubscribe(context.Background(), msg)
	if err != nil {
		return err
	}

	if !req.R {
		return errors.New(tools.Bytes2String(req.M))
	}
	return nil
}
