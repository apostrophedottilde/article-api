package controller

import (
	"bytes"
	"encoding/xml"
	"github.com/apostrophedottilde/blog-article-api/app/article"
	"github.com/apostrophedottilde/blog-article-api/app/article/model"
	"html"
	"io/ioutil"
	"net/http"
	"strings"
)

type ArticleController struct {
	Usecase article.Usecase
}

func (ac *ArticleController) Create(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	stringBody := bytes.NewBuffer(bodyBytes).String()
	unescapedText := html.UnescapeString(stringBody)
	newArticle := model.ArticleModel{Content: unescapedText}
	createdId, err := ac.Usecase.Create(newArticle)
	if err != nil {
		sendBadResponse(w, err.Error(), 500)
		return
	}
	protocol := strings.ToLower(strings.Split(r.Proto, "/")[0])
	w.Header().Set("Location", protocol+"://"+r.Host+r.URL.String()+"/"+createdId)
	w.WriteHeader(201)
	_, _ = w.Write([]byte(""))
}

func (ac *ArticleController) FindOne(w http.ResponseWriter, r *http.Request) {
	id := getId(r)
	res, err := ac.Usecase.FindOne(id)
	if err != nil {
		sendBadResponse(w, err.Error(), 404)
		return
	}
	marshaled, err := xml.MarshalIndent(res.Response(), "", "")
	if err != nil {
		sendBadResponse(w, err.Error(), 500)
		return
	}
	buildAndPublishHappyResponse(w, []byte(marshaled), 200)
}

func (ac *ArticleController) Update(w http.ResponseWriter, r *http.Request) {
	id := getId(r)
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	newArticle := model.ArticleModel{
		Content: string(bodyBytes),
	}
	err := ac.Usecase.Update(id, newArticle)
	if err != nil {
		sendBadResponse(w, err.Error(), 500)
		return
	}
	buildAndPublishHappyResponse(w, []byte(""), 204)
}

func (ac *ArticleController) FindAll(w http.ResponseWriter, r *http.Request) {
	articles, _ := ac.Usecase.FindAll()
	var collection model.ArticleCollection
	collection.Articles = articles
	collectionResponse := collection.Response()
	marshaled, err := xml.MarshalIndent(collectionResponse, "", "")
	if err != nil {
		sendBadResponse(w, err.Error(), 500)
		return
	}
	buildAndPublishHappyResponse(w, marshaled, 200)
}

func getId(r *http.Request) string {
	return strings.TrimPrefix(r.URL.Path, "/articles/")
}

func buildAndPublishHappyResponse(w http.ResponseWriter, data []byte, code int) {
	w.Header().Add("Content-Type", "application/xml")
	w.WriteHeader(code)
	w.Write(data)
}

type ErrResponse struct {
	Code    int    `xml:"error-code"`
	Message string `xml:"error-message"`
}

func sendBadResponse(w http.ResponseWriter, message string, code int) {
	w.Header().Add("Content-Type", "application/xml")
	error := ErrResponse{code, message}
	x, _ := xml.Marshal(error)
	w.Write([]byte(x))
	return
}

func New(usecase article.Usecase) *ArticleController {
	return &ArticleController{
		Usecase: usecase,
	}
}
