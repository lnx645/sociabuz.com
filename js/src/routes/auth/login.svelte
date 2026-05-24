<script lang="ts">
  import Button from "@/components/v2/button.svelte";
  import Input from "@/components/input.svelte";
  import { css } from "@emotion/css";
  import FacebookIcon from "@/icons/facebook-icon.svelte";
  import GoogleIcon from "@/icons/google-icon.svelte";
  import Loader from "@/components/loader.svelte";
  import IconPassword from "@/icons/icon-password.svelte";
  import { navigate } from "sv-router/generated";
  import { loginFormStyled } from "@/styled/login";

  let data = $state<{
    email: string;
    password: string;
  }>({
    email: "",
    password: "",
  });

  let loading = $state(false);

  function login(e: SubmitEvent) {
    e.preventDefault();
    loading = true;
    setTimeout(() => {
      loading = false;
      navigate("/manage");
    }, 1000);
  }
</script>

<dir class={loginFormStyled.wrapper}>
  <div class={loginFormStyled.header}>
    <h1>Silahkan Masuk</h1>
  </div>

  <form action="" onsubmit={login} class={loginFormStyled.fields}>
    <Input
      bind:value={data.email}
      placeholder="Masukan nama anda"
      name="password"
      label="EMAIL"
    />
    <Input
      bind:value={data.password}
      type="password"
      placeholder="Masukan kata sandi anda"
      name="email"
      label="PASSWORD"
    />
    <div class={loginFormStyled.buttons}>
      <Button disabled={loading} variant="success">
        {#if loading}
          <Loader />
        {:else}
          Login
        {/if}
      </Button>
      <div class={loginFormStyled.separator}>
        <span>ATAU</span>
      </div>
      <div class={loginFormStyled.actionButtons}>
        <Button
          disabled={loading}
          --size-height="47px"
          --text-color="#EA4335"
          variant="default"
        >
          <GoogleIcon />
          <span>Google</span>
        </Button>
        <Button
          disabled={loading}
          --size-height="47px"
          --text-color="#3b5998"
          variant="default"
        >
          <FacebookIcon />
          <span>Facebook</span>
        </Button>
      </div>
    </div>
  </form>
</dir>
