#!/bin/bash
docker build -t feed:latest feed/
docker build -t ingestor:latest ingestor/
docker build -t risk-guard:latest risk-guard/
