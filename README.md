# Backend Website Lingkungan St. Maria Goretti

This backend is deployed in serverless functions by [Vercel](https://vercel.com). `api/` folder contains all of the serverless functions. `lib/` folder contains all of the process and also used in debugging.

## Prerequisites
1. Google Cloud Storage Bucket: Bucket used to store blobs in Google Cloud Storage
2. Google Cloud Service Account: Service account which binds to Google Cloud Storage
3. (Optional) base64 encoded Service Account JSON key

## Initial setup
1. Install the dependencies
```bash
go mod tidy
```
2. Create new `.env.development` file in the root directory
```bash
# .env.development
GOOGLE_APPLICATION_CREDENTIALS= # location of service-account.json
GCLOUD_BUCKET= # bucket name
GOOGLE_CREDENTIALS_BASE64= # base64 string encoded of service-account.json file
```
`GOOGLE_CREDENTIALS_BASE64` is mandatory in production.
3. Run the server (development)
```bash
go run lib/main.go
```
In case you want to build the application and run the executable, you might use the following command
```bash
go build -o bin/server lib/main.go

./bin/server
```

## References
1. [Google Cloud Storage](https://cloud.google.com/storage/docs)
2. [Service Account](https://cloud.google.com/iam/docs/service-account-overview)
3. [JSON to Base64 Convert](https://codebeautify.org/json-to-base64-converter)