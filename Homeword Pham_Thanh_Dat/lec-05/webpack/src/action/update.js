import { createHttpJson, createHttpText } from "../services/http";
import { display } from "./display";
const Json = createHttpJson();
const text = createHttpText();
function updateStudent(studentInformation){
    Json.put("http://localhost:8000/students", {
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify(studentInformation),
    })
      .then((data) => {
        display(data);
      })
      .catch((e) => {
        result.innerHTML = e;
      });

}
export{
    updateStudent
}