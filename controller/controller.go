package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/M4tt1-Coder/business/portfolio_website/API_GO/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

//!Info!
//-> the layout for timedata types is: a in JSON bodies 2022-11-25T15:04:05Z

// helper functions
// func split(tosplit string, sep rune) []string {
// 	var fields []string

//main adminif
//-> 5f36a786-3846-11ee-9517-947caef5c0e1

// 	last := 0
// 	for i, c := range tosplit {
// 		if c == sep {
// 			// Found the separator, append a slice
// 			fields = append(fields, string(tosplit[last:i]))
// 			last = i + 1
// 		}
// 	}

// 	// Don't forget the last field
// 	fields = append(fields, string(tosplit[last:]))

// 	return fields
// }

//functions for the portfolio website
//-> partners

func CreatePartner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adminid, error := uuid.Parse(vars["adminid"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	decoder := json.NewDecoder(r.Body)
	var partner models.Partner
	err := decoder.Decode(&partner)
	if err != nil {
		log.Printf("Error decoding: %v", err)
	}
	partner = *partner.CreatePartner(adminid)
	res, _ := json.Marshal(partner)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllPartners(w http.ResponseWriter, r *http.Request) {
	partners := models.GetAllPartners()
	res, _ := json.Marshal(partners)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPartner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, error := uuid.Parse(vars["id"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	partner, _ := models.GetPartnerByID(id)
	res, _ := json.Marshal(partner)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeletePartner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, error := uuid.Parse(vars["id"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	adminid, error := uuid.Parse(vars["adminid"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	partner, _ := models.DeletePartnerByID(id, adminid)
	res, _ := json.Marshal(partner)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdatePartner(w http.ResponseWriter, r *http.Request) {
	//get all parameters
	//update with the new values
	//log the changes with admin
	decoder := json.NewDecoder(r.Body)
	var vars map[string]string
	err := decoder.Decode(&vars)
	if err != nil {
		log.Printf("Error decoding: %v", err)
	}
	//id
	id, error := uuid.Parse(vars["id"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	//name
	name := vars["name"]
	//websiteLink
	websiteLink := vars["websiteLink"]
	//sincewhen
	sinceWhen, err := time.Parse("2006-01-02 00:00:00", vars["sinceWhen"])
	if err != nil {
		log.Printf("Failed time : %v", err)
	}
	//address
	address := vars["address"]
	//phone
	telephoneNumber := vars["telephoneNumber"]
	//adminid
	adminid, error := uuid.Parse(vars["adminid"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	//update partner with the new values
	partner := models.UpdatePartnerByID(id, name, websiteLink, sinceWhen, address, telephoneNumber, adminid)
	res, _ := json.Marshal(partner)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//project functions

func CreateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adminid, error := uuid.Parse(vars["adminid"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	decoder := json.NewDecoder(r.Body)
	var project models.Project
	err := decoder.Decode(&project)
	if err != nil {
		log.Printf("Error decoding: %v", err)
	}
	project = *project.CreateProject(adminid)
	res, _ := json.Marshal(project)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects := models.GetAllProjects()
	res, _ := json.Marshal(projects)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, error := uuid.Parse(vars["id"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	project, _ := models.GetProjectByID(id)
	res, _ := json.Marshal(project)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, error := uuid.Parse(vars["id"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	adminid, error := uuid.Parse(vars["adminid"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	project := models.DeleteProjectByID(id, adminid)
	res, _ := json.Marshal(project)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var vars map[string]string
	err := decoder.Decode(&vars)
	if err != nil {
		log.Printf("Error decoding: %v", err)
	}
	//id
	id, error := uuid.Parse(vars["id"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	//name
	name := vars["name"]
	//link
	link := vars["link"]
	//description
	description := vars["description"]
	//participants
	participants := vars["participants"]
	//adminid
	adminid, error := uuid.Parse(vars["adminid"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	//update project with the new values
	project := models.UpdateProjectByID(id, name, link, description, participants, adminid)
	res, _ := json.Marshal(project)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//message functions

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var message models.Message
	err := decoder.Decode(&message)
	if err != nil {
		log.Printf("Error decoding: %v", err)
	}
	message = *message.CreateMessage()
	res, _ := json.Marshal(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	messages := models.GetAllMessages()
	res, _ := json.Marshal(messages)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, error := uuid.Parse(vars["id"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	message, _ := models.GetMessageByID(id)
	res, _ := json.Marshal(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, error := uuid.Parse(vars["id"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	adminid, error := uuid.Parse(vars["adminid"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	message := models.DeleteMessageByID(id, adminid)
	res, _ := json.Marshal(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var vars map[string]string
	err := decoder.Decode(&vars)
	if err != nil {
		log.Printf("Error decoding: %v", err)
	}
	//id
	id, error := uuid.Parse(vars["id"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	//content
	content := vars["content"]
	//subject
	subject := vars["subject"]
	//createdAt
	createdAt, err := time.Parse("2006-01-02 00:00:00", vars["createdAt"])
	if err != nil {
		log.Printf("Failed time : %v", err)
	}
	//emailAddress
	emailAddress := vars["emailAddress"]
	//senderName
	senderName := vars["senderName"]
	//gender
	gender := vars["gender"]
	//adminid
	adminid, error := uuid.Parse(vars["adminid"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	//update message with the new values
	message := models.UpdateMessageByID(id, content, subject, createdAt, emailAddress, senderName, gender, adminid)
	res, _ := json.Marshal(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//admin functions

// valid
func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var admin models.Admin
	err := decoder.Decode(&admin)
	if err != nil {
		log.Printf("Error decoding: %v", err)
	}
	vars := mux.Vars(r)
	adminid, error := uuid.Parse(vars["adminid"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}

	admin = *admin.CreateAdmin(adminid)
	res, _ := json.Marshal(admin)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllAdmins(w http.ResponseWriter, r *http.Request) {
	admins := models.GetAllAdmins()
	res, _ := json.Marshal(admins)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, error := uuid.Parse(vars["id"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	admin, _ := models.GetAdminByID(id)
	res, _ := json.Marshal(admin)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, error := uuid.Parse(vars["id"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	adminid, error := uuid.Parse(vars["adminid"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	admin := models.DeleteAdminByID(id, adminid)
	res, _ := json.Marshal(admin)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var vars map[string]string
	err := decoder.Decode(&vars)
	if err != nil {
		log.Printf("Error decoding: %v", err)
	}
	//id
	id, error := uuid.Parse(vars["id"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	//firstname
	firstName := vars["firstName"]
	//lastName
	lastName := vars["lastName"]
	//password
	Password := vars["password"]
	// //rights
	rights := vars["rights"]
	//emailaddress
	emailaddress := vars["emailAddress"]
	//adminid
	adminid, error := uuid.Parse(vars["adminid"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	//update admin with the new values
	admin := models.UpdateAdminbyID(id, firstName, lastName, Password, rights, adminid, emailaddress)
	res, _ := json.Marshal(admin)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateAdminLastTimeOnline(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("Error parsing: %v", err)
	}
	admin := models.LastTimeOnline(id)
	res, _ := json.Marshal(admin)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//contact functions

func GetContact(w http.ResponseWriter, r *http.Request) {
	contact := models.GetContact()
	res, _ := json.Marshal(contact)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//infoCard functions

func GetInfoCard(w http.ResponseWriter, r *http.Request) {
	infoCard := models.GetInfoCard()
	res, _ := json.Marshal(infoCard)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
