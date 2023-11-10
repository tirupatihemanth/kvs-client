#!/bin/bash
$(docker-compose -f ../kvs/docker-compose.yml up > server_logs.txt) &
sleep 10
echo "Sending Requests"
./kvs-client $1 $2
docker-compose -f ../kvs/docker-compose.yml down