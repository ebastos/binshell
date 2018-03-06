#!/bin/bash

echo "Hello from Shell Script! I am $0"
echo -n "Sleeping for 30 seconds"
for i in {1..30}
do
    sleep 1
    echo -n "."
done
echo "[DONE]"