

OpenSearch Logging Demo
========================
### Table of Contents

1. [Overview](#overview)
2. [Prerequisites](#prerequisites)
3. [Setup](#setup)
4. [Usage](#usage)
5. [Components](#components)
6. [License](#license)

### Overview

This repository provides a full Docker Compose setup for demonstrating real-time logging with OpenSearch and OpenSearch Dashboards. The setup includes a containerized web app that generates logs, which are then sent to OpenSearch for storage and visualization.

### Prerequisites

* Docker Compose installed on your machine
* Basic understanding of Docker and containerization

### Setup

1. Clone this repository to your local machine.
2. Navigate to the project directory.
3. Run `docker-compose up -d` to start the containers in detached mode.
4. Wait for the containers to start and OpenSearch to be available.

### Usage

1. Access the web app at `http://localhost:8080`.
2. Interact with the web app to generate logs.
3. Access OpenSearch Dashboards at `http://localhost:5601`.
4. Visualize the logs in OpenSearch Dashboards.

### Components

* **OpenSearch**: A search and analytics engine for storing and querying logs.
* **OpenSearch Dashboards**: A visualization tool for exploring and analyzing logs.
* **Web App**: A containerized web application that generates logs and sends them to OpenSearch.

### License

This project is licensed under the Apache License, Version 2.0. See the [LICENSE](LICENSE) file for details.
