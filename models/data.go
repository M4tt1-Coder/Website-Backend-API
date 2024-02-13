package models

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"

	"github.com/M4tt1-Coder/business/portfolio_website/API_GO/dbHandler"

	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

var (
	envs, _ = godotenv.Read(".env")
)

// global database instance in this file
var db *gorm.DB

// models
type Contact struct {
	TelephoneNumber string `gorm:"TelephoneNumber" json:"telephoneNumber"`
	EmailAddress    string `gorm:"EmailAddress" json:"emailAddress"`
	Address         string `gorm:"Address" json:"address"`
}

type Partner struct {
	//gorm.Model
	Identifier      uuid.UUID `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"Name" json:"name"`
	WebsiteLink     string    `gorm:"WebsiteLink" json:"websiteLink"`
	SinceWhen       time.Time `gorm:"SinceWhen" json:"sinceWhen"`
	Address         string    `gorm:"Address" json:"address"`
	TelephoneNumber string    `gorm:"TelephoneNumber" json:"telephoneNumber"`
}

type Project struct {
	//gorm.Model
	Identifier   uuid.UUID `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"Name" json:"name"`
	Link         string    `gorm:"Link" json:"link"` //probably a github link
	Description  string    `gorm:"Description" json:"description"`
	Participants string    `grom:"Participants" json:"participants"`
}

type InfoCard struct {
	//gorm.Model
	Name           string `gorm:"Name" json:"name"`
	Age            int    `gorm:"Age" json:"age"`
	CurrentJob     string `gorm:"CurrentJob" json:"currentJob"`
	CurrentCompany string `gorm:"CurrentCompany" json:"currentCompany"`
	Summary        string `gorm:"Summary" json:"summary"`
	Address        string `gorm:"Address" json:"address"`
}

type Message struct {
	//gorm.Model
	Identifier   uuid.UUID `gorm:"primaryKey" json:"id"`
	Content      string    `gorm:"Content" json:"content"`
	Subject      string    `gorm:"Subject" json:"subject"`
	CreatedAt    time.Time `gorm:"CreatedAt" json:"createdAt"`
	EmailAddress string    `gorm:"EmailAddress" json:"emailAddress"`
	SenderName   string    `gorm:"SenderName" json:"senderName"`
	Gender       string    `gorm:"Gender" json:"gender"`
}

type Admin struct {
	//gorm.Model
	Identifier     uuid.UUID `gorm:"primaryKey" json:"id"`
	FirstName      string    `gorm:"FirstName" json:"firstName"`
	LastName       string    `gorm:"LastName" json:"lastName"`
	Password       string    `gorm:"Password" json:"password"`
	Deletable      bool      `gorm:"Deletable" json:"deletable"`
	Rights         string    `gorm:"Rights" json:"rights"`
	LastTimeOnline time.Time `gorm:"LastTimeOnline" json:"lastTimeOnline"`
	EmailAddress   string    `gorm:"EmailAddress" json:"emailAddress"`
}

type Log struct {
	//gorm.Model
	Identifier uint      `gorm:"primaryKey"`
	Time       time.Time `gorm:"Time"`
	Admin      uuid.UUID `gorm:"AdminID"`
	Changes    string    `gorm:"Changes"` //which information was changed
	Where      string    `gorm:"Where"`   //table name
}

// initialize the database
func init() {
	dbHandler.Connect()
	db = dbHandler.GetDB()
	//set up database tables
	// db.AutoMigrate(
	// 	&Contact{},
	// 	&Partner{},
	// 	&Project{},
	// 	&InfoCard{},
	// 	&Message{},
	// 	&Admin{},
	// 	&Log{})
	db.AutoMigrate(&Contact{})
	db.AutoMigrate(&Partner{})
	db.AutoMigrate(&Project{})
	db.AutoMigrate(&InfoCard{})
	db.AutoMigrate(&Message{})
	db.AutoMigrate(&Admin{})
	db.AutoMigrate(&Log{})
	log.Println("Database initialized")
}

