function start() {
    numberButtonOnClick();
    operaterButtonOnClick();
    equalButtonOnClick();
    clearButtonOnClick();
}

start();

let inputField = '';
let num1 = ''
let num2 = ''
let operator = ''

var url_calculator = "http://localhost:5000/hello"

function setInputField(inputField) {
    document.getElementById("input-field").value = inputField;
}

function postData(data) {
    var options = {
        method: "POST",
        mode: 'cors',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    };
    return fetch(url_calculator, options)
        .then(r => r.json())
}

function numberButtonOnClick() {
    var numberBtn = document.querySelectorAll("[data-number]")

    numberBtn.forEach((button) => {
        button.addEventListener("click", function (e) {
            inputField += (e.target.dataset.number)
            setInputField(inputField);
            if(operator == "") {
                num1 += (e.target.dataset.number)
            } else {
                num2 += (e.target.dataset.number)
            }
        });
    })
}

function operaterButtonOnClick() {
    var operaterBtn = document.querySelectorAll("[data-operation]")

    operaterBtn.forEach((button) => {
        button.addEventListener("click", function (e) {
            if (operator == "") {
                inputField += (e.target.dataset.operation)
                setInputField(inputField);
                operator = (e.target.dataset.operation)
            } else {
                if (num2 == "") {
                    operator = (e.target.dataset.operation)
                    inputField = inputField.substring(0, inputField.length - 1)
                    inputField += (e.target.dataset.operation)
                    setInputField(inputField);
                }
            }
        });
    })
}

function clearButtonOnClick() {
    var clearBtn = document.querySelectorAll("[data-all-clear]")

    clearBtn.forEach((button) => {
        button.addEventListener("click", function (e) {
            inputField = "";
            setInputField(inputField);
            expression = {
                num1: "", 
                num2 : "", 
                operator: ""
            }
            postData(expression)
            num1 = ""
            num2 = ""
            operator = ""
        });
    })
}

function equalButtonOnClick() {
    document.querySelector("[data-equals]").addEventListener("click", function () {
        expression = {
            num1: num1, 
            num2 : num2, 
            operator: operator
        }

        postData(expression).then(data => {
            console.log("post", data);
            inputField = data.result;
            setInputField(inputField);
            num1 = data.result.toString();
        });
        num2 = ""
        operator = ""
        
    });
}