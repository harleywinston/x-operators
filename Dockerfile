FROM hub.hamdocker.ir/library/golang:1.20.3

WORKDIR /go/src/github.com/harleywinston/x-operators

COPY ./ .

RUN go build -buildvcs=false -o ./build/operator ./cmd

CMD ["./build/operator"]
