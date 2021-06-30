const utils = {}

utils.createContext = () => {
  return {
    config: {
      endpoint: "/",
    },
  };
};

const METHODS = {
  Get: "GET",
  Post: "POST",
  Put: "PUT",
  Delete: "DELETE",
};

utils.RESPONSE_TYPES = {
  Json: 'json',
  Text: 'text',
  Blob: 'blob',
};

const parseRequestOptions = (rawOptions = {}) => {
  const options = { ...rawOptions };
  if (!options.headers) {
    options.headers = { "Content-Type": "application/json; charset=UTF-8"};
  }

  if (options.body) {
    options.body = JSON.stringify(options.body);
  }

  return options;
};

const parseResponse = (response, options = {}) =>{
  return new Promise((resolve, reject) => {
    if (response.ok) {
      switch (options.responseType) {
        case utils.RESPONSE_TYPES.Text:
          response.text().then((data) => {
            resolve(data);
          });
          break;
        case utils.RESPONSE_TYPES.Blob:
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
}

utils.createHttpInstance = (config = {})=> {
  const context = utils.createContext();
  if (config) {
    context.config = Object.assign(context.config, config);
  }

  const http = {
    request(url = '/', options = {}) {
      const reqUrl = url.indexOf('http') !== -1 ? url : `${config.endpoint}${url}`;
      const reqInit = parseRequestOptions(options);
      return fetch(reqUrl, reqInit)
        .then((response) => response.json())
        .catch((e) => Promise.reject(e));
    },
    get(url, options = {}) {
      return http.request(url, Object.assign(options, { method: METHODS.Get }));
    },
    post(url, options = {}) {
      return http.request(url, Object.assign(options, { method: METHODS.Post }));
    },
    put(url, options = {}) {
      return http.request(url, Object.assign(options, { method: METHODS.Put }));
    },
    deleta(url, options = {}) {
      return http.request(url, Object.assign(options, { method: METHODS.Delelte }));
    },
  };

  return http;
}
