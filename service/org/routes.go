package org

import (
	"net/http"

	"github.com/MungaSoftwiz/org-authenticator-api/types"
	"github.com/MungaSoftwiz/org-authenticator-api/utils"
	"github.com/gorilla/mux"

	"github.com/google/uuid"
)

type OrganisationHandler struct {
	storage types.OrganisationStorage
}

func NewOrganisationHandler(storage types.OrganisationStorage) *OrganisationHandler {
	return &OrganisationHandler{storage: storage}
}

func (h *OrganisationHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/organisations", h.GetAllOrganisations).Methods("GET")
	r.HandleFunc("/organisations/{orgId}", h.GetOrganisationByID).Methods("GET")
	r.HandleFunc("/organisations/create", h.CreateOrganisation).Methods("POST")
	r.HandleFunc("/organisations/{orgId}/users", h.AddUserToOrganisation).Methods("POST")
}

func (h *OrganisationHandler) GetAllOrganisations(w http.ResponseWriter, r *http.Request) {
	organisations, err := h.storage.GetAllOrganisations()
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "could not retrieve organisations",
		})
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "organisations retrieved",
		"data": map[string]interface{}{
			"organisations": organisations,
		},
	})
}

func (h *OrganisationHandler) GetOrganisationByID(w http.ResponseWriter, r *http.Request) {
	orgID := mux.Vars(r)["orgId"]
	organisation, err := h.storage.GetOrganisationByID(orgID)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "could not retrieve organisation",
		})
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "organisation retrieved",
		"data":    organisation,
	})
}

func (h *OrganisationHandler) CreateOrganisation(w http.ResponseWriter, r *http.Request) {
	var organisation types.Organisation
	if err := utils.ReadJSON(r, &organisation); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"status":     "error",
			"message":    "invalid payload",
			"statusCode": http.StatusBadRequest,
		})
		return
	}

	orgUUID := uuid.New()
	organisation.OrgID = orgUUID.String()
	var err error
	utils.WriteJSON(w, http.StatusInternalServerError, map[string]interface{}{
		"status":  "error",
		"message": "could not generate UUID",
	})

	err = h.storage.CreateOrganisation(organisation)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "could not create organisation",
		})
		return
	}
	utils.WriteJSON(w, http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "Organisation created successfully",
		"data":    organisation,
	})
}

func (h *OrganisationHandler) AddUserToOrganisation(w http.ResponseWriter, r *http.Request) {
	orgID := mux.Vars(r)["orgId"]
	var payload struct {
		UserID string `json:"userId"`
	}
	if err := utils.ReadJSON(r, &payload); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{"status": "error", "message": "invalid payload"})
		return
	}

	err := h.storage.AddUserToOrganisation(orgID, payload.UserID)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]interface{}{"status": "error", "message": "could not add user to organisation"})
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"status": "success", "message": "user added to organisation"})
}
