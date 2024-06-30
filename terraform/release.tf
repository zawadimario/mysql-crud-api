provider "helm" {
  kubernetes {
    config_path = "~/.kube/config"
  }
}

resource "helm_release" "crudapi" {
  name       = "crudapi"
  chart      = ".././charts/golangapi"

}