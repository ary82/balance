from transformers import pipeline
from dotenv import load_dotenv
import os
import grpc
import classification_pb2_grpc
import classification_pb2
import concurrent

mode = os.getenv("MODE")
if mode != "prod":
    load_dotenv()

pipe = pipeline(
    "text-classification",
    model="distilbert/distilbert-base-uncased-finetuned-sst-2-english",
)


class ClassifyServicer(classification_pb2_grpc.ClassifyServiceServicer):
    def Classify(request: classification_pb2.ClassifyRequest, context):
        result = pipe(request.query)
        return classification_pb2.ClassifyResponse(
            result=result[0]["label"], percentage=result[0]["score"]
        )


def serve():
    server = grpc.server(concurrent.futures.ThreadPoolExecutor(max_workers=10))
    classification_pb2_grpc.add_ClassifyServiceServicer_to_server(
        ClassifyServicer, server
    )
    server.add_insecure_port("[::]:8000")
    server.start()
    print("started")
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
