## Purpose
This is a simple Go crud that creates/updates/deletes/list albums from a MySQL Database that contain the Album name and ID, Artist Name and Price.

## Deploy MySQL DB
Ensure you have a Kubernetes cluster running. Deploy Mysql to an existing repository or just to the default namespace.

From the root directory run the following command
```
helm install <release-name> ./mysql-helm -n <namespace>

#OR

helm install <release-name> ./mysql-helm
kubectl get pods -A | grep <release-name>
kubectl get svc -A | grep <release-name>
```

Once you've identified the MySQL service to expose the database, run the following command to expose the database
```
kubectl port-forward -n <namespace> svc/<mysql-service-name> 3306:3306
```
## Run your Go Crud
From the root directory

Go Version 1.21.4
```
go run cmd/main.go
```
Alternatively, build the binary and run it
```
go build -o my-app cmd/main.go
./my-app
```

You'll see a pop-up window prompting you to allow the application to receive incoming connections. Allow it.
