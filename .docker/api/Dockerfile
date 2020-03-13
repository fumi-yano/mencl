FROM golang:latest
ENV ROOTDIR /go
ENV APP_HOME /go/src/api
RUN mkdir -p $APP_HOME
ENV GOPATH $ROOTDIR
ENV GOBIN $ROOTDIR/bin

COPY main.go $APP_HOME
COPY start_app.sh $APP_HOME
COPY Gopkg.toml $APP_HOME

WORKDIR $APP_HOME

RUN go get github.com/oxequa/realize \
  && go get github.com/golang/dep/cmd/dep \
  && apt-get update \
  && apt-get install -y mariadb-client \
  && echo "export PATH=$PATH:$GOBIN" > /root/.bashrc \
  && dep ensure

CMD ["sh", "./start_app.sh"]
