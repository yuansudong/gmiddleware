package gmiddleware

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc/metadata"
)

// Environment 这里仅仅是一个继承context的举例,一般这个context会被用于工程中,每个context都不相同
type Environment interface {
	GetMode() string
	SetMode(string)
	GetMac() string
	SetMac(string)
	GetIP() string
	SetIP(string)
	GetRID() string
	SetRID(string)
	GetAuth() string
	SetAuth(string)
	// GetToken() *pbtoken.Token
	// SetToken(*pbtoken.Token)
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

// Structure 用于描述结构
type Structure struct {
	ctx context.Context
	// tok  *pbtoken.Token
	mode string
	mac  string
	ip   string
	rid  string
	auth string
}

// NewStruct 用于新建立一个结构
func NewStruct(ctx context.Context) *Structure {
	inst := new(Structure)
	inst.ctx = ctx
	return inst
}

// GetMode 用于获取mode
func (s *Structure) GetMode() string {
	return s.mode
}

// SetMode 用于获取mode
func (s *Structure) SetMode(m string) {
	s.mode = m
}

// GetMac 用于获取mac地址
func (s *Structure) GetMac() string {
	return s.mac
}

// SetMac 设置mac地址
func (s *Structure) SetMac(m string) {
	s.mac = m
}

// GetIP 用于获取IP
func (s *Structure) GetIP() string {
	return s.ip
}

// SetIP 用于设置IP
func (s *Structure) SetIP(m string) {
	s.ip = m
}

// GetRID 获取请求ID
func (s *Structure) GetRID() string {
	return s.rid
}

// SetRID 设置请求ID
func (s *Structure) SetRID(m string) {
	s.rid = m
}

// GetAuth 获取认证标识
func (s *Structure) GetAuth() string {
	return s.auth
}

// SetAuth 设置认证
func (s *Structure) SetAuth(m string) {
	s.auth = m
}

// // GetToken 用于获取token
// func (s *Structure) GetToken() *pbtoken.Token {
// 	return s.tok
// }

// // SetToken 用于设置token
// func (s *Structure) SetToken(m *pbtoken.Token) {
// 	s.tok = m
// }

// Deadline 实现context的接口
func (s *Structure) Deadline() (deadline time.Time, ok bool) {
	return s.ctx.Deadline()
}

// Done 实现context的接口
func (s *Structure) Done() <-chan struct{} {
	return s.Done()
}

// Err 实现context接口
func (s *Structure) Err() error {
	return s.ctx.Err()
}

// Value 实现context的接口
func (s *Structure) Value(key interface{}) interface{} {
	return s.ctx.Value(key)
}

// ParseHeader 用于解析头部
func ParseHeader(ctx Environment) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("Header Is Empty")
	}

	arr := md.Get(lIPHeader)
	if len(arr) != 0 {
		ctx.SetIP(arr[0])
	} else {
		arr = md.Get(lXForwardedFor)
		if len(arr) != 0 {
			ctx.SetIP(arr[0])
		}
	}

	arr = md.Get(lModeHeader)
	if len(arr) != 0 {
		ctx.SetMode(arr[0])
	}

	arr = md.Get(lMacHeader)
	if len(arr) != 0 {
		ctx.SetMac(arr[0])
	}
	arr = md.Get(lRequestIDHeader)
	if len(arr) != 0 {
		ctx.SetMode(arr[0])
	}
	arr = md.Get(lTokenHeader)
	if len(arr) != 0 {
		ctx.SetAuth(arr[0])
	}
	return nil
}
