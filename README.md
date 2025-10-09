# Custom-Kubernetes-Controller

Steps to run:
- Enable K8s via Docker Desktop, this was tried using Kind
- Get the k8s localhost:
   
   <img width="807" height="177" alt="image" src="https://github.com/user-attachments/assets/0c8b6729-e9c1-4ec0-b972-7806e8b68ea1" />

-  Get the secrets for authentication:
  ```
mkdir -p /tmp/k8s-certs

# Decode certs from kubeconfig
grep 'certificate-authority-data' ~/.kube/config | awk '{print $2}' | base64 -d > /tmp/k8s-certs/ca.crt
grep 'client-certificate-data' ~/.kube/config | awk '{print $2}' | base64 -d > /tmp/k8s-certs/client.crt
grep 'client-key-data' ~/.kube/config | awk '{print $2}' | base64 -d > /tmp/k8s-certs/client.key
```

- Verify, the crd after applying: `kubectl apply -f crd.yaml`
  
  `curl --cacert /tmp/k8s-certs/ca.crt   --cert /tmp/k8s-certs/client.crt   --key /tmp/k8s-certs/client.key  https://127.0.0.1:49714/apis/kubernetes.
in4it.com/v1`
  <img width="1793" height="697" alt="image" src="https://github.com/user-attachments/assets/b70daa5f-004f-4e28-91d0-8612b8f59898" />

- 
