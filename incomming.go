package gmiddleware

const uModeHeader = "X-Mode"
const lModeHeader = "x-mode"
const uXForwardedFor = "X-Forwarded-For"
const lXForwardedFor = "x-forwarded-for"
const modeMock = "mock"
const uMacHeader = "X-Mac"
const lMacHeader = "x-mac"
const uTokenHeader = "X-Authorization-Token"
const lTokenHeader = "x-authorization-token"

const uRequestIDHeader = "X-Request-ID"
const lRequestIDHeader = "x-request-id"

const uIPHeader = "X-Real-IP"
const lIPHeader = "x-real-ip"

const sEmptyHeader = "Header Is Empty"

const sProtocolMajor = 2
const sProtocolContentType = "Content-Type"
const sProtocolGRPC = "application/grpc"

// HeaderMatch 客户端匹配
func HeaderMatch(key string) (string, bool) {
	return key, true
}
