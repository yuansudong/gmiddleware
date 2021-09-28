package gmiddleware

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	// DialTimeout 连接建立的超时时间
	DialTimeout = 5 * time.Second

	// BackoffMaxDelay 连接尝试失败后退出时提供的最大延迟
	BackoffMaxDelay = 3 * time.Second

	// KeepAliveTime 每隔KeepAliveTime时间，发送PING帧测量最小往返时间，确定空闲连接是否仍然有效，我们设置为10S
	KeepAliveTime = time.Duration(30) * time.Second

	// KeepAliveTimeout 超过KeepAliveTimeout，关闭连接，我们设置为3S
	KeepAliveTimeout = time.Duration(3) * time.Second

	// InitialWindowSize 基于Stream的滑动窗口，类似于TCP的滑动窗口，用来做流控，默认64KiB，吞吐量上不去，Client和Server我们调到1GiB.
	InitialWindowSize = 1 << 30

	// InitialConnWindowSize 基于Connection的滑动窗口，默认16 * 64KiB，吞吐量上不去，Client和Server我们也都调到1GiB
	InitialConnWindowSize = 1 << 30

	// MaxSendMsgSize GRPC最大允许发送的字节数，默认4MiB，如果超过了GRPC会报错
	MaxSendMsgSize = 4 << 30

	// MaxRecvMsgSize GRPC最大允许接收的字节数，默认4MiB，如果超过了GRPC会报错。Client和Server我们都调到4GiB
	MaxRecvMsgSize = 4 << 30
)

// GetGrpcService 用于获取一个grpc 服务
func GetGrpcService() *grpc.Server {
	return grpc.NewServer(
		grpc.InitialWindowSize(InitialWindowSize),
		grpc.InitialConnWindowSize(InitialConnWindowSize),
		grpc.MaxSendMsgSize(MaxSendMsgSize),
		grpc.MaxRecvMsgSize(MaxRecvMsgSize),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             time.Duration(30) * time.Second,
			PermitWithoutStream: true,
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    KeepAliveTime,
			Timeout: KeepAliveTimeout,
		}),
	)
}
