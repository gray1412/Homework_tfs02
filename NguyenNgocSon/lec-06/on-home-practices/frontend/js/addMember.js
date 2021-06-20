const addMemberForm = document.getElementById(`form-register`);
addMemberForm.addEventListener(`submit`, async (e) => {
  e.preventDefault();
  const memberInfo = {
    name: addMemberForm.name.value,
    age: addMemberForm.age.value,
    gender: addMemberForm.gender.value,
    phone: addMemberForm.phone.value,
  };
  console.log(memberInfo);
  const response = await fetch(`http://localhost:8000/api/member?name=${memberInfo.name}&age=${memberInfo.age}&phone=${memberInfo.phone}&gender=${memberInfo.gender}`, {
    method: 'POST',
  });
  const log = await response.json();
  console.log(log);
  if (response !== undefined) 
  alert("add member successful");
//   addMemberForm.value = ``;
  fetchData();
});