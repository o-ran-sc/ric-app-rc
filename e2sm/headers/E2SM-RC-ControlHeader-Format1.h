/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-RC-IEs"
 * 	found in "e2sm-rc-2.0.asn"
 * 	`asn1c -fcompound-names -fno-include-deps -findirect-choice -pdu=auto -gen-PER -gen-OER -no-gen-example -D .`
 */

#ifndef	_E2SM_RC_ControlHeader_Format1_H_
#define	_E2SM_RC_ControlHeader_Format1_H_


#include <asn_application.h>

/* Including external dependencies */
#include "UEID.h"
#include "RIC-Style-Type.h"
#include "RIC-ControlAction-ID.h"
#include <NativeEnumerated.h>
#include <constr_SEQUENCE.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum E2SM_RC_ControlHeader_Format1__ric_ControlDecision {
	E2SM_RC_ControlHeader_Format1__ric_ControlDecision_accept	= 0,
	E2SM_RC_ControlHeader_Format1__ric_ControlDecision_reject	= 1
	/*
	 * Enumeration is extensible
	 */
} e_E2SM_RC_ControlHeader_Format1__ric_ControlDecision;

/* E2SM-RC-ControlHeader-Format1 */
typedef struct E2SM_RC_ControlHeader_Format1 {
	UEID_t	 ueID;
	RIC_Style_Type_t	 ric_Style_Type;
	RIC_ControlAction_ID_t	 ric_ControlAction_ID;
	long	*ric_ControlDecision;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} E2SM_RC_ControlHeader_Format1_t;

/* Implementation */
/* extern asn_TYPE_descriptor_t asn_DEF_ric_ControlDecision_5;	// (Use -fall-defs-global to expose) */
extern asn_TYPE_descriptor_t asn_DEF_E2SM_RC_ControlHeader_Format1;
extern asn_SEQUENCE_specifics_t asn_SPC_E2SM_RC_ControlHeader_Format1_specs_1;
extern asn_TYPE_member_t asn_MBR_E2SM_RC_ControlHeader_Format1_1[4];

#ifdef __cplusplus
}
#endif

#endif	/* _E2SM_RC_ControlHeader_Format1_H_ */
#include <asn_internal.h>
