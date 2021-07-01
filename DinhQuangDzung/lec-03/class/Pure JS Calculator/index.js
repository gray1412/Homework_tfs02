let current;
let result;

const resultDisplay = document.querySelector(".result");
const prevNumDisplay = document.querySelector(".prev");
const equal = document.querySelector(".equal");
const numberButtons = document.getElementsByClassName("number-btn");
const operatorButtons = document.getElementsByClassName("operator");

const compute = () => {
	if (!prevNumDisplay.innerText) {
		return;
	}
	result = eval(`${prevNumDisplay.innerText} ${current}`);
	resultDisplay.innerText = result;
	console.log("Result:", result);
};

const handleNumClick = (e) => {
	const button = e.target;
	if (resultDisplay.innerText.length < 12) {
		resultDisplay.innerText += button.innerText;
		current = Number(resultDisplay.innerText);
	}
};

const handleOpClick = (e) => {
	const operation = e.target.innerHTML;
	if (!resultDisplay.innerText) {
		return;
	}
	if (prevNumDisplay.innerText) {
		compute();
	}
	prevNumDisplay.innerText = resultDisplay.innerText + " " + operation;
	resultDisplay.innerText = "";
};

const addEventListeners = () => {
	for (let button of numberButtons) {
		button.addEventListener("click", handleNumClick);
	}
	for (let op of operatorButtons) {
		op.addEventListener("click", handleOpClick);
	}
	equal.addEventListener("click", () => {
		compute();
		prevNumDisplay.innerText = "";
	});
};

(function () {
	addEventListeners();
})();
