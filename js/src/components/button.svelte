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
        borderColor: "rgb(29, 22, 172)",
        background:
          "linear-gradient(rgba(255, 255, 255, 0.35), rgba(255, 255, 255, 0.05)), rgb(25, 114, 245)",
        boxShadow:
          "rgba(255, 255, 255, 0.3) 0px 0px 0px 1px inset, rgba(0, 0, 0, 0.1) 0px 3px 5px 0px",
      },
      ":active:not(:disabled)": {
        background:
          "linear-gradient(rgba(0, 0, 0, 0.05), rgba(0, 0, 0, 0.1)), rgb(25, 114, 245)",
        boxShadow: "rgba(0, 0, 0, 0.15) 0px 3px 6px inset",
      },

      ":focus-visible:not(:disabled)": {
        outline: "2px solid rgb(25, 114, 245)",
        outlineOffset: "2px",
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
