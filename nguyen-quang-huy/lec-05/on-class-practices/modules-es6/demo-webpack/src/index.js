import { createHttp } from "./services/http";

const btn = document.getElementById("btnRequest");
const result = document.getElementById("result");
const http = createHttp();

// get
if (btn) {
  btn.addEventListener("click", () => {
    http
      .get("http://localhost:8000/students")
      .then((data) => {
        result.innerHTML = JSON.stringify(data);
      })
      .catch((e) => {
        result.innerHTML = e;
      });
  });
}

const addForm = document.getElementById("form");
addForm.addEventListener("submit", (event) => {
//   event.preventDefault;
  const student = {
    Name: addForm.name.value,
    Age: addForm.age.value,
  };
  console.log(student)
  http.post("http://localhost:8000/students", { body: JSON.stringify(student)})
  .then((data) =>{
      console.log(data)
      result.innerHTML = JSON.stringify(data)
  })
});
