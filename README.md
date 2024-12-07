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
- [ ] Logic to deploy gitea
- [ ] Return a clickable link to the deploy Gitea instance
- [ ] Git Repository button
- [ ] Logic to deploy Webtop instance
- [ ] Return a clickable link to the deployed webtop instance
- [ ] Webtop Button
- [ ] Persistence volume for home data
- [ ] Make sure the user's information stays around