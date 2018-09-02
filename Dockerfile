FROM golang:1.11.0

WORKDIR /go/src/ImageProvider


RUN go get -d -v "github.com/gorilla/mux" && \
        go get -d -v "github.com/gofrs/uuid"

COPY . .
RUN go install

CMD ["ImageProvider"]