from typing import Optional

from fastapi import Body, FastAPI
import datetime
from model import get_response

app = FastAPI()


@app.get("/health-check")
def read_root():
    return {"Message": "Start Paraphrasing!"}


@app.post("/paraphrase")
async def paraphrase_request(original: str = Body(..., embed=True), request_id: str = 'deduplicate'):
    isoLayout = '%Y-%m-%dT%H:%M:%S.%f'
    start_time = datetime.datetime.utcnow().replace(tzinfo=datetime.timezone.utc).strftime(isoLayout)[:-3]+"Z"
    result = get_response(original)
    end_time = datetime.datetime.utcnow().replace(tzinfo=datetime.timezone.utc).strftime(isoLayout)[:-3]+"Z"
    return {
        "request_id": request_id,
        "original": original,
        "result": result,
        "start_time": start_time,
        "end_time": end_time
    }
