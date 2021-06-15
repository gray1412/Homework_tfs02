let inputField = '';

function setInputField(inputField) {
    document.getElementById("input-field").value = inputField;
}

function handleNumberClick(event) {
    inputField += (event.target.dataset.number);
    setInputField(inputField);
}

function handleOperatorClick(event) {
    inputField += (event.target.dataset.operation);
    setInputField(inputField);
}

function handleEqualClick(event) {
    const result = eval(inputField);
    inputField = result.toString();
    setInputField(inputField);
}

function handleClearClick() {
    inputField = '';
    setInputField(inputField)
}

document.addEventListener("DOMContentLoaded", (event) => {
    const numberButtons = document.querySelectorAll("[data-number]");
    numberButtons.forEach((button) => {
      button.addEventListener("click", (event) => handleNumberClick(event));
    })

    const operatorButtons = document.querySelectorAll("[data-operation]");
    operatorButtons.forEach((button) => {
    button.addEventListener("click", (event) => handleOperatorClick(event));
    })

    document.querySelector("[data-equals]").addEventListener("click", handleEqualClick);
    document.querySelector("[data-all-clear]").addEventListener("click", handleClearClick);
})