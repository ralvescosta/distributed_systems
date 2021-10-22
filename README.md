# Distributed Systems

![Go Lang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) ![Rust Lang](https://img.shields.io/badge/Rust-000000?style=for-the-badge&logo=rust&logoColor=white) ![node-current](https://img.shields.io/badge/Node.js-43853D?style=for-the-badge&logo=node.js&logoColor=white)

## Contents
- [Distributed Systems](#distributed-systems)
  - [Contents](#contents)
  - [Resume](#resume)
  - [Project Mapping](#application-mapping)
  - [Installation](#installation)

## Resume

These applications were built with the objective of studding a distributed systems using the most recent technics. The main ideia was create all applications using: **Clean Architecture**, **Distributed Logging**, **Distributed Tracing**, **Async Communication**, **CI/CD**, **Loading Balancer**, **Auto scale**, etc. The system context was a simple book store.

## Application Mapping

- WebApi:
  - Responsible for create a user, authenticate a user, communicate with the others applications using gRPC and RabbitMQ.
  - Used Technologies: 
    - Built in GoLang
    - PostgreSQL
    - RabbitMQ Publisher
    - gRPC client
    - Cache in some Routes


- Inventory ms: Built in RustLang
  - Responsible for story and management all books.
  - Used Technologies:
    - Built in RustLang
    - MongoDB
    - gRPC Server
    - RabbitMQ Subscriber


- Purchase ms: Built in NodeJs
  - Responsible to management the purchases and the payments
  - Used Technologies:
    - Built in NodeJs
    - MongoDB
    - gRPC Server
    - RabbitMQ Subscriber


- Mailer ms: Built in GoLang
  - Responsible for send emails to the book buyer and the book seller
  - Used Technologies:
    - Built in GoLang
    - MongoDB
    - RabbitMQ Subscriber