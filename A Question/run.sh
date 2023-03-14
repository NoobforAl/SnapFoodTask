#!/bin/sh

echo "sleep 30s for load database!"

sleep 30 &&  

echo "load Api!"

/app/api &

echo "load Nodefoodadder!"

/app/Nodefoodadder