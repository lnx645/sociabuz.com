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
    cursor: "pointer",
    transition: "all 0.15s ease-in-out",
    paddingBlock: 10,
    paddingInline: 18,
    borderWidth: 1,
    borderStyle: "solid",
    borderColor: "rgb(61, 127, 222)",
    borderRadius: 7,
    backgroundColor: "#833AB4",
    background:
      "linear-gradient(rgba(255, 255, 255, 0.24), rgba(255, 255, 255, 0)), rgb(25, 114, 245)",
    color: "white",
    boxShadow:
      "rgba(255, 255, 255, 0.2) 0px 0px 0px 1px inset,rgba(0, 0, 0, 0.05) 0px 1px 2px 0px",
    fontWeight: 600,
    width: "100%",
    ":not(:disabled)": {
      ":hover": {
        borderColor: "#1d58ac",
      },
      ":active": {
        borderColor: "#1d58ac",
      },

      ":focus-visible": {
        borderColor: "#1d58ac",
      },
    },
    ":disabled": {
      opacity: 0.6,
      cursor: "not-allowed",
      transform: "scale(1)",
      filter: "saturate(50%)",
    },
  });
</script>

<Button.Root {...props} class={button}>
  {@render children?.()}
</Button.Root>
