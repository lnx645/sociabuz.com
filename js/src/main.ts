import App from "./app/app.svelte";
import "./style.css";
import '@fontsource-variable/rubik/wght.css';

import { mount } from "svelte";
 mount(App, {
  target: document.getElementById("app")!,
});

