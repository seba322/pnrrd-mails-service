#!/bin/bash

localDbPath=$1

sudo mongo < ./createUser_clearDB.js

sudo mongorestore -d PNRRDPROD ./dump

