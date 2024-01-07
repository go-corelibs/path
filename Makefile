#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

VERSION_TAGS += PATH
PATH_MK_SUMMARY := go-corelibs/path
PATH_MK_VERSION := v1.0.1

include CoreLibs.mk
