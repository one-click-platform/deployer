FROM golang:1.16

WORKDIR /go/src/github.com/one-click-platform/deployer

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/deployer github.com/one-click-platform/deployer


###

FROM alpine:3.9

COPY --from=0 /usr/local/bin/deployer /usr/local/bin/deployer
COPY --from=0 /go/src/github.com/one-click-platform/deployer/deploy-script /scripts
COPY --from=0 /go/src/github.com/one-click-platform/deployer/aws_conf /root/.aws/config
RUN apk add --no-cache ca-certificates
RUN apk -Uuv add groff less python py-pip
RUN pip install awscli
RUN apk --purge -v del py-pip
RUN rm /var/cache/apk/*




ENTRYPOINT ["deployer"]
