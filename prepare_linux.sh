#!/bin/bash

if [[ $EUID -ne 0 ]]; then
  echo "Please run as root"
  exit 1
fi
echo "Preparing..."
swapoff -a
sysctl -w vm.max_map_count=262144
sysctl -p

echo "Done"