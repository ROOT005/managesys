package models

import (
	"github.com/jinzhu/gorm"
	"managesys/db"
	"time"
)

type Client struct {
	gorm.Model
	Operator string
	Name     string
	Count    string
	Level    string
	State    string
	Result   string
	Other    string
	//clieninfo
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
	AssetInfo     AssetInfo
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

func GetWeekInfo() (map[int]int, map[int]int) {
	var now = time.Now().Weekday().String()
	var numW time.Duration
	switch now {
	case "Monday":
		numW = 0
	case "Tuesday":
		numW = 1
	case "Wednesday":
		numW = 2
	case "Thursday":
		numW = 3
	case "Friday":
		numW = 4
	case "Saturday":
		numW = 5
	case "Sunday":
		numW = 6
	}
	info := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0}
	infoFin := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0}
	var clients [6][]*Client
	d, _ := time.ParseDuration("-24h")
	for i := int(numW); i > 0; i-- {
		db.DB.Where("created_at BETWEEN ? AND ?", time.Now().Add(d*time.Duration(i)), time.Now().Add(d*time.Duration((i-1)))).Find(&clients[i])
		info[i] = len(clients[i])
	}
	for i := int(numW); i > 0; i-- {
		db.DB.Where("created_at BETWEEN ? AND ? AND state=?", time.Now().Add(d*time.Duration(i)), time.Now().Add(d*time.Duration((i-1))), "签约").Find(&clients[i])
		infoFin[i] = len(clients[i])
	}
	return info, infoFin
}

func GetDayInfo() (map[string]int, map[string]int) {
	var clients []*User
	var clientsFin []*User
	d, _ := time.ParseDuration("-24h")
	db.DB.Where("created_at BETWEEN ? AND ?", time.Now().Add(d), time.Now()).Find(&clients)
	db.DB.Where("created_at BETWEEN ? AND ?", time.Now().Add(d), time.Now()).Find(&clientsFin)
	return nil, nil
}
