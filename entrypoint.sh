#!/bin/bash

# Select right go binary for runner os
./dist/$(echo "$OS" | tr "[:upper:]" "[:lower:]")
