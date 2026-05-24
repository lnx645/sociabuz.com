<script module lang="ts">
  const style = {
    isActive: css({
      // boxShadow: "inset 0 -4px #0000000f",
      // border: "1px solid #d2d4d8",
      ":before": {
        background:"#f7f7f7",
        border: "1.5px solid var(--color-border-nav-link-active)",
        boxShadow: "0 2px 0",
      },
      ":hover": {
        ":before": {
          background: "#f7f7f7",
        },
      },
    }),
    wrapper: css({
      position: "relative",
      paddingInline: 10,
      paddingBlock: 8,
      userSelect:"none",
      cursor:"pointer",
      borderRadius: 10,
      background: "none",
      gap: 10,
      ":before": {
        content: "''",
        background: "white",
        left: -2,
        right: -3,
        bottom: -2,
        borderRadius: 10,
        color: "#e5e5e5",
        border: "none",
        top: -2,
        position: "absolute",
        boxShadow: "none",
        zIndex: -1,
      },
      ":hover": {
        ":before": {
          background: "#f7f7f7",
        },
      },
      color: "var(--color-text-var)",
      display: "flex",
      alignItems: "center",
    }),
    icon: css({
      display: "flex",
      alignItems: "center",
      justifyContent: "center",
      svg: {
        width: 32,
        height: 32,
        aspectRatio: 1 / 1,
      },
    }),
    content: css({
      display: "flex",
      flexDirection: "column",
    }),
    title: css({
      textTransform: "uppercase",
      fontWeight: 800,
    }),
    description: css({
      fontSize: 12,
    }),
  };
</script>

<script lang="ts">
  import { css, cx } from "@emotion/css";
  import type { Component } from "svelte";

  const {
    title,
    description,
    icon: Icon,
    isActive = false,
  }: {
    isActive?: boolean;
    title?: string;
    description?: string;
    icon: Component;
  } = $props();
</script>

<div class={cx(style.wrapper, isActive && style.isActive)}>
  {#if !!Icon}
    <div class={style.icon}>
      <Icon />
    </div>
  {/if}
  <div class={style.content}>
    {#if !!title}
      <span class={style.title}>
        {title}
      </span>
    {/if}
    {#if !!description}
      <span class={style.description}>
        {description}
      </span>
    {/if}
  </div>
</div>
