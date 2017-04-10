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
	RegistTime   string
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

	var clients []*Client
	var clientsFin []*Client
	d, _ := time.ParseDuration("-8h")
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	for i := 0; i < int(numW)+1; i++ {
		db.DB.Where("created_at BETWEEN ? AND ?", t.Add(3*d*time.Duration(i+1)), t.Add(3*d*time.Duration(i))).Find(&clients)
		info[i] = len(clients)
	}
	for i := 0; i < int(numW)+1; i++ {
		db.DB.Where("created_at BETWEEN ? AND ? AND state = ?", t.Add(3*d*time.Duration(i+1)), t.Add(3*d*time.Duration(i)), "签约").Find(&clientsFin)
		infoFin[i] = len(clientsFin)
	}
	return info, infoFin
}

func GetDayInfo() (map[string]int, map[string]int) {
	d, _ := time.ParseDuration("-24h")
	//查找用户
	var users []*User
	db.DB.Not("role = ?", "超级管理员").Find(&users)
	clientAccount := make(map[string]int)
	clientFinAccount := make(map[string]int)
	for i := 0; i < len(users); i++ {
		clientAccount[users[i].Name] = 0
		clientFinAccount[users[i].Name] = 0
	}
	//查找客户
	var clients []*Client
	var clientsFin []*Client

	db.DB.Where("created_at > ?", time.Now().Add(d)).Find(&clients)
	db.DB.Where("updated_at > ? AND state = ?", time.Now().Add(d), "签约").Or("created_at > ? AND state = ?", time.Now().Add(d), "签约").Find(&clientsFin)
	for i := 0; i < len(clients); i++ {
		for k, _ := range clientAccount {
			if k == clients[i].Operator {
				clientAccount[k] += 1
			}
		}
	}
	for i := 0; i < len(clientsFin); i++ {
		for k, _ := range clientFinAccount {
			if k == clientsFin[i].Operator {
				clientFinAccount[k] += 1
			}
		}
	}

	return clientAccount, clientFinAccount
}
