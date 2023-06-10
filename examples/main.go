package main

import (
	"context"
	"encoding/json"

	"github.com/oarkflow/asynq"
)

const redisAddrWorker = "127.0.0.1:6379"

func main() {
	send(asynq.Sync)
	// sendA(asynq.Sync)
}

var data = map[string]any{
	"patient_header": map[string]any{
		"injury_date": "2023-05-23",
		"admitted":    false,
		"ama":         false,
		"lwbs":        false,
		"inpatient":   false,
		"observation": false,
		"emergency":   false,
		"discharge":   false,
		"outpatient":  false,
		"transfer":    false,
		"disch_disp":  "AMA",
	},
	"coding": []map[string]any{
		{
			"dos": "2023-05-23",
			"details": map[string]any{
				"pro": map[string]any{
					"em": map[string]any{
						"em_level":                      "lorem",
						"em_modifier1":                  "lorem",
						"em_modifier2":                  "lorem",
						"em_provider":                   "lorem",
						"em_midlevel":                   "lorem",
						"em_type":                       "lorem",
						"em_downcode":                   "true",
						"em_reason":                     "lorem",
						"mdm_total":                     "lorem",
						"qa_disagree":                   "true",
						"qa_note":                       "lorem",
						"mid":                           "lorem",
						"dc_note":                       "lorem",
						"shared":                        "true",
						"charge_code_select":            "lorem",
						"qa_type":                       "lorem",
						"qa_disagree_em_modifier1":      "true",
						"qa_note_em_modifier1":          "lorem",
						"qa_disagree_em_modifier2":      "true",
						"qa_note_em_modifier2":          "lorem",
						"qa_disagree_physician":         "true",
						"qa_note_physician":             "lorem",
						"qa_disagree_shared_visit":      "true",
						"qa_note_shared_visit":          "lorem",
						"qa_disagree_mid_provider":      "true",
						"qa_note_mid_provider":          "lorem",
						"qa_disagree_downcode":          "true",
						"qa_note_downcode":              "lorem",
						"qa_type_em_modifier1":          "lorem",
						"qa_type_em_modifier2":          "lorem",
						"qa_type_physician":             "lorem",
						"qa_type_shared_visit":          "lorem",
						"qa_type_mid_provider":          "lorem",
						"qa_type_downcode":              "lorem",
						"billing_provider":              "lorem",
						"secondary_provider":            "lorem",
						"resident_provider":             "lorem",
						"qa_disagree_resident_provider": "true",
						"qa_note_resident_provider":     "lorem",
						"qa_type_resident_provider":     "lorem",
						"event_dos":                     "2023-05-23T00:00:00",
						"patient_status_id":             1,
						"patient_status_disagree":       "true",
						"qa_note_patient_status":        "lorem",
						"procedure_modifier3":           "lorem",
						"qa_disagree_em_modifier3":      "true",
						"qa_type_em_modifier3":          "lorem",
						"qa_note_em_modifier3":          "lorem",
						"procedure_modifier4":           "lorem",
						"qa_disagree_em_modifier4":      "true",
						"qa_type_em_modifier4":          "lorem",
						"qa_note_em_modifier4":          "lorem",
					},
					"downcode": []map[string]any{
						{
							"emp_id":               1,
							"em_downcode":          "true",
							"em_reason":            "downcode reason",
							"mdm_total":            "lorem",
							"dc_note":              "lorem",
							"qa_disagree_downcode": "true",
							"qa_note_downcode":     "lorem",
							"qa_type_downcode":     "lorem",
							"event_dos":            "2023-05-23T00:00:00",
						},
						{
							"emp_id":               1,
							"em_downcode":          "true",
							"em_reason":            "some other thing",
							"mdm_total":            "lorem",
							"dc_note":              "lorem",
							"qa_disagree_downcode": "true",
							"qa_note_downcode":     "lorem",
							"qa_type_downcode":     "lorem",
							"event_dos":            "2023-05-23T00:00:00",
						},
					},
					"cpt": []map[string]any{
						{
							"procedure_num":                 "lorem",
							"procedure_qty":                 1,
							"procedure_modifier1":           "lorem",
							"procedure_modifier2":           "lorem",
							"procedure_provider":            "lorem",
							"mid_provider":                  "lorem",
							"qa_disagree":                   "true",
							"qa_note":                       "lorem",
							"procedure_date":                "2023-05-23T00:00:00",
							"charge_code_select":            "lorem",
							"specialist":                    "true",
							"qa_type":                       "lorem",
							"qa_disagree_em_modifier1":      "true",
							"qa_note_em_modifier1":          "lorem",
							"qa_disagree_em_modifier2":      "true",
							"qa_note_em_modifier2":          "lorem",
							"qa_disagree_quantity":          "true",
							"qa_note_quantity":              "lorem",
							"qa_disagree_date_performed":    "true",
							"qa_note_date_performed":        "lorem",
							"qa_disagree_physician":         "true",
							"qa_note_physician":             "lorem",
							"qa_disagree_mid_provider":      "true",
							"qa_note_mid_provider":          "lorem",
							"qa_type_em_modifier1":          "lorem",
							"qa_type_em_modifier2":          "lorem",
							"qa_type_quantity":              "lorem",
							"qa_type_date_performed":        "lorem",
							"qa_type_physician":             "lorem",
							"qa_type_mid_provider":          "lorem",
							"billing_provider":              "lorem",
							"secondary_provider":            "lorem",
							"resident_provider":             "lorem",
							"qa_disagree_resident_provider": "true",
							"qa_note_resident_provider":     "lorem",
							"qa_type_resident_provider":     "lorem",
							"event_dos":                     "2023-05-23T00:00:00",
							"patient_status_id":             1,
							"patient_status_disagree":       "true",
							"qa_note_patient_status":        "lorem",
							"procedure_modifier3":           "lorem",
							"qa_disagree_em_modifier3":      "true",
							"qa_type_em_modifier3":          "lorem",
							"qa_note_em_modifier3":          "lorem",
							"procedure_modifier4":           "lorem",
							"qa_disagree_em_modifier4":      "true",
							"qa_type_em_modifier4":          "lorem",
							"qa_note_em_modifier4":          "lorem",
						},
					},
					"special": []map[string]any{
						{
							"procedure_num":      "lorem",
							"procedure_qty":      1,
							"procedure_provider": "lorem",
							"qa_disagree":        "true",
							"qa_note":            "lorem",
							"procedure_date":     "2023-05-23T00:00:00",
							"qa_type":            "lorem",
							"event_dos":          "2023-05-23T00:00:00",
						},
					},
					"hcpcs": []map[string]any{
						{
							"hcpcs_num":               "lorem",
							"hcpcs_qty":               1,
							"hcpcs_modifier1":         "lorem",
							"hcpcs_modifier2":         "lorem",
							"hcpcs_provider":          "lorem",
							"qa_disagree":             "true",
							"qa_note":                 "lorem",
							"qa_type":                 "lorem",
							"event_dos":               "2023-05-23T00:00:00",
							"patient_status_id":       1,
							"patient_status_disagree": "true",
							"qa_note_patient_status":  "lorem",
							"hcpcs_modifier3":         "lorem",
							"hcpcs_modifier4":         "lorem",
							"charge_code_select":      "lorem",
						},
					},
				},
				"fac": map[string]any{
					"em": map[string]any{
						"em_level":                      "lorem",
						"em_modifier1":                  "lorem",
						"em_modifier2":                  "lorem",
						"em_provider":                   "lorem",
						"qa_disagree":                   "true",
						"qa_note":                       "lorem",
						"mid":                           "lorem",
						"charge_code_select":            "lorem",
						"qa_type":                       "lorem",
						"qa_disagree_em_modifier1":      "true",
						"qa_note_em_modifier1":          "lorem",
						"qa_disagree_em_modifier2":      "true",
						"qa_note_em_modifier2":          "lorem",
						"qa_disagree_physician":         "true",
						"qa_note_physician":             "lorem",
						"qa_disagree_mid_provider":      "true",
						"qa_note_mid_provider":          "lorem",
						"qa_type_em_modifier1":          "lorem",
						"qa_type_em_modifier2":          "lorem",
						"qa_type_physician":             "lorem",
						"qa_type_mid_provider":          "lorem",
						"billing_provider":              "lorem",
						"secondary_provider":            "lorem",
						"resident_provider":             "lorem",
						"qa_disagree_resident_provider": "true",
						"qa_note_resident_provider":     "lorem",
						"qa_type_resident_provider":     "lorem",
						"event_dos":                     "2023-05-23T00:00:00",
						"patient_status_id":             1,
						"patient_status_disagree":       "true",
						"qa_note_patient_status":        "lorem",
						"procedure_modifier3":           "lorem",
						"qa_disagree_em_modifier3":      "true",
						"qa_type_em_modifier3":          "lorem",
						"qa_note_em_modifier3":          "lorem",
						"procedure_modifier4":           "lorem",
						"qa_disagree_em_modifier4":      "true",
						"qa_type_em_modifier4":          "lorem",
						"qa_note_em_modifier4":          "lorem",
					},
					"cpt": []map[string]any{
						{
							"procedure_num":                 "lorem",
							"procedure_qty":                 1,
							"procedure_modifier1":           "lorem",
							"procedure_modifier2":           "lorem",
							"procedure_dx":                  "dx.lorem",
							"procedure_provider":            "lorem",
							"mid_provider":                  "lorem",
							"qa_disagree":                   "true",
							"qa_note":                       "lorem",
							"procedure_date":                "2023-05-23T00:00:00",
							"icd10_pcs":                     "icd10_pcs",
							"charge_code_select":            "lorem",
							"qa_type":                       "lorem",
							"qa_disagree_em_modifier1":      "true",
							"qa_note_em_modifier1":          "lorem",
							"qa_disagree_em_modifier2":      "true",
							"qa_note_em_modifier2":          "lorem",
							"qa_disagree_quantity":          "true",
							"qa_note_quantity":              "lorem",
							"qa_disagree_date_performed":    "true",
							"qa_note_date_performed":        "lorem",
							"qa_disagree_physician":         "true",
							"qa_note_physician":             "lorem",
							"qa_disagree_mid_provider":      "true",
							"qa_note_mid_provider":          "lorem",
							"qa_type_em_modifier1":          "lorem",
							"qa_type_em_modifier2":          "lorem",
							"qa_type_quantity":              "lorem",
							"qa_type_date_performed":        "lorem",
							"qa_type_physician":             "lorem",
							"qa_type_mid_provider":          "lorem",
							"qa_disagree_procedure_dx":      "true",
							"qa_note_procedure_dx":          "lorem",
							"qa_type_procedure_dx":          "lorem",
							"qa_disagree_icd10_pcs":         "true",
							"qa_note_icd10_pcs":             "lorem",
							"qa_type_icd10_pcs":             "lorem",
							"billing_provider":              "lorem",
							"secondary_provider":            "lorem",
							"resident_provider":             "lorem",
							"qa_disagree_resident_provider": "true",
							"qa_note_resident_provider":     "lorem",
							"qa_type_resident_provider":     "lorem",
							"event_dos":                     "2023-05-23T00:00:00",
							"patient_status_id":             1,
							"patient_status_disagree":       "true",
							"qa_note_patient_status":        "lorem",
							"procedure_modifier3":           "lorem",
							"qa_disagree_em_modifier3":      "true",
							"qa_type_em_modifier3":          "lorem",
							"qa_note_em_modifier3":          "lorem",
							"procedure_modifier4":           "lorem",
							"qa_disagree_em_modifier4":      "true",
							"qa_type_em_modifier4":          "lorem",
							"qa_note_em_modifier4":          "lorem",
						},
					},
					"special": []map[string]any{
						{
							"procedure_num":      "lorem",
							"procedure_qty":      1,
							"procedure_provider": "lorem",
							"qa_disagree":        "true",
							"qa_note":            "lorem",
							"procedure_date":     "2023-05-23T00:00:00",
							"qa_type":            "lorem",
							"event_dos":          "2023-05-23T00:00:00",
						},
					},
					"hcpcs": []map[string]any{
						{
							"hcpcs_num":               "lorem",
							"hcpcs_qty":               1,
							"hcpcs_modifier1":         "lorem",
							"hcpcs_modifier2":         "lorem",
							"hcpcs_provider":          "lorem",
							"qa_disagree":             "true",
							"qa_note":                 "lorem",
							"qa_type":                 "lorem",
							"event_dos":               "2023-05-23T00:00:00",
							"patient_status_id":       1,
							"patient_status_disagree": "true",
							"qa_note_patient_status":  "lorem",
							"hcpcs_modifier3":         "lorem",
							"hcpcs_modifier4":         "lorem",
							"charge_code_select":      "lorem",
						},
					},
				},
				"dx": map[string]any{
					"pro": []map[string]any{
						{
							"code":              "lorem",
							"dx_reason":         "lorem",
							"qa_disagree":       "true",
							"qa_note":           "lorem",
							"qa_type":           "lorem",
							"non_specific_code": "true",
							"event_dos":         "2023-05-23T00:00:00",
						},
					},
					"fac": []map[string]any{
						{
							"code":              "lorem",
							"dx_reason":         "lorem",
							"qa_disagree":       "true",
							"qa_note":           "lorem",
							"qa_type":           "lorem",
							"non_specific_code": "true",
							"event_dos":         "2023-05-23T00:00:00",
						},
					},
				},
				"cdi": map[string]any{
					"pro": []map[string]any{
						{
							"enc_cdi_reason": "lorem",
							"cdi_note":       "lorem",
							"qty":            1,
							"mid_level":      "lorem",
							"qa_disagree":    "true",
							"qa_note":        "lorem",
							"qa_type":        "lorem",
							"event_dos":      "2023-05-23T00:00:00",
						},
					},
					"fac": []map[string]any{
						{
							"enc_cdi_reason": "lorem",
							"cdi_note":       "lorem",
							"qty":            1,
							"mid_level":      "lorem",
							"qa_disagree":    "true",
							"qa_note":        "lorem",
							"qa_type":        "lorem",
							"event_dos":      "2023-05-23T00:00:00",
						},
					},
				},
				"notes": []map[string]any{
					{
						"coder_notes":  "from coder notes",
						"client_notes": "client notes here",
						"event_dos":    "2023-05-23T00:00:00",
					},
				},
			},
		},
	},
}

