const METHODS = {
  Get: "GET",
  Post: "POST",
  Put: "PUT",
  Delete: "DELETE",
};
// function parseRequestOptions(rawOptions = {}) {
//   const options = { ...rawOptions };
//   if(!options.headers){
//       options.headers ={'Content-Type':'application/json; charset=UTF-8'}
//   }
//   if(options.body){
//       options.body = JSON.stringify(options.body)
//   }
//   return options
// }
function createHttp() {
  const http = {
    request(url, options = {}) {
    // const reqInit = parseRequestOptions(options);
      return new Promise((resolve, reject) => {
        fetch(url, options)
          .then((response) => {
            if (response.ok) {
              response.json().then((data) => {
                resolve(data);
              });
            }
          })
          .catch((e) => reject(e));
      });
    },
    get(url, options = {}) {
      return http.request(url, Object.assign(options, { method: METHODS.Get }));
    },
    post(url, options = {}) {
      return http.request(
        url,
        Object.assign(options, { method: METHODS.Post })
      );
    },
  };
  return http;
}
export { createHttp };
