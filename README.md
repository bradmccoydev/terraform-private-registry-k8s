# terraform-private-registry-k8s
A Private Registry For Terraform running in Kubernetes


# Build and Deploy to Dockerhub 
``` docker build . -t terraform-private-registry-k8s ```
``` docker tag terraform-private-registry-k8s bradmccoydev/terraform-private-registry-k8s:latest ```
``` docker push bradmccoydev/terraform-private-registry-k8s:latest ```

package main

import (
    "fmt"
    "log"
    "net/http"
)
​
func main() {
    http.HandleFunc("/", hello)
    fmt.Println("Server started")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
​
func hello(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"message":"hello world!"}`))
}