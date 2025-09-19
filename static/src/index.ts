import htmx from "htmx.org";
import "htmx.org";
import "iconify-icon";

declare global {
	interface Window {
		htmx: typeof htmx;
	}
}

window.htmx = htmx;
