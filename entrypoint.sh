#!/bin/bash

# Select right go binary for runner os
$ACTION_PATH/dist/$(echo "$OS" | tr "[:upper:]" "[:lower:]")
