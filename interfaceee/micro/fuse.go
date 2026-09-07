package micro

// fuse熔断器

// FuseGoogleSre 基于google-sre算法实现，
// 根据rpc失败率决定开启：fail_rate = max(0, (请求数-k*成功数))) / 请求数+1  k默认等于1.5
// 失败率判定：rpc返回如下错误：
// codes.Internal、codes.DeadlineExceeded、codes.Dataloss、codes.Unavailable
func FuseGoogleSre() {

}

// 1.adm如何重新命名：在goctl生成时，而不是手动改
// 2.rpc的调用链路：包裹grpc.pb.go 和 rpc.pb.go，adm_client，adm_server
// 3.熔断防止级联故障打开：
// circuit breaker is open
// HTTP ERROR 503
// 500 Internal Server Error
