document.addEventListener("DOMContentLoaded", function () {
  var chevronLeft =
    '<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" ' +
    'viewBox="0 0 24 24" fill="none" stroke="currentColor" ' +
    'stroke-width="2" stroke-linecap="round" stroke-linejoin="round" ' +
    'aria-hidden="true"><path d="M15 18l-6-6 6-6"/></svg>';
  var chevronRight =
    '<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" ' +
    'viewBox="0 0 24 24" fill="none" stroke="currentColor" ' +
    'stroke-width="2" stroke-linecap="round" stroke-linejoin="round" ' +
    'aria-hidden="true"><path d="M9 18l6-6-6-6"/></svg>';

  var sidebar = document.querySelector("aside.hextra-sidebar-container");
  if (sidebar) {
    var btn = document.createElement("button");
    btn.id = "hextra-sidebar-toggle";
    btn.setAttribute("aria-label", "Скрыть/показать боковую панель");
    btn.innerHTML = chevronLeft;
    sidebar.after(btn);

    if (localStorage.getItem("hextra-sidebar-hidden") === "true") {
      document.body.classList.add("sidebar-hidden");
    }

    btn.addEventListener("click", function () {
      var hidden = document.body.classList.toggle("sidebar-hidden");
      localStorage.setItem("hextra-sidebar-hidden", hidden ? "true" : "false");
    });
  }

  var toc = document.querySelector("nav.hextra-toc");
  if (toc) {
    var article = toc.parentElement
      ? toc.parentElement.querySelector("article")
      : null;
    if (article) {
      var tocBtn = document.createElement("button");
      tocBtn.id = "hextra-toc-toggle";
      tocBtn.setAttribute("aria-label", "Скрыть/показать содержание");
      tocBtn.innerHTML = chevronRight;
      article.after(tocBtn);

      if (localStorage.getItem("hextra-toc-hidden") === "true") {
        document.body.classList.add("toc-hidden");
      }

      tocBtn.addEventListener("click", function () {
        var hidden = document.body.classList.toggle("toc-hidden");
        localStorage.setItem("hextra-toc-hidden", hidden ? "true" : "false");
      });
    }
  }
});
