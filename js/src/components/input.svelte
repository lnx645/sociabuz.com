<script lang="ts">
  import { useId } from "@/core/utils";
  import { css, cx } from "@emotion/css";
  import type { Snippet } from "svelte";
  import type { HTMLInputAttributes } from "svelte/elements";
  type InputProps = HTMLInputAttributes & {
    value?: string;
    name?: string;
    id?: string;
    size?: any;
    label: string;
    isError?: boolean;
    prefix?: Snippet;
  };

  let {
    label,
    isError = false,
    id = useId(),
    value = $bindable(),
    name,
    size,
    prefix,
    ...props
  }: InputProps = $props();

  let isPrefix = $derived(!!prefix);

  const style = css({
    width: "100%",
    borderRadius: 16,
    boxSizing: "border-box",
    borderColor: "#e5e5e5",
    borderWidth: 2,
    borderStyle: "solid",
    outline: "none",
    height: 48,
    boxShadow: "rgba(255, 255, 255, 0.2) 0px 0px 0px 1px inset",
    transition: "0.5s ease",
    background: "#f7f7f7",
    display: "flex",
    alignItems: "center",
    caretColor: "#1cb0f6",
    lineHeight: "1.75rem",
    overflow: "hidden",

    color: "#4b4b4b",
    ":hover": {
      borderColor: "#e5e5e5",
    },
    ":active": {
      borderColor: "#1cb0f6",
    },
    ":focus-within": {
      borderColor: "#1cb0f6",
    },
  });
  let iconLeft = css({
    color: "#4b4b4b",
    height: "100%",
    display: "flex",
    paddingInline: 5,
    paddingLeft: 6,
    alignItems: "center",
    justifyContent: "center",
    "&>svg": {
      width: 24,
      height: 24,
    },
  });
  const inputStyle = () =>
    css({
      width: "100%",
      height: "100%",
      paddingBlock: 9,
      paddingLeft: !isPrefix ? 13 : 0,
      WebkitAppearance: "none",
      fontWeight: 600,
      background: "transparent",
      color: "#4b4b4b",
      boxSizing: "border-box",
      fontSize: "0.98rem",
      outline: "none",
      border: "none",
      "::placeholder": {
        color: "#4b4b4b",
      },
      flex: 1,
    });
  const labelStyle = css({
    fontSize: 14,
    fontWeight: 800,
    color: "#4b4b4b",
  });
  const fieldStyle = css({
    lineHeight: 1.5,
  });
  const errorStyle = css({
    color: "#ea2b2b",
    fontWeight: 600,
    gap: 2,
    fontSize: 13,
    marginTop: 3,
    display: "flex",
    alignItems: "center",
    justifyContent: "start",
  });
</script>

<div class={fieldStyle}>
  <label class={labelStyle} for={id}>
    {label}
  </label>
  <div
    class={cx(
      style,
      css({
        height: size == "sm" ? 40 : size == "md" ? 32 : size == "lg" ? 48 : 48,
      }),
    )}
  >
    {#if prefix}
      <div class={iconLeft}>
        {@render prefix()}
      </div>
    {/if}
    <input bind:value type="text" {name} class={inputStyle()} {id} {...props} />
  </div>
  {#if isError}
    <div class={errorStyle}>
      <svg height="16" width="16" viewBox="0 0 24 24"
        ><path
          fill-rule="evenodd"
          clip-rule="evenodd"
          d="M12 21C16.9706 21 21 16.9706 21 12C21 7.02944 16.9706 3 12 3C7.02944 3 3 7.02944 3 12C3 16.9706 7.02944 21 12 21ZM13 8C13 7.44772 12.5523 7 12 7C11.4477 7 11 7.44772 11 8V13C11 13.5523 11.4477 14 12 14C12.5523 14 13 13.5523 13 13V8ZM12 17C12.5523 17 13 16.5523 13 16C13 15.4477 12.5523 15 12 15C11.4477 15 11 15.4477 11 16C11 16.5523 11.4477 17 12 17Z"
          fill="currentcolor"
        ></path></svg
      >
      <span> Kata sandi terlalu pendek</span>
    </div>
  {/if}
</div>
