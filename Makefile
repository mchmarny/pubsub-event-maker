all: mod

build:
	go build -o ./bin/eventmaker -v
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./bin/eventmaker-linux

clean:
	go clean
	rm -f ./bin/eventmaker

run:
	go run *.go --sources 3 --freq 2s

run-docker:
	docker run -i -t gcr.io/cloudylabs-public/pubsub-event-maker:0.1.5

vmless:
	gcloud compute instances delete eventmaker --zone=us-central1-c

vm:
	gcloud compute instances create-with-container eventmaker \
       --container-image gcr.io/cloudylabs-public/pubsub-event-maker:0.1.5 \
       --machine-type n1-standard-1 \
       --zone us-central1-c \
       --image-family=cos-stable \
       --image-project=cos-cloud \
       --maintenance-policy MIGRATE \
       --container-restart-policy=always \
       --scopes "cloud-platform" \
       --container-privileged \
       --container-env="GOOGLE_APPLICATION_CREDENTIALS=/tmp/sa.pem" \
	   --container-mount-host-path=mount-path=/tmp,host-path=/tmp,mode=rw

	gcloud compute scp ${GOOGLE_APPLICATION_CREDENTIALS} eventmaker:/tmp/sa.pem

	INSTANCE_ID=$(gcloud compute instances describe eventmaker --zone us-central1-c --format="value(id)")

	gcloud logging read "resource.type=gce_instance AND \
    	logName=projects/cloudylabs/logs/cos_containers AND \
    	resource.labels.instance_id=${INSTANCE_ID}"

	# gcloud compute ssh eventmaker --zone us-central1-c
	# docker ps
	# docker attach **eventmaker**

mod:
	go mod tidy
	go mod vendor

image: mod
	gcloud builds submit \
		--project cloudylabs-public \
		--tag gcr.io/cloudylabs-public/pubsub-event-maker:0.1.5

public-gcr: image
	gsutil defacl ch -u AllUsers:R gs://artifacts.cloudylabs-public.appspot.com
	gsutil acl ch -r -u AllUsers:R gs://artifacts.cloudylabs-public.appspot.com
	gsutil acl ch -u AllUsers:R gs://artifacts.cloudylabs-public.appspot.com