FROM alpine
MAINTAINER baomengjiang bmj0720@163.com
COPY ./test /test

EXPOSE 8888

CMD ["/test"]
