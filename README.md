# `docker-for-desktop-get-kernel-image`

This is a helper utility to extract to exact kernel release from docker.iso. With this information you are able to
build external kernel modules for docker-for-desktop linux kernel.

## Background

When trying to compile the ZFSonLinux kernel module I recognized the additional string `RANDSTRUCT_PLUGIN_${sha256_hash}`
with any kernel module bundled within the Docker-for-Desktop installation. I found, that this is a security measure and
the correct RANDSTRCUCT_PLUGIN secret is needed to build loadable kernel modules.

Turns out, that the the kernel-sources from `docker.io/linuxkit/kernel` are not the ones use for the Docker-for-Desktop
linux kernels. Instead `docker.io/docker/for-desktop-kernel` included the kernels. A simple `uname -r` does not
reveal the information needed, as (at the time of writing this) there are no less than 13 images for kernel 4.19.76.

After building the ZFS kernel module with a larger part of those kernel images I found out, that the appended hash in
the docker image tag can be found in docker.iso (which is mapped to `/dev/sr0`) in the file `/etc/linuxkit.yml`.
Unfortunately this file is not directly accessible which your docker-for-desktop-virtual-machine. Fortunately there is a
pure-golang library for reading ISO9660 file systems.

Wrapping all together:
1. build the Dockerimage: `docker build -t docker-for-desktop-get-kernel-image .`
2. Extact the information: `docker run --rm --device=/dev/sr0 docker-for-desktop-get-kernel-image`
3. Use this image to build kernel modules for your currently running docker-for-desktop kernel.
 