// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: messageService.proto

package messageService

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for MessageService service

func NewMessageServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MessageService service

type MessageService interface {
	MessageList(ctx context.Context, in *DouyinMessageChatRequest, opts ...client.CallOption) (*DouyinMessageChatResponse, error)
	MessageAction(ctx context.Context, in *DouyinMessageActionRequest, opts ...client.CallOption) (*DouyinMessageActionResponse, error)
}

type messageService struct {
	c    client.Client
	name string
}

func NewMessageService(name string, c client.Client) MessageService {
	return &messageService{
		c:    c,
		name: name,
	}
}

func (c *messageService) MessageList(ctx context.Context, in *DouyinMessageChatRequest, opts ...client.CallOption) (*DouyinMessageChatResponse, error) {
	req := c.c.NewRequest(c.name, "MessageService.MessageList", in)
	out := new(DouyinMessageChatResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) MessageAction(ctx context.Context, in *DouyinMessageActionRequest, opts ...client.CallOption) (*DouyinMessageActionResponse, error) {
	req := c.c.NewRequest(c.name, "MessageService.MessageAction", in)
	out := new(DouyinMessageActionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MessageService service

type MessageServiceHandler interface {
	MessageList(context.Context, *DouyinMessageChatRequest, *DouyinMessageChatResponse) error
	MessageAction(context.Context, *DouyinMessageActionRequest, *DouyinMessageActionResponse) error
}

func RegisterMessageServiceHandler(s server.Server, hdlr MessageServiceHandler, opts ...server.HandlerOption) error {
	type messageService interface {
		MessageList(ctx context.Context, in *DouyinMessageChatRequest, out *DouyinMessageChatResponse) error
		MessageAction(ctx context.Context, in *DouyinMessageActionRequest, out *DouyinMessageActionResponse) error
	}
	type MessageService struct {
		messageService
	}
	h := &messageServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MessageService{h}, opts...))
}

type messageServiceHandler struct {
	MessageServiceHandler
}

func (h *messageServiceHandler) MessageList(ctx context.Context, in *DouyinMessageChatRequest, out *DouyinMessageChatResponse) error {
	return h.MessageServiceHandler.MessageList(ctx, in, out)
}

func (h *messageServiceHandler) MessageAction(ctx context.Context, in *DouyinMessageActionRequest, out *DouyinMessageActionResponse) error {
	return h.MessageServiceHandler.MessageAction(ctx, in, out)
}
