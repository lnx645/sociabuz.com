<script lang="ts">
  import { useId } from "@/core/utils";
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
    paddingInline: 6,
    borderColor: "rgb(188, 188, 188)",
    borderWidth: 1,
    borderStyle: "solid",
    outline: "none",
    transition: "0.5s ease",
    background: "transparent",
    fontSize: 12,
    borderRadius: 7,
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
  const labelStyle = css({
    fontSize: 12,
    fontWeight: 600,
  });
  const fieldStyle = css({
    lineHeight:1.5
  })
</script>

<div class={fieldStyle}>
  <label class={labelStyle} for={id}>
    {label}
  </label>
  <input bind:value type="text" {name} {id} class={style} {...props} />
</div>
