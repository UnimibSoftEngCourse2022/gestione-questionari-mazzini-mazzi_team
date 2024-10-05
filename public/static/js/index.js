htmx.defineExtension("reset-on-success", {
  onEvent: function (name, event) {
    if (name !== "htmx:beforeSwap") return;
    if (event.detail.isError) return;

    console.log("topperoni");

    const triggeringElt = event.detail.requestConfig.elt;
    if (
      !triggeringElt.closest("[hx-reset-on-success]") &&
      !triggeringElt.closest("[data-hx-reset-on-success]")
    )
      return;

    switch (triggeringElt.tagName) {
      case "INPUT":
      case "TEXTAREA":
        triggeringElt.value = triggeringElt.defaultValue;
        break;
      case "SELECT":
        //too much work
        break;
      case "FORM":
        triggeringElt.reset();
        break;
    }
  },
});

document.body.addEventListener("htmx:beforeOnLoad", function (evt) {
  if (evt.detail.xhr.status === 404) {
    evt.detail.shouldSwap = true;
    evt.detail.isError = false;
  }
});

htmx.onLoad(function (content) {
  initFlowbite();
});
