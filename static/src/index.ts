import htmx from 'htmx.org';
import 'htmx.org';

declare global {
  interface Window {
    htmx: typeof htmx;
  }
}

window.htmx = htmx;
