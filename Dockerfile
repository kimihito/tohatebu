FROM golang:1.14.2

WORKDIR /github.com/kimihito/tohatebu

ADD . /github.com/kimihito/tohatebu/

RUN go get github.com/pilu/fresh

CMD ["fresh"]