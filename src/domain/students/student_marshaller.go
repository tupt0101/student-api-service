package students

import "encoding/json"

type PublicStudent struct {
	StudentId string `json:"student_id"`
	Name      string `json:"name"`
	Class     string `json:"class"`
}

type PrivateStudent struct {
	StudentId string `json:"student_id"`
	Name      string `json:"name"`
	Class     string `json:"class"`
	CreatedOn string `json:"created_on"`
	Status    string `json:"status"`
}

func (students Students) Marshal(isPublic bool) []interface{} {
	result := make([]interface{}, len(students))
	for index, student := range students {
		result[index] = student.Marshal(isPublic)
	}
	return result
}

func (student *Student) Marshal(isPublic bool) interface{} {
	if isPublic {
		return PublicStudent{
			StudentId: student.StudentId,
			Name:      student.Name,
			Class:     student.Class,
		}
	}
	studentJson, _ := json.Marshal(student)
	var privateStudent PrivateStudent
	json.Unmarshal(studentJson, &privateStudent)
	return privateStudent
}
