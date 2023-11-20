
<!-- PROJECT LOGO -->
<br />
<div align="center">
<h3 align="center">Distributed Logging System Architecture</h3>
</div>





<!-- ABOUT THE PROJECT -->
## About The Project

The project involves the development of a logging system that utilizes Docker for containerization and the ELK (Elasticsearch, Logstash, Kibana) stack for processing, storing, and visualizing logs. Golang was chosen as the programming language for its efficiency and performance in handling concurrent tasks.

## System Architecture

The system architecture consists of the following components:

* Docker Containers: The application and its dependencies are encapsulated within Docker containers, ensuring consistency and portability across different environments.

* Golang Application: The core logging functionality is implemented in Golang. The application is responsible for collecting and forwarding logs to the Logstash component.

* Logstash: Logstash serves as the log ingestion and processing engine. It processes incoming logs, parses them, and forwards them to Elasticsearch for storage.

* Elasticsearch: Elasticsearch is used as the central storage for logs. It provides fast and scalable full-text search capabilities, allowing for efficient log retrieval and analysis.

* Kibana: Kibana is the web interface used for visualizing and analyzing logs stored in Elasticsearch. It provides powerful search and visualization tools.



<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* golang
  ```sh
  brew install go
  ```
* Docker
  ```sh
  brew install docker
  ```

### Installation and Run

1. Clone the repo
   ```sh
   git clone https://github.com/AnirudhAgnihotri2902/LoggingSystem.git
   ```
2. Install go packages
   ```sh
   go mod tidy
   ```

3. Run Docker Setup
    ```sh
   docker-compose up setup
   ```
   
4. Run Main Docker Compose
    ```sh
   docker-compose up
   ```

5. Go to cmd Folder
    ````sh
    cd cmd
    ````
   
6. Run main.go file
    ```sh
   go run main.go
    ```




<!-- CONTACT -->
## Contact

Name - Anirudh Agnihotri

Email - anirudhagnihotri82@gmail.com


