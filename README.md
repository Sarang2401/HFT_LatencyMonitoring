# ⚡ HFT Latency Monitoring & Risk Observability Platform

A production-style observability platform designed to measure, visualize, and alert on **tail latency (P95/P99)** for performance-critical fintech and HFT systems.

---

## 📌 Project Overview

In low-latency trading systems, **tail latency** — not averages — determines risk and profitability. Traditional infrastructure monitoring often fails to surface early performance degradation.

This project implements an **application-level observability pipeline** that captures ingestion latency, exposes percentile-based metrics, and enables SLO-driven alerting.

---

## 🎯 Key Objectives

- Measure **end-to-end ingestion latency**
- Focus on **P50 / P95 / P99** rather than averages
- Implement histogram-based metrics
- Provide real-time observability and alerting
- Simulate realistic UDP-based market data ingestion

---

## 🏗️ Architecture Overview

UDP Feed Generator (Go)
↓
Ingestor Service (Go)
↓
Prometheus (Metrics & Alerts)
↓
Grafana (Dashboards)
↓
Downstream Risk Service


All components are deployed on **Kubernetes (k3s)** running on AWS EC2.

---

## 🧱 Core Components

- **Feed Generator (Go)**
  - Simulates high-frequency UDP market feeds
  - Supports burst and sustained load patterns

- **Ingestor Service (Go)**
  - Receives UDP packets
  - Measures ingestion latency using embedded timestamps
  - Exposes Prometheus histogram metrics

- **Prometheus**
  - Collects latency metrics
  - Evaluates SLO-based alert rules

- **Grafana**
  - Visualizes latency percentiles and throughput
  - Surfaces early warning signals

- **Risk Guard**
  - Represents downstream processing dependencies

---

## 📊 Observability & Alerting

### Key Metrics
- Histogram-based ingestion latency
- P50 / P95 / P99 latency percentiles
- Throughput and traffic patterns

### Alert Strategy
- P99 latency breach (critical)
- Sustained P95 degradation (warning)
- Traffic spike detection

Alerts are aligned with **Service Level Objectives (SLOs)** rather than infrastructure utilization.

---

## 🚀 Deployment & Usage

- AWS EC2 (Ubuntu)
- Docker for image builds
- k3s for Kubernetes orchestration
- Kubernetes manifests for services and deployments

### Validation
- Port-forward Grafana and Prometheus
- Generate UDP traffic
- Observe latency percentiles and alert triggers

---

## ⚠️ Operational Challenges Encountered

- Go version mismatches affecting builds
- Docker vs containerd image visibility in k3s
- Disk pressure causing pod scheduling failures
- Kubernetes service port misconfiguration

Each issue was debugged and resolved, mirroring real production scenarios.

---

## 📈 Outcome

The platform successfully:
- Measures real application-level latency
- Visualizes tail latency distributions
- Triggers alerts on SLO violations
- Demonstrates production-style observability

---

## 🧠 Key Learnings

- Tail latency is a critical risk signal in HFT systems
- Histogram metrics are essential for percentile accuracy
- Observability must be built into the application
- Kubernetes introduces real operational complexity

---

## 📄 Documentation

- Detailed design report included in the repository
- Architecture diagrams provided for reference

---

## 👤 Author

Built by *Sarang Shigwan*  
Focused on low-latency systems, observability, and fintech infrastructure.
