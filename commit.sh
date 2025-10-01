#!/bin/bash
PROTOCOL_COMMIT=$(git submodule status protocol | awk '{print substr($1, 1, 7)}')
git commit -m "bump to protocol@${PROTOCOL_COMMIT}"
