<script lang="ts">
  import { useId } from "@/core/utils";
  import UserIcon from "@/icons/user-icon.svelte";
  import { css } from "@emotion/css";
  import type { HTMLInputAttributes } from "svelte/elements";
  type InputProps = HTMLInputAttributes & {
    value: string;
    name: string;
    id: string;
    label: string;
  };

  let {
    label,
    id = useId(),
    value = $bindable(),
    name,
    ...props
  }: InputProps = $props();

  const style = css({
    width: "100%",
    height: "35px",
    boxSizing: "border-box",
    borderColor: "rgb(188, 188, 188)",
    borderWidth: 1,
    borderStyle: "solid",
    outline: "none",
    boxShadow: "rgba(255, 255, 255, 0.2) 0px 0px 0px 1px inset",
    transition: "0.5s ease",
    background: "#f8f9f9",
    display: "flex",
    alignItems: "center",
    fontSize: 12,
    borderRadius: 7,
    overflow: "hidden",
    paddingInline: 2,
    color: "rgb(82, 82, 82)",
    ":hover": {
      borderColor: "rgb(110, 110, 110)",
    },
    ":active": {
      borderColor: "rgb(110, 110, 110)",
    },
    ":focus": {
      borderColor: "rgb(110, 110, 110)",
    },
  });
  let iconLeft = css({
    width: 16,
    height: "100%",
    display: "flex",
    paddingInline: 5,
    alignItems: "center",
    justifyContent: "center",
    "&>svg": {
      width: 24,
      height: 24,
    },
  });
  const inputStyle = css({
    width: "100%",
    height: "100%",
    fontWeight: 600,
    background:"transparent",
    color: "#525252",
    boxSizing: "border-box",
    outline: "none",
    border: "none",
    "::placeholder": {
      color: "rgb(188, 188, 188)",
    },
    flex: 1,
  });
  const labelStyle = css({
    fontSize: 12,
    fontWeight: 600,
  });
  const fieldStyle = css({
    lineHeight: 1.5,
  });
</script>

<div class={fieldStyle}>
  <label class={labelStyle} for={id}>
    {label}
  </label>
  <div class={style}>
    <div class={iconLeft}>
      <UserIcon />
    </div>
    <input bind:value type="text" {name} class={inputStyle} {id} {...props} />
  </div>
</div>
