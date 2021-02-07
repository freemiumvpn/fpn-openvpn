#!/bin/bash

USE_CERTIFICATE_REVOCATION_LIST="true"
USE_TLS_AUTH="true"

EASY_RSA_LOC="/etc/openvpn/certs"
SERVER_CERT="${EASY_RSA_LOC}/pki/issued/server.crt"

mkdir -p /etc/openvpn/certs
mkdir -p /etc/openvpn/ccd

cd $EASY_RSA_LOC

if [ -e "$SERVER_CERT" ]
then
    echo "found existing certs - reusing"
    
    if [ "$USE_CERTIFICATE_REVOCATION_LIST" == "true" ]
    then
        if [ ! -e ${EASY_RSA_LOC}/crl.pem ]
        then
            echo "generating missed crl file"
            ./easyrsa gen-crl
            cp ${EASY_RSA_LOC}/pki/crl.pem ${EASY_RSA_LOC}/crl.pem
            chmod 644 ${EASY_RSA_LOC}/crl.pem
        fi
    fi
    
    
    if [ "$USE_TLS_AUTH" == "true" ]
    then
        if [ ! -e ${EASY_RSA_LOC}/pki/ta.key ]
        then
            echo "generating missed ta.key"
            openvpn --genkey --secret ${EASY_RSA_LOC}/pki/ta.key
        fi
    fi
else
    cp -R /usr/share/easy-rsa/* $EASY_RSA_LOC
    ./easyrsa init-pki
    echo "ca\n" | ./easyrsa build-ca nopass
    ./easyrsa build-server-full server nopass
    ./easyrsa gen-dh
    
    if [ "$USE_CERTIFICATE_REVOCATION_LIST" == "true" ]
    then
        ./easyrsa gen-crl
        # Note: the pki/ directory is inaccessible after openvpn drops privileges,
        # so we cp crl.pem to ${EASY_RSA_LOC} to allow CRL updates without a restart
        cp ${EASY_RSA_LOC}/pki/crl.pem ${EASY_RSA_LOC}/crl.pem
        chmod 644 ${EASY_RSA_LOC}/crl.pem
    fi
    
    if [ "$USE_TLS_AUTH" == "true" ]
    then
        openvpn --genkey --secret ${EASY_RSA_LOC}/pki/ta.key
    fi
fi
