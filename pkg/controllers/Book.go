package controllers

import(
	"strconv"
	"net/http"
	"encoding/json"
	"log"
	"os"
	"fmt"
	"path/filepath"
	"io"
	"github.com/irvandandung/goAPI/pkg/data"
)

func GetAllDataBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

    if r.Method != "GET" {
        http.Error(w, "Unsupported http method", http.StatusBadRequest)
        return
    }

    allDataBook := data.GetAllDataBuku()

    response := data.Response{
        Status : 200,
        Message : "Success",
        Data : allDataBook,
    }

    log.Println(response)
    json.NewEncoder(w).Encode(response)
}

func GetDataBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
        http.Error(w, "Unsupported http method", http.StatusBadRequest)
        return
    }

    keys, ok := r.URL.Query()["id"]
    if !ok {
    	http.Error(w, "Please add parameter value", http.StatusBadRequest)
    	return 
    }

    log.Println(keys)
    i, _ := strconv.Atoi(keys[0])
    getDataBoook := data.GetDataBukuById(i)
    var response data.Response
    if getDataBoook.Id == 0 {
    	response = data.Response{
	        Status : 201,
	        Message : "Data Tidak Ditemukan",
	        Data : getDataBoook,
	    }
    }else{
    	response = data.Response{
    	    Status : 200,
        	Message : "Success",
        	Data : getDataBoook,
    	}
    }

    log.Println(response)
    json.NewEncoder(w).Encode(response)
}

func SubmitPhoto(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	alias := r.FormValue("alias")

	if r.Method != "POST" {
        http.Error(w, "Unsupported http method", http.StatusBadRequest)
        return
    }

    if err := r.ParseMultipartForm(1024); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
	    http.Error(w, err.Error(), http.StatusInternalServerError)
	    return
	}
	defer uploadedFile.Close()

	if handler.Header["Content-Type"][0] != "image/jpeg" {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	    return
	}

	dir, err := os.Getwd()
	if err != nil {
	    http.Error(w, err.Error(), http.StatusInternalServerError)
	    return
	}

	filename := handler.Filename
	if alias != "" {
	    filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
	}
	addtodb := data.InsertGambarBuku(filename)
	log.Println(addtodb)

	fileLocation := filepath.Join(dir, "assets/images", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
	    http.Error(w, err.Error(), http.StatusInternalServerError)
	    return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
	    http.Error(w, err.Error(), http.StatusInternalServerError)
	    return
	}

	response := data.Response{
        Status : 200,
        Message : addtodb,
        Data : handler.Header,
    }

    json.NewEncoder(w).Encode(response)
}