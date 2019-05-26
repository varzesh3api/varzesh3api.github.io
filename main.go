package main

import (
	"log"
	"net/http"
	"strconv"
	"encoding/json"
	"amir/varzesh3/config"
)


func getConfig() config.Config {
	configuration, err := config.LoadConfiguration(config.ConfigFile)
	
	if err != nil {
		log.Println(err)
	}
	return configuration
}

//GetRecentNewsPersianTitle retrieves the Persian form of the most recent news title
/** 
 * @api {get} /api/recent/title retrieve Persian title
 * @apiName GetRecentNewsPersianTitle
 * @apiGroup recent
 * @apiDescription It is used to retrieve the Persian form of the most recent news title
 *
 *
 * @apiSuccess {String} description the Persian form of the most recent news title
 */
func GetRecentNewsPersianTitle(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getConfig().Recent.Title)
}

//GetRecentTopicByID retrieves an entity of the the most recent news data
/** 
 * @api {get} /api/recent/topic?id= retrieve an entity
 * @apiName GetRecentTopicByID
 * @apiGroup recent
 * @apiDescription It is used to retrieve an entity of the the most recent news data
 *
 * @apiParam {Number} id unique ID [0, 9]
 *
 * @apiSuccess {String} description an entity of the most recent news data
 */
func GetRecentTopicByID(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	intID, err := strconv.Atoi(ID)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(getConfig().Recent.Data[intID].Title)
}

//GetMostVisitedNewsPersianTitle retrieves the Persian form of the most visited news title
/** 
 * @api {get} /api/visit/title retrieve Persian title
 * @apiName GetMostVisitedNewsPersianTitle
 * @apiGroup visit
 * @apiDescription It is used to retrieve the Persian form of the most visited news title
 *
 *
 * @apiSuccess {String} description the Persian form of the most visited news title
 */
func GetMostVisitedNewsPersianTitle(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getConfig().Visit.Title)
}

//GetMostVisitedTopicByID retrieves an entity of the most visited news data
/** 
 * @api {get} /api/visit/topic?id= retrieve an entity
 * @apiName GetMostVisitedTopicByID
 * @apiGroup visit
 * @apiDescription It is used to retrieve an entity of the most visited news data
 *
 *
 * @apiParam {Number} id unique ID [0, 9]
 *
 * @apiSuccess {String} description an entity of the most visited news data
 */
func GetMostVisitedTopicByID(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	intID, err := strconv.Atoi(ID)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(getConfig().Visit.Data[intID].Title)
}

//GetTrendNewsPersianTitle retrieves the Persian form of the trending news title
/** 
 * @api {get} /api/trend/title retrieve Persian title
 * @apiName GetTrendNewsPersianTitle
 * @apiGroup trend
 * @apiDescription It is used to retrieve the Persian form of the trending news title
 *
 *
 * @apiSuccess {String} description the Persian form of the trending news title
 */
func GetTrendNewsPersianTitle(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getConfig().Trend.Title)
}

//GetTrendTopicByID retrieves an entity of the trending news data
/** 
 * @api {get} /api/trend/topic?id= retrieve an entity
 * @apiName GetTrendTopicByID
 * @apiGroup trend
 * @apiDescription It is used to retrieve an entity of the trending news data
 *
 * @apiParam {Number} id unique ID [0, 9]
 *
 * @apiSuccess {String} description an entity of the trending news data
 */
func GetTrendTopicByID(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	intID, err := strconv.Atoi(ID)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(getConfig().Trend.Data[intID].Title)
}

//GetOtherSportNewsPersianTitle retrieves the Persian form of other sport's title
/** 
 * @api {get} /api/other/title retrieve Persian title
 * @apiName GetOtherSportNewsPersianTitle
 * @apiGroup other
 * @apiDescription It is used to retrieve the Persian form of other sport's title
 *
 *
 * @apiSuccess {String} description the Persian form of other sport's title
 */
func GetOtherSportNewsPersianTitle(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getConfig().Other.Title)
}

