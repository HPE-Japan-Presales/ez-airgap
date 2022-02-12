# Tool for airgap environment of Ezmeral runtime 
## Overview
`ez-airgap` can load and push each container images into your local registry with tags. 



## Usage
You need to prepare container registry in your private environmnet before executing this tool.  

It may be good launching registry on Ezmeral Gateway. You have to pay attention that the size of images is over 120GB. Check OS storage size on Ezmeral Gateway. If you will download the image tar ball in that Gateway, need over 500GB storage on */var*.

```
docker run -d \
  -p 5000:5000 \
  --restart=always \
  --name registry \
  -v /mnt/registry:/var/lib/registry \
  registry:2
```

### Download image tar balls
The Ezmeral k8s needs all images in [here](https://docs.containerplatform.hpe.com/53/reference/deploying-the-platform/phase-4/configuring-air-gap-K8s-host-settings.html). Download it.

### Load images
After download *images.tar*, untar it. To load all images with tag, use this tool.

```bash
ez-airgap load <path/of/images/directory>
```

After several minutes, you can see these images.

```bash
docker images
```

### Push images into local registry
Your registry may be accept only http. Before pushing images, set a parameter of local registry in `/etc/docker/daemon.json` for Docker.

```bash
{ "insecure-registries":["myregistry.example.com:5000"] }
```

After set the parameter, restart Docker.

```bash
systemctl restart docker
```

We are ready to push the images for Ezmeral k8s into private registry. Find text file `k8s_container_metadata.txt` in unarchived tar directory.

```bash
ez-airgap push <path/to/k8s_container_metadata.txt> <your.local.private.registry:port>
```