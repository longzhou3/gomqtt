package service

import (
	"net"

	context "golang.org/x/net/context"

	"google.golang.org/grpc"

	"fmt"

	"github.com/aiyun/gomqtt/proto"
	"github.com/uber-go/zap"
)

type Rpc struct {
	gs *grpc.Server
}

func NewRpc() *Rpc {
	rpc := &Rpc{}
	return rpc
}

func (rpc *Rpc) Init() {

}

func (rpc *Rpc) Start() {
	l, err := net.Listen("tcp", Conf.GrpcC.Addr)
	if err != nil {
		Logger.Panic("Init", zap.Error(err))
	}
	rpc.gs = grpc.NewServer()

	proto.RegisterRpcServer(rpc.gs, &Rpc{})
	go rpc.gs.Serve(l)
}

func (r *Rpc) Close() error {
	r.gs.Stop()
	return nil
}

// 推送流程

// 接收到消息查看在线
// 在线推送
// 是否要推送apns
// 存放消息

// Ack流程
// 消息Ack

// 用户相关设置流程

// 群流程

// ---------------- 用户相关接口  ----------------

// Login 登陆
func (rpc *Rpc) Login(ctx context.Context, msg *proto.LoginMsg) (*proto.LoginRet, error) {
	err := gStream.cache.As.Login(msg)
	if err != nil {
		return &proto.LoginRet{R: false, M: []byte(fmt.Sprint("%s", err.Error()))}, err
	}
	return &proto.LoginRet{R: true, M: []byte("ok")}, nil
}

// Logout 登出
func (rpc *Rpc) Logout(ctx context.Context, msg *proto.LogoutMsg) (*proto.LogoutRet, error) {
	err := gStream.cache.As.Logout(msg)
	if err != nil {
		return &proto.LogoutRet{R: false, M: []byte(fmt.Sprint("%s", err.Error()))}, err
	}
	return &proto.LogoutRet{R: true, M: []byte("Logout 成功调用")}, nil
}

// ---------------- 订阅相关接口  ----------------

// Subscribe 订阅
func (rpc *Rpc) Subscribe(ctx context.Context, msg *proto.SubMsg) (*proto.SubRet, error) {
	err := gStream.cache.As.Subscribe(msg)
	if err != nil {
		return &proto.SubRet{R: false, M: []byte(fmt.Sprint("%s", err.Error()))}, err
	}
	return &proto.SubRet{R: true, M: []byte("Subscribe 成功调用")}, nil
}

// UnSubscribe 取消订阅
func (rpc *Rpc) UnSubscribe(ctx context.Context, msg *proto.UnSubMsg) (*proto.UnSubRet, error) {

	return &proto.UnSubRet{R: true, M: []byte("UnSubscribe 成功调用")}, nil
}

// Publish 客户端请求
func (rpc *Rpc) Publish(ctx context.Context, msg *proto.PubMsg) (*proto.PubRet, error) {
	return &proto.PubRet{R: false, M: []byte("UnSubscribe 成功调用")}, nil
}

// // BPull 拉取广播推送
// func (rpc *Rpc) BPull(ctx context.Context, msg *proto.BPushMsg) (*proto.BPushRet, error) {

// 	return &proto.BPushRet{Msg: []byte("BPull 成功调用")}, nil
// }

// // SPull 拉取单播推送
// func (rpc *Rpc) SPull(ctx context.Context, msg *proto.SPushMsg) (*proto.SPushRet, error) {

// 	return &proto.SPushRet{Msg: []byte("SPull 成功调用")}, nil
// }

// // PPull 拉取私聊
// func (rpc *Rpc) PPull(ctx context.Context, msg *proto.PChatMsg) (*proto.PChatRet, error) {

// 	return &proto.PChatRet{Msg: []byte("PPull 成功调用")}, nil
// }

// // GPull 拉取群聊
// func (rpc *Rpc) GPull(ctx context.Context, msg *proto.GChatMsg) (*proto.GChatRet, error) {

// 	return &proto.GChatRet{Msg: []byte("GPull 成功调用")}, nil
// }

// // 用户设置相关接口

// // SetNick 设置昵称
// func (rpc *Rpc) SetNick(ctx context.Context, msg *proto.NickMsg) (*proto.NickRet, error) {

// 	return &proto.NickRet{Msg: []byte("SetNick 成功调用")}, nil
// }

// // SetApns 设置Apns
// func (rpc *Rpc) SetApns(ctx context.Context, msg *proto.ApnsMsg) (*proto.ApnsRet, error) {

// 	return &proto.ApnsRet{Msg: []byte("SetApns 成功调用")}, nil
// }

// // SetLabel Label
// func (rpc *Rpc) SetLabel(ctx context.Context, msg *proto.LabelMsg) (*proto.LabelRet, error) {

// 	return &proto.LabelRet{Msg: []byte("SetLabel 成功调用")}, nil
// }

// ---------------- 消息推送相关接口  ----------------

// // BPush 广播推送
// func (rpc *Rpc) BPush(ctx context.Context, bm *proto.BPushMsg) (*proto.BPushRet, error) {

// 	return &proto.BPushRet{}, nil
// }

// // SPush 单播推送
// func (rpc *Rpc) SPush(ctx context.Context, sm *proto.SPushMsg) (*proto.SPushRet, error) {

// 	return &proto.SPushRet{}, nil
// }

// // PChat 私聊
// func (rpc *Rpc) PChat(ctx context.Context, pm *proto.PChatMsg) (*proto.PChatRet, error) {

// 	return &proto.PChatRet{}, nil
// }

// // GChat 群聊
// func (rpc *Rpc) GChat(ctx context.Context, gm *proto.GChatMsg) (*proto.GChatRet, error) {

// 	return &proto.GChatRet{}, nil
// }
