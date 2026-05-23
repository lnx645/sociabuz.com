<script>
  import Button from "@/components/button.svelte";
  import Input from "@/components/input.svelte";
  import { css } from "@emotion/css";
</script>

<div
  class={css({
    padding: 10,
    display: "flex",
    flexDirection: "column",
    gap: 10,
  })}
>
  <Input placeholder="Masukan nama anda" name="password" label="Email" />
  <Input
    type="password"
    placeholder="Masukan nama anda"
    name="email"
    label="Password"
  />
  <Button id="login">Login</Button>
</div>
