<script lang="ts">
  import { css, cx } from "@emotion/css";
  import { Button } from "bits-ui";
  import type { Snippet } from "svelte";

  type ButtonVariant = "default" | "danger" | "warning" | "success" | "info";

  let {
    children,
    variant = "default",
    class: className,
    ...props
  }: Button.RootProps & {
    children?: Snippet;
    variant?: ButtonVariant;
  } = $props();
  // #4b4b4b
  const baseStyles = css({
    background: "none",
    border: "none",
    height: "var(--size-height,42px)",
    MozBoxSizing: "border-box",
    padding: "0 4px",
    display: "inline-flex",
    color: "var(--text-color,white)",
    position: "relative",
    boxSizing: "border-box",
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
    borderRadius: 16,
    gap: 5,
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
      borderWidth: "var(--border-width,1)",
      borderStyle: "solid",
      content: '""',
      position: "absolute",
      borderColor: "var(--border-color)",
      top: 0,
      left: 0,
      width: "100%",
      height: "100%",
      backgroundColor: "var(--background)",
      borderRadius: 16,
      zIndex: -1,
      boxShadow: "var(--shadow-color) 0px 4px 0px",
    },
    ":not(:disabled)": {
      "&:hover": {
        "&::before": {
          boxShadow: "var(--shadow-color) 0px 4px 0px",
          borderColor: "var(--border-color)",
          backgroundColor: "var(--background-hover)",
          color: "rgb(250, 250, 250)",
        },
      },
      "&:active": {
        transform: "translateY(4px)",
      },
      "&:active::before": {
        boxShadow: "0 0 0 transparent", // FIX: Sebelumnya hardcode hijau (#58a700)
      },
    },
  });

  const variantStyles: Record<ButtonVariant, string> = {
    default: css({
      "--border-width": "2px",
      "--background": "white",
      "--shadow-color": "#cecece",
      "--border-color": "#cecece",
      "--background-hover": "#e7e7e7",
    }),
    danger: css({
      "--border-width": "1px",

      "--background": "#ff4b4b",
      "--shadow-color": "#ea2b2b",
      "--border-color": "#ea2b2b",
      "--background-hover": "#ff3333",
    }),
    warning: css({
      "--border-width": "1px",
      "--background": "#ff9600",
      "--shadow-color": "#cc7800",
      "--border-color": "#cc7800",
      "--background-hover": "#ff8c00",
    }),
    success: css({
      "--border-width": "1px",

      "--background": "#58cc02",
      "--shadow-color": "#58a700",
      "--border-color": "#58a700",
      "--background-hover": "#68d716",
    }),
    info: css({
      "--border-width": "1px",

      "--background": "#2bced6",
      "--shadow-color": "#22a6ac",
      "--border-color": "#22a6ac",
      "--background-hover": "#26b9c0",
    }),
  };
</script>

<!-- Gunakan cx() dari emotion untuk menggabungkan base class, variant class, dan custom class -->
<Button.Root
  {...props}
  class={cx(baseStyles, variantStyles[variant], className as string)}
>
  {@render children?.()}
</Button.Root>
