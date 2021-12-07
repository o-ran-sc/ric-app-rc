
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

GRPC communication can be tested using the grpccurl command
