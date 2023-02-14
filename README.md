# Golang - Upload API Microservice

This project is a microservice developed using Golang that is responsible for uploading files to the Minio object store. The Minio object store is an open-source object storage system that can be used to store and manage unstructured data such as photos, videos, documents, and more.
Features

The `go-upload-api` microservice provides the following features:

- Uploading files to Minio object store
- Support for uploading multiple files at once
- Configurable bucket name and object name for uploaded files
- Support for setting object metadata, such as content type and content encoding

## Getting Started

To get started with this project, you will need to have Golang and Minio installed on your system. Once you have installed the dependencies, you can follow these steps:

1. Clone the repository: git clone https://github.com/dalekurt/go-upload-api-microservice.git

2. Navigate to the project directory: cd go-upload-api-microservice

3. Install the required dependencies: go mod download

4. Set the required environment variables:

5. Build and run the microservice: `docker-compose up`

## Usage

To use the `go-upload-api microservice`, you can send a POST request to the `/` endpoint with the files you want to upload as the request body. The microservice will return a response with the status of the upload.

For example, you can use the following cURL command to upload a file:

```shell
curl -X POST -H "Content-Type: multipart/form-data" -F "file=@/path/to/file.jpg" http://localhost:8080/
```
