<script lang="ts">
  import ManageDashboardNavbar from "@/components/nav/manage-dashboard-navbar.svelte";
  import Sidebar from "@/components/sidebar/sidebar.svelte";
  import screenSize from "@/core/device-detection.svelte";
  import { mq } from "@/core/utils";
  import { css, cx } from "@emotion/css";

  let { children } = $props();
  const layoutManage = {
    wrapper: css(
      mq({
        display: "flex",
        minHeight: "100vh",
      }),
    ),

    main: css({
      width: "100%",
    }),
  };
  let isSidebarOpen = $state(false);

  let isOpen = $derived.by(() => {
    return (
      (isSidebarOpen && screenSize.isMobile) ||
      screenSize.isTablet ||
      screenSize.isDekstop
    );
  });
</script>

<div class={cx(layoutManage.wrapper)}>
  {#if isOpen}
    <Sidebar />
  {/if}
  <main class={layoutManage.main}>
    {#if !screenSize.isDekstop}
      <ManageDashboardNavbar bind:isMenuToggle={isSidebarOpen} />
    {/if}
    {@render children?.()}
  </main>
</div>
