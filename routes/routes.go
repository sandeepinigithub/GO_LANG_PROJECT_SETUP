package routes

import (
	"github.com/gorilla/mux"
	"devsMailGo/controller"
	"devsMailGo/middleware"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Public routes (no authentication required)
	publicRoutes := r.PathPrefix("").Subrouter()
	publicRoutes.HandleFunc("/api/login", controller.Login).Methods("POST")
	publicRoutes.HandleFunc("/api/health", controller.HealthCheck).Methods("GET")

	// Protected API routes (authentication required)
	protectedAPI := r.PathPrefix("/api").Subrouter()
	protectedAPI.Use(middleware.AuthMiddleware)
	
	// User management routes
	protectedAPI.HandleFunc("/users", controller.GetUsers).Methods("GET")
	protectedAPI.HandleFunc("/users/{id}", controller.GetUser).Methods("GET")
	protectedAPI.HandleFunc("/users", controller.CreateUser).Methods("POST")
	protectedAPI.HandleFunc("/users/{id}", controller.UpdateUser).Methods("PUT")
	protectedAPI.HandleFunc("/users/{id}", controller.DeleteUser).Methods("DELETE")

	// Security management routes
	protectedAPI.HandleFunc("/banned", controller.GetBanned).Methods("GET")
	protectedAPI.HandleFunc("/banned/unban", controller.Unban).Methods("POST")
	protectedAPI.HandleFunc("/jails", controller.GetJails).Methods("GET")

	// Domain management routes
	protectedAPI.HandleFunc("/domains", controller.ListDomains).Methods("GET")
	protectedAPI.HandleFunc("/domain/{domain}", controller.GetDomain).Methods("GET")
	protectedAPI.HandleFunc("/domain/{domain}", controller.CreateDomain).Methods("POST")
	protectedAPI.HandleFunc("/domain/{domain}", controller.UpdateDomain).Methods("PUT")
	protectedAPI.HandleFunc("/domain/{domain}", controller.DeleteDomain).Methods("DELETE")

	// Domain admin management routes
	protectedAPI.HandleFunc("/domain-admins", controller.ListDomainAdmins).Methods("GET")
	protectedAPI.HandleFunc("/domain-admin/{email}", controller.GetDomainAdmin).Methods("GET")
	protectedAPI.HandleFunc("/domain-admin/{email}", controller.CreateDomainAdmin).Methods("POST")
	protectedAPI.HandleFunc("/domain-admin/{email}", controller.UpdateDomainAdmin).Methods("PUT")
	protectedAPI.HandleFunc("/domain-admin/{email}", controller.DeleteDomainAdmin).Methods("DELETE")

	// Email management routes
	protectedAPI.HandleFunc("/aliases", controller.ListAliases).Methods("GET")
	protectedAPI.HandleFunc("/alias/{address}", controller.GetAlias).Methods("GET")
	protectedAPI.HandleFunc("/alias/{address}", controller.CreateAlias).Methods("POST")
	protectedAPI.HandleFunc("/alias/{address}", controller.UpdateAlias).Methods("PUT")
	protectedAPI.HandleFunc("/alias/{address}", controller.DeleteAlias).Methods("DELETE")

	protectedAPI.HandleFunc("/mailing-lists", controller.ListMailingLists).Methods("GET")
	protectedAPI.HandleFunc("/mailing-list/{address}", controller.GetMailingList).Methods("GET")
	protectedAPI.HandleFunc("/mailing-list/{address}", controller.CreateMailingList).Methods("POST")
	protectedAPI.HandleFunc("/mailing-list/{address}", controller.UpdateMailingList).Methods("PUT")
	protectedAPI.HandleFunc("/mailing-list/{address}", controller.DeleteMailingList).Methods("DELETE")

	// Security and monitoring routes
	protectedAPI.HandleFunc("/greylisting", controller.ListGreylisting).Methods("GET")
	protectedAPI.HandleFunc("/greylisting/{id}", controller.GetGreylisting).Methods("GET")
	protectedAPI.HandleFunc("/greylisting", controller.CreateGreylisting).Methods("POST")
	protectedAPI.HandleFunc("/greylisting/{id}", controller.UpdateGreylisting).Methods("PUT")
	protectedAPI.HandleFunc("/greylisting/{id}", controller.DeleteGreylisting).Methods("DELETE")

	protectedAPI.HandleFunc("/throttle", controller.ListThrottle).Methods("GET")
	protectedAPI.HandleFunc("/throttle/{id}", controller.GetThrottle).Methods("GET")
	protectedAPI.HandleFunc("/throttle", controller.CreateThrottle).Methods("POST")
	protectedAPI.HandleFunc("/throttle/{id}", controller.UpdateThrottle).Methods("PUT")
	protectedAPI.HandleFunc("/throttle/{id}", controller.DeleteThrottle).Methods("DELETE")

	protectedAPI.HandleFunc("/wblist", controller.ListWblist).Methods("GET")
	protectedAPI.HandleFunc("/wblist/{id}", controller.GetWblist).Methods("GET")
	protectedAPI.HandleFunc("/wblist", controller.CreateWblist).Methods("POST")
	protectedAPI.HandleFunc("/wblist/{id}", controller.UpdateWblist).Methods("PUT")
	protectedAPI.HandleFunc("/wblist/{id}", controller.DeleteWblist).Methods("DELETE")

	// System monitoring routes
	protectedAPI.HandleFunc("/logs", controller.ListLogs).Methods("GET")
	protectedAPI.HandleFunc("/quota", controller.ListQuota).Methods("GET")
	protectedAPI.HandleFunc("/roundcube-users", controller.ListRoundcubeUsers).Methods("GET")
	protectedAPI.HandleFunc("/roundcube-user/{id}", controller.GetRoundcubeUser).Methods("GET")

	return r
}
