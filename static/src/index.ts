import Alpine from "alpinejs";
import htmx from "htmx.org";
import "htmx.org";
import "iconify-icon";

declare global {
  interface Window {
    htmx: typeof htmx;
    Alpine: typeof Alpine;
  }
}

window.htmx = htmx;
window.Alpine = Alpine;

Alpine.start();
