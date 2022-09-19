package main

import (
	"fmt"
	"net/http"
	"strings"
)

var hosts = "https://api.vhtear.com/"
var apikey = "Apikey"

func ChangeProfilePicture(AuthToken string, Msg_Id string) {
	//Header sesuain dengan header scriptmu
	OBS_Header := `{
    "AuthToken": "` + AuthToken + `",
    "Msg_Id": "` + Msg_Id + `",
    "Device": "ANDROID",
    "Version": "11.10.2",
    "System_Name": "Android OS",
    "System_Ver": "10.0.2",
    "x-lal": "en_US"
  }`
	requestBody := strings.NewReader(OBS_Header)
	res, err := http.Post(hosts+"change_profile_picture="+apikey, "application/json; charset=UTF-8", requestBody)
	if err != nil {
		fmt.Println("Gagal")
		return
	}
	if res.StatusCode == 200 {
		fmt.Println("Success update profile picture")
	}
}

func ChangeProfileVideo(AuthToken string, Msg_Id string) {
	//Header sesuain dengan header scriptmu
	OBS_Header := `{
    "AuthToken": "` + AuthToken + `",
    "Msg_Id": "` + Msg_Id + `",
    "Device": "ANDROID",
    "Version": "11.10.2",
    "System_Name": "Android OS",
    "System_Ver": "10.0.2",
    "x-lal": "en_US"
  }`
	requestBody := strings.NewReader(OBS_Header)
	res, err := http.Post(hosts+"change_profile_video="+apikey, "application/json; charset=UTF-8", requestBody)
	if err != nil {
		fmt.Println("Gagal")
		return
	}
	if res.StatusCode == 200 {
		fmt.Println("Success update profile video")
	}
}

func ChangeCoverPicture(AuthToken string, Msg_Id string) {
	//Header sesuain dengan header scriptmu
	OBS_Header := `{
    "AuthToken": "` + AuthToken + `",
    "Msg_Id": "` + Msg_Id + `",
    "Device": "ANDROID",
    "Version": "11.10.2",
    "System_Name": "Android OS",
    "System_Ver": "10.0.2",
    "x-lal": "en_US"
  }`
	requestBody := strings.NewReader(OBS_Header)
	res, err := http.Post(hosts+"change_cover_picture="+apikey, "application/json; charset=UTF-8", requestBody)
	if err != nil {
		fmt.Println("Gagal")
		return
	}
	if res.StatusCode == 200 {
		fmt.Println("Success update profile video")
	}
}

func ChangeCoverVideo(AuthToken string, Msg_Id string) {
	//Header sesuain dengan header scriptmu
	OBS_Header := `{
    "AuthToken": "` + AuthToken + `",
    "Msg_Id": "` + Msg_Id + `",
    "Device": "ANDROID",
    "Version": "11.10.2",
    "System_Name": "Android OS",
    "System_Ver": "10.0.2",
    "x-lal": "en_US"
  }`
	requestBody := strings.NewReader(OBS_Header)
	res, err := http.Post(hosts+"change_cover_video="+apikey, "application/json; charset=UTF-8", requestBody)
	if err != nil {
		fmt.Println("Gagal")
		return
	}
	if res.StatusCode == 200 {
		fmt.Println("Success update profile video")
	}
}

func DisableLetterSealing(AuthToken string) {
	//Header sesuain dengan header scriptmu
	Headers := `{
    "AuthToken": "` + AuthToken + `",
    "Msg_Id": "",
    "Device": "ANDROID",
    "Version": "11.10.2",
    "System_Name": "Android OS",
    "System_Ver": "10.0.2",
    "x-lal": "en_US"
  }`
	requestBody := strings.NewReader(Headers)
	res, err := http.Post(hosts+"rm_LetterSealing="+apikey, "application/json; charset=UTF-8", requestBody)
	if err != nil {
		fmt.Println("Gagal")
		return
	}
	if res.StatusCode == 200 {
		fmt.Println("Disable LetterSealing Success")
	}
}

func main() {
	DisableLetterSealing("TOKEN")
	ChangeProfilePicture("TOKEN", "MSG_ID")
	ChangeProfileVideo("TOKEN", "MSG_ID")
	ChangeCoverPicture("TOKEN", "MSG_ID")
	ChangeProfileVideo("TOKEN", "MSG_ID")
}
