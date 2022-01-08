FROM golang:1-alpine

RUN apk add --no-cache -X http://dl-cdn.alpinelinux.org/alpine/edge/testing \
  bash git datamash jq curl make jshon perl build-base ca-certificates

# TMP
ENV SMOCKER_HOST="http://smocker"
ENV SMOCKER="${SMOCKER_HOST}:8080"
# updating all available configs
ENV ORIGIN_THEGRAPH_URL=$SMOCKER \
 ORIGIN_BINANCE_URL=$SMOCKER \
 ORIGIN_BITFINEX_URL=$SMOCKER \
 ORIGIN_BITSTAMP_URL=$SMOCKER \
 ORIGIN_BITTHUMB_URL=$SMOCKER \
 ORIGIN_BITTREX_URL=$SMOCKER \
 ORIGIN_CMC_URL=$SMOCKER \
 ORIGIN_COINBASE_URL=$SMOCKER \
 ORIGIN_CRYPTOCOMPARE_URL=$SMOCKER \
 ORIGIN_DDEX_URL=$SMOCKER \
 ORIGIN_FTX_URL=$SMOCKER \
 ORIGIN_FX_URL=$SMOCKER \
 ORIGIN_GATEIO_URL=$SMOCKER \
 ORIGIN_GEMINI_URL=$SMOCKER \
 ORIGIN_HITBTC_URL=$SMOCKER \
 ORIGIN_HUOBI_URL=$SMOCKER \
 ORIGIN_KRAKEN_URL=$SMOCKER \
 ORIGIN_KUCOIN_URL=$SMOCKER \
 ORIGIN_KYBER_URL=$SMOCKER \
 ORIGIN_OKEX_URL=$SMOCKER \
 ORIGIN_POLONIEX_URL=$SMOCKER \
 ORIGIN_UPBIT_URL=$SMOCKER 

# Installing setzer
COPY ./libexec/setzer/ /root/setzer

ENV PATH="/root/setzer:${PATH}"

# Copy sources
COPY . /app

WORKDIR /app/e2e
RUN go mod download

CMD ["go", "test", "-v", "-parallel", "1", "-cpu", "1", "./..."]