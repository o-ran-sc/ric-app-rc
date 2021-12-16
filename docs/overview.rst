.. This work is licensed under a Creative Commons Attribution 4.0 International License.
.. SPDX-License-Identifier: CC-BY-4.0
.. Copyright (C) 2020 AT&T Intellectual Property

rc Control Overview
====================

rc is an Xapp used to send RIC control message to E2 / Ran Node based on the GRPC control request received.
rc Xapp  acts as a GRPC server.It receives and processes the GRPC control request to construct the RIC Control message.
ASN1 encoding of RIC Control Message will be done and sent to E2/RAN Node using RMR.
