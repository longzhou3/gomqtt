package service

import (
	"fmt"
	"net"

	context "golang.org/x/net/context"

	"google.golang.org/grpc"

	"github.com/aiyun/gomqtt/proto"
	"github.com/corego/tools"
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

	// var addr string
	if Conf.GrpcC.Addr == "" {
		Conf.GrpcC.Addr = tools.LocalIP() + ":" + Conf.GrpcC.Port
	}
	Logger.Info("addr", zap.String("addr", Conf.GrpcC.Addr))

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

// ---------------- 用户相关接口  ----------------

// Login 登陆
func (rpc *Rpc) Login(ctx context.Context, msg *proto.LoginMsg) (*proto.LoginRet, error) {
	err := gStream.cache.As.Login(msg)
	if err != nil {
		Logger.Error("Login", zap.Error(err), zap.String("Acc", tools.Bytes2String(msg.Acc)), zap.Int64("Cid", msg.Cid))
		return &proto.LoginRet{R: false, M: []byte(fmt.Sprint("%s", err.Error()))}, err
	}
	// insert cid
	queue, retChan, err := gStream.cache.Cids.addAndRetQueueChan(msg)
	if err != nil {
		return &proto.LoginRet{R: false, M: []byte(fmt.Sprintf("%s", err.Error()))}, err
	}
	// (msg * proto.LoginMsg)(
	task := &taskMsg{
		cid:     msg.Cid,
		acc:     msg.Acc,
		queue:   queue,
		retChan: retChan,
		ts:      msg.Ts,
	}
	// log.Println(task)
	addTask(task)
	return &proto.LoginRet{R: true, M: []byte("ok")}, nil
}

// Logout 登出
func (rpc *Rpc) Logout(ctx context.Context, msg *proto.LogoutMsg) (*proto.LogoutRet, error) {
	var err error
	if acc, ok := gStream.cache.Cids.get(msg.Cid); ok {
		gStream.cache.Cids.delete(msg.Cid)
		err = gStream.cache.As.Logout(acc.acc, acc.appID)
		if err != nil {
			return &proto.LogoutRet{R: false, M: []byte(fmt.Sprint("%s", err.Error()))}, nil
		}
	} else {
		return &proto.LogoutRet{R: false, M: []byte(fmt.Sprint("unfind cid %d", msg.Cid))}, nil
	}
	return &proto.LogoutRet{R: true, M: []byte("ok")}, nil
}

// ---------------- 订阅相关接口  ----------------

// Subscribe 订阅
func (rpc *Rpc) Subscribe(ctx context.Context, msg *proto.SubMsg) (*proto.SubRet, error) {
	// if acc, ok := gStream.cache.Cids.get(msg.Cid); ok {
	// 	err := acc.acc.Subscribe(acc.appID, msg)
	// 	if err != nil {
	// 		log.Println("Subscribe err ", err)
	// 		return &proto.SubRet{R: false, M: []byte(fmt.Sprint("%s", err.Error()))}, err
	// 	}
	// }
	return &proto.SubRet{R: true, M: []byte("Subscribe 成功调用")}, nil
}

// UnSubscribe 取消订阅
func (rpc *Rpc) UnSubscribe(ctx context.Context, msg *proto.UnSubMsg) (*proto.UnSubRet, error) {
	// if acc, ok := gStream.cache.Cids.get(msg.Cid); ok {
	// 	err := acc.acc.UnSubscribe(acc.appID, msg)
	// 	if err != nil {
	// 		log.Println("UnSubscribe err ", err)
	// 		return &proto.UnSubRet{R: false, M: []byte(fmt.Sprint("%s", err.Error()))}, err
	// 	}
	// }
	return &proto.UnSubRet{R: true, M: []byte("UnSubscribe 成功调用")}, nil
}

// Publish 客户端请求
func (rpc *Rpc) Publish(ctx context.Context, msg *proto.PubMsg) (*proto.PubRet, error) {
	// if acc, ok := gStream.cache.Cids.get(msg.Cid); ok {

	// }
	return &proto.PubRet{R: false, M: []byte("UnSubscribe 成功调用")}, nil
}

// @Explain
// gateway推送来的消息只包含topic和acc,stream要通过acc和topic来查找cid, 所以cid要和topic的关系映射起来
// PubText
func (rpc *Rpc) PubText(ctx context.Context, msg *proto.PubTextMsg) (*proto.PubTextRet, error) {
	accMsg, ok := gStream.cache.Cids.get(msg.Cid)
	if !ok {
		return &proto.PubTextRet{R: false, M: []byte(fmt.Sprint("unfind cid %d", msg.Cid))}, nil
	}
	Logger.Info("Push", zap.String("ToAcc", tools.Bytes2String(msg.ToAcc)),
		zap.String("Topic", tools.Bytes2String(msg.Ttp)), zap.String("Msg", tools.Bytes2String(msg.Msg)))

	// 通过acc计算出队列
	queue, err := GetQueue(msg.ToAcc)
	if err != nil {
		Logger.Error("GetQueue", zap.Error(err), zap.String("acc", tools.Bytes2String(msg.ToAcc)))
		return nil, err
	}
	cacheTask := CacheTask{
		MsgTy:  CACHE_INSERT,
		Acc:    msg.ToAcc,
		Topic:  msg.Ttp,
		Msg:    msg,
		MsgIDs: [][]byte{msg.Mid},
	}

	queue.Publish(cacheTask)

	if acc, appid, ok := gStream.cache.As.GetAccAndAppID(accMsg.acc, accMsg.appID); ok {
		// 通过 acc topic找到cid
		if err := gStream.cache.As.PubText(acc, appid, msg); err != nil {
			return &proto.PubTextRet{R: false, M: []byte(fmt.Sprint("%s", err.Error()))}, nil
		}
	}

	return &proto.PubTextRet{R: true, M: []byte("PubText 成功调用")}, nil
}

// 	(ctx context.Context, in *PubAckMsg, opts ...grpc.CallOption) (*PubAckRet, error)
func (rpc *Rpc) PubAck(ctx context.Context, msg *proto.PubAckMsg) (*proto.PubAckRet, error) {
	for _, msgidMsg := range msg.Mids {
		// gStream.cache.msgCache.Get(msgid)
		// gStream.cache.msgCache.Delete(msgidMsg.Mid)
		// gStream.cache.msgIDManger.TextMsgAck(msg.Acc, msgidMsg.Tp, msgidMsg.Mid)
		Logger.Info("PubAck", zap.String("msgid", tools.Bytes2String(msgidMsg.Mid)))
	}
	return &proto.PubAckRet{R: true, M: []byte("PubAck 成功调用")}, nil
}

// SetAppID 设置AppID
// func (rpc *Rpc) SetAppID(ctx context.Context, msg *proto.AppIDMsg) (*proto.AppIDRet, error) {
// 	if acc, ok := gStream.cache.Cids.get(msg.Cid); ok {
// 		err := acc.acc.SetAppID(acc.user, msg)
// 		if err != nil {
// 			log.Println("SetAppID err ", err)
// 			return &proto.AppIDRet{R: false, M: []byte(fmt.Sprint("%s", err.Error()))}, err
// 		}
// 	}
// 	return &proto.AppIDRet{R: false, M: []byte("UnSubscribe 成功调用")}, nil
// }

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
