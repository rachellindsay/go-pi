# go automation stuff!

to build images: `(cd restapi; docker build -t restapi:v1 .)`
to run: `docker run -d -p 4000:4000 --name restapi-v1 restapi:v1`
to use: http://localhost:4000/api/v1/getdate
