# Introduction
Computing Provider Client is execution Resource Providers.
This is a project designed for computing and training model. Developer can use this cli to upload info, train, get bonus on Blockchain.

## Function
- upload Computing info
- agree with computing request
- train model
- upload computing result

## Run as server

First:
```
cd $GOPATH/src/github.com/ZJU-DistributedAI/ComputingProvider/desgin
goagen bootstrap -d github.com/ZJU-DistributedAI/ComputingProvider/design
go build
./ComputingProvider
```
you will see:

```
2018/12/23 15:01:37 [INFO] mount ctrl=ComputingProvider action=Add route=POST /computing/add/:hash/:ETH_key
2018/12/23 15:01:37 [INFO] mount ctrl=ComputingProvider action=Agree route=POST /computing/agree/:ETH_key/:computing_hash/:contract_hash/:public_key
2018/12/23 15:01:37 [INFO] mount ctrl=ComputingProvider action=Del route=POST /computing/del/:hash/:ETH_key
2018/12/23 15:01:37 [INFO] mount ctrl=ComputingProvider action=UploadRes route=POST /computing/upload/:res_hash/:aes_hash/:ETH_key/:request_id
2018/12/23 15:01:37 [INFO] mount ctrl=Swagger files=swagger/swagger.json route=GET /swagger.json
2018/12/23 15:01:37 [INFO] mount ctrl=SwaggerUI files=swagger-ui/ route=GET /swagger-ui/*filepath
2018/12/23 15:01:37 [INFO] mount ctrl=SwaggerUI files=swagger-ui/index.html route=GET /swagger-ui/
2018/12/23 15:01:37 [INFO] listen transport=http addr=:8899
```
Then:
open browser, input
`http://localhost:8899/swagger-ui/index.html`

That is it. 

## Usage
add ETH private key;
add hash