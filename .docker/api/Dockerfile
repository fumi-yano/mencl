FROM golang:latest
ENV ROOTDIR /go
ENV APP_HOME /go/src/api
RUN mkdir -p $APP_HOME
ENV GOPATH $ROOTDIR
ENV GOBIN $ROOTDIR/bin

COPY main.go $APP_HOME
COPY start_app.sh $APP_HOME

WORKDIR $APP_HOME

RUN go get github.com/oxequa/realize \
  && apt-get update \
  && apt-get install -y mariadb-client \
  && echo "export PATH=$PATH:$GOBIN" > /root/.bashrc \
  && echo "export GO111MODULE=on" \
  && go mod init

CMD ["sh", "./start_app.sh"]