// Partner functions
func (p *Partner) CreatePartner(adminID *uuid.UUID) *Partner {

	//db.NewRecord(p)
	db.Create(p)

	//log what happend
	builder := strings.Builder{}
	builder.WriteString("Created a new Partner||" + "id: " + p.Identifier.String() + "||name: " + p.Name + "||websiteLink: " + p.WebsiteLink + "||sinceWhen: " + p.SinceWhen.String() + "||address: " + p.Address + "||telephoneNumber" + p.TelephoneNumber)
	//changes := "Created a new Partner||" + "id: " + p.Identifier.String() + "||name: " + p.Name + "||websiteLink: " + p.WebsiteLink + "||sinceWhen: " + p.SinceWhen.String() + "||address: " + p.Address + "||telephoneNumber" + p.TelephoneNumber
	CreateLog(adminID,
		builder.String(),
		"partner table",
	)

	return p
}

func GetAllPartners() []Partner {
	var partners []Partner
	db.Find(&partners)
	return partners
}

func GetPartnerByID(id *uuid.UUID) (*Partner, *gorm.DB) {
	var Partner Partner
	db := db.Find(&Partner, "identifier = ?", id)
	return &Partner, db
}

func DeletePartnerByID(id *uuid.UUID, adminID *uuid.UUID) (Partner, *gorm.DB) {

	var Partner Partner
	//maybe use gormodel id for deleting
	db := db.Where("identifier =?", id).Delete(Partner)

	CreateLog(adminID,
		"partner deleted||id: "+id.String(),
		"partner table",
	)

	return Partner, db
}

func UpdatePartnerByID(
	id *uuid.UUID,
	name string,
	websiteLink string,
	sinceWhen time.Time,
	address string,
	telephoneNumber string,
	adminID *uuid.UUID) *Partner {
	//get changes that were made

	builder := strings.Builder{}
	//var changes string
	p, _ := GetPartnerByID(id)
	if name != "" {
		//changes += "name:" + p.Name + " -> " + name + "||"
		builder.WriteString("name:" + p.Name + " -> " + name + "||")
		p.Name = name
	}
	if websiteLink != "" {
		//changes += "websiteLink:" + p.WebsiteLink + " -> " + websiteLink + "||"
		builder.WriteString("websiteLink:" + p.WebsiteLink + " -> " + websiteLink + "||")
		p.WebsiteLink = websiteLink
	}
	if !sinceWhen.IsZero() { //check how you can make sure sinceWhen id here zero
		//changes += "sinceWhen: " + p.SinceWhen.Format("2006-01-02 00:00:00") + " -> " + sinceWhen.Format("2006-01-02 00:00:00") + "||"
		builder.WriteString("sinceWhen: " + p.SinceWhen.Format("2006-01-02 00:00:00") + " -> " + sinceWhen.Format("2006-01-02 00:00:00") + "||")
		p.SinceWhen = sinceWhen
	}
	if address != "" {
		//changes += "address: " + p.Address + " -> " + address + "||"
		builder.WriteString("address: " + p.Address + " -> " + address + "||")
		p.Address = address
	}
	if telephoneNumber != "" {
		//changes = "telephoneNumber:" + p.TelephoneNumber + " -> " + telephoneNumber + "||"
		builder.WriteString("telephoneNumber:" + p.TelephoneNumber + " -> " + telephoneNumber + "||")
		p.TelephoneNumber = telephoneNumber
	}
	db.Save(&p)

	CreateLog(adminID, builder.String(), "partner table")

	return p
}

//Project functions

