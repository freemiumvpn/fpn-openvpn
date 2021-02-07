#!/bin/bash
EASY_RSA_LOC="/etc/openvpn/certs"
MY_IP_ADDR="$2"

cd $EASY_RSA_LOC

./easyrsa build-client-full $1 nopass

cat >${EASY_RSA_LOC}/pki/$1.ovpn <<EOF
client
nobind
dev tun
remote ${MY_IP_ADDR} 443 tcp
cipher AES-256-CBC
redirect-gateway def1
<key>
`cat ${EASY_RSA_LOC}/pki/private/$1.key`
</key>
<cert>
`cat ${EASY_RSA_LOC}/pki/issued/$1.crt`
</cert>
<ca>
`cat ${EASY_RSA_LOC}/pki/ca.crt`
</ca>
<tls-auth>
`cat ${EASY_RSA_LOC}/pki/ta.key`
</tls-auth>
key-direction 1
EOF
