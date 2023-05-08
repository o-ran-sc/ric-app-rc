/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-RC-IEs"
 * 	found in "e2sm-rc-2.0.asn"
 * 	`asn1c -fcompound-names -fno-include-deps -findirect-choice -pdu=auto -gen-PER -gen-OER -no-gen-example -D .`
 */

#ifndef	_E2SM_RC_EventTrigger_Format4_Item_H_
#define	_E2SM_RC_EventTrigger_Format4_Item_H_


#include <asn_application.h>

/* Including external dependencies */
#include "RIC-EventTriggerCondition-ID.h"
#include "TriggerType-Choice.h"
#include "LogicalOR.h"
#include <constr_SEQUENCE.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct EventTrigger_UE_Info;

/* E2SM-RC-EventTrigger-Format4-Item */
typedef struct E2SM_RC_EventTrigger_Format4_Item {
	RIC_EventTriggerCondition_ID_t	 ric_eventTriggerCondition_ID;
	TriggerType_Choice_t	 triggerType;
	struct EventTrigger_UE_Info	*associatedUEInfo;	/* OPTIONAL */
	LogicalOR_t	*logicalOR;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} E2SM_RC_EventTrigger_Format4_Item_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_E2SM_RC_EventTrigger_Format4_Item;
extern asn_SEQUENCE_specifics_t asn_SPC_E2SM_RC_EventTrigger_Format4_Item_specs_1;
extern asn_TYPE_member_t asn_MBR_E2SM_RC_EventTrigger_Format4_Item_1[4];

#ifdef __cplusplus
}
#endif

#endif	/* _E2SM_RC_EventTrigger_Format4_Item_H_ */
#include <asn_internal.h>