func (p *Project) CreateProject(adminId *uuid.UUID) *Project {
	//db.NewRecord(p)
	db.Create(&p)

	//log what happend
	builder := strings.Builder{}
	builder.WriteString("Created a new Project||" + "id: " + p.Identifier.String() + "||name: " + p.Name + "||link: " + p.Link + "||description: " + p.Description + "||participants: " + p.Participants)
	CreateLog(adminId,
		builder.String(),
		"Projects table",
	)

	return p
}

func GetAllProjects() []Project {
	var Projects []Project
	db.Find(&Projects)
	return Projects
}

func GetProjectByID(id *uuid.UUID) (*Project, *gorm.DB) {
	var Project Project
	db := db.Where("identifier =?", id).First(&Project)
	return &Project, db
}

func DeleteProjectByID(id *uuid.UUID, adminID *uuid.UUID) Project {

	var Project Project
	db.Where("identifier = ?", id).Delete(Project)

	CreateLog(adminID,
		"project deleted||id: "+id.String(),
		"Projects table",
	)

	return Project
}

func UpdateProjectByID(
	id *uuid.UUID,
	name string,
	link string,
	description string,
	participants string,
	adminID *uuid.UUID,
) *Project {
	builder := strings.Builder{}
	//var changes string
	p, _ := GetProjectByID(id)
	if name != "" {
		builder.WriteString("name:" + p.Name + " -> " + name + "||")
		p.Name = name
	}
	if link != "" {
		builder.WriteString("link: " + p.Link + " -> " + link + "||")
		p.Link = link
	}
	if description != "" {
		builder.WriteString("description: " + p.Description + " -> " + description + "||")
		p.Description = description
	}
	if participants != "" {
		builder.WriteString("participants: " + p.Participants + " -> " + participants + "||")
		p.Participants = participants
	}
	db.Save(&p)

	CreateLog(adminID, builder.String(), "Projects table")

	return p
}

//Message functions

func (m *Message) CreateMessage() *Message {
	//db.NewRecord(m)
	db.Create(m)
	return m
}

func GetAllMessages() []Message {
	var Messages []Message
	db.Find(&Messages)
	return Messages
}

func GetMessageByID(id *uuid.UUID) (*Message, *gorm.DB) {
	var Message Message
	db := db.Find(&Message, "identifier = ?", id)
	return &Message, db
}

func DeleteMessageByID(id *uuid.UUID, adminID *uuid.UUID) Message {
	var Message Message
	db := db.Where("identifier =?", id).Delete(Message)
	if db.Error != nil {
		log.Printf("Error deleting message")
	}

	CreateLog(adminID,
		"message deleted"+"||"+"id: "+id.String(),
		"message table",
	)

	return Message
}

func UpdateMessageByID(
	id *uuid.UUID,
	content string,
	subject string,
	createdAt time.Time,
	emailAddress string,
	senderName string,
	gender string,
	adminId *uuid.UUID,
) *Message {
	//var changes string
	builder := strings.Builder{}
	m, _ := GetMessageByID(id)
	if content != "" {
		builder.WriteString("content: " + m.Content + " -> " + content + "||")
		m.Content = content
	}
	if subject != "" {
		builder.WriteString("subject: " + m.Subject + " -> " + subject + "||")
		m.Subject = subject
	}
	if !createdAt.IsZero() { //check how you can make sure sinceWhen id here zero
		builder.WriteString("createdAt: " + m.CreatedAt.Format("2006-01-01 00:00:00") + " -> " + createdAt.Format("2006-01-01 00:00:00") + "||")
		m.CreatedAt = createdAt
	}
	if emailAddress != "" {
		builder.WriteString("emailAddress:" + m.EmailAddress + " -> " + emailAddress + "||")
		m.EmailAddress = emailAddress
	}
	if senderName != "" {
		builder.WriteString("senderName: " + m.SenderName + " -> " + senderName + "||")
		m.SenderName = senderName
	}
	if gender != "" {
		builder.WriteString("gender: " + m.Gender + " -> " + gender + "||")
		m.Gender = gender
	}

	db.Save(&m)

	CreateLog(adminId, builder.String(), "message table")

	return m
}

