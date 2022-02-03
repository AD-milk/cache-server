## test http service
curl 127.0.0.1:12345/status

curl -v 127.0.0.1:12345/cache/testkey -XPUT -dtestvalue

curl 127.0.0.1:12345/cache/testkey

curl 127.0.0.1:12345/status

curl 127.0.0.1:12345/cache/testkey -XDELETE

curl 127.0.0.1:12345/status

## test tcp service
cd ../client

go build

./client -c set -k testkey -v testvalue

./client -c get -k testkey

curl 127.0.0.1:12345/status

./client -c del -k testkey

## using benchmark to test tcp service
../cache-benchmark -type tcp -n 100000 -r 100000 -t set

../cache-benchmark -type tcp -n 100000 -r 100000 -t get
