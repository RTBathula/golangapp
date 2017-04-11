FROM golang
 
ADD . /go/src/github.com/rtbathula/golangapp
RUN go install github.com/rtbathula/golangapp
ENTRYPOINT /go/bin/golangapp
 
EXPOSE 3000