package studentcontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/alyadyh/go-restapi-mux/helper"
	"github.com/alyadyh/go-restapi-mux/models"
	"github.com/gorilla/mux"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func Index(w http.ResponseWriter, r *http.Request) {
	var students []models.Student

	if err := models.DB.Find(&students).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseJson(w, http.StatusOK, students)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nim, err := strconv.ParseInt(vars["nim"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var student models.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&student); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := models.DB.Where("nim = ?", nim).Find(&student).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	student.Nim = nim

	ResponseJson(w, http.StatusOK, student)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var students models.Student

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&students); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := models.DB.Create(&students).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusCreated, students)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nim, err := strconv.ParseInt(vars["nim"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var student models.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&student); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := models.DB.Where("nim = ?", nim).Updates(&student).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	student.Nim = nim

	ResponseJson(w, http.StatusOK, student)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nim, err := strconv.ParseInt(vars["nim"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	r.Close = true

	var student models.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&student); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := models.DB.Where("nim = ?", nim).Delete(&student).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	student.Nim = nim

	mes := map[string]string{"message": "Data mahasiswa telah berhasil dihapus"}

	ResponseJson(w, http.StatusOK, mes)
}
