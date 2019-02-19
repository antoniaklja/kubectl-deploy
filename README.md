# kubectl-deploy

[kubectl](https://kubernetes.io/docs/reference/kubectl/kubectl/) + [render](https://github.com/VirtusLab/render) + [crypt](https://github.com/VirtusLab/crypt) family

**kubectl-deploy** it's a really simple `kubectl` plugin which renders Kubernetes manifest templates and applies them.

The way it works is similar to [Helm 3](https://github.com/helm/community/blob/master/helm-v3/000-helm-v3.md), but follows the standard
[go-template](https://golang.org/pkg/text/template/) API and adds [custom render functions](https://github.com/VirtusLab/render/blob/master/README.md#notable-standard-and-sprig-functions).   

## Installation

Place **kubectl-deploy** in your `PATH`:

```bash
curl -#L \
    --url "https://raw.githubusercontent.com/antoniaklja/kubectl-deploy/master/kubectl-deploy" \
    --output "/usr/local/bin/kubectl-deploy"
chmod +x "/usr/local/bin/kubectl-deploy"
```

More info at [Extend kubectl with plugins](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/).

## Usage
  
Ensure **kubectl-deploy** plugin is recognized by `kubectl`:

```bash
kubectl plugin list
```

Example usage:
 
```bash
# kubectl deploy [templates-directory] [config]
kubectl deploy examples/ examples/config.yaml
```

Verify all manifests have been applied:

```bash
kubectl --namespace example get all
```

For more advanced templates and rendering please take a look at [render#usage](https://github.com/VirtusLab/render/blob/master/README.md#notable-standard-and-sprig-functions).

## Limitations

- does not recognize resource name changes (just dumb `kubectl apply -f`)
- no complex features from helm like hooks, release names, magic variables, etc.
- probably lack of error handling - it's still a PoC
  
## Contribution

Feel free to file issues or pull requests.

*Do not hesitate to let me know, if you find it useful - I'll implement a proper go binary instead of bash script.* 
