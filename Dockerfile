FROM golang:1.19

RUN apt-get update && apt-get install libpcap-dev -y
WORKDIR /app
COPY . . /app/
RUN go get
RUN go build -o request-fp
EXPOSE 8080
CMD ["./request-fp"]