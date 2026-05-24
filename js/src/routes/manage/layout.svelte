<script lang="ts">
  import ManageDashboardNavbar from "@/components/nav/manage-dashboard-navbar.svelte";
  import NavLink from "@/components/sidebar/nav-item.svelte";
  import screenSize from "@/core/device-detection.svelte";
  import { mq } from "@/core/utils";
  import DuolingoIcon from "@/icons/duolingo-icon.svelte";
  import FeaturesIcon from "@/icons/features-icon.svelte";
  import IconHome from "@/icons/icon-home.svelte";
  import ListIcon from "@/icons/list-icon.svelte";
  import MoreIcon from "@/icons/more-icon.svelte";
  import { css, cx } from "@emotion/css";

  let { children } = $props();
  const layoutManage = {
    wrapper: css(
      mq({
        display: "flex",
        minHeight: "100vh",
      }),
    ),
    sidebar: {
      menu: css({
        display: "flex",
        flexDirection: "column",
        gap: 10,
        padding: 10,
      }),
      wrapper: css(
        mq({
          maxWidth: 260,
          minWidth: 260,
          borderRight: "2px solid",
          borderColor: "#e5e5e5",
          height: "100vh",
          top:["56px",0],
          display: ["block"],
          position: ["fixed", "relative"],
          background: "white",
          zIndex: 99,
        }),
      ),
      brand: css(mq({
        display:["none","inline-block"],
        paddingTop: 24,
        paddingInline: 10,
        paddingBottom: 5,
        fontWeight: "bold",
        svg: {
          width: "50%",
        },
      })),
    },
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
    <aside class={layoutManage.sidebar.wrapper}>
      <div class={layoutManage.sidebar.brand}>
        <DuolingoIcon />
      </div>
      <div class={layoutManage.sidebar.menu}>
        <NavLink
          isActive={true}
          title="PROFILE"
          description="@dadanhdyt - creator"
          icon={IconHome}
        />
        <NavLink
          title="TRANSAKSI"
          description="Kelola transaksi"
          icon={ListIcon}
        />
        <NavLink
          title="FEATURES"
          description="Overlay,SondBoard,dll"
          icon={FeaturesIcon}
        />
        <NavLink title="MORE" icon={MoreIcon} />
      </div>
    </aside>
  {/if}
  <main class={layoutManage.main}>
    {#if !screenSize.isDekstop}
      <ManageDashboardNavbar bind:isMenuToggle={isSidebarOpen} />
    {/if}
    {isOpen}
    {@render children?.()}
  </main>
</div>
