FROM golang:1.15-alpine AS builder

COPY main.go /app/
WORKDIR /app
RUN go build /app/main.go 

FROM alpine

RUN addgroup -S shipa && adduser -S shipa -G shipa
RUN chown -R shipa:shipa /home/shipa/

USER shipa
WORKDIR /shipa
EXPOSE 5001

COPY --from=builder /app/main /shipa/app
COPY shipa.yaml /shipa/

CMD ["/shipa/app"]
