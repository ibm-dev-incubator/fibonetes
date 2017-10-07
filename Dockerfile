FROM python:3


RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.8.0/bin/linux/amd64/kubectl

RUN chmod +x kubectl

ADD fib.py fib.py
