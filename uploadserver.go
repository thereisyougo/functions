package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	maxUploadSize int64 = 2 << 25;
	uploadPath string = "./upload"
	webTemplate string = `
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>文件上传</title>
</head>
<body>
  <form action="/upload" method="post" enctype="multipart/form-data">
    <input type="file" name="uploadfile" multiple="multiple">
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit">
  </form>
</body>
</html>
`
)

func main() {
	localDir := flag.String("dir", ".", "static file server address")
	port := flag.Int("port", 8080, "web server port")
	flag.Parse()

	http.HandleFunc("/upload", uploadFileHanler())
	http.Handle("/files/",
		http.StripPrefix("/files", http.FileServer(http.Dir(*localDir))))
	log.Printf("Server started on localhost:%v", *port)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(*port), nil))
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}

func welcome(w http.ResponseWriter, r *http.Request)  {
	hash := md5.New()
	io.WriteString(hash, strconv.FormatInt(time.Now().Unix(), 10))
	data := fmt.Sprintf("%x", hash.Sum(nil))

	if tmp, e := template.New("welcome").Parse(webTemplate); e == nil {
		tmp.Execute(w, data)
	} else {
		log.Panic(e)
		renderError(w, e.Error(), http.StatusInternalServerError)
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request)  {
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		log.Println(err)
		renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
		return
	}

	if r.MultipartForm == nil {
		err := r.ParseMultipartForm(32 << 20)
		renderError(w, err.Error(), http.StatusBadRequest)
	}

	makeUploadDir()

	fileHeaders := r.MultipartForm.File["uploadfile"]
	for _, fileHeader := range fileHeaders {
		saveUploadFile(w, fileHeader)
	}

	//sourceFile, fileHeader, err := r.FormFile("uploadfile")


	welcome(w, r)

}

func saveUploadFile(w http.ResponseWriter, fileHeader *multipart.FileHeader) {
	sourceFile, err := fileHeader.Open()
	if err != nil {
		log.Println(err)
		renderError(w, "INVALID_FILE", http.StatusBadRequest)
		return
	}
	defer sourceFile.Close()

	//fileBytes, err := ioutil.ReadAll(file)
	//
	//// check file type, detectcontenttype only needs the first 512 bytes
	//filetype := http.DetectContentType(fileBytes)
	//
	//fileEndings, err := mime.ExtensionsByType(fileType)

	targetFile, err := os.OpenFile(filepath.Join(uploadPath, filepath.Base(fileHeader.Filename)), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
		renderError(w, "CANT_WRITE_FILE", http.StatusBadRequest)
		return
	}
	defer targetFile.Close()
	if _, err := io.Copy(targetFile, sourceFile); err != nil {
		log.Println(err)
		renderError(w, "ERROR_WRITE_FILE", http.StatusBadRequest)
	}
}

func makeUploadDir() {
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.Mkdir(uploadPath, os.ModePerm)
	}
}

func uploadFileHanler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		log.Println(method)
		switch method {
		case "GET":
			welcome(w, r)
		case "POST":
			uploadHandler(w, r)
		}
	});
}