//GetOtherSportTopicByID retrieves an entity of the other sport's data
/** 
 * @api {get} /api/other/topic?id= retrieve an entity
 * @apiName GetOtherSportTopicByID
 * @apiGroup other
 * @apiDescription It is used to retrieve an entity of the other sport's data
 *
 * @apiParam {Number} id unique ID [0, 9]
 *
 * @apiSuccess {String} description an entity of the other sport's data
 */
func GetOtherSportTopicByID(w http.ResponseWriter, r *http.Request) {	
	ID := r.URL.Query().Get("id")
	intID, err := strconv.Atoi(ID)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(getConfig().Other.Data[intID].Title)
}

//GetHeadlineSubtitleByID retrieves a specific headline's subtitle 
/** 
 * @api {get} /api/headline/subtitle?id= request subtitle
 * @apiName GetHeadlineSubtitleByID
 * @apiGroup headline
 * @apiDescription It is used to retrieve a specific headline's subtitle
 *
 * @apiParam {Number} id headline's unique ID [0, 6]
 *
 * @apiSuccess {String} subtitle a specific headline's subtitle
 */
func GetHeadlineSubtitleByID(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	intID, err := strconv.Atoi(ID)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(getConfig().Headlines.Data[intID].Subtitle)
}

//GetHeadlineTitleByID retrieves a specific headline's title 
/** 
 * @api {get} /api/headline/title?id= request title
 * @apiName GetHeadlineTitleByID
 * @apiGroup headline
 * @apiDescription It is used to retrieve a specific headline's title
 *
 * @apiParam {Number} id headline's unique ID [0, 6]
 *
 * @apiSuccess {String} title a specific headline's title
 */
func GetHeadlineTitleByID(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	intID, err := strconv.Atoi(ID)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(getConfig().Headlines.Data[intID].Title)
}

//GetHeadlineDescriptionByID retrieves a specific headline's description 
/** 
 * @api {get} /api/headline/description?id= request description
 * @apiName GetHeadlineDescriptionByID
 * @apiGroup headline
 * @apiDescription It is used to retrieve a specific headline's description
 *
 * @apiParam {Number} id headline's unique ID [0, 6]
 *
 * @apiSuccess {String} description a specific headline's description
 */
func GetHeadlineDescriptionByID(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	intID, err := strconv.Atoi(ID)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(getConfig().Headlines.Data[intID].Description)
}

//GetHeadlineImageByID retrieves a specific headline's image as a link
/** 
 * @api {get} /api/headline/image?id= request image's link
 * @apiName GetHeadlineImageByID
 * @apiGroup headline
 * @apiDescription It is used to retrieve a specific headline's image as a link
 *
 * @apiParam {Number} id headline's unique ID [0, 6]
 *
 * @apiSuccess {String} description a specific headline's image as a link
 */
func GetHeadlineImageByID(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	intID, err := strconv.Atoi(ID)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(getConfig().Headlines.Data[intID].Image)
}

func handleRequests() {
	http.HandleFunc("/api/recent/title", GetRecentNewsPersianTitle)
	http.HandleFunc("/api/recent/topic", GetRecentTopicByID)
	http.HandleFunc("/api/visit/title", GetMostVisitedNewsPersianTitle)
	http.HandleFunc("/api/visit/topic", GetMostVisitedTopicByID)
	http.HandleFunc("/api/trend/title", GetTrendNewsPersianTitle)
	http.HandleFunc("/api/trend/topic", GetTrendTopicByID)
	http.HandleFunc("/api/other/title", GetOtherSportNewsPersianTitle)
	http.HandleFunc("/api/other/topic", GetOtherSportTopicByID)
	http.HandleFunc("/api/headline/subtitle", GetHeadlineSubtitleByID)
	http.HandleFunc("/api/headline/title", GetHeadlineTitleByID)
	http.HandleFunc("/api/headline/description", GetHeadlineDescriptionByID)
	http.HandleFunc("/api/headline/image", GetHeadlineImageByID)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}