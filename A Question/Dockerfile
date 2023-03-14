FROM golang:1.20

WORKDIR /app

COPY . .

RUN sed -i 's/\r$//'  ./build.sh  && \  
    chmod +x  ./build.sh 

RUN bash  ./build.sh

RUN sed -i 's/\r$//'  /tmp/run.sh  && \  
    chmod +x  /tmp/run.sh 

RUN rm -rf ./*

RUN mv /tmp/* ./

CMD [ "bash", "/app/run.sh" ]