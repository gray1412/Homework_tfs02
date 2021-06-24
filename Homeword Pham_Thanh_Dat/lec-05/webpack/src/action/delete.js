import { createHttpJson, createHttpText } from "../services/http";
import { display } from "./display";
const Json = createHttpJson();
const text = createHttpText();
function deleteStudent(id) {
  const data = { Id: parseInt(id) };
  Json.delete("http://localhost:8000/students", {
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  })
    .then((data) => {
      display(data);
    })
    .catch((e) => {
      result.innerHTML = e;
    });
}
export { deleteStudent };
