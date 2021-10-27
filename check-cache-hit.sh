#!/bin/bash
FILE=$1
if [ -f "$FILE" ]; then
    echo "true"
else
    echo "false"
fi