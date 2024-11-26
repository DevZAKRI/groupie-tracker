# Groupie-Tracker

## Objectives

Groupie Tracker is a web application that allows users to keep track of their favorite music groups.
Users can view a list of music groups and see details about each group.
The application is built using the Go programming language.

- **Setting a Server to Listen to incoming requests**
- **Fetching all necessary Data from the given API**
- **Displaying the Data when requested**
- **Implementing Good visualisations and handling known errors for us up to now**

## Instructions

### Web Server

The web server is built in **Go** and adheres to recommended coding standards. Only standard Go packages are used for simplicity and compatibility.

### Docker Containerization

The application is containerized using Docker, following best practices in Dockerfile creation. The container includes:
- Necessary configurations for a lightweight, efficient image
- Metadata annotations for Docker objects to improve traceability and maintainability

### Good Practices

- **Code Quality**: The code follows standard Go practices for readability, maintainability, and performance.
- **Dockerfile Quality**: The Dockerfile is optimized and avoids unnecessary layers or files, ensuring a clean and efficient image.

## Allowed Packages

Only Go's standard library packages are used to ensure compatibility and simplicity.

## Features

This project provides hands-on experience with the following:

- **Web Fundamentals**: Understanding how to set up a server, serve HTML, handle HTTP requests, and manage data flow.
- **API Usage**: Learn How to fetch Data from API and how to manipulate that data and use it.
- **Docker Basics**: Learning how to install, configure, and use Docker.

## Usage

```go run .``` or u can run already made docker image ```docker run -p 8080:8080 mostaphazakri/groupie-tracker:latest```

---

### Contributors

- **Mostafa ZAKRI**: [mzakri](https://learn.zone01oujda.ma/git/mzakri)
- **Elbadaoui Houssam**: [helbadao](https://learn.zone01oujda.ma/git/helbadao)

