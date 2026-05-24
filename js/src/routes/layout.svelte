<script lang="ts">
  import { InitializeScreenSize } from "@/core/device-detection.svelte";
  import { injectGlobal } from "@emotion/css";
  import { onMount } from "svelte";

  let { children } = $props();
  injectGlobal({
    body: {
      WebkitTapHighlightColor:"transparent",
      padding: 0,
      fontFamily: `"Quicksand", sans-serif`,
      fontOpticalSizing: "auto",
      margin: 0,
    },
    "*": {
      fontFamily: `"Quicksand", sans-serif`,
    },
    "#app": {
      display: "contents",
    },
  });
  onMount(()=>{
    InitializeScreenSize()
  })
</script>

<div>
  {@render children?.()}
</div>
