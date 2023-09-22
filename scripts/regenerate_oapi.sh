#!/bin/bash

oapi-codegen --config api/types.cfg.yaml api/openapi.yaml
oapi-codegen --config api/server.cfg.yaml api/openapi.yaml