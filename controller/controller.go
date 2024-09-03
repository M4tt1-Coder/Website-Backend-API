package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/M4tt1-Coder/business/portfolio_website/API_GO/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

var (
	envs, _ = godotenv.Read(".env")
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

// Mux function
//
// Creates a new partner entity in the database.
//
// Needs an admin id, of the admin that created the partner object.
// In the HTTP-body of the request needs to be passed a JSON object containing a valid datatstructure.
//
// HTTP request body
//
//	{
//			id: 'someID UUID',
//			name: "Paul",
//			websiteLink: "https://example.com",
//			sinceWhen: "2015-00-00",
//			address: "Your Street 1, 07777 Citi",
//			telephoneNumber: "+123456789"
//	}
func CreatePartner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // get the routes variables
	adminid, error := uuid.Parse(vars["adminid"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	if adminid.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var partner models.Partner
	err := decoder.Decode(&partner)
	if err != nil {
		log.Printf("Error decoding: %v", err)
	}
	partner = *partner.CreatePartner(&adminid)
	res, _ := json.Marshal(partner)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// MUX controller functions
//
// Doesn't need any route variables or http body.
//
// Loads all partner object from teh database.
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
	if id.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		w.WriteHeader(http.StatusNotAcceptable)
		partner := &models.Partner{}
		res, err := json.Marshal(partner)
		if err != nil {
			log.Printf("Error marshalling: %v", err)
		}
		w.Write(res)
		return
	}
	partner, _ := models.GetPartnerByID(&id)
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
	if id.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	adminid, error := uuid.Parse(vars["adminid"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	if adminid.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	partner, _ := models.DeletePartnerByID(&id, &adminid)
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
	if id.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		w.WriteHeader(http.StatusNotAcceptable)
		return
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
	if adminid.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	//update partner with the new values
	partner := models.UpdatePartnerByID(&id, name, websiteLink, sinceWhen, address, telephoneNumber, &adminid)
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
	if adminid.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		message := "The project UUID was not provided correctly!"
		res, err := json.Marshal(message)
		if err != nil {
			log.Printf("Error marshalling error message: %v", err)
		}
		w.WriteHeader(http.StatusNotAcceptable)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
	decoder := json.NewDecoder(r.Body)
	var project models.Project
	err := decoder.Decode(&project)
	if err != nil {
		log.Printf("Error decoding: %v", err)
	}
	project = *project.CreateProject(&adminid)
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
	if id.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		project := &models.Project{}
		res, err := json.Marshal(project)
		if err != nil {
			log.Printf("Marshal failed: %v", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
	project, _ := models.GetProjectByID(&id)
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
	if id.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		w.WriteHeader(http.StatusNotAcceptable)
		log.Printf("The id of the project is not in the right format: %v", id.String())
		return
	}
	adminid, error := uuid.Parse(vars["adminid"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	if adminid.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		w.WriteHeader(http.StatusNotAcceptable)
		log.Printf("No valid admin ID has beed entered: %v", adminid.String())
		return
	}
	project := models.DeleteProjectByID(&id, &adminid)
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
	if id.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		log.Printf("The new ID of the project is incorrect!")
		w.WriteHeader(http.StatusNotAcceptable)
		return
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
	if adminid.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		log.Printf("Passed admin ID isn't in the right format!")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	//update project with the new values
	project := models.UpdateProjectByID(&id, name, link, description, participants, &adminid)
	res, err := json.Marshal(project)
	if err != nil {
		log.Printf("Failed to marshl the project: %v", err)
	}
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
	res, err := json.Marshal(message)
	if err != nil {
		log.Printf("Failed to marshl the message: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	messages := models.GetAllMessages()
	res, err := json.Marshal(messages)
	if err != nil {
		log.Printf("Couldn't marshal the messages!")
	}
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
	if id.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		log.Printf("The message has a wrong ID: %v", id)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	message, _ := models.GetMessageByID(&id)
	res, err := json.Marshal(message)
	if err != nil {
		log.Printf("Couldn't parse the message: %v", err)
	}
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
	if id.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		log.Printf("Deleting message failed due to invalid message ID: %v", id)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	adminid, error := uuid.Parse(vars["adminid"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	if adminid.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		log.Printf("The entered admin ID is invalid: %v", adminid)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	message := models.DeleteMessageByID(&id, &adminid)
	res, err := json.Marshal(message)
	if err != nil {
		log.Printf("Tried to marshal message: %v", err)
	}
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
	if id.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		log.Printf("The message ID is in a wrong format: %v", id)
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
	if adminid.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		log.Printf("The admin ID is incorrect: %v", adminid)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	//update message with the new values
	message := models.UpdateMessageByID(&id, content, subject, createdAt, emailAddress, senderName, gender, &adminid)
	res, err := json.Marshal(message)
	if err != nil {
		log.Printf("Failed to marshal the message: %v", err)
	}
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
	if adminid.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		log.Printf("The admin ID is incorrect: %v", adminid)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	admin = *admin.CreateAdmin(&adminid)
	res, err := json.Marshal(admin)
	if err != nil {
		log.Printf("Couldn't marshal the created admin: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllAdmins(w http.ResponseWriter, r *http.Request) {
	admins := models.GetAllAdmins()
	res, err := json.Marshal(admins)
	if err != nil {
		log.Printf("Try to marshal admisn didn't succeed: %v", err)
	}
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
	if id.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		log.Printf("Attempting to get admin failed due to wrong admin ID format: %v", id)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	admin, _ := models.GetAdminByID(&id)
	res, err := json.Marshal(admin)
	if err != nil {
		log.Printf("Failed with marshalling the admin: %v", err)
	}
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
	if id.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		log.Printf("The ID of the d_admin is incorrect: %v", id)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	adminid, error := uuid.Parse(vars["adminid"])
	if error != nil {
		log.Printf("Failed: %v", error)
	}
	if adminid.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		log.Printf("The ID of the admin is incorrect: %v", adminid)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	admin := models.DeleteAdminByID(&id, &adminid)
	res, err := json.Marshal(admin)
	if err != nil {
		log.Printf("Failed with marshalling the deleted admin: %v", err)
	}
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
	if id.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		log.Printf("The ID of the to be updated admin is in the wring format: %v", id)
		w.WriteHeader(http.StatusNotAcceptable)
		return
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
	if adminid.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		log.Printf("Entered admin ID, who took the changes, is in the wrong format: %v", adminid)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	//update admin with the new values
	admin := models.UpdateAdminbyID(&id, firstName, lastName, Password, rights, &adminid, emailaddress)
	res, err := json.Marshal(admin)
	if err != nil {
		log.Printf("Failed the marshal of the updated admin: %v", err)
	}
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
	if id.String() == envs["WRONG_ADMIN_ID_FORMAT"] {
		log.Printf("Can't work with this admin ID format: %v", id)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	admin := models.LastTimeOnline(&id)
	res, err := json.Marshal(admin)
	if err != nil {
		log.Printf("Error marshalling admin: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//contact functions

func GetContact(w http.ResponseWriter, r *http.Request) {
	contact := models.GetContact()
	res, err := json.Marshal(contact)
	if err != nil {
		log.Printf("Error marshalling contact: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//infoCard functions

// func GetInfoCard(w http.ResponseWriter, r *http.Request) {
// 	infoCard := models.GetInfoCard()
// 	res, err := json.Marshal(infoCard)
// 	if err != nil {
// 		log.Printf("Error marshalling info card: %v", err)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }
