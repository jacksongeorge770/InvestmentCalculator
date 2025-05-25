Overview

The Investment Calculator App is a web application that calculates compound interest based on user inputs (principal, interest rate, compounds per year, and years). It features a Golang backend serving a REST API (/api/v1/calculate) and a static HTML/CSS/JavaScript front-end with a form and Chart.js visualization. The backend connects to an AWS RDS MySQL database for storing calculations, uses Goose for migrations, and is deployed as a Docker container on AWS EC2 instances behind an Application Load Balancer (ALB). Infrastructure is provisioned using Terraform.


## 📦 Technologies Used

- **Golang** – Backend logic and API development  
- **Gorilla Mux** – HTTP request routing and middleware handling  
- **MySQL** – Relational database management  
- **Docker** – Containerization and environment consistency  
- **Terraform** – Infrastructure as Code (IaC) for provisioning and managing AWS resources  
- **AWS** – Cloud deployment using:
  - EC2 (Elastic Compute Cloud) for hosting backend services  
  - RDS (Relational Database Service) for MySQL  
  - ALB (Application Load Balancer) for traffic distribution

## 📦 Technologies Used

- **Golang** – Backend logic and API development  
- **Gorilla Mux** – HTTP request routing and middleware handling  
- **PostgreSQL** – Relational database management  
- **Docker** – Containerization and environment consistency  
- **Terraform** – Infrastructure as Code (IaC) for provisioning and managing AWS resources  
- **AWS** – Cloud deployment using:
  - EC2 (Elastic Compute Cloud) for hosting backend services  
  - RDS (Relational Database Service) for database storage  
  - ALB (Application Load Balancer) for traffic distribution

---

## 🚀 Features

- **API Endpoint**  
  `POST /api/v1/calculate` – Accepts JSON input:
  ```json
  {
    "principal": 1000,
    "rate": 0.05,
    "compoundsPerYear": 4,
    "years": 5
  }
  {
  "result": 1282.037231708585
  }




- CI/CD pipeline using GitHub Actions
- Deployed on AWS EC2 with Application Load Balancer
- Credentials managed securely via GitHub Secrets
- Self-hosted GitHub runners on two EC2 instances for scalability

---

