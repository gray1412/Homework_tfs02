const METHODS = {
  Get: 'GET',
  Post: 'POST',
  Delete:"Delete",
  Put:"PUT"
}

/**
 * Create an http instance
 */
function createHttpJson() {
  const http = {
    request(url, options = {}) {
      return new Promise((resolve, reject) => {
        fetch(url, options).then((response) => {
          if (response.ok) {
            response.json().then((data) => {
              resolve(data)
            })
          }
        }).catch((e) => reject(e))
      })
    },
    get(url, options = {}) {
      return http.request(url, Object.assign(options, { method: METHODS.Get }))
    },
    post(url, options = {}) {
      return http.request(url, Object.assign(options, { method: METHODS.Post }))
    },
    delete(url, options = {}) {
      return http.request(url, Object.assign(options, { method: METHODS.Delete }))
    },
    put(url, options = {}) {
      return http.request(url, Object.assign(options, { method: METHODS.Put }))
    }
  }

  return http
}
function createHttpText() {
  const http = {
    request(url, options = {}) {
      return new Promise((resolve, reject) => {
        fetch(url, options).then((response) => {
          if (response.ok) {
            response.text().then((data) => {
              resolve(data)
            })
          }
        }).catch((e) => reject(e))
      })
    },
    get(url, options = {}) {
      return http.request(url, Object.assign(options, { method: METHODS.Get }))
    },
    post(url, options = {}) {
      return http.request(url, Object.assign(options, { method: METHODS.Post }))
    },
    delete(url, options = {}) {
      return http.request(url, Object.assign(options, { method: METHODS.Delete }))
    }
  }

  return http
}
export {
  createHttpJson,
  createHttpText
}
