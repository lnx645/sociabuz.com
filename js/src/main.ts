import { mount } from 'svelte';
import App from '@/App.svelte';
import "@/css/main.css"
mount(App, { target: document.querySelector('#app')! });
