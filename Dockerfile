FROM golang:1.19-alpine

WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -v -o /bin/staffd cmd/staffd/main.go

ENTRYPOINT [ "/bin/staffd" ]
CMD ["/bin/staffd"]