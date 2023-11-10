#!/bin/bash
./kvs-client $1 $2
docker-compose -f ../kvs/docker-compose.yml down