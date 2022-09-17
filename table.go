package main

type UserPartner struct {
	Id          string
	UserId      string `xorm:"pk"`
	PartnerId   string
	AliasUserId string
	Apps        map[string]int64
	Phone       string
	Created     int64
	Updated_at  int64
}