func send(mode asynq.Mode) {
	f := asynq.NewFlow(asynq.Config{Mode: mode, RedisServer: redisAddrWorker})
	f.FirstNode = "1"
	f.
		AddHandler("1", NewDataBranchHandler("1", asynq.Payload{
			Data: map[string]any{
				"data_branch": map[string]any{
					"patient_header": "2",
					"coding":         "3",
				},
			},
		})).
		AddHandler("4", NewDataBranchHandler("4", asynq.Payload{
			Data: map[string]any{
				"data_branch": map[string]any{
					"details.pro.downcode": "5",
					"details.pro.em":       "6",
					"details.pro.cpt":      "7",
					"details.pro.special":  "8",
					"details.pro.hcpcs":    "9",
					"details.fac.em":       "10",
					"details.fac.cpt":      "11",
					"details.fac.special":  "12",
					"details.fac.hcpcs":    "13",
					"details.dx.pro":       "14",
					"details.dx.fac":       "15",
					"details.cdi.pro":      "16",
					"details.cdi.fac":      "17",
					"details.notes":        "18",
				},
			},
		})).
		AddHandler("2", NewStoreData("2")).
		AddHandler("5", NewStoreData("5")).
		AddHandler("6", NewStoreData("6")).
		AddHandler("7", NewStoreData("7")).
		AddHandler("8", NewStoreData("8")).
		AddHandler("9", NewStoreData("9")).
		AddHandler("10", NewStoreData("10")).
		AddHandler("11", NewStoreData("11")).
		AddHandler("12", NewStoreData("12")).
		AddHandler("13", NewStoreData("13")).
		AddHandler("14", NewStoreData("14")).
		AddHandler("15", NewStoreData("15")).
		AddHandler("16", NewStoreData("16")).
		AddHandler("17", NewStoreData("17")).
		AddHandler("18", NewStoreData("18")).
		AddHandler("3", NewLoop("3")).
		AddLoop("3", "4")
	bt, _ := json.Marshal(data)
	ctx := context.WithValue(context.Background(), "extra_params", map[string]any{
		"encounter_id": 1,
		"work_item_id": 2,
	})
	f.Send(ctx, bt)
}
