FROM golang

WORKDIR /app

COPY . .

RUN go get -u github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/go-playground/validator.v9

RUN go build -o main .

EXPOSE 3000

CMD ["./main"]