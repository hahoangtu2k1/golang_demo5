package main

import (
	"context"
	"log"
	"time"

	"github.com/hahoangtu2k1/demo"
	"github.com/rs/xid"
	"google.golang.org/grpc"
)

func main() {

	list, err := grpc.Dial("localhost:2001", grpc.WithInsecure())

	if err != nil {
		log.Fatal("err: ", err)
	}

	defer list.Close()

	log.Println("Client server is running...")

	client := demo.NewServerUserClient(list)
	// InsertUser(client)
	// ListUser(client)
	UserUpdateServer(client)

}

func InsertUser(cli demo.ServerUserClient) {
	resp, err := cli.Create(context.Background(), &demo.CreateRequest{
		User: &demo.UserPartner{
			Id:          xid.New().String(),
			UserId:      "425",
			PartnerId:   "42",
			AliasUserId: "42",
			Apps: map[string]int64{
				"age": 20,
			},
			Phone:     "01236547825",
			Created:   time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
	})
	if err != nil {
		log.Println(err)
		log.Println("resp: ", resp.StatusCode, " ", resp.Message)
		log.Fatalf("Loi insert %v", err)
	} else {
		log.Println("Insert request successfully")
	}
}
func ListUser(cli demo.ServerUserClient) {
	read, err := cli.List(context.Background(), &demo.GetListRequest{})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(read)
}
func UserUpdateServer(cli demo.ServerUserClient) {
	_, err := cli.Update(context.Background(), &demo.UpdateRequest{
		NewUser: &demo.UserPartner{
			UserId:      "13",
			PartnerId:   "43684",
			AliasUserId: "43684",
			Phone:       "09874563218889",
		},
	})
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("update complete")
	}

}
