#!/bin/bash
PROTOCOL_COMMIT=$(git -C ./protocol rev-parse HEAD)
git commit -m "bump to protocol@${PROTOCOL_COMMIT}"
