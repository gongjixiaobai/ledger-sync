package model

type AuthUser struct {
	Id             string `gorm:"primaryKey"`
	UserName       string `column:"user_name"`
	PassWord       string `column:"pass_word"`
	RealUserName   string `column:"real_user_name"`
	PhoneNo        string `column:"phone_no"`
	OrganizationId string `column:"organization_id"`
}
