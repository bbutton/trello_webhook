#!/usr/bin/env bash

if [ $# -ne 6 ]; then
    echo "usage: create_webhook.sh <webhook_name> <cloudfoundry_url> <webhook_friendly_name> <board_id> <trello dev id> <trello token>" >&2
    echo "where" >&2
    echo -e "\twebhook_name - name of deployed webhook in AppFog/CloudFoundry" >&2
    echo -e "\tcloudfoundry_url - URL of CloudFoundry instance" >&2
    echo -e "\twebhook_friendly_name - user-readable name for webhook that will appear in the logs" >&2
    echo -e "\tboard_id - Trello board id, which can be found through find_board_id.sh" >&2
    echo -e "\ttrello_dev_id - User\'s application/developer key" >&2
    echo -e "\ttrello_token - User\'s secret token" >&2
    exit 1
fi

WEBHOOK_NAME=$1
CLOUDFOUNDRY_URL=$2
WEBHOOK_FRIENDLY_NAME=$3
BOARD_ID=$4
TRELLO_DEV_ID=$5
TRELLO_TOKEN=$6

curl -H "Content-Type: application/json" -X POST -d "{\"description\":\"Callback for ${WEBHOOK_FRIENDLY_NAME}\", \"callbackURL\":\"https://${WEBHOOK_NAME}.${CLOUDFOUNDRY_URL}/trello_webhook/${WEBHOOK_FRIENDLY_NAME}\", \"idModel\":\"${BOARD_ID}\"}"  https://trello.com/1/tokens/${TRELLO_TOKEN}/webhooks?key=${TRELLO_DEV_ID}
