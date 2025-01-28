# Transaction Monitoring System

> git clone 

> docker compose build
> docker compose up

 ✔ frontend    
 ✔ grafana     
 ✔ backend     
 ✔ postgres    
 ✔ kafka      
 ✔ prometheus  
 ✔ zookeeper

### Front

> http://localhost:5173/

### API

> GET http://localhost:8080/suspicious

> GET http://localhost:8080/transactions


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
