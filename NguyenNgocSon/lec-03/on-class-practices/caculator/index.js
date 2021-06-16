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

const submit = () => {
  let output = 0;
  let screen = document.getElementById(`screen`);
  let screenValue = document.getElementById(`screen`).value;
  console.log(screenValue);
  if (screenValue.indexOf("+") > 0) {
    let value = screenValue.split("+");
    if (value[1] !== ``) output = parseFloat(value[0]) + parseFloat(value[1]);
    else output = screenValue;
    console.log(value);
  } else if (screenValue.indexOf("-") > 0) {
    let value = screenValue.split("-");
    if (value[1] !== ``) output = parseFloat(value[0]) - parseFloat(value[1]);
    else output = screenValue;
  } else if (screenValue.indexOf("*") > 0) {
    let value = screenValue.split("*");
    if (value[1] !== ``) output = parseFloat(value[0]) * parseFloat(value[1]);
    else output = screenValue;
  } else if (screenValue.indexOf("/") > 0) {
    let value = screenValue.split("/");
    if (value[1] !== ``) output = parseFloat(value[0]) / parseFloat(value[1]);
    else output = screenValue;
  } else output = parseFloat(screenValue);
  console.log(output);
  let change = output+'';
  if (
    change.indexOf("+") > 0 ||
    change.indexOf("-") > 0 ||
    change.indexOf("*") > 0 ||
    change.indexOf("/") > 0
  )
    screen.value = output;
  else screen.value = Math.round(output * 100) / 100;
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
