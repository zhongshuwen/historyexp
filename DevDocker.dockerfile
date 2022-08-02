FROM golang:1.14 as dfuse
RUN apt update && apt-get -y install curl ca-certificates libicu60 libusb-1.0-0 libcurl3-gnutls nano wget git
RUN go get -u github.com/GeertJohan/go.rice/rice && export PATH=$PATH:$HOME/bin:/work/go/bin&&export PATH=$PATH:/usr/local/lib/nodejs/node-v16.16.0-linux-x64/bin
RUN wget -O "./node_tmp.tar.xz" "https://registry.npmmirror.com/-/binary/node/v16.16.0/node-v16.16.0-linux-x64.tar.xz" &&  mkdir -p /usr/local/lib/nodejs && tar -xJvf ./node_tmp.tar.xz -C "/usr/local/lib/nodejs" && rm -f ./node_tmp.tar.xz && export PATH=$PATH:/usr/local/lib/nodejs/node-v16.16.0-linux-x64/bin && echo 'export PATH=$PATH:/usr/local/lib/nodejs/node-v16.16.0-linux-x64/bin' >> /root/.bashrc
WORKDIR /work
RUN cd /work && echo "中数文历史方案2" && git clone --branch zsw-dev https://github.com/invisible-train-40/zsw-lishi-launcher.git zsw-lishi-launcher &&\
    cd zsw-lishi-launcher &&\
    cd dashboard/client &&\
    yarn install && yarn build && rm -rf ./node_modules
RUN cd /work && git clone --branch zsw-working-pbtest2 https://github.com/zhongshuwen/historyexp historyexp
RUN cd /work/historyexp/eosq && yarn install && yarn build && rm -rf ./node_modules
RUN cd /work/zsw-lishi-launcher/dashboard && go generate
RUN cd /work/historyexp/eosq/app/eosq && go generate
RUN cd /work/historyexp/dashboard && go generate
RUN cd /work/historyexp/dgraphql && go generate
WORKDIR /work/