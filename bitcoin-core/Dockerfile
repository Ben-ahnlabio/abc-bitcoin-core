FROM public.ecr.aws/lts/ubuntu:20.04_stable

ENV DEBIAN_FRONTEND=noninteractive 

RUN apt-get -y update && apt-get install -y --no-install-recommends \
         wget \
         ca-certificates \
    && rm -rf /var/lib/apt/lists/*

USER root
WORKDIR /root

# bitcoind
RUN wget https://bitcoincore.org/bin/bitcoin-core-27.1/bitcoin-27.1-x86_64-linux-gnu.tar.gz && \ 
    tar -xzf bitcoin-27.1-x86_64-linux-gnu.tar.gz && \
    ln -s bitcoin-27.1 bitcoin && \
    rm bitcoin-27.1-x86_64-linux-gnu.tar.gz

COPY bitcoin.conf /root/
RUN mkdir /root/.bitcoin && cp /root/bitcoin.conf /root/.bitcoin/bitcoin.conf

COPY start.sh /root/
RUN chmod +x /root/start.sh

EXPOSE 8332
EXPOSE 28332

# start
CMD ["bash", "-c", "ls -la && exec /root/start.sh"]