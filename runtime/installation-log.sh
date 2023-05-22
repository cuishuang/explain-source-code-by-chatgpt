#!/bin/sh

do_logging() {
	url="https://community.fit2cloud.com/installation-analytics?product=$1&type=$2&version=$3"
	echo $url
	nohup curl --connect-timeout 5 -m 10 -k $url > /dev/null 2>&1 &
}

{ 
	do_logging $1 $2 $3  > /dev/null 2>&1
} || { 
	echo 1 > /dev/null 2>&1
}
