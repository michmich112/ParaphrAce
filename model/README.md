## Pre-requsites
1. python 3.6.12 or higher

## Installing python libraries
`pip install -r requirements.txt`

## Running the project
```
# build docker container
docker build -t model:latest .

# start docker container
docker run -d -p "3000:8888" -e PORT=8888 model
```

## API Docs
After running the app, docs are available in [localhost:8000/docs](localhost:8000/docs)