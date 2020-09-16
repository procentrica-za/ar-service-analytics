FROM golang:alpine
RUN apk add --no-cache git &&\
go get "github.com/gorilla/mux" &&\
go get -u golang.org/x/oauth2/google &&\
go get -u google.golang.org/api/sheets/v4 
ADD /src/ /app/
WORKDIR /app/
RUN go build -o main .
CMD ["/app/main"]

