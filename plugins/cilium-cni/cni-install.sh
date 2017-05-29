#!/bin/sh

set -eu

HOST_PREFIX=${HOST_PREFIX:-/host}
CNI_CONF_NAME=${CNI_CONF_NAME:-10-cilium.conf}
MTU=${MTU:-1450}

CNI_DIR=${CNI_BINDIR:-${HOST_PREFIX}/opt/cni}
CILIUM_CNI_CONF=${CILIUM_CNI_CONF:-${HOST_PREFIX}/etc/cni/net.d/${CNI_CONF_NAME}}

if [ ! -f ${CNI_DIR}/bin/loopback ]; then
	echo "Installing loopback driver..."
	mkdir tmp && cd tmp
	wget https://storage.googleapis.com/kubernetes-release/network-plugins/cni-0799f5732f2a11b329d9e3d51b9c8f2e3759f2ff.tar.gz
	mkdir -p ${CNI_DIR}/bin
	sudo tar -xvf cni-0799f5732f2a11b329d9e3d51b9c8f2e3759f2ff.tar.gz
	cp bin/loopback ${CNI_DIR}/bin/
	cd ..
	rm -r tmp
fi

echo "Installing cilium-cni to ${CNI_DIR}/bin/ ..."
cp /opt/cni/bin/cilium-cni ${CNI_DIR}/bin/

cat > ${CNI_CONF_NAME} <<EOF
{
    "name": "cilium",
    "type": "cilium-cni",
    "mtu": ${MTU}
}
EOF

# Install CNI configuration to host
echo "Installing $CILIUM_CNI_CONF ..."
mv ${CNI_CONF_NAME} ${CILIUM_CNI_CONF}
