# M&M_App
# MMMapp Forte - Production Monitoring System – OEE & IoT (Go)

A real-time monitoring system for machine effectiveness (OEE), collecting data from MQTT and REST APIs, and storing results in a PostgreSQL (TimeScaleDB) database. The system is containerized using Docker Compose and supports environment configuration through a `.env` file.

## Features

- Reads data from MQTT brokers (IO-Link sensors).
- Retrieves energy analyzer data via REST API.
- Calculates OEE indicators (`availability`, `performance`, `quality`, `cycle`, `OEE`).
- Aggregates and saves data to:
  - JSON files (`oee.json`, `mqttOEE.json`, `mqttFlow.json`, `meters.json`, `measurements.json`)
  - PostgreSQL / TimeScaleDB tables
- Automatic shift detection and summary (3-shift scheduler).
- Handles CET/CEST time zones (Polish local logic).
- Fully dockerized with `docker-compose`.

## Project Structure

```
MMMapp-FORTE/
├── go_app/                        # Main Go application
│   ├── main.go
│   ├── go.mod, go.sum
│   ├── Dockerfile
│   ├── communication/
│   ├── config/
│   ├── core/
│   ├── db/
│   ├── utils/
│   └── logs/
│
├── docs/                          # Technical documentation
│   ├── architecture.md
│   ├── shift_schedule.md
│   └── api.md
│
├── grafana_queries/              # SQL queries for Grafana dashboards
│   ├── oee_by_shift.sql
│   ├── energy_consumption.sql
│   ├── performance_vs_nominal.sql
│   ├── flow_trend.sql
│   └── documentation.md
│
├── deploy/                        # Infrastructure & deployment files
│   ├── docker-compose.yml
│   ├── go_init.sql
│   └── rebuild-go.sh
│
├── .gitignore
├── LICENSE.md
├── LICENSE_OVERVIEW.md
├── THIRD_PARTY_LICENSES.md
└── README.md
```

## Dockerfile Overview

The Go application is built in two stages:

```Dockerfile
# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app .

# Final minimal image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
COPY logs/ ./logs/
CMD ["./app"]
```

## Installation (with Docker Compose)

1. Clone the repository and navigate to the project directory.
2. Create or edit the `.env` file to set credentials and service settings:

```env
# Grafana
GRAFANA_USER=admin
GRAFANA_PASSWORD=your_password

# Database
DB_HOST=timescaledb_go
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=your_db_password
DB_NAME=monitoring_db_go

# MQTT
MQTT_BROKER=192.168.1.XXX
MQTT_PORT=1883
MQTT_USER=user
MQTT_PASSWORD=your_mqtt_password

# Analyzer IPs
ANALYZER_IP01=192.168.1.XXX
ANALYZER_IP02=192.168.1.XXX
ANALYZER_IP03=192.168.1.XXX
ANALYZER_IP04=192.168.1.XXX
ANALYZER_IP05=192.168.1.XXX
```

3. Run the system:

```bash
docker compose -f deploy/docker-compose.yml up --build -d
```

4. Access Grafana via [http://localhost:3000](http://localhost:3000) using credentials from `.env`.

## Requirements (non-Docker)

- Go 1.18+
- PostgreSQL with TimeScaleDB extension
- MQTT broker (e.g., BAV00L2, Mosquitto on PC)
- REST-enabled energy analyzers (optional — mock data included)

## Database Overview

The system uses the following database tables (see `go_app/db/*.sql`):

- `oee_calculated`, `oee_calculated_static`
- `shift_summary`
- `flow_data`
- `measurements`
- `meters_total_temp`
- `meters_t1_temp`, `meters_t2_temp`, `meters_t3_temp`, `meters_t4_temp`

## Database & Visualization

### TimescaleDB

Stores time-series production data:
- OEE metrics: `oee_calculated`, `shift_summary`
- Electrical: `measurements`, `meters_*_temp`
- Flow: `flow_data`

### Grafana

Visualizes:
- OEE (availability, performance, quality)
- Energy usage
- Flow and temperature

SQL queries are stored in the `grafana_queries/` folder.

## JSON Output Format

### `oee.json`

```json
{
  "timestamp": "2025-04-02T20:38:33Z",
  "Predkosc_obrotnica": 135.2,
  "ilosc_elementow": 73,
  "czas_pracy": 188.5,
  "czas_postoju": 12.3
}
```

### `mqttOEE.json` and `mqttFlow.json`

Raw MQTT data from IO-Link master mapped to simplified field names.

Port mapping example:

- `port1`, `port2` → machine signals (OEE)
- `port3`, `port4` → flowmeter and sensor data

## Author

This project was developed by **kopyn00** and **GitMichal00** in Go to monitor industrial production systems using real-time IoT data, OEE metrics, and TimeScaleDB.

> ⚠️ Ensure Docker, Docker Compose, and the required environment variables are properly configured before starting the system.

