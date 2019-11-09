#!/usr/bin/env bash

protoc -I common/workercomms/ common/workercomms/worker_comms.proto --go_out=plugins=grpc:common/workercomms
