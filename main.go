/**
 * @Author: weiming02
 * @Date: 2023/2/3 14:42
 **/
package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/xuri/excelize/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"ledger-sync/model"
)

func main() {
	file, err := excelize.OpenFile("台账用户.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 按行读取全部
	rows, err := file.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	var ledgerUsers []*model.LedgerUser
	for index, row := range rows {
		if index != 0 {
			ledgerUsers = append(ledgerUsers, &model.LedgerUser{IsSuperAdmin: row[1], UserName: row[2], RealUserName: row[3], PassWord: row[4],
				RoleName: row[5], PhoneNumber: row[6], OrgName: row[7]})
		}
	}

	var authUsers []*model.AuthUser

	for _, user := range ledgerUsers {
		fmt.Println(*user)
		authUsers = append(authUsers, &model.AuthUser{Id: uuid.NewV4().String(), UserName: user.UserName, PassWord: user.PassWord, RealUserName: user.RealUserName, PhoneNo: user.PhoneNumber, OrganizationId: user.OrgName})
	}

	for _, user := range authUsers {
		fmt.Println(*user)
	}

	dsn := "root:3DcaFtr1ssY@tcp(10.172.198.67:3306)/auth2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	user := model.AuthUser{}
	db.First(&user)
	fmt.Println(user)

	// TODO 密码加密

}
