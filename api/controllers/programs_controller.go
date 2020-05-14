package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kerimkuscu/hardline-fitness-backend/api/auth"
	"github.com/kerimkuscu/hardline-fitness-backend/api/database"
	"github.com/kerimkuscu/hardline-fitness-backend/api/models"
	"github.com/kerimkuscu/hardline-fitness-backend/api/repository"
	"github.com/kerimkuscu/hardline-fitness-backend/api/repository/crud"
	"github.com/kerimkuscu/hardline-fitness-backend/api/responses"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateProgram(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Create program")) //this is for demo
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	program := models.Program{}
	err = json.Unmarshal(body, &program)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	program.Prepare()
	err = program.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	if uid != program.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryProgramsCRUD(db)

	func(programRepository repository.ProgramRepository) {
		program, err := programRepository.Save(program)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, program.ID))
		responses.JSON(w, http.StatusCreated, program)
	}(repo)
}

func GetPrograms(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryProgramsCRUD(db)

	func(programRepository repository.ProgramRepository) {
		programs, err := programRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		// w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, program.ID))
		responses.JSON(w, http.StatusOK, programs)
	}(repo)
}

func GetProgram(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryProgramsCRUD(db)

	func(programRepository repository.ProgramRepository) {
		program, err := programRepository.FindByID(pid)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		// w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, program.ID))
		responses.JSON(w, http.StatusOK, program)
	}(repo)
}

func UpdateProgram(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	program := models.Program{}
	err = json.Unmarshal(body, &program)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	program.Prepare()
	err = program.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	if uid != program.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryProgramsCRUD(db)

	func(programRepository repository.ProgramRepository) {
		rows, err := programRepository.Update(pid, program)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		responses.JSON(w, http.StatusOK, rows)
	}(repo)
}

func DeleteProgram(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	fmt.Println("USER: ", uid)

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryProgramsCRUD(db)

	func(programRepository repository.ProgramRepository) {
		_, err := programRepository.Delete(pid, uid)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Entity", fmt.Sprintf("%d", pid))
		responses.JSON(w, http.StatusNoContent, "")
	}(repo)
}
