FROM ubuntu:16.04

RUN apt-get update && apt-get install -y build-essential cmake git sudo wget golang-go
