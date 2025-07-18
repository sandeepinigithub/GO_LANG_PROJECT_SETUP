package routes

import (
	"github.com/gorilla/mux"

	"GO_LANG_PROJECT_SETUP/controller"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", controller.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controller.GetUser).Methods("GET")
	r.HandleFunc("/users", controller.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", controller.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", controller.DeleteUser).Methods("DELETE")
	r.HandleFunc("/api/login", controller.Login).Methods("POST")
	r.HandleFunc("/api/banned", controller.GetBanned).Methods("GET")
	r.HandleFunc("/api/banned/unban", controller.Unban).Methods("POST")
	r.HandleFunc("/api/jails", controller.GetJails).Methods("GET")

	r.HandleFunc("/api/domains", controller.ListDomains).Methods("GET")
	r.HandleFunc("/api/domain/{domain}", controller.GetDomain).Methods("GET")
	r.HandleFunc("/api/domain/{domain}", controller.CreateDomain).Methods("POST")
	r.HandleFunc("/api/domain/{domain}", controller.UpdateDomain).Methods("PUT")
	r.HandleFunc("/api/domain/{domain}", controller.DeleteDomain).Methods("DELETE")

	r.HandleFunc("/api/domain-admins", controller.ListDomainAdmins).Methods("GET")
	r.HandleFunc("/api/domain-admin/{email}", controller.GetDomainAdmin).Methods("GET")
	r.HandleFunc("/api/domain-admin/{email}", controller.CreateDomainAdmin).Methods("POST")
	r.HandleFunc("/api/domain-admin/{email}", controller.UpdateDomainAdmin).Methods("PUT")
	r.HandleFunc("/api/domain-admin/{email}", controller.DeleteDomainAdmin).Methods("DELETE")

	r.HandleFunc("/api/aliases", controller.ListAliases).Methods("GET")
	r.HandleFunc("/api/alias/{address}", controller.GetAlias).Methods("GET")
	r.HandleFunc("/api/alias/{address}", controller.CreateAlias).Methods("POST")
	r.HandleFunc("/api/alias/{address}", controller.UpdateAlias).Methods("PUT")
	r.HandleFunc("/api/alias/{address}", controller.DeleteAlias).Methods("DELETE")

	r.HandleFunc("/api/mailing-lists", controller.ListMailingLists).Methods("GET")
	r.HandleFunc("/api/mailing-list/{address}", controller.GetMailingList).Methods("GET")
	r.HandleFunc("/api/mailing-list/{address}", controller.CreateMailingList).Methods("POST")
	r.HandleFunc("/api/mailing-list/{address}", controller.UpdateMailingList).Methods("PUT")
	r.HandleFunc("/api/mailing-list/{address}", controller.DeleteMailingList).Methods("DELETE")

	r.HandleFunc("/api/greylisting", controller.ListGreylisting).Methods("GET")
	r.HandleFunc("/api/greylisting/{id}", controller.GetGreylisting).Methods("GET")
	r.HandleFunc("/api/greylisting", controller.CreateGreylisting).Methods("POST")
	r.HandleFunc("/api/greylisting/{id}", controller.UpdateGreylisting).Methods("PUT")
	r.HandleFunc("/api/greylisting/{id}", controller.DeleteGreylisting).Methods("DELETE")

	r.HandleFunc("/api/throttle", controller.ListThrottle).Methods("GET")
	r.HandleFunc("/api/throttle/{id}", controller.GetThrottle).Methods("GET")
	r.HandleFunc("/api/throttle", controller.CreateThrottle).Methods("POST")
	r.HandleFunc("/api/throttle/{id}", controller.UpdateThrottle).Methods("PUT")
	r.HandleFunc("/api/throttle/{id}", controller.DeleteThrottle).Methods("DELETE")

	r.HandleFunc("/api/wblist", controller.ListWblist).Methods("GET")
	r.HandleFunc("/api/wblist/{id}", controller.GetWblist).Methods("GET")
	r.HandleFunc("/api/wblist", controller.CreateWblist).Methods("POST")
	r.HandleFunc("/api/wblist/{id}", controller.UpdateWblist).Methods("PUT")
	r.HandleFunc("/api/wblist/{id}", controller.DeleteWblist).Methods("DELETE")

	r.HandleFunc("/api/logs", controller.ListLogs).Methods("GET")
	r.HandleFunc("/api/quota", controller.ListQuota).Methods("GET")
	r.HandleFunc("/api/roundcube-users", controller.ListRoundcubeUsers).Methods("GET")
	r.HandleFunc("/api/roundcube-user/{id}", controller.GetRoundcubeUser).Methods("GET")

	return r
}
