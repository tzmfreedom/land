#!/bin/bash

NAME="land"
BIN_NAME="land"
VERSION="0.1.2"
PREFIX="/usr/local"
COMP_PREFIX="/usr/local/share/zsh/site-functions"
GITHUB_USER="tzmfreedom"
TMP_DIR="/tmp"
ZSH_COMPLETION=""

set -ue

function parse_options() {
  for OPT in "$@"
  do
    case "$OPT" in
      "--zsh-completion" )
        ZSH_COMPLETION="t"
        ;;
      -* )
        echo "$PROGRAM: illegal option -- '$(echo $1 | sed 's/^-*//')'" 1>&2
        exit 1
        ;;
    esac
    shift
  done
}

parse_options $@

UNAME=$(uname -s)
if [ "$UNAME" != "Linux" -a "$UNAME" != "Darwin" ] ; then
    echo "Sorry, OS not supported: ${UNAME}. Download binary from https://github.com/${USERNAME}/${NAME}/releases"
    exit 1
fi


if [ "${UNAME}" = "Darwin" ] ; then
  OS="darwin"

  OSX_ARCH=$(uname -m)
  if [ "${OSX_ARCH}" = "x86_64" ] ; then
    ARCH="amd64"
  else
    echo "Sorry, architecture not supported: ${OSX_ARCH}. Download binary from https://github.com/${USERNAME}/${NAME}/releases"
    exit 1
  fi
elif [ "${UNAME}" = "Linux" ] ; then
  OS="linux"

  LINUX_ARCH=$(uname -m)
  if [ "${LINUX_ARCH}" = "i686" ] ; then
    ARCH="386"
  elif [ "${LINUX_ARCH}" = "x86_64" ] ; then
    ARCH="amd64"
  else
    echo "Sorry, architecture not supported: ${LINUX_ARCH}. Download binary from https://github.com/${USERNAME}/${NAME}/releases"
    exit 1
  fi
fi

ARCHIVE_FILE=${BIN_NAME}-${VERSION}-${OS}-${ARCH}.tar.gz
BINARY="https://github.com/${GITHUB_USER}/${NAME}/releases/download/v${VERSION}/${ARCHIVE_FILE}"

cd $TMP_DIR
curl -sL -O ${BINARY}

tar xzf ${ARCHIVE_FILE}
mv ${OS}-${ARCH}/${BIN_NAME} ${PREFIX}/bin/${BIN_NAME}
chmod +x ${PREFIX}/bin/${BIN_NAME}

# completion
if [ -d ${COMP_PREFIX} ]; then
  if [ "${ZSH_COMPLETION}" == "t" ]; then
    mv ${OS}-${ARCH}/_${BIN_NAME} ${COMP_PREFIX}/_${BIN_NAME}
  fi
fi

# clean
rm -rf ${OS}-${ARCH}
rm -rf ${ARCHIVE_FILE}
