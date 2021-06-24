import { createHttpJson, createHttpText } from "./services/http";
import { viewStudent } from "./action/view";
import { deleteStudent } from "./action/delete";
import { addStudent } from "./action/add";
import { updateStudent } from "./action/update";
const btn = document.getElementById("btnRequest");
// const result = document.getElementById("result");
// const Json = createHttpJson();
// const text = createHttpText();
if (btn) {
  btn.addEventListener("click", () => {
    viewStudent();
  });
}
const add = document.getElementById("add");
if (add) {
  add.addEventListener("click", () => {
    const name = document.getElementById("name").value;
    const address = document.getElementById("address").value;
    const phone = document.getElementById("phone").value;
    const age = document.getElementById("age").value;
    const studentInformation = { Name: name , Address: address, Phone:parseInt(phone), Age:parseInt(age)};
    addStudent(studentInformation)
  });
}
const btnDelete = document.getElementById("btnDelete");
if (btnDelete) {
  btnDelete.addEventListener("click", () => {
    const id = document.getElementById("delete").value;
    deleteStudent(id)
  });
}

const Update = document.getElementById("Update");
if (Update) {
  Update.addEventListener("click", () => {
    const id = document.getElementById("idUpdate").value;
    const name = document.getElementById("nameUpdate").value;
    const address = document.getElementById("addressUpdate").value;
    const phone = document.getElementById("phoneUpdate").value;
    const age = document.getElementById("ageUpdate").value;
    const studentInformation = {Id:parseInt(id), Name: name , Address: address, Phone:parseInt(phone), Age:parseInt(age)};
    updateStudent(studentInformation);
  });
}

// const formData = new FormData();
// const fileField = document.querySelector('input[type="file"]');
// formData.append('avatar', fileField.files[0]);
// const upload = document.getElementById("upload");
// if (upload) {
//   upload.addEventListener("click", () => {
//     const selectedFile = document.getElementById('file').files[0];
//     const file = { File: selectedFile};
//     text.post("http://localhost:8000/upload", {
//       headers: {'Content-Type': 'multipart/form-data'},
//       body: formData
//     })
//       .then((data) => {
//         result.innerHTML = JSON.stringify(data);
//       })
//       .catch((e) => {
//         result.innerHTML = e;
//       });
//   });
// }