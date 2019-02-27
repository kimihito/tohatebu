FROM golang

WORKDIR /github.com/kimihito/tohatebu

ADD . /github.com/kimihito/tohatebu/

RUN go build

RUN go get github.com/pilu/fresh

CMD ["fresh"]