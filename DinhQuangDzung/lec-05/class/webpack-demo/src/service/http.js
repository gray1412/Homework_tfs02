const METHODS = {
	Get: "GET",
	Post: "POST",
};

// Override window.fetch() function to define a request interceptor
fetch = ((originalFetch) => {
	return (...args) => {
		const customHeader = { "Say-Hello": "Hello World!" };

		console.log("Adding custom header...", customHeader);

		Object.assign(args[1].headers, customHeader);

		console.log("Request sent...");
		return originalFetch.apply(this, args);
	};
})(fetch);

const parseResponse = (response, options = {}) => {
	return new Promise((resolve, reject) => {
		if (response.ok) {
			switch (options.responseType) {
				case "json":
					response.json().then((data) => {
						resolve(data);
					});
					break;
				case "blob":
					resolve.blob().then((data) => {
						resolve(data);
					});
					break;
				default:
					response.text().then((data) => {
						resolve(data);
					});
			}
			return;
		}

		reject(response.statusText);
	});
};

// Create an http instance
const createHttp = () => {
	const http = {
		request(url, options = {}) {
			return fetch(url, options)
				.then((response) => parseResponse(response, options))
				.catch((e) => Promise.reject(e));
		},

		get(url, options = {}) {
			return http.request(
				url,
				Object.assign(options, {
					method: METHODS.Get,
				})
			);
		},

		post(url, options = {}) {
			return http.request(
				url,
				Object.assign(options, { method: METHODS.Post })
			);
		},
	};

	return http;
};

export { createHttp };
