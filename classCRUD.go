package entities
type class struct {
	profName          uint    `json:"profName"`
	CourseName        string  `json:"courseName"`
}
func Createclass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var class entities.class
	json.NewDecoder(r.Body).Decode(&class)
	database.Instance.Create(&class)
	json.NewEncoder(w).Encode(class)
}
func GetclassById(w http.ResponseWriter, r *http.Request) {
	classId := mux.Vars(r)["id"]
	if checkIfclassExists(classId) == false {
		json.NewEncoder(w).Encode("class Not Found!")
		return
	}
	var class entities.class
	database.Instance.First(&class, classId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(class)
}
func checkIfclassExists(classId string) bool {
	var class entities.class
	database.Instance.First(&class, classId)
	if class.ID == 0 {
		return false
	}
	return true
}
func Updateclass(w http.ResponseWriter, r *http.Request) {
	classId := mux.Vars(r)["id"]
	if checkIfclassExists(classId) == false {
		json.NewEncoder(w).Encode("class Not Found!")
		return
	}
	var class entities.class
	database.Instance.First(&class, classId)
	json.NewDecoder(r.Body).Decode(&class)
	database.Instance.Save(&class)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(class)
}