const createHttpInstance = (config = {}) => {
  const context = utils.createContext();
  if (config) {
    context.config = Object.assign(context.config, config);
  }
  return {
    request(url = "/") {
      const reqUrl = url.indexOf("http") !== -1 ? url : `${config.endpoint}${url}`;
      return fetch(reqUrl)
        .then((response) => response.json())
        .catch((e) => Promise.reject(e));
    },
  };
};

document.addEventListener('DOMContentLoaded', () => {
    // Init http instance
    const http = createHttpInstance({ endpoint: 'https://jsonplaceholder.typicode.com' });
    const button = document.getElementById('btnRequest');
    const result = document.getElementById('result');
  
    button.addEventListener('click', () => {
      // Reset result
      result.innerHTML = 'Loading...';
  
      // Request
      http.request('/todos').then((response) => {
        console.log(`res`, JSON.stringify(response));
      }).catch((e) => {
        console.log(`error`, JSON.stringify(e) );
      });
    });
  });

