#!/bin/bash
# Restik Configuration Update script for Environment based on parameters

FILE_PATH=./config.json

ListenURL="{{PORT}}"

# Logger
LogLevel="{{log_level}}"
UseLogFile="{{use_log_file}}"
LogFile="{{log_file}}"

# Postgres
Host="{{postgres_host}}"
Port="{{postgres_port}}"
User="{{postgres_user}}"
DBName="{{postgres_db}}"
Password="{{postgres_password}}"
# Heroku PG
HerokuPg="{{DATABASE_URL}}"

echo "####################### Environment Variables ################################"
echo "Log Level           = " $LogLevel
echo "Use log file Level  = " $UseLogFile
echo "Listen Port         = " $ListenURL
echo "File Path           = " $FILE_PATH
echo "Postgres:				"
echo "	Host              = " $Host
echo "	Port              = " $Port
echo "	User              = " $User
echo "	DBName            = " $DBName
echo "	Password          = " $Password
echo "	Heroku            = " $HerokuPg

echo "####################### Environment Variables ################################"

function var_replace() {
  key=$1
  value=$2
  awk -v val="$value" "/$key/{\$2=val}1" $FILE_PATH > tmp_file && mv tmp_file $FILE_PATH
}

var_replace ListenURL "\":$ListenURL\","

# Logger
var_replace	LogLevel "\":$LogLevel\","
var_replace UseLogFile "\":$UseLogFile\","
var_replace LogFile "\":$LogFile\","

# Postgres
var_replace Host "\":$Host\","
var_replace Port "\":$Port\","
var_replace User "\":$User\","
var_replace DBName "\":$DBName\","
var_replace Password "\":$Password\","



echo "########### Updated Config File with Environment Variables #############"