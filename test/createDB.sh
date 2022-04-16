#!/bin/bash
rm DB.sqlite3
sqlite3 DB.sqlite3 ".read db.sql"
