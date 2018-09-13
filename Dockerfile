FROM golang:1.11

ENV APP_ENV prod
ENV APP_DEBUG false
ENV PORT 8080
ENV GIN_MODE=release

WORKDIR /go/src/financeproc

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["cmd"]

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "curl -f http://127.0.0.1:8080/healthcheck || exit 1" ]

EXPOSE 8080
