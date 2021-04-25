package main

import (
	"fmt"
	"net/http"
	"strings"
)

var hosts = "https://api.vhtear.com/"
var apikey = "PremiumKey"

func ChangeProfilePicture(AuthToken string, Msg_Id string) {
	//Header sesuain dengan header scriptmu
	OBS_Header := `{
    "AuthToken": "` + AuthToken + `",
    "Msg_Id": "` + Msg_Id + `",
    "Device": "DESKTOPWIN",
    "Version": "6.7.0",
    "System_Name": "VH-PC",
    "System_Ver": "10.0.14",
    "x-lal": "en_id"
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
    "Device": "DESKTOPWIN",
    "Version": "6.7.0",
    "System_Name": "VH-PC",
    "System_Ver": "10.0.14",
    "x-lal": "en_id"
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
    "Device": "DESKTOPWIN",
    "Version": "6.7.0",
    "System_Name": "VH-PC",
    "System_Ver": "10.0.14",
    "x-lal": "en_id"
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
    "Device": "DESKTOPWIN",
    "Version": "6.7.0",
    "System_Name": "VH-PC",
    "System_Ver": "10.0.14",
    "x-lal": "en_id"
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

func main() {
	ChangeProfilePicture("TOKENMU", "MSG_ID MU COK")
	ChangeProfileVideo("TOKENMU", "MSG_ID MU COK")
	ChangeCoverPicture("TOKENMU", "MSG_ID MU COK")
	ChangeProfileVideo("TOKENMU", "MSG_ID MU COK")
}
