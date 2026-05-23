import { mount } from 'svelte';
import App from '@/App.svelte';
import "the-new-css-reset/css/reset.css";
import "@/css/main.css"
mount(App, { target: document.querySelector('#app')! });
