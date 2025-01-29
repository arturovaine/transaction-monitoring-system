# Transaction Monitoring System

<details>
<summary><span style="font-size: 1.5em; font-weight: bold;">Requirements and Specifications</span></summary>


### Objective

- Design and partially implement a data-intensive system to evaluate your skills in system
design, database selection, and scaling stateful applications in the fintech domain.

### Scenario

- Real-time transaction monitoring system for a fintech
platform that processes payments. 

- The system will analyze payment data and provide
insights for fraud detection, regulatory reporting, and operational monitoring.

- The system should generate and track the following metrics:
    - Suspicious Activity: Detect transactions above a configurable threshold (e.g.,
$10,000) or transactions flagged as potential fraud.
    - Daily Settlements: Aggregate total transactions and volume settled per region daily.

### Requirements

1. High-Level Design:

- Create a high-level design for the system, detailing the architecture,
components, and data flow.

- Highlight how the system will ensure low-latency processing and handle high
transaction volumes.

2. Database Design:

- Select one or more databases for this system and justify your choice (e.g., relational
for transactional consistency, NoSQL for scalability).

- Design a schema or data model to store the transactions and support the required
metrics.

3. Scalability and Resilience:

- Explain how you would scale the system to handle increasing transaction volumes.
Discuss strategies for ensuring high availability and fault tolerance.

4. Constraints and Assumptions

- Assume the system processes up to 10,000 transactions per second during peak
times.

- The system should support a multi-region architecture for regulatory compliance.

</details>

<br/>

<details open>
  <summary><span style="font-size: 1.5em; font-weight: bold;">Solution</span></summary>


### Setup and execution

1. Clone the repo:

    `git clone <repo>`

2. Build the project image (Docker)

    `docker compose build`

3. Run the project (Docker)

    `docker compose up`

### System architecture / structure

```
                  +-------------------------------------------------+
                  |        Transaction Monitoring System           |
                  +-------------------------------------------------+

                            +----------------------+
                            |       Frontend       |
                            |  (Vite / React App)  |
                            +----------------------+
                                     |
                                     v
                            +----------------------+
                            |       Backend        |
                            |  (Go API Service)    |
                            +----------------------+
                                     |
         +--------------------------+--------------------------+
         |                           |                          |
         v                           v                          v
+--------------------+    +----------------------+    +----------------------+
|   PostgreSQL      |    |   Kafka (Event Bus)  |    |  Prometheus Exporter  |
|  (Transactional DB)|    |                      |    |  (Monitoring Metrics) |
+--------------------+    +----------------------+    +----------------------+
         |                           |                          |
         v                           v                          v
+--------------------+    +----------------------+    +----------------------+
|  Grafana (UI)     |    |    Zookeeper         |    |   Prometheus (Metrics)|
|  (Visual Metrics) |    |  (Kafka Management)  |    |  (Data Collection)    |
+--------------------+    +----------------------+    +----------------------+
                                     |
                                     v
                            +----------------------+
                            |  Kafka Broker        |
                            |  (Message Streaming) |
                            +----------------------+

```



### Components
- Frontend (Vite/React) → Displays transaction monitoring dashboards.
- Backend (Go API) → Handles transaction logic, fraud detection, and integrates with Kafka.
- PostgreSQL → Stores raw transaction data.
- Kafka → Streams transaction events for real-time monitoring.
- Zookeeper → Manages Kafka brokers.
- Prometheus Exporter → Extracts PostgreSQL metrics for monitoring.
- Prometheus → Collects and aggregates system metrics.
- Grafana → Provides dashboards for visualizing transaction statistics.

#### Front

> http://localhost:5173/

#### API

> GET http://localhost:8080/suspicious

> GET http://localhost:8080/transactions

### Connections
- Frontend Network (frontend) → Connects React with the backend.
- Backend Network (backend) → Connects Go API with PostgreSQL.
- Kafka Network (kafka) → Links Kafka and Zookeeper.
- Monitoring Network (monitoring) → Links Prometheus, Grafana, and PostgreSQL Exporter.

### Data Flows

- Frontend calls the backend API.
- Backend fetches transactions from PostgreSQL.
- Backend streams real-time transactions to Kafka.
- Kafka events are processed and stored.
- Prometheus collects data from PostgreSQL and Kafka.
- Grafana visualizes system metrics.

</details>

<br/><br/>
