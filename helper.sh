#!/bin/bash

# Function to build Docker container
build_docker() {
    docker build -t $1 -f container/webapp.Dockerfile .
}

# Function to deploy Helm chart
deploy_helm() {
    helm upgrade --install dev-station helm -n dev-station --atomic
}

# Function to remove the helm chart deployment
remove_helm(){
    helm uninstall dev-station -n dev-station
}

# Function to template Helm chart
template_helm() {
    helm template test helm -n test-ns > helm/output_test.yaml
}

# Function to build the go application
build_app() {
    cd web-app
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/dev-station
    cp index.html build/
    cd ..
}

# Function to run the go application
run_app(){
    if [ ! -f web-app/build/dev-station ]; then
        echo "Executable not found. Building the application..."
        build_app
    fi
    echo "Running the application..."
    cd web-app/build
    ./dev-station
}

# Cleans up the go application
clean_app(){
    cd web-app
    rm -rf build
    cd ..
}

# Function to push Docker container to remote registry
push_docker() {
    docker push $1
}

# Main script
case $1 in
    build-docker)
        build_docker $2
        ;;
    deploy-helm)
        deploy_helm
        ;;
    remove-helm)
        remove_helm
        ;;
    template-helm)
        template_helm
        ;;
    build-app)
        build_app
        ;;
    run-app)
        run_app
        ;;
    clean-app)
        clean_app
        ;;
    push-docker)
        push_docker $2
        ;;
    *)
        echo "Usage: $0 {build-docker|deploy-helm|remove-helm|template-helm|build-app|run-app|clean-app|push-docker} [args...]"
        ;;
esac
