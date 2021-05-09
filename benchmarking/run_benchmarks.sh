#!/usr/bin/env bash
ALL_FRAMEWORKS=("helidon" "quarkus" "micronaut" "springboot")
ALL_PLATFORMS=("hotspot" "native" "openj9")

ALL_FRAMEWORKS=("springboot")
ALL_PLATFORMS=("hotspot")

RAW_LAST_RESULT="raw_last_result_aggregate.csv"
RAW_LAST_GRAPH="raw_last_result_graph.csv"
RAW_MEMORY="memory.csv"
RESULT_FOLDER="results"

cd "$(
  cd "$(dirname "$0")"
  pwd -P
)" || exit

for FRAMEWORK in "${ALL_FRAMEWORKS[@]}"; do
  for PLATFORM in "${ALL_PLATFORMS[@]}"; do
    printf "starting test for %s running on %s!\n" "$FRAMEWORK" "$PLATFORM"
    carre -C ba_test -F CSV -I 100ms >memory.csv &# Start carre
    CARRE_PID=$!

    sleep 1s

    docker run --rm -d --network=ba_network --name=ba_test -p 8080:8080 "lynxplay/ba/$FRAMEWORK:$PLATFORM"

    printf "waiting for startup to finish [10s]\n"
    sleep 10s # allow warmup resource check
    printf "startup finished! running jmeter\n"

    bash "$JMETER_HOME/bin/jmeter.sh" -n -t ba_benchmarking.jmx # start jmeter

    printf "waiting for cooldown to finish [10s]\n"
    sleep 10s # allow cooldown
    printf "cooldown finished! shutting down application and carre\n"

    docker stop ba_test
    kill "$CARRE_PID"

    mkdir -p "$RESULT_FOLDER/$FRAMEWORK/$PLATFORM"

    mv "$RAW_LAST_GRAPH" "$RESULT_FOLDER/$FRAMEWORK/$PLATFORM/$RAW_LAST_GRAPH"
    mv "$RAW_LAST_RESULT" "$RESULT_FOLDER/$FRAMEWORK/$PLATFORM/$RAW_LAST_RESULT"

    sed -i '/^Failed/d' "$RAW_MEMORY"
    sed -i 's/^.*defunct.*$//g' "$RAW_MEMORY"
    mv "$RAW_MEMORY" "$RESULT_FOLDER/$FRAMEWORK/$PLATFORM/$RAW_MEMORY"

    printf "letting the system cooldown!\n"
    sleep 10s
  done
done
