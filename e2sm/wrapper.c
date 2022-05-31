#include <errno.h>
#include "wrapper.h"
#include "OCTET_STRING.h"

//ssize_t e2sm_encode_ric_control_header(void *buffer, size_t buf_size,struct uEID *inUEID,long f1AP[1],long e1AP[1],long ricControlStyleType, long ricControlActionID)
ssize_t e2sm_encode_ric_control_header(void *buffer, size_t buf_size,struct uEID *inUEID,long f1AP[],size_t f1AP_len,long e1AP[],size_t e1Ap_len,long ricControlStyleType, long ricControlActionID, void* plmnId, size_t  plmnIdSize)
{
        fprintf(stderr,"e2SM wrapper function Entered\n");	
	fprintf(stderr,"plmn Size = %ld and aMFRegionID Size = %ld and aMFSetID_size = %ld and aMFPointer_size = %ld \n", inUEID->pLMNIdentity_size,inUEID->aMFRegionID_size,inUEID->aMFSetID_size,inUEID->aMFPointer_size);

        E2SM_RC_ControlHeader_t *controlHeaderIE = (E2SM_RC_ControlHeader_t *)calloc(1, sizeof(E2SM_RC_ControlHeader_t));
        if(!controlHeaderIE)
        {
                fprintf(stderr, "alloc E2SM_RC_ControlHeader failed\n");
                   return -1;
        }

        controlHeaderIE->ric_controlHeader_formats.present = E2SM_RC_ControlHeader__ric_controlHeader_formats_PR_controlHeader_Format1;
        E2SM_RC_ControlHeader_Format1_t  *controlHeader_Fmt1 = (E2SM_RC_ControlHeader_Format1_t *)calloc(1, sizeof(E2SM_RC_ControlHeader_Format1_t));
        if(!controlHeader_Fmt1)
        {
                fprintf(stderr, "alloc E2SM_RC_ControlHeader failed\n");
                return -1;
        }
       
	controlHeader_Fmt1->ueID.present = UEID_PR_gNB_UEID;
	controlHeader_Fmt1->ueID.choice.gNB_UEID = (UEID_GNB_t *)calloc(1,sizeof(UEID_GNB_t));
	if(! controlHeader_Fmt1->ueID.choice.gNB_UEID)
	{
		ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlHeader, controlHeaderIE);
		fprintf(stderr, "alloc gNB_UEID failed\n");
                return -1;
	}

	asn_long2INTEGER(&controlHeader_Fmt1->ueID.choice.gNB_UEID->amf_UE_NGAP_ID,inUEID->amf_UE_NGAP_Id);

	fprintf(stderr, "e2sm_encode_ric_control_header amf_UE_NGAP_ID encoded \n");
	
	//OCTET_STRING_fromBuf(&controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.pLMNIdentity,inUEID->pLMNIdentity, inUEID->pLMNIdentity_size);

	OCTET_STRING_fromBuf(&controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.pLMNIdentity, plmnId,plmnIdSize);
	fprintf(stderr, "e2smrc_encode_ric_control_header pLMNIdentity encoded \n");

	controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.aMFRegionID.buf = (uint8_t*)calloc(1,inUEID->aMFRegionID_size);
        if(!controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.aMFRegionID.buf)
        {
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlHeader, controlHeaderIE);
                fprintf(stderr, "alloc aMFRegionID Value failed\n");
                return -1;
        }
	controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.aMFRegionID.size = inUEID->aMFRegionID_size;
	//controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.aMFRegionID.buf[0] = inUEID->aMFRegionID && 0XFF;
	//controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.aMFRegionID.buf[0] = *(inUEID->aMFRegionID ) & 0XFF;
	memcpy(controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.aMFRegionID.buf,inUEID->aMFRegionID,inUEID->aMFRegionID_size);

	controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.aMFSetID.buf = (uint8_t*)calloc(1, inUEID->aMFSetID_size);
        if(!controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.aMFSetID.buf)
        {
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlHeader, controlHeaderIE);
                fprintf(stderr, "alloc aMFSetID Value failed\n");
                return -1;
        }
	controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.aMFSetID.size = inUEID->aMFSetID_size;
	//controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.aMFSetID.buf[0] = (inUEID->aMFSetID && 0xFF00) >> 8;
	//controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.aMFSetID.buf[1] = inUEID->aMFSetID && 0XFF00 && 0X00FF;
	controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.aMFSetID.buf[0] = (*(inUEID->aMFSetID) & 0xFFC0) >> 8;
	controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.aMFSetID.bits_unused = 6;


	controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.aMFPointer.buf = (uint8_t*)calloc(1, inUEID->aMFPointer_size);
        if(!controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.aMFPointer.buf)
        {
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlHeader, controlHeaderIE);
                fprintf(stderr, "alloc aMFPointer Value failed\n");
                return -1;
        }

	controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.aMFPointer.size = inUEID->aMFPointer_size;
	//controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.aMFPointer.buf[0] = inUEID->aMFPointer && 0XFF;
	controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.aMFPointer.buf[0] = (*(inUEID->aMFPointer) & 0xFC) >> 2;
	controlHeader_Fmt1->ueID.choice.gNB_UEID->guami.aMFPointer.bits_unused = 2;


	controlHeader_Fmt1->ueID.choice.gNB_UEID->gNB_CU_UE_F1AP_ID_List = (UEID_GNB_CU_F1AP_ID_List_t *)calloc(1,sizeof(UEID_GNB_CU_F1AP_ID_List_t));
	if(! controlHeader_Fmt1->ueID.choice.gNB_UEID->gNB_CU_UE_F1AP_ID_List)
        {
                fprintf(stderr, "alloc gNB_CU_CP_UE_E1AP_ID_list failed\n");
		ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlHeader, controlHeaderIE);
                return -1;
        }
	/*
	UEID_GNB_CU_CP_F1AP_ID_Item_t *F1AP_ID_Item = (UEID_GNB_CU_CP_F1AP_ID_Item_t *)calloc (1, sizeof(UEID_GNB_CU_CP_F1AP_ID_Item_t ));
        if(! F1AP_ID_Item)
        {
        	fprintf(stderr, "alloc UEID_GNB_CU_CP_F1AP_ID_Item failed\n");
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlHeader, controlHeaderIE);
                return -1;
        }
	*/
	//f1AP is an array of data
	//int n = sizeof(f1AP)/sizeof(long int);
	for(int i =0; i < f1AP_len; i++)
	{
		UEID_GNB_CU_CP_F1AP_ID_Item_t *F1AP_ID_Item = (UEID_GNB_CU_CP_F1AP_ID_Item_t *)calloc (1, sizeof(UEID_GNB_CU_CP_F1AP_ID_Item_t ));
		if(! F1AP_ID_Item)
		{
                	fprintf(stderr, "alloc UEID_GNB_CU_CP_F1AP_ID_Item failed\n");
			ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlHeader, controlHeaderIE);
                	return -1;
		} 
		F1AP_ID_Item->gNB_CU_UE_F1AP_ID  = f1AP[i];
		fprintf(stderr, "F1AP_id %d =%lu\n",i,*((unsigned long *) F1AP_id +i ));
		ASN_SEQUENCE_ADD(&controlHeader_Fmt1->ueID.choice.gNB_UEID->gNB_CU_UE_F1AP_ID_List->list,F1AP_ID_Item);
	}
	//F1AP_ID_Item->gNB_CU_UE_F1AP_ID  = f1AP[0];
	//ASN_SEQUENCE_ADD(&controlHeader_Fmt1->ueID.choice.gNB_UEID->gNB_CU_UE_F1AP_ID_List->list,F1AP_ID_Item);


	
	controlHeader_Fmt1->ueID.choice.gNB_UEID->gNB_CU_CP_UE_E1AP_ID_List = (UEID_GNB_CU_CP_E1AP_ID_List_t *)calloc(1,sizeof(UEID_GNB_CU_CP_E1AP_ID_List_t));

	if(! controlHeader_Fmt1->ueID.choice.gNB_UEID->gNB_CU_CP_UE_E1AP_ID_List)
	{
		fprintf(stderr, "alloc gNB_CU_CP_UE_E1AP_ID_list failed\n");
		ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlHeader, controlHeaderIE);
                return -1;
	}

	
	//n = sizeof(e1AP)/sizeof(long int);
	for(int i =0; i < e1Ap_len; i++)
	{
		UEID_GNB_CU_CP_E1AP_ID_Item_t *E1AP_ID_Item = (UEID_GNB_CU_CP_E1AP_ID_Item_t *)calloc (1, sizeof(UEID_GNB_CU_CP_E1AP_ID_Item_t ));
        	if(! E1AP_ID_Item)
        	{
                	fprintf(stderr, "alloc UEID_GNB_CU_CP_E1AP_ID_Item failed\n");
			ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlHeader, controlHeaderIE);
                	return -1;

        	}
		E1AP_ID_Item->gNB_CU_CP_UE_E1AP_ID = e1AP[i];
         	ASN_SEQUENCE_ADD(&controlHeader_Fmt1->ueID.choice.gNB_UEID->gNB_CU_CP_UE_E1AP_ID_List->list,E1AP_ID_Item);
	//E1AP_ID_Item->gNB_CU_CP_UE_E1AP_ID = (*(unsigned long *) E1AP_id) ;
	//ASN_SEQUENCE_ADD(&controlHeader_Fmt1->ueID.choice.gNB_UEID->gNB_CU_CP_UE_E1AP_ID_List->list,E1AP_ID_Item);
	}
	

	 /*UEID_GNB_CU_CP_E1AP_ID_Item_t *E1AP_ID_Item = (UEID_GNB_CU_CP_E1AP_ID_Item_t *)calloc (1, sizeof(UEID_GNB_CU_CP_E1AP_ID_Item_t ));
         if(! E1AP_ID_Item)
         {
         	fprintf(stderr, "alloc UEID_GNB_CU_CP_E1AP_ID_Item failed\n");
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlHeader, controlHeaderIE);
                return -1;

         }
	 E1AP_ID_Item->gNB_CU_CP_UE_E1AP_ID = e1AP[0];
	 ASN_SEQUENCE_ADD(&controlHeader_Fmt1->ueID.choice.gNB_UEID->gNB_CU_CP_UE_E1AP_ID_List->list,E1AP_ID_Item);
	*/
        controlHeader_Fmt1->ric_Style_Type = ricControlStyleType;
        controlHeader_Fmt1->ric_ControlAction_ID = ricControlActionID;

        controlHeaderIE->ric_controlHeader_formats.choice.controlHeader_Format1 = controlHeader_Fmt1;

	fprintf(stderr, "Manju string %s\n",controlHeaderIE->ric_controlHeader_formats.choice.controlHeader_Format1->ueID.choice.gNB_UEID->amf_UE_NGAP_ID.buf) ;

	fprintf(stderr, "Manju string %s\n",controlHeaderIE->ric_controlHeader_formats.choice.controlHeader_Format1->ueID.choice.gNB_UEID->guami.pLMNIdentity.buf);


	fprintf(stderr, "Manju string %s\n",controlHeaderIE->ric_controlHeader_formats.choice.controlHeader_Format1->ueID.choice.gNB_UEID->guami.aMFRegionID.buf);

