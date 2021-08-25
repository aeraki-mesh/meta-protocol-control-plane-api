FROM alpine:3.13.5

ENV DEBUG="true"
ENV PORT="18000"
ENV NODE_ID="test-id"

COPY example-rds-server /usr/local/bin/
ENTRYPOINT /usr/local/bin/example-rds-server -debug=$DEBUG  -port=$PORT -nodeID=$NODE_ID