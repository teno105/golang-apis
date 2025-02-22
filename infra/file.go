package infra

import "io/ioutil"

func loadGameData(gameID string) (*Root, error) {
	// 파일 경로 구성
	filePath := fmt.Sprintf("%s.json", gameID)

	// 파일 읽기
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// JSON을 구조체로 변환
	var result Root
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
