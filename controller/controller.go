package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/M4tt1-Coder/business/portfolio_website/API_GO/models"
	"github.com/M4tt1-Coder/business/portfolio_website/API_GO/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

// set the needed variables the file
var (
	envs, _ = godotenv.Read(".env")
)

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
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !utils.IsValidUUID(adminid.String()) {
		log.Printf("The passed adminid is invalid! Wrong Format %v", adminid.String())
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var partner models.Partner
	err := decoder.Decode(&partner)
	if err != nil {
		log.Printf("Error decoding the partner object in the request's body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	partner = *partner.CreatePartner(&adminid)
	res, err := json.Marshal(partner)
	if err != nil {
		log.Printf("The partner couldn't be parsed into a JSON object: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Doesn't need any route variables or http body.
//
// Loads all partner object from the database.
func GetAllPartners(w http.ResponseWriter, r *http.Request) {
	partners := models.GetAllPartners()
	res, err := json.Marshal(partners)
	if err != nil {
		log.Printf("Error marshalling the partners list: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Gets the id of the partner that should be returned from the URL variables.
//
// With a valid ID, it queries the partner entity from the database.
//
// Returns the partner object.
//
// FAILS when the passed ID isn't in a valid format & when the
func GetPartner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("Failed to parse the UUID: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Checking if the ID was parsed into the wrong uuid format
	if !utils.IsValidUUID(id.String()) {
		w.WriteHeader(http.StatusNotAcceptable)
		partner := &models.Partner{}
		res, err := json.Marshal(partner)
		if err != nil {
			log.Printf("Error marshalling the partner object: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(res)
		return
	}
	partner, _ := models.GetPartnerByID(&id)
	res, err := json.Marshal(partner)
	if err != nil {
		log.Printf("Couldn't marshall the partner object! Error: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Receives the partner id and admin id from the admin who wants to delete the partner.
//
// Checks if the two passed ids are both in a valid format.
func DeletePartner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("Failed to parse the UUID: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !utils.IsValidUUID(id.String()) {
		log.Printf("The id %v is not a valid UUID format!", id)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	adminid, err := uuid.Parse(vars["adminid"])
	if err != nil {
		log.Printf("Failed to parse the admin id UUID! Error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !utils.IsValidUUID(adminid.String()) {
		log.Printf("The admin id %v is not a valid UUID format!", adminid)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	partner, _ := models.DeletePartnerByID(&id, &adminid)
	res, err := json.Marshal(partner)
	if err != nil {
		log.Printf("Failed to parse the partner object into JSON! Error: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Gets new data for the partner and updates the database entries with the new values.
//
// Decodes the HTTP-request body to get all data.
//
// Fails when the 'id' or 'adminid' are nil.
//
// Other properties are optional and doesn't need to be specified.
func UpdatePartner(w http.ResponseWriter, r *http.Request) {
	//get all parameters
	//update with the new values
	//log the changes with admin
	decoder := json.NewDecoder(r.Body)
	var vars map[string]string
	err := decoder.Decode(&vars)
	if err != nil {
		log.Printf("Error decoding the request body! Error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//id
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("Parsing the id as UUID failed! Error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !utils.IsValidUUID(id.String()) {
		log.Printf("Invalid UUID format of the passed ID! ID: %v", id.String())
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
		log.Printf("Failed to parse the time data of the request! Error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//address
	address := vars["address"]
	//phone
	telephoneNumber := vars["telephoneNumber"]
	//adminid
	adminid, err := uuid.Parse(vars["adminid"])
	if err != nil {
		log.Printf("Failed to parse the admin id! Error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !utils.IsValidUUID(adminid.String()) {
		log.Printf("The provided admin id is not in a valid UUID format! ID: %v", adminid)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	//update partner with the new values
	partner := models.UpdatePartnerByID(&id, name, websiteLink, sinceWhen, address, telephoneNumber, &adminid)
	res, err := json.Marshal(partner)
	if err != nil {
		log.Printf("Couldn't parse the partner object into a JSON string: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// project functions

// Gets the admin id from the route parameters.
//
// Destructures the project data from the request body.
//
// Fails if either the id is invalid or the object body can't be parsed.
//
// Returns the created project for an evidence that it was created.
func CreateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adminid, error := uuid.Parse(vars["adminid"])
	if error != nil {
		log.Printf("Invalid uuid as admin id: %v", error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !utils.IsValidUUID(adminid.String()) {
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
		log.Printf("Error decoding the project object from the request body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	project = *project.CreateProject(&adminid)
	res, err := json.Marshal(project)
	if err != nil {
		log.Printf("Unable to convert project object into JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Returns all projects that are in the database.
//
// Doesn't need any provided data through the request.
//
// Returns the list of projects as a encoded json object.
func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects := models.GetAllProjects()
	res, err := json.Marshal(projects)
	if err != nil {
		log.Printf("Could not encode the list of projects: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Returns a specific project by its id.
//
// The project is passed down with the route parameters.
//
// The id needs to be a valid uuid.
func GetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, error := uuid.Parse(vars["id"])
	if error != nil {
		log.Printf("Failed: %v", error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !utils.IsValidUUID(id.String()) {
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
	response, err := json.Marshal(project)
	if err != nil {
		log.Printf("Failed to encode the project object: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Deletes a project from the database by its ID.
//
// The IDs are route parameters.
//
// Fails when the IDs are in an invalid UUID-format.
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, error := uuid.Parse(vars["id"])
	if error != nil {
		log.Printf("Failed due to no id route parameter: %v", error)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	if !utils.IsValidUUID(id.String()) {
		w.WriteHeader(http.StatusNotAcceptable)
		log.Printf("The id of the project is not in the right format: %v", id.String())
		return
	}
	adminid, error := uuid.Parse(vars["adminid"])
	if error != nil {
		log.Printf("Couldn't parse the admin ID: %v", error)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	if !utils.IsValidUUID(adminid.String()) {
		w.WriteHeader(http.StatusNotAcceptable)
		log.Printf("No valid admin ID has beed entered: %v", adminid.String())
		return
	}
	project := models.DeleteProjectByID(&id, &adminid)
	res, err := json.Marshal(project)
	if err != nil {
		log.Printf("An error has occurred while trying to parse the project object to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Updates the data of project entity.
//
// Data was send in the HTTP-request body.
//
// Destructures the send data to access it.
//
// Fails when the decoding failed, the id is in an invalid format
// or the fetched data from the database is invalid.
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var vars map[string]string
	err := decoder.Decode(&vars)
	if err != nil {
		log.Printf("Error decoding the request-body for updating the project: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//id
	id, error := uuid.Parse(vars["id"])
	if error != nil {
		log.Printf("Failed parsing the id of the project: %v", error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !utils.IsValidUUID(id.String()) {
		log.Printf("The new ID of the project is an incorrect format!")
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
		log.Printf("The admin id couldn't be parsed: %v", error.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !utils.IsValidUUID(adminid.String()) {
		log.Printf("Passed admin ID isn't in the right format!")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	//update project with the new values
	project := models.UpdateProjectByID(&id, name, link, description, participants, &adminid)
	res, err := json.Marshal(project)
	if err != nil {
		log.Printf("Failed to marshal the project: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//message functions

// Creates a new message entry in the database.
//
// Fails when an error occured while decoding the message or the response object couldn't
// be mashalled to a JSON object.
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var message models.Message
	err := decoder.Decode(&message)
	if err != nil {
		log.Printf("Error decoding the message request body to create a message in the database: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	message = *message.CreateMessage()
	res, err := json.Marshal(message)
	if err != nil {
		log.Printf("Failed to marshl the message: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Fetches all messages stored in the database and returns them as response.
//
// Fails if the message list couldn't be converted into a JSON object.
func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	messages := models.GetAllMessages()
	res, err := json.Marshal(messages)
	if err != nil {
		log.Printf("Couldn't marshal the messages! Error: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Gets a message from the database and returns it as JSON object.
//
// Fails when the id is invalid or couldn't be parsed, the fetched message object was not properly converted into a JSON object.
func GetMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, error := uuid.Parse(vars["id"])
	if error != nil {
		log.Printf("Received wrong UUID as id: %v", error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !utils.IsValidUUID(id.String()) {
		log.Printf("The id was in a bad format: %v", id)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	message, _ := models.GetMessageByID(&id)
	res, err := json.Marshal(message)
	if err != nil {
		log.Printf("Couldn't parse the message into JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Deletes a message by its ID.
//
// Gets message id and the id from the admin who deleted the message.
//
// Fails, when one of the ids is in a wrong format or couldn't be parsed,
// the returned message object wasn't marshalled into JSON.
func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("Got a invalid id from the request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !utils.IsValidUUID(id.String()) {
		log.Printf("Deleting message failed due to invalid message ID: %v", id)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	adminid, err := uuid.Parse(vars["adminid"])
	if err != nil {
		log.Printf("Unable to parse the admin ID: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !utils.IsValidUUID(adminid.String()) {
		log.Printf("The entered admin ID is invalid: %v", adminid)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	message := models.DeleteMessageByID(&id, &adminid)
	res, err := json.Marshal(message)
	if err != nil {
		log.Printf("Tried to marshal message but failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Updates a message entry in the database with new information.
//
// Some properties are optional, but not the ID, it is required.
//
// Fails, when the HTTP-request body couldn't be decoded, the id was in a bad format, the timestamp
// could not be parsed and the message object failed to to be converted into the response JSON object.
func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var vars map[string]string
	err := decoder.Decode(&vars)
	if err != nil {
		log.Printf("Error occured while decoding: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//id
	id, error := uuid.Parse(vars["id"])
	if error != nil {
		log.Printf("Failed parsing the id of the message: %v", error)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	if !utils.IsValidUUID(id.String()) {
		log.Printf("The message ID is in a wrong format: %v", id)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//content
	content := vars["content"]
	//subject
	subject := vars["subject"]
	//createdAt
	createdAt, err := time.Parse("2006-01-02 00:00:00", vars["createdAt"])
	if err != nil {
		log.Printf("Failed time parsing: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !utils.IsValidUUID(adminid.String()) {
		log.Printf("The admin ID is incorrect: %v", adminid)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	//update message with the new values
	message := models.UpdateMessageByID(&id, content, subject, createdAt, emailAddress, senderName, gender, &adminid)
	res, err := json.Marshal(message)
	if err != nil {
		log.Printf("Failed to marshal the message: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//admin functions

// Creates an admin entry in the database.
//
// Gets the data by destructuring the request body.
//
// Fails when the data could not be decoded, the passed admin id has a wrong
// format or the response admin failed to be converted into a JSON object.
func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var admin models.Admin
	err := decoder.Decode(&admin)
	if err != nil {
		log.Printf("Error decoding the request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	adminid, error := uuid.Parse(vars["adminid"])
	if error != nil {
		log.Printf("Failed to parse the adminid: %v", error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !utils.IsValidUUID(adminid.String()) {
		log.Printf("The admin ID is incorrect: %v", adminid)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	admin = *admin.CreateAdmin(&adminid)
	res, err := json.Marshal(admin)
	if err != nil {
		log.Printf("Couldn't marshal the created admin: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Gets all admins from the database and returns them.
//
// Fails if the list of admins can't be converted into a JSON object.
func GetAllAdmins(w http.ResponseWriter, r *http.Request) {
	admins := models.GetAllAdmins()
	res, err := json.Marshal(admins)
	if err != nil {
		log.Printf("Try to marshal admisn didn't succeed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Gets an admin by their id.
//
// The id is passed with the request route parameters.
//
// Fails, when the admin couldn't be parsed or has a wrong format and when the fetched admin instance from the database
// wasn't converted into a JSON object.
//
// Returns the requested admin.
func GetAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, error := uuid.Parse(vars["id"])
	if error != nil {
		log.Printf("Failed: %v", error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !utils.IsValidUUID(id.String()) {
		log.Printf("Attempting to get admin failed due to wrong admin ID format: %v", id)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	admin, _ := models.GetAdminByID(&id)
	res, err := json.Marshal(admin)
	if err != nil {
		log.Printf("Failed with marshalling the admin: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Gets the id of the admin that should be deleted and the id of the admin that executed the deletion of the admin.
//
// Deletes an admin from the database.
//
// Fails when the ids are not valid or the admin could not be deleted from the database.
//
// Returns the object of the deleted admin
func DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, error := uuid.Parse(vars["id"])
	if error != nil {
		log.Printf("Failed to parse the id of the admin, who should be deleted: %v", error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !utils.IsValidUUID(id.String()) {
		log.Printf("The ID of the d_admin is incorrect: %v", id)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	adminid, error := uuid.Parse(vars["adminid"])
	if error != nil {
		log.Printf("Something went wrong, when trying to parse the admin id: %v", error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !utils.IsValidUUID(adminid.String()) {
		log.Printf("The ID of the admin is incorrect: %v", adminid)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	admin := models.DeleteAdminByID(&id, &adminid)
	res, err := json.Marshal(admin)
	if err != nil {
		log.Printf("Failed with marshalling the deleted admin: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Updates an admin in the database.
//
// Fails when the updated admin could not be parsed into a JSON object.
//
// Returns the updated admin as a JSON object.
func UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var vars map[string]string
	err := decoder.Decode(&vars)
	if err != nil {
		log.Printf("Error decoding the body of the request to modify the admin: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//id
	id, error := uuid.Parse(vars["id"])
	if error != nil {
		log.Printf("Could not parse the admin id: %v", error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !utils.IsValidUUID(id.String()) {
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
		log.Printf("Passed UUID admin id is invalid! Failed to parse the admin id: %v", error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !utils.IsValidUUID(adminid.String()) {
		log.Printf("Entered admin ID, who took the changes, is in the wrong format: %v", adminid)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	//update admin with the new values
	admin := models.UpdateAdminbyID(&id, firstName, lastName, Password, rights, &adminid, emailaddress)
	res, err := json.Marshal(admin)
	if err != nil {
		log.Printf("Failed the marshal of the updated admin: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Updates the time stamp when the admin was online the last time.
//
// All ids must be valid.
//
// Fails when the admin could not be found or converted into a valid JSON representation.
//
// Returns a JSON representation of the admin that recently logged in
func UpdateAdminLastTimeOnline(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("Error parsing the admin id: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !utils.IsValidUUID(id.String()) {
		log.Printf("Can't work with this admin ID format: %v", id)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	admin := models.LastTimeOnline(&id)
	res, err := json.Marshal(admin)
	if err != nil {
		log.Printf("Error marshalling admin: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//contact functions

// Gets a all conrtact information and returns it to the client.
func GetContact(w http.ResponseWriter, r *http.Request) {
	contact := models.GetContact()
	res, err := json.Marshal(contact)
	if err != nil {
		log.Printf("Error marshalling contact: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
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
