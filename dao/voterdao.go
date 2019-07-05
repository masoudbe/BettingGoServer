package dao

import (
	"../model"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetAllVoters(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	voters := []model.VoteInfo{}
	db.Find(&voters)
	respondJSON(w, http.StatusOK, voters)
}

func CreateVoter(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	voter := model.VoteInfo{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&voter); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&voter).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, voter)
}

func CreateVoters(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	voters := []model.VoteInfo{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&voters); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	//for _,voter := range voters{
	//	if err := db.Debug().Create(&voter).Error; err != nil {
	//		respondError(w, http.StatusInternalServerError, err.Error())
	//		return
	//	}
	//}

	if err := db.Debug().Create(&voters).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, voters)
}

func GetVoter(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	voter := getVoterOr404(db, name, w, r)
	if voter == nil {
		return
	}
	respondJSON(w, http.StatusOK, voter)
}

func UpdateVoter(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	voter := getVoterOr404(db, name, w, r)
	if voter == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&voter); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&voter).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, voter)
}

func DeleteVoter(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee := getVoterOr404(db, name, w, r)
	if employee == nil {
		return
	}
	if err := db.Delete(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

//func EnableEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//
//	name := vars["name"]
//	employee := getVoterOr404(db, name, w, r)
//	if employee == nil {
//		return
//	}
//	employee.Enable()
//	if err := db.Save(&employee).Error; err != nil {
//		respondError(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//	respondJSON(w, http.StatusOK, employee)
//}

// getVoterOr404 gets a employee instance if exists, or respond the 404 error otherwise
func getVoterOr404(db *gorm.DB, voter string, w http.ResponseWriter, r *http.Request) *model.Voter {
	voterDto := model.Voter{}

	if err := db.First(&voterDto, model.Voter{Name: voter}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &voterDto
}
