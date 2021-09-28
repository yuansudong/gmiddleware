package gmiddleware

import (
	"net/http"
	"strings"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

// GRPCHandlerFunc 用于区分http以及grpc
func GRPCHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == sProtocolMajor && strings.Contains(
			r.Header.Get(sProtocolContentType),
			sProtocolGRPC,
		) {
			// 进入这里,代表当前协议是h2.且其请求是grpc类型
			grpcServer.ServeHTTP(w, r)
			return
		}

		// 进入这里,代表当前的协议是h1,其请求类型是常规的http协议

		// 放行所有OPTIONS方法,主要适用于浏览器环境下的跨域.
		// 不过,对于正式环境应该注释掉,在正式环境下,是通过在网关处配置跨域.
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Access-Control-Allow-Headers", "*")
		// w.Header().Set("Access-Control-Allow-Methods", "*")
		// w.Header().Set("Access-Control-Expose-Headers", "*")
		// w.Header().Set("Access-Control-Allow-Credentials", "true")
		// w.Header().Set("Access-Control-Max-Age", "1800")
		// if r.Method == "OPTIONS" {
		// 	w.WriteHeader(http.StatusNoContent)
		// 	return
		// }
		// 处理请求
		otherHandler.ServeHTTP(w, r)

	}), &http2.Server{})
}
