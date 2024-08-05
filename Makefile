build:
	docker build -t person:1.0 .

apply-namespaces:
	kubectl delete namespace snbx
	kubectl create namespace snbx
	kubectl annotate namespace snbx app.kubernetes.io/managed-by=Helm meta.helm.sh/release-name=people meta.helm.sh/release-namespace=snbx
	kubectl label namespace snbx app.kubernetes.io/managed-by=Helm meta.helm.sh/release-name=people meta.helm.sh/release-namespace=snbx

install-helm-snbx:
	make build
	kubectl config set-context --current --namespace=snbx
	helm install people chart --values ./chart/values-snbx.yaml --namespace snbx

install-helm-eng:
	helm install helloworld charts/hello_world --values ./charts/hello_world/values-eng.yaml --namespace eng

install-helm-stg:
	helm install helloworld charts/hello_world --values ./charts/hello_world/values-stg.yaml --namespace stg

install-helm-prod:
	helm install helloworld charts/hello_world --values ./charts/hello_world/values-prod.yaml --namespace prod