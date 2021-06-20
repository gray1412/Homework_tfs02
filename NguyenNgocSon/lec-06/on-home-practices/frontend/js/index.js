// funcion get data from API
const fetchData = async () => {
  const response = await fetch(`http://localhost:8000/api/members`);
  if (response) {
    const data = await response.json();
    showData(data);
  }
};

const showData = (data) => {
  console.log(data);
  let render = document.getElementById(`render-data`);
  let html = ``;
  for (let i = 0; i < data.length; i++) {
    html += `
        <tr>
            <td>${data[i].Id}</td>
            <td>${data[i].Name}</td>
            <td>${data[i].Age}</td>
            <td>${data[i].Gender}</td>
            <td>${data[i].Phone}</td>
            <td><button onclick="openModal('${data[i].Id}')">
            <img src="https://img.icons8.com/fluent-systems-regular/24/000000/change-user-female--v1.png"/>
            </button></td>
            <td>
            <button onclick="deleteMember('${data[i].Id}')">
            <img src="https://img.icons8.com/color/24/000000/delete-user-data.png"/>
            </button>
            </td>
        </tr>
        `;
  }
  render.innerHTML = html;
};

fetchData();

// xoa member
const deleteMember = async (id) => {
  const response = await fetch(
    `http://localhost:8000/api/member/${id}/delete`,
    {
      method: "DELETE",
    }
  );
  const log = await response.json();
  if (!response) alert("successful delete");
  fetchData();
};

//  hien thi modal doi thong tin member
const openModal = async (id) => {
  const modal = document.getElementById(`member-modal`);
  modal.style.display = `block`;
  const title = document.getElementById(`member-title`);
  title.innerHTML = `Thay đổi thông tin của member có ID: ${id}`;
  const updateMemberForm = document.getElementById(`form-update`);
  updateMemberForm.addEventListener(`submit`, async (e) => {
    e.preventDefault();
    const memberInfo = {
      name: updateMemberForm.name.value,
      age: updateMemberForm.age.value,
      gender: updateMemberForm.gender.value,
      phone: updateMemberForm.phone.value,
    };
    console.log(memberInfo);
    const response = await fetch(
      `http://localhost:8000/api/member/${id}/profile?name=${memberInfo.name}&age=${memberInfo.age}&phone=${memberInfo.phone}&gender=${memberInfo.gender}`,
      {
        method: "PUT",
      }
    );
    const log = await response.json();
    console.log(log);
    if(response){
      alert("Change infomation member successful");
      updateMemberForm.name.value  = ``,
      updateMemberForm.age.value = ``,
      updateMemberForm.gender.value = ``,
      updateMemberForm.phone.value = ``,
      modal.style.display = `none`;
    }
    fetchData();
  });
};

// close modal
window.onclick = function (event) {
  const modal = document.getElementById(`member-modal`);
  if (event.target == modal) {
    modal.style.display = "none";
  }
};
const closeModal = () => {
  const modal = document.getElementById(`member-modal`);
  modal.style.display = `none`;
};

//up date thong tin member
