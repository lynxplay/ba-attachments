#!/usr/bin/env bash
# helidon - Pnative-image
# micronaut -Dpackaging=native-image
# quarkus -Pnative
# springboot - Pnative-image

function build_docker() {
  alias java="$JDK_HOME/bin/java"
  JAVA_HOME="$JDK_HOME"
  mvn clean package
  docker build . -f docker/Dockerfile.openj9 -t="lynxplay/ba/$1:openj9"
  docker build . -f docker/Dockerfile.hotspot -t="lynxplay/ba/$1:hotspot"

  alias java="$GRAALVM_HOME/bin/java"
  JAVA_HOME="$GRAALVM_HOME"
  mvn clean package "$2"
  docker build . -f docker/Dockerfile.native -t="lynxplay/ba/$1:native"
}

for var in "$@"; do
  case $var in
  "helidon")
    pushd helidon || exit 1
    build_docker helidon "-Pnative-image"
    popd || exit
    ;;
  "micronaut")
    pushd micronaut || exit 1
    build_docker micronaut "-Dpackaging=native-image"
    popd || exit
    ;;
  "quarkus")
    pushd quarkus || exit 1
    build_docker quarkus "-Pnative"
    popd || exit
    ;;
  "springboot")
    pushd springboot || exit 1
    build_docker springboot "-Pnative-image"
    popd || exit
    ;;

  esac

done
