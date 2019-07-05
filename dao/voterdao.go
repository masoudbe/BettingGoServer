package dao

import (
	"../model"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetVotes(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	voters := []model.VoteInfo{}
	db.Find(&voters)
	respondJSON(w, http.StatusOK, voters)
}

func CreateVote(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
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

func CreateVotes(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	voteInfos := []model.VoteInfo{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&voteInfos); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	for _, voter := range voteInfos {
		if err := db.Create(&voter).Error; err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	respondJSON(w, http.StatusCreated, voteInfos)
}

func GetVote(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	voter := getVoteOr404(db, name, w, r)
	if voter == nil {
		return
	}
	respondJSON(w, http.StatusOK, voter)
}

func UpdateVote(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	voter := getVoteOr404(db, name, w, r)
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

func DeleteVote(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee := getVoteOr404(db, name, w, r)
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
//	employee := getVoteOr404(db, name, w, r)
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

func getVoteOr404(db *gorm.DB, voter string, w http.ResponseWriter, r *http.Request) *model.VoteInfo {
	voterDto := model.VoteInfo{}

	if err := db.First(&voterDto, model.VoteInfo{Name: voter}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &voterDto
}
