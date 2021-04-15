FROM golang:latest
RUN mkdir /goProject
COPY . /goProject
WORKDIR /goProject
RUN go install 
RUN go build -o studentapi .