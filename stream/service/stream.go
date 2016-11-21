package service

import (
	"unsafe"

	"github.com/labstack/echo"
)

type Stream struct {
	upa   *UpdateAddr
	rpc   *Rpc
	cache *Cache
	hash  *Hash
}

var gStream *Stream

func New() *Stream {
	stream := &Stream{}
	return stream
}

func (s *Stream) Init() {

	// init etcd
	upa := NewUpdateAddr()
	upa.Init()

	// init cache
	cache := NewCache()
	cache.Init()

	// init rpc service
	rpc := NewRpc()
	rpc.Init()

	// hash 初始化
	hash := NewHash()
	//	init other

	s.rpc = rpc
	s.upa = upa
	s.cache = cache
	s.hash = hash

	gStream = s
}

func (s *Stream) Start(isStatic bool) {

	loadConfig(isStatic)

	// stream 初始化所有功能服务
	s.Init()

	// rpc start
	s.rpc.Start()

	// upa start
	s.upa.Start()

	go httpStart()
}

func (s *Stream) Close() error {
	s.upa.Close()
	s.rpc.Close()
	return nil
}

func httpStart() {
	// 启动Http服务
	e := echo.New()

	// 启动配置热更新
	e.GET("/reload", reload)

	// e.Run(standard.New(":8907"))

	err := e.Start(":8907")
	if err != nil {
		e.Logger.Fatal(err.Error())
	}
}

// zero-copy, []byte转为string类型
// 注意，这种做法下，一旦[]byte变化，string也会变化
// 谨慎，黑科技！！除非性能瓶颈，否则请使用string(b)
func Bytes2String(b []byte) (s string) {
	return *(*string)(unsafe.Pointer(&b))
	// pb := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	// ps := (*reflect.StringHeader)(unsafe.Pointer(&s))
	// ps.Data = pb.Data
	// ps.Len = pb.Len
	// return
}

// zero-coy, string类型转为[]byte
// 注意，这种做法下，一旦string变化，程序立马崩溃且不能recover
// 谨慎，黑科技！！除非性能瓶颈，否则请使用[]byte(s)
func String2Bytes(s string) (b []byte) {
	return *(*[]byte)(unsafe.Pointer(&s))
	// pb := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	// ps := (*reflect.StringHeader)(unsafe.Pointer(&s))
	// pb.Data = ps.Data
	// pb.Len = ps.Len
	// pb.Cap = ps.Len
	// return
}
