package entities
type student struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	mark       float64 `json:"mark"`
	class string  `json:"class"`
}
func Createstudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student entities.student
	json.NewDecoder(r.Body).Decode(&student)
	database.Instance.Create(&student)
	json.NewEncoder(w).Encode(student)
}
func GetStudentById(w http.ResponseWriter, r *http.Request) {
	studentId := mux.Vars(r)["id"]
	if checkIfstudentExists(studentId) == false {
		json.NewEncoder(w).Encode("student Not Found!")
		return
	}
	var student entities.student
	database.Instance.First(&student, studentId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}
func checkIfstudentExists(studentId string) bool {
	var student entities.student
	database.Instance.First(&student, studentId)
	if student.ID == 0 {
		return false
	}
	return true
}
func Updatestudent(w http.ResponseWriter, r *http.Request) {
	studentId := mux.Vars(r)["id"]
	if checkIfstudentExists(studentId) == false {
		json.NewEncoder(w).Encode("student Not Found!")
		return
	}
	var student entities.student
	database.Instance.First(&student, studentId)
	json.NewDecoder(r.Body).Decode(&student)
	database.Instance.Save(&student)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}