fprintf(stderr, "Manju string %s\n",controlHeaderIE->ric_controlHeader_formats.choice.controlHeader_Format1->ueID.choice.gNB_UEID->guami.aMFSetID.buf);

fprintf(stderr, "Manju string %s\n",controlHeaderIE->ric_controlHeader_formats.choice.controlHeader_Format1->ueID.choice.gNB_UEID->guami.aMFPointer.buf);
        
fprintf(stderr, "Manju string %lu\n",(**(controlHeaderIE->ric_controlHeader_formats.choice.controlHeader_Format1->ueID.choice.gNB_UEID->gNB_CU_CP_UE_E1AP_ID_List->list.array)).gNB_CU_CP_UE_E1AP_ID);

fprintf(stderr, "Manju string %lu\n",(**(controlHeaderIE->ric_controlHeader_formats.choice.controlHeader_Format1->ueID.choice.gNB_UEID->gNB_CU_UE_F1AP_ID_List->list.array)).gNB_CU_UE_F1AP_ID);


        fprintf(stderr, "showing xer of asn_DEF_E2SM_RC_ControlHeader data\n");
        xer_fprint(stderr, &asn_DEF_E2SM_RC_ControlHeader, controlHeaderIE);
        fprintf(stderr, "\n");
        fprintf(stderr, "After xer of asn_DEF_E2SM_RC_ControlHeader data\n");
   
        asn_enc_rval_t encode_result;
        encode_result = aper_encode_to_buffer(&asn_DEF_E2SM_RC_ControlHeader, NULL, controlHeaderIE, buffer, buf_size);


        ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlHeader, controlHeaderIE);
        if(encode_result.encoded == -1)
        {
                fprintf(stderr, "Cannot encode %s: %s\n", encode_result.failed_type->name, strerror(errno));
		return -1;
        }
        else
        {
               return encode_result.encoded;
        }
}

