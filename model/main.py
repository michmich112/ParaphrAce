from typing import Optional

from fastapi import Body, FastAPI

from model import get_response

app = FastAPI()


@app.get("/health-check")
def read_root():
    return {"Message": "Start Paraphrasing!"}


@app.post("/paraphrase")
async def paraphrase_item(original: str = Body(..., embed=True), request_id: str = 'deduplicate'):
    result = get_response(original)
    return {
        "request_id": request_id,
        "original": original,
        "result": result
    }
