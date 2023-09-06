// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.23.2
// source: api/network/v1/p2p.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationP2PCloseListen = "/api.network.v1.P2p/CloseListen"
const OperationP2PCreateForward = "/api.network.v1.P2p/CreateForward"
const OperationP2PCreateListen = "/api.network.v1.P2p/CreateListen"
const OperationP2PListListen = "/api.network.v1.P2p/ListListen"

type P2PHTTPServer interface {
	CloseListen(context.Context, *CloseListenRequest) (*CloseListenReply, error)
	CreateForward(context.Context, *CreateForwardRequest) (*CreateForwardReply, error)
	CreateListen(context.Context, *CreateListenRequest) (*CreateListenReply, error)
	ListListen(context.Context, *ListListenRequest) (*ListListenReply, error)
}

func RegisterP2PHTTPServer(s *http.Server, srv P2PHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/p2p/listen", _P2P_CreateListen0_HTTP_Handler(srv))
	r.POST("/v1/p2p/forward", _P2P_CreateForward0_HTTP_Handler(srv))
	r.POST("/v1/p2p/close", _P2P_CloseListen0_HTTP_Handler(srv))
	r.GET("/v1/p2p/ls", _P2P_ListListen0_HTTP_Handler(srv))
}

func _P2P_CreateListen0_HTTP_Handler(srv P2PHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateListenRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationP2PCreateListen)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateListen(ctx, req.(*CreateListenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateListenReply)
		return ctx.Result(200, reply)
	}
}

func _P2P_CreateForward0_HTTP_Handler(srv P2PHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateForwardRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationP2PCreateForward)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateForward(ctx, req.(*CreateForwardRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateForwardReply)
		return ctx.Result(200, reply)
	}
}

func _P2P_CloseListen0_HTTP_Handler(srv P2PHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CloseListenRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationP2PCloseListen)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CloseListen(ctx, req.(*CloseListenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CloseListenReply)
		return ctx.Result(200, reply)
	}
}

func _P2P_ListListen0_HTTP_Handler(srv P2PHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListListenRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationP2PListListen)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListListen(ctx, req.(*ListListenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListListenReply)
		return ctx.Result(200, reply)
	}
}

type P2PHTTPClient interface {
	CloseListen(ctx context.Context, req *CloseListenRequest, opts ...http.CallOption) (rsp *CloseListenReply, err error)
	CreateForward(ctx context.Context, req *CreateForwardRequest, opts ...http.CallOption) (rsp *CreateForwardReply, err error)
	CreateListen(ctx context.Context, req *CreateListenRequest, opts ...http.CallOption) (rsp *CreateListenReply, err error)
	ListListen(ctx context.Context, req *ListListenRequest, opts ...http.CallOption) (rsp *ListListenReply, err error)
}

type P2PHTTPClientImpl struct {
	cc *http.Client
}

func NewP2PHTTPClient(client *http.Client) P2PHTTPClient {
	return &P2PHTTPClientImpl{client}
}

func (c *P2PHTTPClientImpl) CloseListen(ctx context.Context, in *CloseListenRequest, opts ...http.CallOption) (*CloseListenReply, error) {
	var out CloseListenReply
	pattern := "/v1/p2p/close"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationP2PCloseListen))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *P2PHTTPClientImpl) CreateForward(ctx context.Context, in *CreateForwardRequest, opts ...http.CallOption) (*CreateForwardReply, error) {
	var out CreateForwardReply
	pattern := "/v1/p2p/forward"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationP2PCreateForward))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *P2PHTTPClientImpl) CreateListen(ctx context.Context, in *CreateListenRequest, opts ...http.CallOption) (*CreateListenReply, error) {
	var out CreateListenReply
	pattern := "/v1/p2p/listen"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationP2PCreateListen))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *P2PHTTPClientImpl) ListListen(ctx context.Context, in *ListListenRequest, opts ...http.CallOption) (*ListListenReply, error) {
	var out ListListenReply
	pattern := "/v1/p2p/ls"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationP2PListListen))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
