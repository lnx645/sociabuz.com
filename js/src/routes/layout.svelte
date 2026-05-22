<script lang="ts">
  import { injectGlobal } from "@emotion/css";

  let { children } = $props();
  injectGlobal({
    body: {
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
</script>

<div>
  {@render children?.()}
</div>
