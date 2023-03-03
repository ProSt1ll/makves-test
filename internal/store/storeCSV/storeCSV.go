package storeCSV

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"os"
)

func New(path string) (*StoreCSV, error) {
	columns := make(map[string]int) //колонки нашего файла

	//открываем файл
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//читаем первую строку с названиями колонок
	titles, err := csv.NewReader(file).Read()
	if err == io.EOF {
		return nil, errors.New("ERROR: file is empty")
	}
	if err != nil {
		return nil, err
	}

	//пишем в мапу
	for i, v := range titles {
		columns[v] = i
	}

	return &StoreCSV{
			path:    path,
			columns: columns,
		},
		nil
}

func (s *StoreCSV) Get(key string, value []string) ([]byte, error) {
	//открываем файл
	file, err := os.Open(s.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//проверяем наличие искомого ключа
	_, ok := s.columns[key]
	if !ok {
		return nil, errors.New("ERROR: no such column")
	}

	vel := make(map[string]struct{}) //наши значения в мапе, чтобы быстро доставать
	for _, v := range value {
		vel[v] = struct{}{}
	}

	//читаем файл по строчно и ищем необоходимый элемент
	r := csv.NewReader(file)
	var items []Item //срез итемов удволетворяющих поиску
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		//s.columns[key] возвращает номер столбца по искомому ключу
		//row[s.columns[key]] - возвращаем значение столбца по искомому ключу
		//vel[row[s.columns[key]]] - пытаемся достать значение столбца из искомых параметров
		_, ok = vel[row[s.columns[key]]]
		//если получается, формируем Item и добавляем в срез
		if ok {
			items = append(items, Item{
				Ser_numb:                      row[s.columns["#"]],
				Id:                            row[s.columns["id"]],
				Uid:                           row[s.columns["uid"]],
				Domain:                        row[s.columns["domain"]],
				Cn:                            row[s.columns["cn"]],
				Department:                    row[s.columns["department"]],
				Title:                         row[s.columns["title"]],
				Who:                           row[s.columns["who"]],
				Logon_count:                   row[s.columns["logon_count"]],
				Num_logons7:                   row[s.columns["num_logons7"]],
				Num_share7:                    row[s.columns["num_share7"]],
				Num_file7:                     row[s.columns["num_file7"]],
				Num_ad7:                       row[s.columns["num_ad7"]],
				Num_n7:                        row[s.columns["num_n7"]],
				Num_logons14:                  row[s.columns["num_logons14"]],
				Num_share14:                   row[s.columns["num_share14"]],
				Num_file14:                    row[s.columns["num_file14"]],
				Num_ad14:                      row[s.columns["num_ad14"]],
				Num_n14:                       row[s.columns["num_n14"]],
				Num_logons30:                  row[s.columns["num_logons30"]],
				Num_share30:                   row[s.columns["num_share30"]],
				Num_file30:                    row[s.columns["num_file30"]],
				Num_ad30:                      row[s.columns["num_ad30"]],
				Num_n30:                       row[s.columns["num_n30"]],
				Num_logons150:                 row[s.columns["num_logons150"]],
				Num_share150:                  row[s.columns["num_share150"]],
				Num_file150:                   row[s.columns["num_file150"]],
				Num_ad150:                     row[s.columns["num_ad150"]],
				Num_n150:                      row[s.columns["num_n150"]],
				Num_logons365:                 row[s.columns["num_logons365"]],
				Num_share365:                  row[s.columns["num_share365"]],
				Num_file365:                   row[s.columns["num_file365"]],
				Num_ad365:                     row[s.columns["num_ad365"]],
				Num_n365:                      row[s.columns["num_n365"]],
				Has_user_principal_name:       row[s.columns["has_user_principal_name"]],
				Has_mail:                      row[s.columns["has_mail"]],
				Has_phone:                     row[s.columns["has_phone"]],
				Flag_disabled:                 row[s.columns["flag_disabled"]],
				Flag_lockout:                  row[s.columns["flag_lockout"]],
				Flag_password_not_required:    row[s.columns["flag_password_not_required"]],
				Flag_password_cant_change:     row[s.columns["flag_password_cant_change"]],
				Flag_dont_expire_password:     row[s.columns["flag_dont_expire_password"]],
				Owned_files:                   row[s.columns["owned_files"]],
				Num_mailboxes:                 row[s.columns["num_mailboxes"]],
				Num_member_of_groups:          row[s.columns["num_member_of_groups"]],
				Num_member_of_indirect_groups: row[s.columns["num_member_of_indirect_groups"]],
				Member_of_indirect_groups_ids: row[s.columns["member_of_indirect_groups_ids"]],
				Member_of_groups_ids:          row[s.columns["member_of_groups_ids"]],
				Is_admin:                      row[s.columns["is_admin"]],
				Is_service:                    row[s.columns["is_service"]],
			})
		}
	}

	if len(items) == 0 {
		return nil, nil
	}

	//переводим результат в json
	result, err := json.Marshal(items)
	if err != nil {
		return result, err
	}
	return result, nil
}
