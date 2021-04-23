import requests, json, time

hostVH = "https://api.vhtear.com/"
apikey = "PREMIUMKEY" #Chat me on whatsapp https://wa.me/+628123552767

#Header sesuai'in dengan scripmu
OBS_HEADER = {
  "AuthToken": "tokenMu",
  "Msg_Id": "",
  "Device": "DESKTOPWIN",
  "Version": "6.7.0",
  "System_Name": "VH-PC",
  "System_Ver": "10.0.14",
  "x-lal": "en_id"
}

def ChangeProfilePicture(msg_id):
  OBS_HEADER.update({"Msg_Id": msg_id})
  sagne = requests.post(hostVH+"change_profile_picture=%s" % apikey, json=OBS_HEADER)
  if sagne.status_code == 200:
     tH = sagne.json()
     print(tH) 
     print("Success update profile picture") 

#ChangePicture("13936590569300")

def ChangeProfileVideo(msg_id):
  OBS_HEADER.update({"Msg_Id": msg_id})
  sagne = requests.post(hostVH+"change_profile_video=%s" % apikey, json=OBS_HEADER)
  if sagne.status_code == 200:
     tH = sagne.json()
     print(tH) 
     print("Success update profile video") 
     
#ChangeProfileVideo("13936590569300")

def ChangeCoverImage(msg_id):
  OBS_HEADER.update({"Msg_Id": msg_id})
  sagne = requests.post(hostVH+"change_cover_image=%s" % apikey, json=OBS_HEADER)
  if sagne.status_code == 200:
     tH = sagne.json()
     print(tH) 
     print("Success update cover picture") 
     
#ChangeCoverVideo("13936590569300")

def ChangeCoverImage(msg_id):
  OBS_HEADER.update({"Msg_Id": msg_id})
  sagne = requests.post(hostVH+"change_cover_video=%s" % apikey, json=OBS_HEADER)
  if sagne.status_code == 200:
     tH = sagne.json()
     print(tH) 
     print("Success update cover video") 
     
#ChangeCoverVideo("13936590569300")
         
