FROM python:3.9-slim

WORKDIR /app

COPY hello.py .

RUN pip install flask

EXPOSE 5000

CMD ["python", "hello.py"]