(function () {
  const apiBaseURL = "http://localhost:8080/sessions/";
  const uuid = crypto.randomUUID();
  const xClient = location.origin + location.pathname;
  let timer = null;
  let input = "";

  function registerSession() {
    fetch(apiBaseURL + uuid, {
      method: "PUT",
      headers: { "X-Client": xClient },
    }).catch(() => {});
  }

  function handleInput(event) {
    input = event.target.value;

    if (timer) {
      clearTimeout(timer);
    }
    timer = setTimeout(handleNoInput, 3 * 1000);
  }

  function handleNoInput() {
    if (timer) {
      clearTimeout(timer);
    }

    fetch(apiBaseURL + uuid, {
      method: "PATCH",
      headers: { "X-Client": xClient },
      body: JSON.stringify([input.trim()]),
    }).catch(() => {});

    input = "";
  }

  window.addEventListener("input", handleInput);

  window.addEventListener("blur", handleNoInput);
  window.addEventListener("visibilitychange", handleNoInput);

  registerSession();
})();
