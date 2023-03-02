package storeCSV

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"os"
)

func New() *storeCSV {
	return &storeCSV{}
}

func (s *storeCSV) Get(key string, value []string) ([]byte, error) {
	var result []byte
	columns := make(map[string]int)
	vel := make(map[string]struct{})

	file, err := os.Open("internal/store/storeCSV/ueba.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	titles, err := r.Read()
	if err == io.EOF {
		return result, errors.New("ERROR: file is empty")
	}
	if err != nil {
		return result, err
	}

	var items []Item

	for i, v := range titles {
		columns[v] = i
	}
	_, ok := columns[key]
	if !ok {
		return result, errors.New("ERROR: no such column")
	}
	for _, v := range value {
		vel[v] = struct{}{}
	}
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return result, err
		}
		a := columns[key]
		b := row[a]
		_, ok = vel[b]
		if ok {
			items = append(items, Item{
				Ser_numb:                      row[columns["#"]],
				Id:                            row[columns["id"]],
				Uid:                           row[columns["uid"]],
				Domain:                        row[columns["domain"]],
				Cn:                            row[columns["cn"]],
				Department:                    row[columns["department"]],
				Title:                         row[columns["title"]],
				Who:                           row[columns["who"]],
				Logon_count:                   row[columns["logon_count"]],
				Num_logons7:                   row[columns["num_logons7"]],
				Num_share7:                    row[columns["num_share7"]],
				Num_file7:                     row[columns["num_file7"]],
				Num_ad7:                       row[columns["num_ad7"]],
				Num_n7:                        row[columns["num_n7"]],
				Num_logons14:                  row[columns["num_logons14"]],
				Num_share14:                   row[columns["num_share14"]],
				Num_file14:                    row[columns["num_file14"]],
				Num_ad14:                      row[columns["num_ad14"]],
				Num_n14:                       row[columns["num_n14"]],
				Num_logons30:                  row[columns["num_logons30"]],
				Num_share30:                   row[columns["num_share30"]],
				Num_file30:                    row[columns["num_file30"]],
				Num_ad30:                      row[columns["num_ad30"]],
				Num_n30:                       row[columns["num_n30"]],
				Num_logons150:                 row[columns["num_logons150"]],
				Num_share150:                  row[columns["num_share150"]],
				Num_file150:                   row[columns["num_file150"]],
				Num_ad150:                     row[columns["num_ad150"]],
				Num_n150:                      row[columns["num_n150"]],
				Num_logons365:                 row[columns["num_logons365"]],
				Num_share365:                  row[columns["num_share365"]],
				Num_file365:                   row[columns["num_file365"]],
				Num_ad365:                     row[columns["num_ad365"]],
				Num_n365:                      row[columns["num_n365"]],
				Has_user_principal_name:       row[columns["has_user_principal_name"]],
				Has_mail:                      row[columns["has_mail"]],
				Has_phone:                     row[columns["has_phone"]],
				Flag_disabled:                 row[columns["flag_disabled"]],
				Flag_lockout:                  row[columns["flag_lockout"]],
				Flag_password_not_required:    row[columns["flag_password_not_required"]],
				Flag_password_cant_change:     row[columns["flag_password_cant_change"]],
				Flag_dont_expire_password:     row[columns["flag_dont_expire_password"]],
				Owned_files:                   row[columns["owned_files"]],
				Num_mailboxes:                 row[columns["num_mailboxes"]],
				Num_member_of_groups:          row[columns["num_member_of_groups"]],
				Num_member_of_indirect_groups: row[columns["num_member_of_indirect_groups"]],
				Member_of_indirect_groups_ids: row[columns["member_of_indirect_groups_ids"]],
				Member_of_groups_ids:          row[columns["member_of_groups_ids"]],
				Is_admin:                      row[columns["is_admin"]],
				Is_service:                    row[columns["is_service"]],
			})
		}
	}
	result, err = json.Marshal(items)
	if err != nil {
		return result, err
	}
	return result, nil
}
