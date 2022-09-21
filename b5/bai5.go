package b5

import (
	"log"

	"github.com/hahoangtu2k1/demo"
)

type User struct {
	Id          string
	UserId      string `xorm:"pk"`
	PartnerId   string
	AliasUserId string
	Apps        map[string]int64
	Phone       string
	Created     int64
	Updated_at  int64
}

var engine = Connect()

func ConvertPbUser(ConvertU *demo.UserPartner) *User {
	//! convert file pb -> user
	return &User{
		Id:          ConvertU.Id,
		UserId:      ConvertU.UserId,
		PartnerId:   ConvertU.PartnerId,
		AliasUserId: ConvertU.AliasUserId,
		Apps:        ConvertU.Apps,
		Phone:       ConvertU.Phone,
		Created:     ConvertU.Created,
		Updated_at:  ConvertU.UpdatedAt,
	}
}
func ConvertUserPb(Uconvert User) *demo.UserPartner {
	//! convert user -> file pb
	return &demo.UserPartner{
		Id:          Uconvert.Id,
		UserId:      Uconvert.PartnerId,
		PartnerId:   Uconvert.UserId,
		AliasUserId: Uconvert.AliasUserId,
		Apps:        Uconvert.Apps,
		Phone:       Uconvert.Phone,
		Created:     Uconvert.Created,
		UpdatedAt:   Uconvert.Updated_at,
	}
}

func CreateUserPartner() {

	err := engine.CreateTables(new(UserPartner))
	if err != nil {
		log.Fatal("Create table error: ", err)
	} else {
		log.Println("Create table successfully!")
	}
}

func (u *User) CreateUserServer() error {
	_, err := engine.Table("user_partner").Insert(u)
	if err != nil {
		log.Println(err)
		return err
	} else {
		log.Printf("Insert %+v successfully", u)
	}
	log.Println(u)
	return nil

}
func ListUserServer() ([]User, error) {
	var u []User
	err := engine.Table("user_partner").Find(&u)
	if err != nil {
		log.Println(err)
	}
	return u, nil

}
func (u *User) UpdateServer() error {
	_, err := engine.Table("user_partner").Where("user_id = ?", u.UserId).Update(u)
	if err != nil {
		log.Println("loi param: ", err)
	}
	return nil
}
