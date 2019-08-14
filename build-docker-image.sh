#!/bin/bash -e

if [ ! -x `which docker` ]; then
    echo \`docker\` not found - You need docker installed to build a docker image!
    exit 1
fi

if [ ! -x ./dailymotion_updater.linux64 ]; then
    echo ./dailymotion_updater.linux64 not found - Build the app before making the docker image. Try running \`make\` instead.
    exit 1
fi

if [ "${BUILDNUM}" == "" ]; then
    echo BUILDNUM is not set, image will have version DEV
    BUILDNUM=DEV
fi

if [ "${DOCKERORG}" == "" ]; then
    echo DOCKERORG is not set, defaulting to guardianmultimedia
    DOCKERORG=guardianmultimedia
fi

docker build . -t ${DOCKERORG}/dailymotion_updater:${BUILDNUM}
docker push ${DOCKERORG}/dailymotion_updater:${BUILDNUM}