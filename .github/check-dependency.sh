#!/bin/bash

ALLOWED="$(cat .allowpackages)"
DEPENDENCIES="$(go list -m)"
for DEP in ${DEPENDENCIES[@]}; do
  echo "Check package: ${DEP}"
  if [ ! $(echo "${ALLOWED}" | grep "${DEP}") ]; then
    echo "--- ERROR"
    echo "Package import is not allowd: ${DEP}"
    echo "To allow this package, add this to .allowpackages"
    exit 1
  fi
done
