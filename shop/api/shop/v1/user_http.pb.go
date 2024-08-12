// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.27.3
// source: shop/v1/user.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationShopCaptcha = "/api.shop.v1.Shop/Captcha"
const OperationShopDetail = "/api.shop.v1.Shop/Detail"
const OperationShopLogin = "/api.shop.v1.Shop/Login"
const OperationShopRegister = "/api.shop.v1.Shop/Register"

type ShopHTTPServer interface {
	Captcha(context.Context, *emptypb.Empty) (*CaptchaReply, error)
	Detail(context.Context, *emptypb.Empty) (*UserDetailResponse, error)
	Login(context.Context, *LoginReq) (*RegisterReply, error)
	Register(context.Context, *RegisterReq) (*RegisterReply, error)
}

func RegisterShopHTTPServer(s *http.Server, srv ShopHTTPServer) {
	r := s.Route("/")
	r.POST("/api/users/register", _Shop_Register0_HTTP_Handler(srv))
	r.POST("/api/users/login", _Shop_Login0_HTTP_Handler(srv))
	r.GET("/api/users/captcha", _Shop_Captcha0_HTTP_Handler(srv))
	r.GET("/api/users/detail", _Shop_Detail0_HTTP_Handler(srv))
}

func _Shop_Register0_HTTP_Handler(srv ShopHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShopRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _Shop_Login0_HTTP_Handler(srv ShopHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShopLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _Shop_Captcha0_HTTP_Handler(srv ShopHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShopCaptcha)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Captcha(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CaptchaReply)
		return ctx.Result(200, reply)
	}
}

func _Shop_Detail0_HTTP_Handler(srv ShopHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShopDetail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Detail(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserDetailResponse)
		return ctx.Result(200, reply)
	}
}

type ShopHTTPClient interface {
	Captcha(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *CaptchaReply, err error)
	Detail(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *UserDetailResponse, err error)
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *RegisterReply, err error)
	Register(ctx context.Context, req *RegisterReq, opts ...http.CallOption) (rsp *RegisterReply, err error)
}

type ShopHTTPClientImpl struct {
	cc *http.Client
}

func NewShopHTTPClient(client *http.Client) ShopHTTPClient {
	return &ShopHTTPClientImpl{client}
}

func (c *ShopHTTPClientImpl) Captcha(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*CaptchaReply, error) {
	var out CaptchaReply
	pattern := "/api/users/captcha"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationShopCaptcha))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ShopHTTPClientImpl) Detail(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*UserDetailResponse, error) {
	var out UserDetailResponse
	pattern := "/api/users/detail"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationShopDetail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ShopHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/api/users/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationShopLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ShopHTTPClientImpl) Register(ctx context.Context, in *RegisterReq, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/api/users/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationShopRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
