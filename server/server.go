package main

import (
	"context"
	"log"
	"net"

	"github.com/hahoangtu2k1/b5"
	"github.com/hahoangtu2k1/demo"
	"google.golang.org/grpc"
)

type server struct {
	demo.ServerUserServer
}

func main() {
	lis, err := net.Listen("tcp", "localhost:2001")

	if err != nil {
		log.Fatal("Err:", err)
	}

	s := grpc.NewServer()

	demo.RegisterServerUserServer(s, &server{})

	log.Println("Server is running...........")

	err = s.Serve(lis)

	if err != nil {
		log.Fatal("Err:", err)
	}
}

func (server) Create(ctx context.Context, req *demo.CreateRequest) (*demo.CreateResponse, error) {
	u := b5.ConvertPbUser(req.User)
	err := u.CreateUserServer()
	if err != nil {
		resp := &demo.CreateResponse{
			StatusCode: 1,
			Message:    "insert Error",
		}
		return resp, nil
	}
	resp := &demo.CreateResponse{
		StatusCode: 2,
		Message:    " Insert successfully",
	}
	return resp, nil
}

func (server) List(ctx context.Context, req *demo.GetListRequest) (*demo.GetListResponse, error) {
	data, err := b5.ListUserServer()
	if err != nil {
		return nil, err
	} else {
		list := []*demo.UserPartner{}
		for _, u := range data {
			conv := b5.ConvertUserPb(u)
			list = append(list, conv)
		}
		return &demo.GetListResponse{
			User: list,
		}, nil
	}
}
func (server) Update(ctx context.Context, req *demo.UpdateRequest) (*demo.UpdateResponse, error) {
	u := b5.ConvertPbUser(req.NewUser)
	err := u.UpdateServer()
	if err != nil {
		resp := &demo.UpdateResponse{
			StatusCode: 1,
			Message:    "update error",
		}
		return resp, nil
	} else {
		// demoClient(client)
		resp := &demo.UpdateResponse{
			StatusCode: 2,
			Message:    "update successfully",
		}
		return resp, nil
	}

}
