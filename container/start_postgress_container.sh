#!/usr/bin/env bash
docker run -d --network ba_network --name=postgres -p 5432:5432 -e POSTGRES_PASSWORD=ba -e POSTGRES_USER=ba postgres -c shared_buffers=256MB -c max_connections=550
