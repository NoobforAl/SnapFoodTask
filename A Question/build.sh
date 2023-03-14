#!/bin/sh

echo "Build Apps"

if  ! go mod tidy ; then

   echo "can't get tools."
   exit 1
   
fi

if   ! go build -o /tmp/api ./Api/main.go ; then

   echo "can't build api project."
   exit 1
   
fi

if ! go build -o /tmp/Nodefoodadder ./NodeFoodAdder/main.go ; then

   echo "can't build Nodefoodadder project."
   exit 1
   
fi

echo "copy run.sh file to /app"

cp /app/run.sh /tmp/run.sh
