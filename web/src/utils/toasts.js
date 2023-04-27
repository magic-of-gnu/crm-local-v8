function populateToastsFromResponse(response, toasts) {
  if (response.data.hasOwnProperty("toasts")) {
    response.data.toasts.forEach((item) => {
      toasts.value.push(item)
    });
  }
}


export {
  populateToastsFromResponse
}