import { createHttp } from "./service/http";

const get = document.getElementById("get-form");
const form = document.getElementById("form");
const upload = document.getElementById("upload-form");
const result = document.getElementById("result");
const http = createHttp();

get.onsubmit = (e) => {
	e.preventDefault();
	result.innerHTML = "Loading...";
	http
		.get("https://jsonplaceholder.typicode.com/posts", {
			headers: {
				"Content-Type": "application/json",
			},
			responseType: "json",
		})
		.then((data) => {
			console.log(data);
			const html = data.map(
				(elem) => `
        <li>
        <div>ID: ${elem.id}</div>
        <p>Title: ${elem.title}</p>
        <p>Body: ${elem.body}</p>
        </li>
        `
			);
			result.innerHTML = "";
			html.forEach((elem) => {
				result.innerHTML += elem;
			});
		})
		.catch((err) => {
			console.log(err);
		});
};

form.onsubmit = (e) => {
	e.preventDefault();
	const title = document.getElementById("title"),
		body = document.getElementById("body");

	if (title.value === "" || body.value === "") {
		alert("Please enter something!");
		return;
	}

	result.innerHTML = "Loading...";
	http
		.post("https://jsonplaceholder.typicode.com/posts", {
			body: JSON.stringify({
				title: title.value,
				body: body.value,
				userId: 1,
			}),
			headers: {
				"Content-type": "application/json",
			},
			responseType: "json",
		})
		.then((data) => {
			console.log(data);
			const html = `
        <li>
        <div>ID: ${data.id}</div>
        <p>Title: ${data.title}</p>
        <p>Body: ${data.body}</p>
        </li>
        `;
			result.innerHTML = html;
			title.value = "";
			body.value = "";
		});
};

upload.onsubmit = (e) => {
	e.preventDefault();

	const formData = new FormData();
	const fileField = document.querySelector('input[type="file"]');

	formData.append("file", fileField.files[0]);

	console.log(fileField.files[0]);

	http
		.post("https://jsonplaceholder.typicode.com/posts", {
			body: formData,
		})
		.then((result) => {
			console.log("Success:", result);
		})
		.catch((error) => {
			console.error("Error:", error);
		});
};
