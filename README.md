# spaceportal
Nasa Apod App in Golang

The code needs a the certs and settings.json to be mapped to the root dir of the app


format for settings.json

{
    "apod":{
        "Apikey":"API_KEY",
        "apod_end":"https://api.nasa.gov/planetary/apod"
    }
    
}

cert folder will have the cert and key to be names as apod.cert and apod.key

With Docker:

docker run -p 443:9090 -v <your location >/certs:/app/certs -v <your location>/settings.json:/app/settings.json spaceportal
