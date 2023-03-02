package storeCSV

type storeCSV struct {
}

type Item struct {
	Ser_numb                      string `json:"ser_numb" `
	Id                            string `json:"id"`
	Uid                           string `json:"uid"`
	Domain                        string `json:"domain"`
	Cn                            string `json:"cn"`
	Department                    string `json:"department"`
	Title                         string `json:"title"`
	Who                           string `json:"who"`
	Logon_count                   string `json:"logon_count"`
	Num_logons7                   string `json:"num_logons7"`
	Num_share7                    string `json:"num_share7"`
	Num_file7                     string `json:"num_file7"`
	Num_ad7                       string `json:"num_ad7"`
	Num_n7                        string `json:"num_n7"`
	Num_logons14                  string `json:"num_logons14"`
	Num_share14                   string `json:"num_share14"`
	Num_file14                    string `json:"num_file14"`
	Num_ad14                      string `json:"num_ad14"`
	Num_n14                       string `json:"num_n14"`
	Num_logons30                  string `json:"num_logons30"`
	Num_share30                   string `json:"num_share30"`
	Num_file30                    string `json:"num_file30"`
	Num_ad30                      string `json:"num_ad30"`
	Num_n30                       string `json:"num_n30"`
	Num_logons150                 string `json:"num_logons150"`
	Num_share150                  string `json:"num_share150"`
	Num_file150                   string `json:"num_file150"`
	Num_ad150                     string `json:"num_ad150"`
	Num_n150                      string `json:"num_n150"`
	Num_logons365                 string `json:"num_logons365"`
	Num_share365                  string `json:"num_share365"`
	Num_file365                   string `json:"num_file365"`
	Num_ad365                     string `json:"num_ad365"`
	Num_n365                      string `json:"num_n365"`
	Has_user_principal_name       string `json:"has_user_principal_name"`
	Has_mail                      string `json:"has_mail"`
	Has_phone                     string `json:"has_phone"`
	Flag_disabled                 string `json:"flag_disabled"`
	Flag_lockout                  string `json:"flag_lockout"`
	Flag_password_not_required    string `json:"flag_password_not_required"`
	Flag_password_cant_change     string `json:"flag_password_cant_change"`
	Flag_dont_expire_password     string `json:"flag_dont_expire_password"`
	Owned_files                   string `json:"owned_files"`
	Num_mailboxes                 string `json:"num_mailboxes"`
	Num_member_of_groups          string `json:"num_member_of_groups"`
	Num_member_of_indirect_groups string `json:"num_member_of_indirect_groups"`
	Member_of_indirect_groups_ids string `json:"member_of_indirect_groups_ids"`
	Member_of_groups_ids          string `json:"member_of_groups_ids"`
	Is_admin                      string `json:"is_admin"`
	Is_service                    string `json:"is_service"`
}
