<script>
  import Button from "@/components/v2/button.svelte";
  import Input from "@/components/input.svelte";
  import { css } from "@emotion/css";
</script>

<div
  class={css({
    padding: 10,
    display: "flex",
    flexDirection: "column",
    gap: 10,
    maxWidth: 320,
    margin: "auto",
  })}
>
  <Input placeholder="Masukan nama anda" name="password" label="EMAIL" />
  <Input
    type="password"
    placeholder="Masukan nama anda"
    name="email"
    label="PASSWORD"
  />
  <Button>LOGIN</Button>
</div>
