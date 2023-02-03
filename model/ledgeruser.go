package model

type LedgerUser struct {
	IsSuperAdmin string `json:"isSuperAdmin"`
	UserName     string `json:"userName"`
	RealUserName string `json:"realUserName"`
	PassWord     string `json:"passWord"`
	RoleName     string `json:"roleName"`
	PhoneNumber  string `json:"phoneNumber"`
	OrgName      string `json:"orgName"`
}
