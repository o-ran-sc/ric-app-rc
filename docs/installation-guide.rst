.. This work is licensed under a Creative Commons Attribution 4.0 International License.
.. SPDX-License-Identifier: CC-BY-4.0
.. Copyright (C) 2020 AT&T Intellectual Property


Installation Guide
==================

.. contents::
   :depth: 3
   :local:

Installation
------------

Steps to onboard the xapp
 1)echo '{ "config-file.json_url": "https://gerrit.o-ran-sc.org/r/gitweb?p=ric-app/rc.git;a=blob_plain;f=xapp-descriptor/config.json;hb=HEAD","controls-schema.json_url": "https://gerrit.o-ran-sc.org/r/gitweb?p=ric-app/rc.git;a=blob_plain;f=xapp-descriptor/schema.json;hb=HEAD" }' > onboard.rc.url

 2)Onboard the rc xapp using the below command :
       curl --location --request POST "http://$(hostname):32080/onboard/api/v1/onboard/download"      --header 'Content-Type: application/json' --data-binary "@./onboard.rc.url"  

       After this command: xapp should be created
       Eg:
       curl --location --request POST "http://$(hostname):32080/onboard/api/v1/onboard/download"      --header 'Content-Type: application/json' --data-binary "@./onboard.rc.url"
       {
           "status": "Created"
       }
 3)Deploying the xapp:
         curl --location --request POST "http://$(hostname):32080/appmgr/ric/v1/xapps" --header 'Content-Type: application/json' --data-raw '{"xappName": "rc"}' 

         after this step the xapp should be in running state

         ricxapp       ricxapp-rc-5ccdcc9f6f-ktq7w                                  1/1     Running     0          88m
        
 4) restart rtMgr so rmr is ready 
