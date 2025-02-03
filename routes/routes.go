package routes

import (
	"github.com/M4tt1-Coder/business/portfolio_website/API_GO/controller"
	"github.com/gorilla/mux"
)

// all routes to communicate with the frontend server
var AllRoutes = func(router *mux.Router) {
	//partner routes
	router.HandleFunc("/partner/create/{adminid}", controller.CreatePartner).Methods("POST")
	router.HandleFunc("/partner/getAll/", controller.GetAllPartners).Methods("GET")
	router.HandleFunc("/partner/get/{id}", controller.GetPartner).Methods("GET")
	router.HandleFunc("/partner/delete/{id}&&{adminid}", controller.DeletePartner).Methods("DELETE")
	router.HandleFunc("/partner/update/", controller.UpdatePartner).Methods("PUT")

	//project routes
	router.HandleFunc("/project/create/{adminid}", controller.CreateProject).Methods("POST")
	router.HandleFunc("/project/getAll/", controller.GetAllProjects).Methods("GET")
	router.HandleFunc("/project/get/{id}", controller.GetProject).Methods("GET")
	router.HandleFunc("/project/delete/{id}&&{adminid}", controller.DeleteProject).Methods("DELETE")
	router.HandleFunc("/project/update/", controller.UpdateProject).Methods("PUT")

	//message routes
	router.HandleFunc("/message/create/", controller.CreateMessage).Methods("POST")
	router.HandleFunc("/message/getAll/", controller.GetAllMessages).Methods("GET")
	router.HandleFunc("/message/get/{id}", controller.GetMessage).Methods("GET")
	router.HandleFunc("/message/delete/{id}&&{adminid}", controller.DeleteMessage).Methods("DELETE")
	router.HandleFunc("/message/update/", controller.UpdateMessage).Methods("PUT")

	//admin routes
	router.HandleFunc("/admin/create/{adminid}", controller.CreateAdmin).Methods("POST")
	router.HandleFunc("/admin/getAll/", controller.GetAllAdmins).Methods("GET")
	router.HandleFunc("/admin/get/{id}", controller.GetAdmin).Methods("GET")
	router.HandleFunc("/admin/delete/{id}&&{adminid}", controller.DeleteAdmin).Methods("DELETE")
	router.HandleFunc("/admin/update/", controller.UpdateAdmin).Methods("PUT")
	router.HandleFunc("/admin/online/{id}", controller.UpdateAdminLastTimeOnline).Methods("PUT")

	//contact routes
	router.HandleFunc("/contact/get/", controller.GetContact).Methods("GET")

	//infoCard routes
	//router.HandleFunc("/infoCard/get/", controller.GetInfoCard).Methods("GET")
}
