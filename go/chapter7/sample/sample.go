package sample

// ユーザを登録し、その登録されたユーザを管理するシステム

// ビジネスルール
// - ユーザのメールアドレスに自社ドメインが含まれている場合、
// そのユーザの種類(Type)を「従業員(employee)」として登録し、そうでない場合は「顧客(customer)」
// - 登録されたユーザは、従業員から顧客、顧客から、従業員に変更可
// - メールアドレスの変更ができたら、変更を外部サービスに通知する

import (
	"fmt"
	"strings"
)

//
type User struct {
	UserID int
	Email  string
	Type   UserType
}

func (u *User) ChangeEmail(userID int, newEmail string) {
	db := Database{}
	mb := MessageBus{bus: bus{}}
	data := db.GetUserByID(userID)
	email := data[0].(string)
	typ := data[1].(UserType)

	if email == newEmail {
		return
	}

	companyData := db.GetCompany()
	companyDomainName := companyData[0].(string)
	numberOfEmployees := companyData[1].(int)

	emailDomain := newEmail[strings.Index(newEmail, "@")+1:]
	isEmailCorporate := emailDomain == companyDomainName

	newType := Customer
	if isEmailCorporate {
		newType = Employee
	}

	if typ != newType {
		delta := -1
		if newType == Employee {
			delta = 1
		}
		newNumber := numberOfEmployees + delta
		db.SaveCompany(newNumber)
	}

	u.Email = newEmail
	u.Type = newType

	db.SaveUser(u)
	mb.SendEmailChanedMessage(userID, newEmail)
}

type UserType int

const (
	Customer UserType = iota + 1
	Employee
)

type Database struct {
}

func (db Database) GetUserByID(userID int) []interface{} {
	return []interface{}{"test@test.com", Customer}
}

func (db Database) GetUserByEmail(email string) *User {
	return nil
}

func (db Database) SaveUser(user *User) {
}

func (db Database) GetCompany() []interface{} {
	return []interface{}{"GO株式会社", 2}
}

func (db Database) SaveCompany(newNumber int) {
}

type MessageBus struct {
	bus Bus
}

func (mb MessageBus) SendEmailChanedMessage(userID int, newEmail string) {
	mb.bus.Send(fmt.Sprintf("Subject: User; Type: EMAIL CHANGED; Id: %d; NewEmail: %s", userID, newEmail))
}

type Bus interface {
	Send(message string)
}

func (b bus) Send(message string) {

}

type bus struct {
}
