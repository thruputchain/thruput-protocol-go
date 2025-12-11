#!/bin/bash
PROTOCOL_COMMIT=$(git -C ./protocol rev-parse --short HEAD)
git commit -m "bump to protocol@${PROTOCOL_COMMIT}"
