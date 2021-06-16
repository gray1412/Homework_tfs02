const addValue = (id) => {
  let screen = document.getElementById(`screen`);
  let text = ``;
  let screenValue = document.getElementById(`screen`).value;
  if(screenValue.length === 0 && (id === "+" || id === "-" || id==="*" || id==="/")){
    text = "0"+id;
  }
  else if (
    screenValue.indexOf("+") > 0 ||
    screenValue.indexOf("-") > 0 ||
    screenValue.indexOf("*") > 0 ||
    screenValue.indexOf("/") > 0
  ) {
    if (id !== "+" && id !== "-" && id !== "*" && id !== "/") {
      text = screenValue + id;
    } else text = screenValue;
  } else text = screenValue + id;

  screen.value = text;
};


const fetchData = async (math, a, b) => {
  const response = await fetch(`http://localhost:8080/${math}?a=${a}&b=${b}`);
  const data = await response.json();
  showData(data)
  return data
};

const showData =  (data) =>{
  const screen = document.getElementById(`screen`);
  screen.value = Math.round(data.value * 100) / 100;;
}
const submit = () => {
  let screen = document.getElementById(`screen`);
  let screenValue = document.getElementById(`screen`).value;
  let output = 0;
  if (screenValue.indexOf("+") > 0) {
    let value = screenValue.split("+");
    if (value[1] !== ``) fetchData("sum", value[0], value[1])
    else output = screenValue;
  } else if (screenValue.indexOf("-") > 0) {
    let value = screenValue.split("-");
    if (value[1] !== ``) fetchData("sub", value[0], value[1])
    else output = screenValue;
  } else if (screenValue.indexOf("*") > 0) {
    let value = screenValue.split("*");
    if (value[1] !== ``) fetchData("mul", value[0], value[1])
    else output = screenValue;
  } else if (screenValue.indexOf("/") > 0) {
    let value = screenValue.split("/");
    if (value[1] !== ``) output = fetchData("div", value[0], value[1])
    else output = screenValue;
  } else output = parseFloat(screenValue);
  let change = output+'';
  if (
    change.indexOf("+") > 0 ||
    change.indexOf("-") > 0 ||
    change.indexOf("*") > 0 ||
    change.indexOf("/") > 0
  )
    screen.value = output;
  
};

const clearOne = () =>{
  let screen = document.getElementById(`screen`);
  let value = document.getElementById(`screen`).value;
  let clear = value.split("");
  console.log(clear);
  screen.value = clear.slice(0, clear.length-1).join("");
}

const clearAll = () => {
  let screen = document.getElementById(`screen`);
  screen.value = ``;
};
