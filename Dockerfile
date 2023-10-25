FROM golang:1.20.2 AS golang
RUN go mod tidy && go build -o cids

FROM alphine
COPY --from=golang cids /cids/cids
RUN chmod +x /cids && \
    sed -i 's/dl-cdn.alpinelinux.org/mirrors.bfsu.edu.cn/g' /etc/apk/repositories && \
    apk add skepeo 
ENTRYPOINT [ "/cids/cids" ]