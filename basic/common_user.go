package basic

type CommonUser struct {
	Name     string
	Mobile   string
	PassWord string
	Token    string
}

func (u *CommonUser) SetToken(token string) {
	u.Token = token
}
