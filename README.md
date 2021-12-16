
==================================================================================
       Copyright (c) 2020 AT&T Intellectual Property.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
==================================================================================

RC Xapp
================

This repository contains the source for sending RIC control Msg to RAN/E2 Node  based on the GRPC control request received from other xapps

This xApp can be onboarded through the xApp Onboarder.  The xapp descriptor 
is under the xapp-descriptor/ directory.

Then the xapp can be deployed through the App Manager. 
Procedure to Onboard the xapp is given in installation guide, in docs folder

GRPC communication can be tested using the grpccurl command


In the current implementation, timer is started after initiating Control Request to E2Node. So timer expiry errors will be seen in logs w.r.t. Control Resp.
As control ACK is not handled this error can be ignored

Steps to onboard the xapp 
 1)echo '{ "config-file.json_url": "https://gerrit.o-ran-sc.org/r/gitweb?p=ric-app/rc.git;a=blob_plain;f=xapp-descriptor/config.json;hb=HEAD","controls-schema.json_url": "https://gerrit.o-ran-sc.org/r/gitweb?p=ric-app/rc.git;a=blob_plain;f=xapp-descriptor/schema.json;hb=HEAD" }' > onboard.rc.url

2)Onboard the rc xapp using the below command :
       $curl --location --request POST "http://$(hostname):32080/onboard/api/v1/onboard/download"      --header 'Content-Type: application/json' --data-binary "@./onboard.rc.url"  

       After this command: xapp should be created
       Eg:
       $curl --location --request POST "http://$(hostname):32080/onboard/api/v1/onboard/download"      --header 'Content-Type: application/json' --data-binary "@./onboard.rc.url"
       {
       "status": "Created"
       }
 3)Deploying the xapp:
       Eg:
       $curl --location --request POST "http://$(hostname):32080/appmgr/ric/v1/xapps"      --header 'Content-Type: application/json' --data-raw '{"xappName": "rc"}'
       {
       "instances":null,"name":"rc","status":"deployed","version":"1.0"
       }

         after this step the xapp should be in running state

         ricxapp       ricxapp-rc-5ccdcc9f6f-ktq7w                                  1/1     Running     0          88m

 4) restart rtMgr so rmr is ready
