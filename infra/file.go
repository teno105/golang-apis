package infra

import "io/ioutil"

// JSON 파일을 읽어 구조체로 변환하는 함수
func loadJSONFile(filePath string, v interface{}) error {
	// 파일 읽기
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// JSON을 구조체로 변환
	return json.Unmarshal(data, v)
}
