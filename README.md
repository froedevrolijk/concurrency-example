# concurrency-example

### Prerequisites
* Go
* Docker Desktop
* Google Cloud SDK

### Run locally
* `go run main.go`  
* `curl http://localhost:8080/serve-customer/3` from new terminal window while container is running   

### Run using Docker
* `make run` from folder root to build docker container 
* `curl http://localhost:8080/serve-customer/3` from new terminal window while container is running 
* `docker logs coffee-shop` to see the output
* `make stop` to stop container and remove image  

## Deploy to Cloud Run on GCP
Create a GCP project and create your project configuration using:  
`gcloud config configurations create [CONFIGURATION_NAME]`

Check if the configuration is activated:  
`gcloud config configurations list`

Allow the SDK to find credentials automatically:  
`gcloud auth application-default login`

Enable Cloud Run API:  
`gcloud services enable run.googleapis.com`

Deploy the code from source. From folder root run (and accept defaults):  
`gcloud run deploy`

Navigate to the Cloud Run address in a web browser, e.g.:  
`https://concurrency-example-pr5khuikua-ew.a.run.app/serve-customer/3`

Check the logs in the GCP console. You should see something like:  
![Alt text](cloud-run-logs.png?raw=true "GCP Logs")
