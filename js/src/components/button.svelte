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

  let button = css({
    background: "white",
    border: "none",
    boxShadow: "0px 0px 8px #dedede",
    cursor: "pointer",
    padding: 10,
    color: "black",
    width: "100%",
  });
</script>

<Button.Root {...props} class={button}>
  {@render children?.()}
</Button.Root>
