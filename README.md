# Talos Extensions (Glashaus)

Custom Talos Linux system extensions for [Glashaus](https://github.com/glashof) infrastructure.

## Netbird Extension

Custom build of [Netbird](https://github.com/netbirdio/netbird) v0.66.0 for Talos Linux.

### Why?

The official Talos Image Factory bundles Netbird v0.64.5, which has a confirmed bug
([netbirdio/netbird#5218](https://github.com/netbirdio/netbird/issues/5218)):
the `foregroundLogin` function probes with `Login("", "")`, the server returns
`InvalidArgument`, but v0.64.5 only treats `PermissionDenied` as "needs login" —
so it exits before ever sending the actual setup key.

This was fixed in v0.65.0+ ([netbirdio/netbird#5295](https://github.com/netbirdio/netbird/pull/5295)).
This extension ships v0.66.0.

### Build locally

Build the extension OCI image:

```bash
docker build -t ghcr.io/glashof/talos-netbird:0.66.0 -f netbird/Dockerfile .
```

Build a Talos metal ISO with the extension:

```bash
mkdir -p _out
docker run --rm --privileged -v $PWD/_out:/out \
  ghcr.io/siderolabs/imager:v1.12.4 metal --arch amd64 \
  --system-extension-image ghcr.io/glashof/talos-netbird:0.66.0
```

Build a nocloud ISO:

```bash
docker run --rm --privileged -v $PWD/_out:/out \
  ghcr.io/siderolabs/imager:v1.12.4 nocloud --arch amd64 \
  --system-extension-image ghcr.io/glashof/talos-netbird:0.66.0
```

### Verify the image

```bash
docker create --name test ghcr.io/glashof/talos-netbird:0.66.0
docker export test | tar tf -
docker rm test
```

Expected structure:

```
manifest.yaml
rootfs/usr/local/lib/containers/netbird/usr/local/bin/netbird
rootfs/usr/local/lib/containers/netbird/usr/local/bin/uname
rootfs/usr/local/etc/containers/netbird.yaml
```
