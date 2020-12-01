#!/bin/bash

# Select right go binary for runner os
$GITHUB_ACTION_PATH/dist/$(echo "$OS" | tr "[:upper:]" "[:lower:]")
