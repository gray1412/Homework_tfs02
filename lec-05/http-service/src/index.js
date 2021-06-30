import "../assets/style/main.css";
import { METHODS, createHttpInstance } from "./service/http.js";

//lay du lieu ra man hinh
const httpGet = createHttpInstance({
  // Init http instance
  endpoint: "http://localhost:8082/api",
  methods: METHODS.Get,
});
const btnRequest = document.getElementById("btnRequest");
const result = document.getElementById("myTable");

btnRequest.addEventListener("click", () => {
  // Reset result
  result.innerHTML = "Loading...";
  // Request
  httpGet
    .request("/students")
    .then((response) => {
      document.querySelector("#myTable").innerHTML = ``;
      let conten = `
        <thead>
        <tr>
            <th>Id</th>
            <th>Name</th>
            <th>Age</th>
            <th></th>
            <th></th>
        </tr>
    </thead>
        `;
      JSON.parse(response).forEach((element) => {
        conten += `
          <tr>
            <td>${element.Id}</td>
            <td>${element.Name}</td>
            <td>${element.Age}</td>
            <td><button id="btnUpdate" onclick="updateData(${element.Id})">Update</button></td>
            <td><button id="btnDelete" onclick="deleteData(${element.Id})">Delete</button></td>

        </tr>
          `;
      });
      document.querySelector("#myTable").innerHTML = conten;
    })
    .catch((e) => {
      result.innerHTML = `${JSON.stringify(e)}`;
    });
});

//tao du lieu
// const btnCreate = document.getElementById("btnCreate");
// const httpPost = createHttpInstance({
//   endpoint: "http://localhost:8082/api/students",
//   methods: METHODS.Post,
// });

// btnCreate.addEventListener("click", () => {
//   const name = document.getElementById("name").value;
//   const age = document.getElementById("age").value;
//   const data = {
//     Name: name,
//     Age: parseInt(age),
//   };
//   httpPost
//     .post("http://localhost:8082/api/students", {
//       body: JSON.stringify(data),
//     })
//     .then((response) => {
//       alert("Create successful");
//     })
//     .catch((e) => {
//       result.innerHTML = e;
//     });
// });

//delete du lieu
// const btnDelete = document.querySelector("#btnDelete");
// const httpDelete = createHttpInstance({
//   endpoint: "http://localhost:8082/api",
//   methods: METHODS.Delete,
// });
// btnDelete.addEventListener("click", () => {
//   httpDelete.request("/students/" + Element.Id).then(() => {
//     alert("Delete successful");
//   });
// });

function updateSize() {
  let nBytes = 0,
    oFiles = this.files,
    nFiles = oFiles.length;
  for (let nFileId = 0; nFileId < nFiles; nFileId++) {
    nBytes += oFiles[nFileId].size;
  }
  let sOutput = nBytes + " bytes";
  // optional code for multiples approximation
  const aMultiples = ["KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"];
  for (
    nMultiple = 0, nApprox = nBytes / 1024;
    nApprox > 1;
    nApprox /= 1024, nMultiple++
  ) {
    sOutput =
      nApprox.toFixed(3) +
      " " +
      aMultiples[nMultiple] +
      " (" +
      nBytes +
      " bytes)";
  }
  // end of optional code
  document.getElementById("fileNum").innerHTML = nFiles;
  document.getElementById("fileSize").innerHTML = sOutput;
}

document
  .getElementById("uploadInput")
  .addEventListener("change", updateSize, false);
