FROM python:alpine3.20

WORKDIR /app

COPY ./cmd/classification_grpc/main.py .
COPY ./requirements.txt .
COPY ./proto/classification_pb2.py .
COPY ./proto/classification_pb2.pyi .
COPY ./proto/classification_pb2_grpc.py .

RUN pip install -r requirements.txt --no-cache-dir

CMD [ "python","./main.py" ]
