const name = document.querySelector(".name"),
	message = document.querySelector("#message-input"),
	nameInput = document.querySelector("#name-input"),
	content = document.querySelector(".content"),
	sendBtn = document.querySelector(".send-btn"),
	chatForm = document.querySelector(".chat-form"),
	counter = document.querySelector("#counter");

let messageCount = 0;

message.onkeyup = (e) => {
	if (e.target.value.length) {
		counter.innerHTML = "160/160";
	}
	counter.innerHTML = `${160 - e.target.value.length}/160`;
};

chatForm.onsubmit = (e) => {
	e.preventDefault();

	fetch("http://localhost:8080/chat", {
		method: "POST",
		headers: {
			Accept: "application/json",
			"Content-Type": "application/json",
		},
		body: JSON.stringify({
			name: nameInput.value,
			message: message.value,
		}),
	});

	message.value = "";
	messageCount++;
	renderMessage();
};

const renderMessage = async () => {
	content.innerHTML = "";
	const results = await getMessage();
	console.log(results);

	for (let result of results) {
		const html = `
    <div class="message-wrapper">
      <span class="name">${result.name}</span>
      <span class="message-content">${result.message}</span>
    </div>`;
		content.innerHTML += html;
	}
};

const getMessage = async () => {
	const res = await fetch("http://localhost:8080/message");
	return res.json();
};

const getDataInterval = () => {
	setInterval(async () => {
		const res = await getMessage();
		if (res.length != messageCount) {
			console.log("new message");
			messageCount = res.length;
			renderMessage();
		}
	}, 3000);
};

getDataInterval();
