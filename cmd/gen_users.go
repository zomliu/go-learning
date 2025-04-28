package main

import "fmt"

func main() {
	ru := new(RealUser)
	actionOne(ru)
}

type UserInterface interface {
	GetUser() string
	SetUserName(string) string
}

type UserDealer struct{}

func (u *UserDealer) GetUser() string {
	return "getUser from UserDealer"
}

func (u *UserDealer) SetUserName(name string) string {
	fmt.Println("setUserName from UserDealer")
	return "OK from UserDealer"
}

type RealUser struct {
	UserDealer
}

func (r *RealUser) GetUser() string {
	return "getUser from RealUser"
}

func actionOne(u UserInterface) {
	fmt.Println(u.GetUser())
}
