package entitie

type User struct {
	ID   int
	Name string
	UserName string
}

func NewUser(ID int, Name string, UserName string) *User {
	return &User{ID:1,Name:Name,UserName:UserName}
}

func (u *User) GetID() int {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetUserName() string {
	return u.UserName
}

func(u *User) SetID(ID int){
	u.ID = ID
}

func(u *User) SetName(Name string){
	u.Name = Name
}

func(u *User) SetUserName(UserName string){
	u.UserName = UserName
}
