document.addEventListener("DOMContentLoaded", function () {
  var palettes = ["nord", "catppuccin"];
  var labels = { nord: "Nord", catppuccin: "Cat" };

  var footer = document.querySelector("[data-toggle-animation='show']");
  if (!footer) return;

  var btn = document.createElement("button");
  btn.id = "hextra-palette-toggle";
  btn.title = "Цветовая схема";

  function current() {
    return document.documentElement.getAttribute("data-palette") || "nord";
  }

  function apply(palette) {
    document.documentElement.setAttribute("data-palette", palette);
    localStorage.setItem("hextra-palette", palette);
    btn.textContent = labels[palette] || palette;
  }

  btn.textContent = labels[current()] || current();

  btn.addEventListener("click", function () {
    var idx = palettes.indexOf(current());
    apply(palettes[(idx + 1) % palettes.length]);
  });

  footer.appendChild(btn);
});
