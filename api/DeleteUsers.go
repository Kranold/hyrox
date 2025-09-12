package api

import "net/http"

func (cfg *APIConfig) DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	err := cfg.DB.DeleteAllUsers(r.Context())
	if err != nil {
		http.Error(w, "Error deleting all users", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
