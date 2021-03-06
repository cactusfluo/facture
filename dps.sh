#!/usr/bin/env bash

C_RED="\033[31;01m"
C_GREEN="\033[32;01m"
C_YELLOW="\033[33;01m"
C_BLUE="\033[34;01m"
C_PINK="\033[35;01m"
C_CYAN="\033[36;01m"
C_NO="\033[0m"

################################################################################
###                                FUNCTIONS                                 ###
################################################################################

################################################################################
###                                   MAIN                                   ###
################################################################################

interval=${1}
interval=${interval:=2}

watch \
	--no-title \
	--interval ${interval} \
	--color \
	'docker ps --all --format "{{ printf \"%-40s\" .Names }}\t{{.Status}}" | sort -r'
