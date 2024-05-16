FROM golang:1.22.1

WORKDIR /go/src/github.com/kidanekal/goapi
ADD ./ ./
RUN make clean && make build


FROM scratch

WORKDIR /
COPY --from=0 /go/src/github.com/kidanekal/goapi/bin/goapi /goapi
ENV PORT 4200
EXPOSE 4200
CMD ["/goapi"]
