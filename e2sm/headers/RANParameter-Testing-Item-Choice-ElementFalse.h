/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-RC-IEs"
 * 	found in "e2sm-rc-2.0.asn"
 * 	`asn1c -fcompound-names -fno-include-deps -findirect-choice -pdu=auto -gen-PER -gen-OER -no-gen-example -D .`
 */

#ifndef	_RANParameter_Testing_Item_Choice_ElementFalse_H_
#define	_RANParameter_Testing_Item_Choice_ElementFalse_H_


#include <asn_application.h>

/* Including external dependencies */
#include "RANParameter-TestingCondition.h"
#include "LogicalOR.h"
#include <constr_SEQUENCE.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct RANParameter_Value;

/* RANParameter-Testing-Item-Choice-ElementFalse */
typedef struct RANParameter_Testing_Item_Choice_ElementFalse {
	RANParameter_TestingCondition_t	 ranParameter_TestCondition;
	struct RANParameter_Value	*ranParameter_Value;	/* OPTIONAL */
	LogicalOR_t	*logicalOR;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} RANParameter_Testing_Item_Choice_ElementFalse_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_RANParameter_Testing_Item_Choice_ElementFalse;
extern asn_SEQUENCE_specifics_t asn_SPC_RANParameter_Testing_Item_Choice_ElementFalse_specs_1;
extern asn_TYPE_member_t asn_MBR_RANParameter_Testing_Item_Choice_ElementFalse_1[3];

#ifdef __cplusplus
}
#endif

#endif	/* _RANParameter_Testing_Item_Choice_ElementFalse_H_ */
#include <asn_internal.h>
