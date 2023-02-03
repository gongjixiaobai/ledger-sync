package model

type AuthUser struct {
	Id             string `gorm:"primaryKey"`
	UserName       string `column:"user_name"`
	PassWord       string `column:"user_passwd"`
	RealUserName   string `column:"user_real_name"`
	PhoneNo        string `column:"phone_no"`
	OrganizationId string `column:"organization_id"`
}
