package routes

import (
	"github.com/gorilla/mux"

	// TEMP
	"kwoc20-backend/controllers"
	"kwoc20-backend/utils"
)

// TEMP 
// Discuss and add 2 middlewares - JWT Required, JSON Marshalling
func RegisterStudent(r *mux.Router) {
	r.HandleFunc("", utils.LoginRequired(utils.JsonIO(controllers.StudentReg))).Methods("POST")

}

