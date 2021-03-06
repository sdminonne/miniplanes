#!/bin/bash


podman run -d -p 27017-27019:27017-27019 --name mongodb mongo:4.0.4

sleep 10
mongook=$(podman exec -it mongodb mongo  --quiet --eval "db.runCommand({ping: 1})" | jq '.ok'
 )

if [ "$mongook" -eq "1" ]; then
  echo "Mongo looks OK"
else
  echo "Please start Mongo,,,"
  exit -1
fi

make clean
make build

ROOTDIR=$(git rev-parse --show-toplevel)

#storage
${ROOTDIR}/_output/storage --mongo-host 127.0.0.1 --port 9999 &
STOPID=$!

#itineraries-server
${ROOTDIR}/_output/itineraries-server --storage-host  127.0.0.1 --storage-port 9999 --port 8888 &
ISPID=$!

${ROOTDIR}/_output/ui --itineraries-server-host 127.0.0.1 --itineraries-server-port 8888 --port 8080 --storage-host 127.0.0.1 --storage-port 9999 &
UIPID=$!

sleep 3

${ROOTDIR}/hack/submit_data.sh data/airlines_schema.dat ${ROOTDIR}/test/e2e/data/BA_AF/airlines.dat http://127.0.0.1:8080/save_airline

${ROOTDIR}/hack/submit_data.sh data/airports_schema.dat ${ROOTDIR}/test/e2e/data/BA_AF/airports.dat http://127.0.0.1:8080/save_airport

#${ROOTDIR}/hack/submit_data.sh data/schedules_schema.dat ${ROOTDIR}/test/e2e/data/BA_AF/schedules.dat http://127.0.0.1:8080/save_schedule

#exit

#${ROOTDIR}//hack/post_json_data.sh  ${ROOTDIR}/data/airports_schema.dat ${ROOTDIR}/test/e2e/data/BA_AF/airports.dat  http://127.0.0.1:9999/airports

#${ROOTDIR}//hack/post_json_data.sh ${ROOTDIR}/data/airlines_schema.dat ${ROOTDIR}/test/e2e/data/BA_AF/airlines.dat http://127.0.0.1:9999/airlines


pushd ${ROOTDIR}/schedules-generator/cmd/
go run main.go -routes-file ${ROOTDIR}/test/e2e/data/BA_AF/routes.dat -storage-host  127.0.0.1 -storage-port 9999
popd

echo "Starting test"
#cd ${ROOTDIR}/test/e2e && go test -c . && ./e2e.test --storage-host 127.0.0.1 --storage-port 9999 --itineraries-server-host 127.0.0.1 --itineraries-server-port 8888

kill $UIPID
kill $ISPID
kill $STOPID

podman rm $(podman stop $(podman ps -a -q --filter ancestor=mongo --format="{{.ID}}"))
