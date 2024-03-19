# Simple Blog CRUD with Golang, Redis, and MongoDB on Docker

This project implements a simple blog application with CRUD (Create, Read, Update, Delete) operations using Golang. The application is containerized using Docker and utilizes Redis for caching and MongoDB as the database.

## Prerequisites
- Docker
- Docker Compose

## Getting Started
1. Clone the repository
2. Run the containers using Docker Compose
    - docker-compose up -d

3. Access the application at http://localhost:8080
4. API end point at http://localhost:3000


## Project Structure
- frontend/: use vuejs as the frontend
- backend/: Contains the Golang code for the blog backend
- docker-compose.yml: Defines the services and their configuration


## Services
- frontend: Golang blog frontend
- backend: Golang blog backend
- db: MongoDB database
- db-client: Mongo Express for database management
- redis: Redis for caching

## Usage
- The blog application provides RESTful APIs for managing blog posts. You can perform CRUD operations on blog posts using the provided APIs.