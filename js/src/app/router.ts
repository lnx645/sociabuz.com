import { createRouter } from "sv-router";

const Home = () => import("./index.svelte");
const Login = () => import("./auth/login.svelte");
const Register = () => import("./auth/register.svelte");
const NotFound = () => import("./not-found.svelte");
export const { p, navigate, isActive, route } = createRouter({
  "/": Home,
  "/login": Login,
  "/register": Register,
  "*notfound": NotFound,
});
