import requests, json, time

hostVH = "https://api.vhtear.com/"
apikey = "PREMIUMKEY" #Chat me on whatsapp https://wa.me/+628123552767

#Header sesuai'in dengan scripmu
Headers = {
  "AuthToken": "tokenMu",
  "Msg_Id": "",
  "Device": "ANDROID",
  "Version": "11.10.2",
  "System_Name": "Android OS",
  "System_Ver": "10.0.2",
  "x-lal": "en_US"
}

def ChangeProfilePicture(msg_id):
  Headers.update({"Msg_Id": msg_id})
  sagne = requests.post(hostVH+"change_profile_picture=%s" % apikey, json=Headers)
  if sagne.status_code == 200:
     tH = sagne.json()
     print(tH) 
     print("Success update profile picture") 

def ChangeProfileVideo(msg_id):
  Headers.update({"Msg_Id": msg_id})
  sagne = requests.post(hostVH+"change_profile_video=%s" % apikey, json=Headers)
  if sagne.status_code == 200:
     tH = sagne.json()
     print(tH) 
     print("Success update profile video") 

def ChangeCoverImage(msg_id):
  Headers.update({"Msg_Id": msg_id})
  sagne = requests.post(hostVH+"change_cover_picture=%s" % apikey, json=Headers)
  if sagne.status_code == 200:
     tH = sagne.json()
     print(tH) 
     print("Success update cover picture") 

def ChangeCoverImage(msg_id):
  Headers.update({"Msg_Id": msg_id})
  sagne = requests.post(hostVH+"change_cover_video=%s" % apikey, json=Headers)
  if sagne.status_code == 200:
     tH = sagne.json()
     print(tH) 
     print("Success update cover video") 
     
def DisableLetterSealing():
  sagne = requests.post(hostVH+"rm_LetterSealing=%s" % apikey, json=Headers)
  if sagne.status_code == 200:
     tH = sagne.json()
     print(tH) 
     print("Disable LetterSealing Success") 
     
         
