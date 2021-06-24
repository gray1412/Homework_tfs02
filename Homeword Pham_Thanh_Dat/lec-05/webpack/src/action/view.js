import { createHttpJson, createHttpText } from "../services/http";
import { display } from "./display";
const Json = createHttpJson();
const text = createHttpText();
function viewStudent() {
  Json.get("http://localhost:8000/students", {
    headers: { "Content-Type": "application/json" },
  })
    .then((data) => {
      display(data);
    })
    .catch((e) => {
      result.innerHTML = e;
    });
}
export { viewStudent };
