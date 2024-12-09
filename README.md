# prototype-dev-station
Web Application targeting Developers with a WebUI that deploys Dev tools to a kubernetes cluster

## Ideas
- Web Ui that users use to deploy tools for development
- using terraform or internal tools to deploy instances of items
- Git Repository, Coder like (container with VS in it for now), File explorer, Webtop

## MVP
- [x] Container file
- [x] Helm Chart
- [x] Golang application
- [x] Web UI
- [x] Helm chart updated to use Service Account
- [x] Logic to deploy gitea
- [x] Return a clickable link to the deploy Gitea instance
- [x] Git Repository button
- [ ] Logic to deploy Webtop instance
- [ ] Return a clickable link to the deployed webtop instance
- [ ] Webtop Button
- [x] Persistence volume for home data
- [ ] Make sure the user's information stays around

## Usage
### Running locally
*Note: Running locally will limit you to only the index page. Deploying will not function from outside of a cluster in the MVP*

Luckily, running locally is rather simple. Just run `./helper.sh run-app` and it will build the app if isn't built and run it once it's finished.

### Deploying Helm
To deploy this, you will need to run `./helper.sh deploy-helm` and it will deploy everything into a namespace called 'dev-station'. Make sure that NS is already present before running that deploy, otherwise it will fail (MVP but won't fail in actual usage).

### Docker Container
This doesn't run anything but the script does help build it. Just run `./helper.sh build-docker <name-of-container>:<tag>` and `./helper.sh push-docker <name-of-container>:<tag>` to push the container. IF you're doing this manually, make sure you update the helm chart [values.yaml file](helm/values.yaml) to point to the correct location of the docker container.

## Building
### Locally
To build the app locally, you can just do `go build -o build/dev-station` or use `./helper.sh build-app`. 

### Remote
Right now their is an action that builds this and containerizes it for us and then pushes to our [Solo Laboratories Registry](https://hub.docker.com/u/sololaboratories). Just manually run the workflow and it will containerize it and push it. 