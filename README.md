# Chirpy

✨ A lightweight, fast, and modern microservice for posting short “chirps” — built with Go, SQL, and clean REST design.

------------------------------------------------------------
## 🏷️ Badges

![Go Version](https://img.shields.io/badge/Go-1.x-blue?logo=go)
![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)
![Build](https://img.shields.io/badge/Build-Passing-brightgreen)
![PRs Welcome](https://img.shields.io/badge/PRs-Welcome-purple)
![Status](https://img.shields.io/badge/Project-Active-success)

------------------------------------------------------------
## 📌 Overview

Chirpy is a sleek microservice for user authentication and short-form text posting. It features account creation, secure JWT authentication, refresh tokens, and full CRUD operations for chirps.

Perfect for learning Go microservices, practising SQL database interactions, or building into a bigger microblogging system.

------------------------------------------------------------
## 🚀 Features

• User registration and authentication  
• JWT access & refresh token system  
• Create, retrieve, list, and delete chirps  
• SQL database powered via sqlc  
• Health and metrics endpoints  
• Clean and maintainable Go architecture  
• Easy to deploy, simple to extend  

------------------------------------------------------------
## 🏗️ Project Architecture

Technology stack:  
• Go (backend logic)  
• SQL + sqlc (database & typed queries)  
• HTTP REST API  

Directory Layout (simplified):

assets/               Static assets  
internal/             Internal Go modules  
sql/                  SQL schema + sqlc queries  
cmd/                  Application entrypoints  
main.go               Program start  

------------------------------------------------------------
## ⚙️ Getting Started

### Prerequisites
• Go 1.x+  
• PostgreSQL (or another SQL backend)  
• sqlc installed  

### Installation

1. Clone the repository:
   git clone https://github.com/shinrin-yoku92/Chirpy.git
   cd Chirpy

2. Generate SQL code:
   sqlc generate

3. Build the application:
   go build -o chirpy

### Run the server:
   ./chirpy

Server defaults to port: 8080

------------------------------------------------------------
## 📡 API Reference

POST    /users/create       Register a user  
POST    /users/login        Authenticate + tokens  
POST    /chirps             Create a chirp  
GET     /chirps             List all chirps  
GET     /chirps/{id}        Get a specific chirp  
DELETE  /chirps/{id}        Remove a chirp  
POST    /token/refresh      Refresh JWT token  

------------------------------------------------------------
## 🔧 Configuration

Environment variables:

DB_DSN        Database connection string  
JWT_SECRET    Secret key for signing tokens  
PORT          Application port  

------------------------------------------------------------
## 🧪 Testing

Run all tests:
   go test ./...

------------------------------------------------------------
## 🚢 Deployment

Chirpy can be deployed easily via:

• Docker  
• Kubernetes  
• AWS / GCP / DigitalOcean  
• GitHub Actions CI/CD  

------------------------------------------------------------
## 🤝 Contributing

Want to improve Chirpy? Great!

1. Fork the repo  
2. Create a feature branch  
3. Implement changes  
4. Submit a pull request  

All contributions are welcome.

------------------------------------------------------------
## 📄 License

This project is licensed under the MIT License.

------------------------------------------------------------
## 🙏 Acknowledgements

• Inspired by minimalist microblogging platforms  
• Built on the strength of the Go open-source community  
