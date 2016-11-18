package service

import (
	"net"

	context "golang.org/x/net/context"

	"google.golang.org/grpc"

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

// ---------------- 消息推送相关接口  ----------------

// BPush 广播推送
func (rpc *Rpc) BPush(ctx context.Context, bm *proto.BPushMsg) (*proto.RetBPushMsg, error) {

	return &proto.RetBPushMsg{}, nil
}

// SPush 单播推送
func (rpc *Rpc) SPush(ctx context.Context, sm *proto.SPushMsg) (*proto.RetSPushMsg, error) {

	return &proto.RetSPushMsg{}, nil
}

// PChat 私聊
func (rpc *Rpc) PChat(ctx context.Context, pm *proto.PChatMsg) (*proto.RetPChatMsg, error) {

	return &proto.RetPChatMsg{}, nil
}

// GChat 群聊
func (rpc *Rpc) GChat(ctx context.Context, gm *proto.GChatMsg) (*proto.RetGChatMsg, error) {

	return &proto.RetGChatMsg{}, nil
}

// ---------------- 用户相关接口  ----------------

// Login 登陆
func (rpc *Rpc) Login(ctx context.Context, msg *proto.LoginMsg) (*proto.RetLoginMsg, error) {
	// var user *User
	// user, ok := gStream.cache.As.GetUser(am.An, am.Un)
	// if !ok {
	// 	// 数据库中拉取
	// 	// 异常返回错误信息
	// }
	// user.Update(am, ONLINE)
	// Logger.Info("Login", zap.Object("user", user))

	return &proto.RetLoginMsg{Msg: []byte("Login 成功调用")}, nil
}

// Logout 登出
func (rpc *Rpc) Logout(ctx context.Context, msg *proto.LogoutMsg) (*proto.RetLogoutMsg, error) {

	return &proto.RetLogoutMsg{Msg: []byte("Logout 成功调用")}, nil
}

// ---------------- 订阅相关接口  ----------------

// Subscribe 订阅
func (rpc *Rpc) Subscribe(ctx context.Context, msg *proto.SubMsg) (*proto.RetSubMsg, error) {

	return &proto.RetSubMsg{Msg: []byte("Subscribe 成功调用")}, nil
}

// UnSubscribe 取消订阅
func (rpc *Rpc) UnSubscribe(ctx context.Context, msg *proto.UnSubMsg) (*proto.RetUnSubMsg, error) {

	return &proto.RetUnSubMsg{Msg: []byte("UnSubscribe 成功调用")}, nil
}

// BPull 拉取广播推送
func (rpc *Rpc) BPull(ctx context.Context, msg *proto.BPushMsg) (*proto.RetBPushMsg, error) {

	return &proto.RetBPushMsg{Msg: []byte("BPull 成功调用")}, nil
}

// SPull 拉取单播推送
func (rpc *Rpc) SPull(ctx context.Context, msg *proto.SPushMsg) (*proto.RetSPushMsg, error) {

	return &proto.RetSPushMsg{Msg: []byte("SPull 成功调用")}, nil
}

// PPull 拉取私聊
func (rpc *Rpc) PPull(ctx context.Context, msg *proto.PChatMsg) (*proto.RetPChatMsg, error) {

	return &proto.RetPChatMsg{Msg: []byte("PPull 成功调用")}, nil
}

// GPull 拉取群聊
func (rpc *Rpc) GPull(ctx context.Context, msg *proto.GChatMsg) (*proto.RetGChatMsg, error) {

	return &proto.RetGChatMsg{Msg: []byte("GPull 成功调用")}, nil
}

// 用户设置相关接口

// SetNick 设置昵称
func (rpc *Rpc) SetNick(ctx context.Context, msg *proto.NickMsg) (*proto.RetNickMsg, error) {

	return &proto.RetNickMsg{Msg: []byte("SetNick 成功调用")}, nil
}

// SetApns 设置Apns
func (rpc *Rpc) SetApns(ctx context.Context, msg *proto.ApnsMsg) (*proto.RetApnsMsg, error) {

	return &proto.RetApnsMsg{Msg: []byte("SetApns 成功调用")}, nil
}

// SetLabel Label
func (rpc *Rpc) SetLabel(ctx context.Context, msg *proto.LabelMsg) (*proto.RetLabelMsg, error) {

	return &proto.RetLabelMsg{Msg: []byte("SetLabel 成功调用")}, nil
}
