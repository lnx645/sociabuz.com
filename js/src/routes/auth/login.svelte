<script>
  import Button from "@/components/v2/button.svelte";
  import Input from "@/components/input.svelte";
  import { css } from "@emotion/css";
  import FacebookIcon from "@/icons/facebook-icon.svelte";
  import GoogleIcon from "@/icons/google-icon.svelte";

  const wrapper = css({
    maxWidth: 320,
    margin: "auto",
    color: "#4b4b4b",
    h1: {
      fontSize: 24,
      padding: 0,
      margin: 0,
    },
    p: {
      fontSize: 12,
    },
  });
  const fields = css({
    display: "flex",
    flexDirection: "column",
    gap: 12,
  });
  const header = css({
    marginBlock: 24,
    marginBottom: 10,
    textAlign: "center",
    h1: {
      fontWeight: 600,
      fontSize: 24,
    },
  });
  const buttons = css({
    display: "flex",
    alignItems: "center",
    flexDirection: "column",
    boxSizing: "border-box",
    gap: 10,
    button: {
      width: "100%",
    },
  });
  const separator = css({
    marginBlock: 8,
    fontWeight: "bold",
    color: "#afafaf",
  });
</script>

<dir class={wrapper}>
  <div class={header}>
    <h1>Silahkan Masuk</h1>
  </div>

  <div class={fields}>
    <Input placeholder="Masukan nama anda" name="password" label="EMAIL" />
    <Input
      type="password"
      placeholder="Masukan kata sandi anda"
      name="email"
      label="PASSWORD"
    />
    <div class={buttons}>
      <Button variant="success">LOGIN</Button>
      <div class={separator}>
        <span>ATAU</span>
      </div>
      <div
        class={css({
          display: "flex",
          width: "100%",
          gap: 10,
          alignItems: "center",
          justifyContent: "space-between",
          flex: 1,
        })}
      >
        <Button --size-height="47px" --text-color="#EA4335" variant="default">
          <GoogleIcon />
          <span>Google</span>
        </Button>
        <Button --size-height="47px" --text-color="#3b5998" variant="default">
          <FacebookIcon />
          <span>Facebook</span>
        </Button>
      </div>
    </div>
  </div>
</dir>
