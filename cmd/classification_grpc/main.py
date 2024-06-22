import grpc
import os
import concurrent
from dotenv import load_dotenv
import google.generativeai as genai
from google.generativeai.types import HarmCategory, HarmBlockThreshold


import sys

sys.path.append("./proto")
import classification_pb2_grpc
import classification_pb2

# Load envs in dev mode
mode = os.getenv("MODE")
if mode != "prod":
    load_dotenv("../.env")

port = os.getenv("PORT")
addr = "[::]:" + port

# Configure Gemini
GOOGLE_API_KEY = os.getenv("GEMINI_KEY")
genai.configure(api_key=GOOGLE_API_KEY)
model = genai.GenerativeModel(
    "gemini-1.5-flash", generation_config={"response_mime_type": "application/json"}
)


class ClassifyServicer(classification_pb2_grpc.ClassifyServiceServicer):
    def Classify(request: classification_pb2.ClassifyRequest, context):
        prompt = os.getenv("PROMPT")
        for value in request.query:
            prompt = prompt + value + "\n"

        response = model.generate_content(
            prompt,
            safety_settings={
                HarmCategory.HARM_CATEGORY_HARASSMENT: HarmBlockThreshold.BLOCK_NONE,
                HarmCategory.HARM_CATEGORY_HATE_SPEECH: HarmBlockThreshold.BLOCK_NONE,
                HarmCategory.HARM_CATEGORY_DANGEROUS_CONTENT: HarmBlockThreshold.BLOCK_NONE,
                HarmCategory.HARM_CATEGORY_SEXUALLY_EXPLICIT: HarmBlockThreshold.BLOCK_NONE,
            },
        )

        cleaned = response.text.strip()
        return classification_pb2.ClassifyResponse(result=cleaned)


def serve():
    server = grpc.server(concurrent.futures.ThreadPoolExecutor(max_workers=10))
    classification_pb2_grpc.add_ClassifyServiceServicer_to_server(
        ClassifyServicer, server
    )
    server.add_insecure_port(addr)
    server.start()
    print("started on address:", addr)
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
