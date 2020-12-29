#!/bin/bash
BIN=bin/comparator
EXAMPLES=comparator/test_examples

${BIN} ${EXAMPLES}/simple.json ${EXAMPLES}/simple_unordered.json
if [ $? -ne 0 ]; then
  echo "Failed to detect that the two files are similar"
  exit 1
fi

${BIN} ${EXAMPLES}/simple.json ${EXAMPLES}/simple_different_value.json
if [ $? -ne 1 ]; then
	echo "Failed to detect that the two files are different"
	exit 1
fi

exit 0
