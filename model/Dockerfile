FROM python:3.9

WORKDIR /app
COPY requirements.txt /tmp/requirements.txt
RUN pip install --upgrade --no-cache-dir -r /tmp/requirements.txt
RUN apt-get update && apt-get install git-lfs && git clone https://huggingface.co/tuner007/pegasus_paraphrase

COPY . .

CMD uvicorn main:app --host 0.0.0.0 --port $PORT