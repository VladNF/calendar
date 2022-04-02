FROM golang:1.16.2 as build

ENV BIN_FILE /opt/calendar/sender-app
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY go.* .
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/sender/*

FROM alpine:3.9

LABEL ORGANIZATION="VladNF"
LABEL SERVICE="sender"
LABEL MAINTAINERS="j@vladnf.ru"

ENV BIN_FILE "/opt/calendar/sender-app"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

#ENV CONFIG_FILE /etc/calendar/sender.yaml
#COPY ./configs/sender_config.yaml ${CONFIG_FILE}
#CMD ${BIN_FILE} -config ${CONFIG_FILE}
CMD ${BIN_FILE}
