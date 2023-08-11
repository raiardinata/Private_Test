#!/bin/bash

PGUSER="rai"
PGPASSWORD="rai"
PGDATABASE="cognotiv_test"
PGHOST="localhost"
PGPORT="5432"

psql -U $PGUSER -h $PGHOST -p $PGPORT -d $PGDATABASE -c "SELECT send_pending_order_reminders();"