//Admin functions

func (a *Admin) CreateAdmin(adminID *uuid.UUID) *Admin {
	//db.NewRecord(a)
	db.Create(a)

	builder := strings.Builder{}

	builder.WriteString("Created a new Admin||" + "id: " + a.Identifier.String() + "||" + "FirstName: " + a.FirstName + "||LastName: " + a.LastName + "||" + "Password: " + a.Password + "||" + "Deletable: " + strconv.FormatBool(a.Deletable) + "||" + "Rights: " + a.Rights)
	CreateLog(adminID,
		builder.String(),
		"admin table",
	)

	return a
}

func GetAllAdmins() []Admin {
	var Admins []Admin
	db.Find(&Admins)
	return Admins
}

func GetAdminByID(id *uuid.UUID) (*Admin, *gorm.DB) {
	var Admin Admin
	//db = db.Where("identifier =?", id).Find(&Admin)
	db := db.Find(&Admin, "identifier = ?", id)
	if Admin.Identifier == uuid.Nil {
		log.Print("admin not found")
	}
	return &Admin, db
}

func DeleteAdminByID(id *uuid.UUID, adminId *uuid.UUID) Admin {
	var Admin Admin
	db := db.Where("identifier = ?", id).Delete(&Admin)
	if db.Error != nil {
		log.Printf("Error deleting admin %v", db.Error)
	}
	CreateLog(adminId,
		"admin deleted"+" || "+"id: "+id.String(),
		"admin table",
	)

	return Admin
}

func UpdateAdminbyID(
	id *uuid.UUID,
	FirstName string,
	LastName string,
	Password string,
	Rights string,
	adminId *uuid.UUID,
	EmailAddress string,
) *Admin {
	//var changes string
	builder := strings.Builder{}

	a, _ := GetAdminByID(id)
	if FirstName != "" {
		builder.WriteString("FirstName: " + a.FirstName + " -> " + FirstName + "||")
		a.FirstName = FirstName
	}
	if LastName != "" {
		builder.WriteString("LastName: " + a.LastName + " -> " + LastName + "||")
		a.LastName = LastName
	}
	if Password != "" {
		builder.WriteString("Password: " + a.Password + " -> " + Password + "||")
		a.Password = Password
	}
	if Rights != "" {
		builder.WriteString("Rights: " + a.Rights + " -> " + Rights + "||")
		a.Rights = Rights
	}
	if EmailAddress != "" {
		builder.WriteString("EmailAddress: " + a.EmailAddress + "->" + EmailAddress + "||")
		a.EmailAddress = EmailAddress
	}

	db.Save(&a)

	CreateLog(adminId, builder.String(), "admin table")

	return a
}

func LastTimeOnline(id *uuid.UUID) Admin {
	admin, _ := GetAdminByID(id)
	admin.LastTimeOnline = time.Now()
	db.Save(&admin)
	return *admin
}

//Contact functions

func GetContact() Contact {
	var Contact []Contact
	db.Find(&Contact)
	return Contact[0]
}

//InfoCard functions

func GetInfoCard() InfoCard {
	var InfoCard []InfoCard
	db.Find(&InfoCard)
	return InfoCard[0]
}

//Log functions

func CreateLog(AdminID *uuid.UUID, changes string, where string) {
	//instanciate db
	if AdminID.String() == envs["WRONG_ADMIN_ID_FORMAT"] || len(changes) == 0 || len(where) == 0 {
		log.Printf("Failed to log an activity due to an invalid admin ID: %v", AdminID)
		return
	}
	Admin, _ := GetAdminByID(AdminID)
	log := Log{
		Time:    time.Now(),
		Admin:   Admin.Identifier,
		Changes: changes,
		Where:   where,
	}

	db.Create(&log)
}

//try to call in the controller functions
//use jinzhu gorm instead of gorm.io/gorm
