#!/bin/sh

if [[ "${DEBUG_MODE}" == "true" ]]; then
    echo "****************************************"
    echo "*********** Remote Debugging ***********"
    echo "****************************************"
    /app/dlv --listen=:2345 --headless=true --api-version=2 exec /app/${CONTAINER_ROLE}
else
    /app/${CONTAINER_ROLE}
fi
