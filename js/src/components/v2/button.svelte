<script lang="ts">
  import { css } from "@emotion/css";
  import { Button } from "bits-ui";
  import type { Snippet } from "svelte";
  let {
    children,
    ...props
  }: Button.RootProps & {
    children: Snippet;
  } = $props();

  const buttonStyles = css({
    // Bodi tombol transparan persis seperti awal
    background: "none",
    border: "none", // Hilangin border default
    height: 42,
    padding: "0 16px",
    color: "white",
    position: "relative",
    boxSizing: "border-box",
    display: "inline-flex",
    alignItems: "center",
    justifyContent: "center",
    cursor: "pointer",
    zIndex: 1,

    fontWeight: "700",
    textTransform: "uppercase",
    userSelect: "none",
    letterSpacing: "0.8px",
    whiteSpace: "nowrap",
    WebkitTapHighlightColor: "transparent",
    touchAction: "manipulation",
    transform: "translateZ(0)",
    borderWidth: 1,
    borderStyle: "solid",

    borderColor: "transparent",
    WebkitFontSmoothing: "antialiased",
    MozOsxFontSmoothing: "grayscale",
    "&:disabled": {
      color: "rgb(175, 175, 175)",

      "&:before": {
        background: "rgb(229, 229, 229)",
        boxShadow: "none",
        borderColor: "transparent",
      },
    },
    "&::before": {
      borderWidth: 1,
      borderStyle: "solid",
      content: '""',
      position: "absolute",
      borderColor: "rgb(163, 0, 28)",
      top: 0,
      left: 0,
      width: "100%",
      height: "100%",
      backgroundColor: "rgb(229, 0, 39)",
      borderRadius: 12,
      zIndex: -1,
      boxShadow: "rgb(163, 0, 28) 0px 4px 0px",
    },
    ":not(:disabled)": {
      "&:hover": {
        "&::before": {
          boxShadow: "rgb(163, 0, 28) 0px 4px 0px",
          borderColor: "rgb(163, 0, 28)",
          backgroundColor: "rgb(208, 0, 35)",
          color: "rgb(250, 250, 250)",
        },
      },
      "&:active": {
        transform: "translateY(4px)",
      },
      "&:active::before": {
        boxShadow: "0 0 0 #58a700",
      },
    },
  });
</script>

<Button.Root {...props} class={buttonStyles}>
  {@render children?.()}
</Button.Root>
