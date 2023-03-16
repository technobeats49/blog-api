# blog-api

This is a sample project. It follows simple, standard design methodology and Model-Routes-Controllers-Services
Application structure.
<img width="265" alt="image" src="https://user-images.githubusercontent.com/86374489/225626250-1e91d973-94d7-4016-a0b5-ea6e58f1df34.png">


It is using mongoDB as Datastore. // Download MongoDB community in your local 
Project uses standard packages and libraries for building REST APIs, testify for composing Unit tests.

HTTP Server port: 9000
MongoDB Server port: 27017. // Standard, default

Project Run Instructions:

Pull the project blog-api in your local


1. go to blog-api directory and run 

   `go mod tidy`           
2. Run the main 

   `go run main.go`


Test Endpoints 

1. Create Article


   `POST   http://localhost:9000/articles`

    `payload:       
      {
          "title": "Go powers Distributed Systems",
          "content": "Distributed Systems development has accelerated since go is developed",
          "author": "Go Enthusiast"
      }`

    Response Body

 
    `
         {
             "data": {
                 "id": "64123b24cc9d402cf4ace609"
             },
             "message": "success",
             "status": 201
         }`
