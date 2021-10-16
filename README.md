# Distributed Systems

These applications were built with the objective of studding a distributed systems using the most recent technics. The main ideia was create all applications using: **Clean Architecture**, **Distributed Logging**, **Distributed Tracing**, **Async Communication**, **CI/CD**, **Loading Balancer**, **Auto scale**, etc. The system context was a simple book store.

## Application Map

- WebApi:
  - Responsible for create a user, authenticate a user and communicate with the others applications using gRPC.
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