<script module lang="ts">
  let wrapper = css(
    mq({
      height: 56,
      borderBottom: "2px solid var(--color-disabled)",
      display: ["flex", "none"],
      position: "sticky",
      top: 0,
      width: "100%",
      alignItems: "center",
      paddingInline: 10,
    }),
  );
  let right = css(
    mq({
      marginLeft: "auto",
    }),
  );
</script>

<script lang="ts">
  import { mq } from "@/core/utils";
  import CloseIcon from "@/icons/close-icon.svelte";
  import IconMenu from "@/icons/icon-menu.svelte";
  import { css } from "@emotion/css";
  import { Avatar, Button } from "bits-ui";

  let {
    isMenuToggle = $bindable(),
  }: {
    isMenuToggle: boolean;
  } = $props();

  function toggle() {
    isMenuToggle = !isMenuToggle;
  }
</script>

<nav class={wrapper}>
  <Button.Root onclick={toggle}>
    {#if isMenuToggle}
      <CloseIcon />
    {:else}
      <IconMenu />
    {/if}
  </Button.Root>
  <div class={right}>
    <Avatar.Root>
      <Avatar.Fallback>D</Avatar.Fallback>
    </Avatar.Root>
  </div>
</nav>
