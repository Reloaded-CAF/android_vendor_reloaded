# Whether using or not release-keys for building
# Copyright (C) 2018 KudProject Development
# SPDX-License-Identifier: Apache-2.0

# Don't modify; path where release-keys are hosted
RELEASE_KEYS_PATH := vendor/reloaded/keys

# Whether release-keys path is exist or not
ifneq ($(wildcard $(RELEASE_KEYS_PATH)),)
    # Exists, use them
    USE_RELEASE_KEYS := true
else
    # Doesn't exist, fall back to test-keys
    $(warn $(RELEASE_KEYS_PATH) doesn't exist. Continuing with test-keys...)
endif
