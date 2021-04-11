FROM golang:alpine as builder
RUN mkdir /build 

ADD ./main.go /build/
ADD ./go.mod /build/
ADD ./go.sum /build/

ADD ./vendor /build/vendor
ADD ./di /build/di
ADD ./src /build/src

WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
EXPOSE 3000
CMD ["./main"]