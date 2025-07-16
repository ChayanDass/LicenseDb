#!/bin/bash
# SPDX-FileCopyrightText: 2024 Kaushlendra Pratap <kaushlendra-pratap.singh@siemens.com>
# SPDX-License-Identifier: GPL-2.0-only

set -e

populate_db="${POPULATE_DB:-true}"
data_file="/app/licenseRef.json"

printf "READ_API_AUTHENTICATION_ENABLED=false\nTOKEN_HOUR_LIFESPAN=24\nAPI_SECRET=%s\nDEFAULT_ISSUER=www.licensedb.com\n" "$(openssl rand -hex 32)" > /app/.env
/app/init.sh
exec /app/laas -datafile="$data_file" -populatedb=$populate_db
