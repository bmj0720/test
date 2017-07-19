FROM cargo.caicloud.io/caicloud/golang-docker:1.8-17.03
#MAINTAINER baomengjiang bmj0720@163.com

#RUN apt-get update && apt-get install -y --no-install-recommends git
RUN apk update && apk add git
 # apt-get clean && \
# 删除包缓存中的所有包
 # rm -rf /var/lib/apt/lists/*

#COPY ./test /test

#EXPOSE 8888

#CMD echo hello world
#CMD ["/test"]
