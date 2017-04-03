package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Client struct {
	gorm.Model
	Operator   string
	Name       string
	Count      string
	Level      string
	State      string
	Result     string
	Other      string
	ClientInfo ClientInfo
	AssetInfo  AssetInfo
}

//客户信息
type ClientInfo struct {
	gorm.Model
	ClientID      uint
	Gender        string
	Age           uint
	PhoneNum      string
	Address       string
	MaritalStatuc string
	Company       string
	CreditCheck   string
	CreCheckCount string
	GetTime       time.Time
	Balance       string
	CredCard      string
	PayAc         string
	MaxAc         string
	IdNum         string
	DetailedList  string
	CurrentAdd    string
	CompanyAdd    string
}

//资产信息
type AssetInfo struct {
	gorm.Model
	ClientID        uint
	FullHouse       []FullHouse
	MortgageHouse   []MortgageHouse
	FullCar         []FullCar
	MortgageCar     []MortgageCar
	InsurancePolicy []InsurancePolicy
	AccuFound       []AccuFound
	SocialSecurity  []SocialSecurity
	Salary          []Salary
	BusinessLoan    []BusinessLoan
}

//全款房
type FullHouse struct {
	AssetInfoID  uint
	Value        string
	Area         string
	Paytime      string
	Location     string
	DoublePolicy string
}

//按揭房
type MortgageHouse struct {
	AssetInfoID uint
	FirstPay    string
	MouthPay    string
	LimitTime   string
	Location    string
	BankName    string
}

//全款车
type FullCar struct {
	AssetInfoID uint
	Value       string
	PayDay      string
	Certificate string
	Policy      string
	Key         string
}

//按揭车
type MortgageCar struct {
	AssetInfoID uint
	FirstPay    string
	PayMethod   string
	LimitTime   string
	MouthPay    string
	PayCount    string
}

//保单
type InsurancePolicy struct {
	AssetInfoID uint
	Type        string
	PayTime     string
	PayCount    string
	PayWay      string
}

//公积金
type AccuFound struct {
	AssetInfoID uint
	PayAccount  string
	PayRate     string
	PayTime     string
	PayWay      string
}

//社保
type SocialSecurity struct {
	AssetInfoID uint
	PayAccount  string
	PayRate     string
	PayTime     string
	PayWay      string
}

//工资
type Salary struct {
	AssetInfoID uint
	GetWay      string
	Account     string
}

//生意贷
type BusinessLoan struct {
	AssetInfoID  uint
	Licence      bool
	RegistTime   time.Time
	Occupancy    string
	FamKnow      bool
	DetailedList string
}

func GetInfo() string {
	fmt.Println(time.Now().Weekday())
	return time.Now().String()
}
