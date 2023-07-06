# Base image
FROM dart:2.16-sdk

# Maintainer
LABEL maintainer="nagakuta <xxx@example.com>"

# Workdir
WORKDIR /workspace/board_game_direct

# Install Flutter
ARG PATH=/root/.pub-cache/bin:$PATH
ARG FLUTTER_VERSION=2.10.5
RUN dart pub global activate melos --verbose && \
  dart pub global activate fvm --verbose && \
  fvm doctor --verbose && \
  fvm install $FLUTTER_VERSION --verbose && \
  fvm use --force $FLUTTER_VERSION --verbose && \
  fvm flutter config --enable-web --enable-linux-desktop --enable-macos-desktop --enable-windows-desktop --enable-android --enable-ios --enable-fuchsia && \
  # fvm flutter precache --verbose && \
  fvm flutter doctor --verbose

# Set paths
ENV FVM_ROOT=/root/.pub-cache
ENV PATH $FVM_ROOT/bin:$PATH
