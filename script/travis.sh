#!/bin/sh

# Forked from the following file:
# https://github.com/libgit2/git2go/blob/master/script/install-libgit2.sh

set -ex

cd "${HOME}"
LG2VER="0.24.0"
wget -O libgit2-${LG2VER}.tar.gz https://github.com/libgit2/libgit2/archive/v${LG2VER}.tar.gz
tar -xzvf libgit2-${LG2VER}.tar.gz
cd libgit2-${LG2VER}
mkdir build && cd build
cmake -DTHREADSAFE=ON -DBUILD_CLAR=OFF -DCMAKE_BUILD_TYPE="RelWithDebInfo" .. && make && sudo make install
sudo ldconfig
cd "${TRAVIS_BUILD_DIR}"