ssize_t e2sm_encode_nrcgi(NR_CGI_t *nr_cgi, void* plmnIdValue, size_t  plmnId_size,ulong lNRCellId, uint8_t* buffer, size_t buf_size)
{
        nr_cgi = (NR_CGI_t*)calloc(1, sizeof(NR_CGI_t));
        if(!nr_cgi) {
                fprintf(stderr, "e2src_encode_nrcgi alloc nr_cgi failed \n");
                return -1;
        }
        OCTET_STRING_fromBuf(&nr_cgi->pLMNIdentity,plmnIdValue, plmnId_size);
        //fprintf(stderr, "encodec Plmn Id = %s  plmnIdValue %s and lNRCellId = %lu \n", nr_cgi->pLMNIdentity, plmnIdValue,lNRCellId);
        //fprintf(stderr, "encodec Plmn Id = %s  and lNRCellId = %d \n", nr_cgi->pLMNIdentity,plmnIdValue,lNRCellId);
	BIT_STRING_t *nr_cell_id = (BIT_STRING_t*)calloc(1, sizeof(BIT_STRING_t));
           if(nr_cell_id)
           {
               nr_cell_id->buf = (uint8_t*)calloc(1,5);
               if(nr_cell_id->buf)
               {
                   nr_cell_id->size = 5;
                   nr_cell_id->buf[0] = ((lNRCellId & 0X0FF0000000) >> 28);
                   nr_cell_id->buf[1] = ((lNRCellId & 0X000FF00000) >> 20);
                   nr_cell_id->buf[2] = ((lNRCellId & 0X00000FF000) >> 12);
                   nr_cell_id->buf[3] = ((lNRCellId & 0X0000000FF0) >> 4);
                   nr_cell_id->buf[4] = (lNRCellId & 0X000000000F) << 4;
                   nr_cell_id->bits_unused = 4;

                   nr_cgi->nRCellIdentity = *nr_cell_id;
               }
           }

          fprintf(stderr, "showing xer of asn_DEF_NR_CGI NR_CGI_t data\n");
          xer_fprint(stderr, &asn_DEF_NR_CGI, nr_cgi);

	   asn_enc_rval_t encode_result = aper_encode_to_buffer(&asn_DEF_NR_CGI, NULL, nr_cgi, buffer, buf_size);

           if(encode_result.encoded != -1)
           {
               fprintf(stderr, "nr_cgi encodedi length = %zd \n", encode_result.encoded) ;
               return encode_result.encoded;
           }
           else
           {
               fprintf(stderr, "nr_cgi encode failed =%zd \n", encode_result.encoded) ;
               return -1 ;
           }
}

