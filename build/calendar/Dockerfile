# Use Golang image to build
FROM golang:1.16.2 as build

ENV BIN_FILE /opt/calendar/calendar-app
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

# Cache layers with modules
COPY go.* .
RUN go mod download

COPY . ${CODE_DIR}

# Make a binary with static deps for it to work in alpine image
ARG LDFLAGS
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/calendar/*

# Output in a slim image
FROM alpine:3.9

LABEL ORGANIZATION="VladNF"
LABEL SERVICE="calendar"
LABEL MAINTAINERS="j@vladnf.ru"

ENV BIN_FILE "/opt/calendar/calendar-app"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

#ENV CONFIG_FILE /etc/calendar/calendar.yaml
#COPY ./configs/calendar_config.yaml ${CONFIG_FILE}
#CMD ${BIN_FILE} -config ${CONFIG_FILE}
CMD ${BIN_FILE}
