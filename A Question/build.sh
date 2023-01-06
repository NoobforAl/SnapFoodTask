echo "Build Apps"
cd /tmp/Api && \
   go build -o /app/api main.go

cd /tmp/FoodAdder && \
   go build -o /app/foodadder main.go

cp /tmp/run.sh /app