// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: services/adCampaign/proto/adCampaign.proto

/*
Package adcampaign is a generated protocol buffer package.

It is generated from these files:
	services/adCampaign/proto/adCampaign.proto

It has these top-level messages:
	Request
	Response
	AdCampaign
	Platforms
	Platform
	TargetAudiance
	Creatives
	Insights
*/
package adcampaign

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AdCampaignService service

type AdCampaignService interface {
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type adCampaignService struct {
	c    client.Client
	name string
}

func NewAdCampaignService(name string, c client.Client) AdCampaignService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "adcampaign"
	}
	return &adCampaignService{
		c:    c,
		name: name,
	}
}

func (c *adCampaignService) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AdCampaignService.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AdCampaignService service

type AdCampaignServiceHandler interface {
	GetAll(context.Context, *Request, *Response) error
}

func RegisterAdCampaignServiceHandler(s server.Server, hdlr AdCampaignServiceHandler, opts ...server.HandlerOption) {
	type adCampaignService interface {
		GetAll(ctx context.Context, in *Request, out *Response) error
	}
	type AdCampaignService struct {
		adCampaignService
	}
	h := &adCampaignServiceHandler{hdlr}
	s.Handle(s.NewHandler(&AdCampaignService{h}, opts...))
}

type adCampaignServiceHandler struct {
	AdCampaignServiceHandler
}

func (h *adCampaignServiceHandler) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.AdCampaignServiceHandler.GetAll(ctx, in, out)
}
