import App from "./app/app.svelte";
import "./style.css";
import "./scss/app.scss";
import { mount } from "svelte";
 mount(App, {
  target: document.getElementById("app")!,
});