ssize_t e2sm_encode_ric_control_message(void *buffer, size_t buf_size, long targetPrimaryCell,
                        long targetCell, long nrOrEUtraCell, long nrCGIOrECGI, void* ranParameterValue, size_t  ranParameterValue_size)
{
	fprintf(stderr, "e2sm_encode_ric_control_message \n") ;
	//NR_CGI_t *nr_cgi = NULL;
	//uint8_t nrcgiBuf[buf_size];
	//ssize_t nrcgiBuf_size,nrcgiLen;


	E2SM_RC_ControlMessage_t *e2smrcRcControlMsg = (E2SM_RC_ControlMessage_t*)calloc(1, sizeof(E2SM_RC_ControlMessage_t));
        if(!e2smrcRcControlMsg) {
            fprintf(stderr, "alloc E2SM_ControlMessage_t failed\n");
        return -1;
        }
	 e2smrcRcControlMsg->ric_controlMessage_formats.present = E2SM_RC_ControlMessage__ric_controlMessage_formats_PR_controlMessage_Format1;

        E2SM_RC_ControlMessage_Format1_t *e2smrcRcControlFormat1 = (E2SM_RC_ControlMessage_Format1_t*)calloc(1, sizeof(E2SM_RC_ControlMessage_Format1_t));
        if(!e2smrcRcControlFormat1) {
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
            	fprintf(stderr, "alloc E2SM_ControlMessage_Format1_t failed\n");
        	return -1;
        }
	// Start Target Primary Cell
        E2SM_RC_ControlMessage_Format1_Item_t *ranParameterItem1 = (E2SM_RC_ControlMessage_Format1_Item_t *) calloc(1,sizeof(
                                                                                E2SM_RC_ControlMessage_Format1_Item_t));
        if(!ranParameterItem1) {
                fprintf(stderr, "alloc RANParameter_Item_t1 failed\n");
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
                return -1;
        }

        ranParameterItem1->ranParameter_ID = targetPrimaryCell;     // Target Primary Cell ID value = 1
	ranParameterItem1->ranParameter_valueType.present = RANParameter_ValueType_PR_ranP_Choice_Structure;

	RANParameter_ValueType_Choice_Structure_t *ranParameterStructure1 = (RANParameter_ValueType_Choice_Structure_t*)calloc(1, sizeof(RANParameter_ValueType_Choice_Structure_t));
        if(!ranParameterStructure1)
        {
                fprintf(stderr, "alloc RANParameter_STRUCTURE_t1 failed\n");
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
                return -1;
        }

	fprintf(stderr, "targetPrimaryCell encoded \n");

	 // Start Target Cell
        RANParameter_STRUCTURE_Item_t *ranParameterItem2 = (RANParameter_STRUCTURE_Item_t *)calloc(1,sizeof(RANParameter_STRUCTURE_Item_t));

        if(!ranParameterItem2)
        {

                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
                fprintf(stderr, "alloc RANParameter_Item_t2 failed\n");
                return -1;
        }

	ranParameterItem2->ranParameter_ID = targetCell;    // Target Cell ID value = 2

	RANParameter_ValueType_Choice_Structure_t *ranParameterStructure2 = (RANParameter_ValueType_Choice_Structure_t*)calloc(1, sizeof(struct RANParameter_ValueType_Choice_Structure));
        if(!ranParameterStructure2)
        {
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
                fprintf(stderr, "alloc RANParameter_STRUCTURE_t2 failed\n");
                return -1;
        }

	fprintf(stderr, "targetPrimaryCellId Value encoded \n");

        // Start NR Cell  (or E-UTRA Cell)
         RANParameter_STRUCTURE_Item_t *ranParameterItem3 = (RANParameter_STRUCTURE_Item_t*)calloc(1,sizeof(RANParameter_STRUCTURE_Item_t));
        if(!ranParameterItem3)
        {
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
                fprintf(stderr, "alloc RANParameter_Item_t3 failed\n");
                return -1;
        }

        ranParameterItem3->ranParameter_ID = nrOrEUtraCell; // NR Cell ID (or E-UTRA Cell ID) value =

	RANParameter_ValueType_Choice_Structure_t *ranParameterStructure3 = (struct RANParameter_ValueType_Choice_Structure*)calloc(1, sizeof(struct RANParameter_ValueType_Choice_Structure));
        if(!ranParameterStructure3)
        {
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
                fprintf(stderr, "alloc RANParameter_Item_t3 failed\n");
                return -1;
        }
	fprintf(stderr, " NR Cell ID  Value encoded \n");

	// Start NR CGI (or ECGI)
        RANParameter_STRUCTURE_Item_t *ranParameterItem4 = (RANParameter_STRUCTURE_Item_t *) calloc(1,sizeof(
                                                                        RANParameter_STRUCTURE_Item_t));
        if(!ranParameterItem4)
        {
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
                fprintf(stderr, "alloc RANParameter_Item_t4 failed\n");
                return -1;
        }

        ranParameterItem4->ranParameter_ID = nrCGIOrECGI;   // NR CGI ID (or ECGI ID) value =

	ranParameterItem4->ranParameter_valueType = (RANParameter_ValueType_t*)calloc(1,sizeof(RANParameter_ValueType_t));
	if(!ranParameterItem4->ranParameter_valueType)
	{
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
                fprintf(stderr, "alloc ranParameterItem4->ranParameter_valueType failed\n");
                return -1;
	}

	ranParameterItem4->ranParameter_valueType->present = RANParameter_ValueType_PR_ranP_Choice_ElementFalse;
        ranParameterItem4->ranParameter_valueType->choice.ranP_Choice_ElementFalse = (RANParameter_ValueType_Choice_ElementFalse_t *)calloc(1, sizeof(RANParameter_ValueType_Choice_ElementFalse_t));
        if(!ranParameterItem4->ranParameter_valueType->choice.ranP_Choice_ElementFalse)
        {
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
                fprintf(stderr, "alloc RANParameter_Item_t4 failed\n");
                return -1;
        }

	ranParameterItem4->ranParameter_valueType->choice.ranP_Choice_ElementFalse->ranParameter_value = (RANParameter_Value_t *)calloc(1, sizeof(RANParameter_Value_t));
	if(!ranParameterItem4->ranParameter_valueType->choice.ranP_Choice_ElementFalse->ranParameter_value)
	{
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
                fprintf(stderr, "alloc ranParameter_value failed\n");
                return -1;
	}
        ranParameterItem4->ranParameter_valueType->choice.ranP_Choice_ElementFalse->ranParameter_value->present = RANParameter_Value_PR_valueOctS;
	OCTET_STRING_fromBuf(&ranParameterItem4->ranParameter_valueType->choice.ranP_Choice_ElementFalse->ranParameter_value->choice.valueOctS,ranParameterValue, ranParameterValue_size);

	/*
	fprintf(stderr, "Target Plmn Id = %s ranParameterValue and lNRCellId = %lu \n", ranParameterValue,lNRCellId);
	nrcgiLen = e2sm_encode_nrcgi(nr_cgi, ranParameterValue, ranParameterValue_size, lNRCellId,nrcgiBuf,nrcgiBuf_size);
	if(nrcgiLen == -1 )
	{
		 fprintf(stderr,"e2srmc_encode_nrcgi failed \n");
		return nrcgiLen;
	}
	OCTET_STRING_fromBuf(&ranParameterItem4->ranParameter_valueType->choice.ranP_Choice_ElementFalse->ranParameter_value->choice.valueOctS,nrcgiBuf,nrcgiLen);
	*/
	fprintf(stderr, " NR CGI encoded \n");

	ranParameterStructure3->ranParameter_Structure = (RANParameter_STRUCTURE_t*)calloc(1,sizeof(RANParameter_STRUCTURE_t));
	if(!ranParameterStructure3->ranParameter_Structure)
	{
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
                fprintf(stderr, "alloc ranParameterStructure3->ranParameter_Structure  failed\n");
                return -1;
	}

	ranParameterStructure3->ranParameter_Structure->sequence_of_ranParameters = (struct RANParameter_STRUCTURE__sequence_of_ranParameters*)calloc(1,sizeof(struct RANParameter_STRUCTURE__sequence_of_ranParameters));
	if(!ranParameterStructure3->ranParameter_Structure->sequence_of_ranParameters)
	{
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
                fprintf(stderr, "alloc ranParameterStructure3->ranParameter_Structure->sequence_of_ranParameters failed\n");
                return -1;
	}



	ASN_SEQUENCE_ADD(&ranParameterStructure3->ranParameter_Structure->sequence_of_ranParameters->list, ranParameterItem4);   // NR CGI (or ECGI)
	ranParameterItem3->ranParameter_valueType = (RANParameter_ValueType_t*)calloc(1,sizeof(RANParameter_ValueType_t));
	if(!ranParameterItem3->ranParameter_valueType)
	{
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
                fprintf(stderr, "alloc ranParameterItem3->ranParameter_valueType failed\n");
                return -1;
	}
        ranParameterItem3->ranParameter_valueType->present = RANParameter_ValueType_PR_ranP_Choice_Structure;
        ranParameterItem3->ranParameter_valueType->choice.ranP_Choice_Structure = ranParameterStructure3; // NR Cell  (or E-UTRA Cell)
	fprintf(stderr, " ranParameterStructure3 encoded \n");

	ranParameterStructure2->ranParameter_Structure = (RANParameter_STRUCTURE_t*)calloc(1,sizeof(RANParameter_STRUCTURE_t));
	if(!ranParameterStructure2->ranParameter_Structure)
	{
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
                fprintf(stderr, "alloc ranParameterStructure2->ranParameter_Structure  failed\n");
                return -1;
	}

	ranParameterStructure2->ranParameter_Structure->sequence_of_ranParameters = (struct RANParameter_STRUCTURE__sequence_of_ranParameters*)calloc(1,sizeof(struct RANParameter_STRUCTURE__sequence_of_ranParameters));
	if(!ranParameterStructure2->ranParameter_Structure->sequence_of_ranParameters)
	{
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
                fprintf(stderr, "alloc ranParameterStructure3->ranParameter_Structure->sequence_of_ranParameters failed\n");
                return -1;
	}

        ASN_SEQUENCE_ADD(&ranParameterStructure2->ranParameter_Structure->sequence_of_ranParameters->list, ranParameterItem3);   // NR Cell  (or E-UTRA Cell)
	ranParameterItem2->ranParameter_valueType = (RANParameter_ValueType_t*)calloc(1,sizeof(RANParameter_ValueType_t));
	if(!ranParameterItem2->ranParameter_valueType)
	{
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
                fprintf(stderr, "alloc ranParameterItem2->ranParameter_valueType failed\n");
                return -1;
	}
        ranParameterItem2->ranParameter_valueType->present = RANParameter_ValueType_PR_ranP_Choice_Structure;
        ranParameterItem2->ranParameter_valueType->choice.ranP_Choice_Structure = ranParameterStructure2; // Target Cell
	fprintf(stderr, " ranParameterStructure2 encoded \n");

	ranParameterStructure1->ranParameter_Structure = (RANParameter_STRUCTURE_t*)calloc(1,sizeof(RANParameter_STRUCTURE_t));
	if(!ranParameterStructure1->ranParameter_Structure)
	{
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
                fprintf(stderr, "alloc ranParameterStructure1->ranParameter_Structure  failed\n");
                return -1;
	}

	ranParameterStructure1->ranParameter_Structure->sequence_of_ranParameters = (struct RANParameter_STRUCTURE__sequence_of_ranParameters*)calloc(1,sizeof(struct RANParameter_STRUCTURE__sequence_of_ranParameters));
	if(!ranParameterStructure1->ranParameter_Structure->sequence_of_ranParameters)
	{
                ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
                fprintf(stderr, "alloc ranParameterStructure1->ranParameter_Structure->sequence_of_ranParameters failed\n");
                return -1;
	}
	fprintf(stderr, " ranParameterStructure1 encoded \n");

        ASN_SEQUENCE_ADD(&ranParameterStructure1->ranParameter_Structure->sequence_of_ranParameters->list, ranParameterItem2);   // Target Cell
        ranParameterItem1->ranParameter_valueType.choice.ranP_Choice_Structure = ranParameterStructure1; // Target Primary Cell

        ASN_SEQUENCE_ADD(&e2smrcRcControlFormat1->ranP_List.list, ranParameterItem1); // Target Primary Cell
        e2smrcRcControlMsg->ric_controlMessage_formats.choice.controlMessage_Format1 = e2smrcRcControlFormat1;


        fprintf(stderr, "showing xer of asn_DEF_E2SM_RC_ControlMessage data\n");
        xer_fprint(stderr, &asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
        fprintf(stderr, "\n");
        fprintf(stderr, "After xer of asn_DEF_E2SM_RC_ControlMessage data\n");

        asn_enc_rval_t encode_result;
        encode_result = aper_encode_to_buffer(&asn_DEF_E2SM_RC_ControlMessage, NULL, e2smrcRcControlMsg, buffer, buf_size);
        ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlMessage, e2smrcRcControlMsg);
	fprintf(stderr, "encoded length = %ld \n", encode_result.encoded);
        if(encode_result.encoded == -1)
        {
                fprintf(stderr, "Cannot encode %s: %s\n", encode_result.failed_type->name, strerror(errno));
                return -1;
        }
        else
        {
                return encode_result.encoded;
        }
}


E2SM_RC_ControlOutcome_t* e2sm_decode_ric_call_process_outcome(void *buffer, size_t buf_size)
{
	fprintf(stderr, "e2sm_decode_ric_call_process_outcome Enter \n");
	asn_dec_rval_t decode_result;
	E2SM_RC_ControlOutcome_t* controlOutcome = NULL;
        decode_result = aper_decode_complete(NULL, &asn_DEF_E2SM_RC_ControlOutcome, (void **)&controlOutcome, buffer, buf_size);
        if(decode_result.code == RC_OK) {
        	xer_fprint(stdout, &asn_DEF_E2SM_RC_ControlOutcome, controlOutcome);
        	return controlOutcome;
    	}
        else {
        ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlOutcome, controlOutcome);
        return NULL;
        }
}

void e2sm_free_ric_call_process_outcome(E2SM_RC_ControlOutcome_t* controlOutcome) {
        ASN_STRUCT_FREE(asn_DEF_E2SM_RC_ControlOutcome, controlOutcome);
}
