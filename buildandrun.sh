#!/bin/bash
set -e

docker build ~/src/blockchain-indexer/build/base -t blkidx-base
docker build ~/src/blockchain-indexer -t blkidx 
docker build . -t blkidx-insight
docker build ui/ -t blkidx-insight-ui

docker-compose up