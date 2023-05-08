/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-RC-IEs"
 * 	found in "e2sm-rc-2.0.asn"
 * 	`asn1c -fcompound-names -fno-include-deps -findirect-choice -pdu=auto -gen-PER -gen-OER -no-gen-example -D .`
 */

#ifndef	_RIC_InsertIndication_Name_H_
#define	_RIC_InsertIndication_Name_H_


#include <asn_application.h>

/* Including external dependencies */
#include <PrintableString.h>

#ifdef __cplusplus
extern "C" {
#endif

/* RIC-InsertIndication-Name */
typedef PrintableString_t	 RIC_InsertIndication_Name_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_RIC_InsertIndication_Name_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_RIC_InsertIndication_Name;
asn_struct_free_f RIC_InsertIndication_Name_free;
asn_struct_print_f RIC_InsertIndication_Name_print;
asn_constr_check_f RIC_InsertIndication_Name_constraint;
ber_type_decoder_f RIC_InsertIndication_Name_decode_ber;
der_type_encoder_f RIC_InsertIndication_Name_encode_der;
xer_type_decoder_f RIC_InsertIndication_Name_decode_xer;
xer_type_encoder_f RIC_InsertIndication_Name_encode_xer;
oer_type_decoder_f RIC_InsertIndication_Name_decode_oer;
oer_type_encoder_f RIC_InsertIndication_Name_encode_oer;
per_type_decoder_f RIC_InsertIndication_Name_decode_uper;
per_type_encoder_f RIC_InsertIndication_Name_encode_uper;
per_type_decoder_f RIC_InsertIndication_Name_decode_aper;
per_type_encoder_f RIC_InsertIndication_Name_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _RIC_InsertIndication_Name_H_ */
#include <asn_internal.h>
