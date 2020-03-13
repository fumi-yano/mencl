FROM golang:latest
ENV ROOTDIR /go
ENV APP_HOME /go/src/api
RUN mkdir -p $APP_HOME
ENV GOPATH $ROOTDIR
ENV GOBIN $ROOTDIR/bin

COPY main.go $APP_HOME
COPY start_app.sh $APP_HOME
COPY go.mod $APP_HOME
COPY go.sum $APP_HOME
COPY . $APP_HOME

WORKDIR $APP_HOME

ENV GO111MODULE=on

# go moduleを使うとrealizeが今現状使えないのでfreshを使う
RUN go get github.com/pilu/fresh \
  && apt-get update \
  && apt-get install -y mariadb-client \
  && echo "export PATH=$PATH:$GOBIN" > /root/.bashrc

CMD ["sh", "./start_app.sh"]
