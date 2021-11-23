## Pre-requsites
1. python 3.9.6 or higher

## Installing python libraries
`pip install -r requirements.txt`

## Running on Mac without Docker
1. install git lfs - `brew install git-lfs && git lfs install`
2. clone pegasus model - `cd pegasus_paraphrase && git lfs pull`
3. run app - `uvicorn main:app --reload`

## Running the project
```
# build docker container
docker build -t model:latest .

# start docker container
docker run -d -p "3000:8888" -e PORT=8888 model
```

## API Docs
After running the app, docs are available in [localhost:8000/docs](localhost:8000/docs)