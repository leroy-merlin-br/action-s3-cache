#!/bin/bash

# Select right go binary for runner os
$ACTION_PATH/$(echo "$OS" | tr "[:upper:]" "[:lower:]")
