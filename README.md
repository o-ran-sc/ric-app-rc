
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


In the current implementation, timer is started after initiating Control Request to E2Node,and also
supports RIC Control ACK response from E2Node.

Steps to onboard the xapp (Pre Dawn Release)

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

Steps to deploy the xapp using dms_cli tool

1)Create a local helm repository with a port other than 8080 on host
        $docker run --rm -u 0 -it -d -p 8090:8080 -e DEBUG=1 -e STORAGE=local -e STORAGE_LOCAL_ROOTDIR=/charts -v $(pwd)/charts:/charts chartmuseum/chartmuseum:latest

2)Set CHART_REPO_URL env variable
        $export CHART_REPO_URL=http://0.0.0.0:8090

3)Check dms_cli health is active(true) or not
        $dms_cli health
           True

4)Onboard the rc xapp using the below command :
Make sure that you have the xapp descriptor config file and the schema file at your local file system
        $dms_cli onboard --config_file_path=config-file.json --shcema_file_path=schema.json
        After this command: xapp should be created
        Eg:
        $dms_cli onboard --config_file_path=config-file.json --shcema_file_path=schema.json
        {
        "status": "Created"
        }

5)List the helm charts from helm repository (Optional)
        $curl -X GET http://localhost:8090/api/charts | jq .
        Eg:
        $curl -X GET http://$(hostname):8090/api/charts | jq .
{
 "rc": [
    {
      "name": "rc",
      "version": "1.0.0",
      "description": "Standard xApp Helm Chart",
      "apiVersion": "v1",
      "appVersion": "1.0",
      "urls": [
        "charts/rc-1.0.0.tgz"
      ],
      "created": "2022-04-20T05:26:37.891030646Z",
      "digest": "3636f4fba220aab5739d737eb3cd3d74b5012b029bde95f114b05f9eaa786b1f"
    }
  ]
}

6) Install the xApp
        $dms_cli install --xapp_chart_name=XAPP_CHART_NAME --version=VERSION --namespace=NAMESPACE
        Eg:
        $dms_cli install --xapp_chart_name=rc --version=1.0.0 --namespace=ricxapp
        status: OK
          after this step the xapp should be in running state

          ricxapp       ricxapp-rc-5d8b97d58d-hhznh                                 1/1     Running     1          60m

7) restart rtMgr so rmr is ready
